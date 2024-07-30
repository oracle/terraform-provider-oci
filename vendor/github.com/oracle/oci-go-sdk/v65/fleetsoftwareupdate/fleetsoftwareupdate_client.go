// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// FleetSoftwareUpdateClient a client for FleetSoftwareUpdate
type FleetSoftwareUpdateClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewFleetSoftwareUpdateClientWithConfigurationProvider Creates a new default FleetSoftwareUpdate client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewFleetSoftwareUpdateClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client FleetSoftwareUpdateClient, err error) {
	if enabled := common.CheckForEnabledServices("fleetsoftwareupdate"); !enabled {
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
	return newFleetSoftwareUpdateClientFromBaseClient(baseClient, provider)
}

// NewFleetSoftwareUpdateClientWithOboToken Creates a new default FleetSoftwareUpdate client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewFleetSoftwareUpdateClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client FleetSoftwareUpdateClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newFleetSoftwareUpdateClientFromBaseClient(baseClient, configProvider)
}

func newFleetSoftwareUpdateClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client FleetSoftwareUpdateClient, err error) {
	// FleetSoftwareUpdate service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("FleetSoftwareUpdate"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = FleetSoftwareUpdateClient{BaseClient: baseClient}
	client.BasePath = "20220528"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *FleetSoftwareUpdateClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("fleetsoftwareupdate", "https://fleet-software-update.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *FleetSoftwareUpdateClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *FleetSoftwareUpdateClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AbortFsuDiscovery Aborts Exadata Fleet Update Discovery in progress.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/AbortFsuDiscovery.go.html to see an example of how to use AbortFsuDiscovery API.
// A default retry strategy applies to this operation AbortFsuDiscovery()
func (client FleetSoftwareUpdateClient) AbortFsuDiscovery(ctx context.Context, request AbortFsuDiscoveryRequest) (response AbortFsuDiscoveryResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.abortFsuDiscovery, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AbortFsuDiscoveryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AbortFsuDiscoveryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AbortFsuDiscoveryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AbortFsuDiscoveryResponse")
	}
	return
}

// abortFsuDiscovery implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) abortFsuDiscovery(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fsuDiscoveries/{fsuDiscoveryId}/actions/abort", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AbortFsuDiscoveryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuDiscovery/AbortFsuDiscovery"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "AbortFsuDiscovery", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddFsuCollectionTargets Adds targets to an existing Exadata Fleet Update Collection.
// Targets that are already part of a different Collection with an active Fleet Software Update Cycle cannot be added.
// This operation can only be performed on Collections that do not have an Action executing under an active Fleet Software Update Cycle.
// Additionally, during an active Fleet Software Update Cycle, targets can be added only prior to executing an Apply Action. This will require running a new Stage Action for the active Cycle.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/AddFsuCollectionTargets.go.html to see an example of how to use AddFsuCollectionTargets API.
// A default retry strategy applies to this operation AddFsuCollectionTargets()
func (client FleetSoftwareUpdateClient) AddFsuCollectionTargets(ctx context.Context, request AddFsuCollectionTargetsRequest) (response AddFsuCollectionTargetsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.addFsuCollectionTargets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddFsuCollectionTargetsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddFsuCollectionTargetsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddFsuCollectionTargetsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddFsuCollectionTargetsResponse")
	}
	return
}

// addFsuCollectionTargets implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) addFsuCollectionTargets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fsuCollections/{fsuCollectionId}/targets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddFsuCollectionTargetsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuCollection/AddFsuCollectionTargets"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "AddFsuCollectionTargets", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CancelFsuAction Cancels a scheduled Action. Only applicable for Actions that have not started executing.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/CancelFsuAction.go.html to see an example of how to use CancelFsuAction API.
// A default retry strategy applies to this operation CancelFsuAction()
func (client FleetSoftwareUpdateClient) CancelFsuAction(ctx context.Context, request CancelFsuActionRequest) (response CancelFsuActionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.cancelFsuAction, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelFsuActionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelFsuActionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelFsuActionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelFsuActionResponse")
	}
	return
}

// cancelFsuAction implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) cancelFsuAction(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fsuActions/{fsuActionId}/actions/cancel", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelFsuActionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuAction/CancelFsuAction"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "CancelFsuAction", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeFsuActionCompartment Moves a Exadata Fleet Update Action resource from one compartment identifier to another.
// When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/ChangeFsuActionCompartment.go.html to see an example of how to use ChangeFsuActionCompartment API.
// A default retry strategy applies to this operation ChangeFsuActionCompartment()
func (client FleetSoftwareUpdateClient) ChangeFsuActionCompartment(ctx context.Context, request ChangeFsuActionCompartmentRequest) (response ChangeFsuActionCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeFsuActionCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeFsuActionCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeFsuActionCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeFsuActionCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeFsuActionCompartmentResponse")
	}
	return
}

// changeFsuActionCompartment implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) changeFsuActionCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fsuActions/{fsuActionId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeFsuActionCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuAction/ChangeFsuActionCompartment"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "ChangeFsuActionCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeFsuCollectionCompartment Moves a Exadata Fleet Update Collection resource from one compartment identifier to another.
// When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/ChangeFsuCollectionCompartment.go.html to see an example of how to use ChangeFsuCollectionCompartment API.
// A default retry strategy applies to this operation ChangeFsuCollectionCompartment()
func (client FleetSoftwareUpdateClient) ChangeFsuCollectionCompartment(ctx context.Context, request ChangeFsuCollectionCompartmentRequest) (response ChangeFsuCollectionCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeFsuCollectionCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeFsuCollectionCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeFsuCollectionCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeFsuCollectionCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeFsuCollectionCompartmentResponse")
	}
	return
}

// changeFsuCollectionCompartment implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) changeFsuCollectionCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fsuCollections/{fsuCollectionId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeFsuCollectionCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuCollection/ChangeFsuCollectionCompartment"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "ChangeFsuCollectionCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeFsuCycleCompartment Moves a Exadata Fleet Update Cycle resource from one compartment identifier to another.
// When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/ChangeFsuCycleCompartment.go.html to see an example of how to use ChangeFsuCycleCompartment API.
// A default retry strategy applies to this operation ChangeFsuCycleCompartment()
func (client FleetSoftwareUpdateClient) ChangeFsuCycleCompartment(ctx context.Context, request ChangeFsuCycleCompartmentRequest) (response ChangeFsuCycleCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeFsuCycleCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeFsuCycleCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeFsuCycleCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeFsuCycleCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeFsuCycleCompartmentResponse")
	}
	return
}

// changeFsuCycleCompartment implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) changeFsuCycleCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fsuCycles/{fsuCycleId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeFsuCycleCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuCycle/ChangeFsuCycleCompartment"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "ChangeFsuCycleCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeFsuDiscoveryCompartment Moves a Exadata Fleet Update Discovery resource from one compartment identifier to another.
// When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/ChangeFsuDiscoveryCompartment.go.html to see an example of how to use ChangeFsuDiscoveryCompartment API.
// A default retry strategy applies to this operation ChangeFsuDiscoveryCompartment()
func (client FleetSoftwareUpdateClient) ChangeFsuDiscoveryCompartment(ctx context.Context, request ChangeFsuDiscoveryCompartmentRequest) (response ChangeFsuDiscoveryCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeFsuDiscoveryCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeFsuDiscoveryCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeFsuDiscoveryCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeFsuDiscoveryCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeFsuDiscoveryCompartmentResponse")
	}
	return
}

// changeFsuDiscoveryCompartment implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) changeFsuDiscoveryCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fsuDiscoveries/{fsuDiscoveryId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeFsuDiscoveryCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuDiscovery/ChangeFsuDiscoveryCompartment"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "ChangeFsuDiscoveryCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CloneFsuCycle Clones existing Exadata Fleet Update Cycle details into a new Exadata Fleet Update Cycle resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/CloneFsuCycle.go.html to see an example of how to use CloneFsuCycle API.
// A default retry strategy applies to this operation CloneFsuCycle()
func (client FleetSoftwareUpdateClient) CloneFsuCycle(ctx context.Context, request CloneFsuCycleRequest) (response CloneFsuCycleResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.cloneFsuCycle, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CloneFsuCycleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CloneFsuCycleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CloneFsuCycleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CloneFsuCycleResponse")
	}
	return
}

// cloneFsuCycle implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) cloneFsuCycle(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fsuCycles/{fsuCycleId}/actions/clone", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CloneFsuCycleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuCycle/CloneFsuCycle"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "CloneFsuCycle", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &fsucycle{})
	return response, err
}

// CreateFsuAction Creates a new Exadata Fleet Update Action.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/CreateFsuAction.go.html to see an example of how to use CreateFsuAction API.
// A default retry strategy applies to this operation CreateFsuAction()
func (client FleetSoftwareUpdateClient) CreateFsuAction(ctx context.Context, request CreateFsuActionRequest) (response CreateFsuActionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createFsuAction, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateFsuActionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateFsuActionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateFsuActionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateFsuActionResponse")
	}
	return
}

// createFsuAction implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) createFsuAction(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fsuActions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateFsuActionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "CreateFsuAction", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &fsuaction{})
	return response, err
}

// CreateFsuCollection Creates a new Exadata Fleet Update Collection.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/CreateFsuCollection.go.html to see an example of how to use CreateFsuCollection API.
// A default retry strategy applies to this operation CreateFsuCollection()
func (client FleetSoftwareUpdateClient) CreateFsuCollection(ctx context.Context, request CreateFsuCollectionRequest) (response CreateFsuCollectionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createFsuCollection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateFsuCollectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateFsuCollectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateFsuCollectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateFsuCollectionResponse")
	}
	return
}

// createFsuCollection implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) createFsuCollection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fsuCollections", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateFsuCollectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "CreateFsuCollection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &fsucollection{})
	return response, err
}

// CreateFsuCycle Creates a new Exadata Fleet Update Cycle.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/CreateFsuCycle.go.html to see an example of how to use CreateFsuCycle API.
// A default retry strategy applies to this operation CreateFsuCycle()
func (client FleetSoftwareUpdateClient) CreateFsuCycle(ctx context.Context, request CreateFsuCycleRequest) (response CreateFsuCycleResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createFsuCycle, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateFsuCycleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateFsuCycleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateFsuCycleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateFsuCycleResponse")
	}
	return
}

// createFsuCycle implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) createFsuCycle(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fsuCycles", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateFsuCycleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "CreateFsuCycle", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &fsucycle{})
	return response, err
}

// CreateFsuDiscovery Creates a new Exadata Fleet Update Discovery.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/CreateFsuDiscovery.go.html to see an example of how to use CreateFsuDiscovery API.
// A default retry strategy applies to this operation CreateFsuDiscovery()
func (client FleetSoftwareUpdateClient) CreateFsuDiscovery(ctx context.Context, request CreateFsuDiscoveryRequest) (response CreateFsuDiscoveryResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createFsuDiscovery, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateFsuDiscoveryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateFsuDiscoveryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateFsuDiscoveryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateFsuDiscoveryResponse")
	}
	return
}

// createFsuDiscovery implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) createFsuDiscovery(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fsuDiscoveries", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateFsuDiscoveryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "CreateFsuDiscovery", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteFsuAction Deletes a Exadata Fleet Update Action resource by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/DeleteFsuAction.go.html to see an example of how to use DeleteFsuAction API.
// A default retry strategy applies to this operation DeleteFsuAction()
func (client FleetSoftwareUpdateClient) DeleteFsuAction(ctx context.Context, request DeleteFsuActionRequest) (response DeleteFsuActionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteFsuAction, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteFsuActionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteFsuActionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteFsuActionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteFsuActionResponse")
	}
	return
}

// deleteFsuAction implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) deleteFsuAction(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/fsuActions/{fsuActionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteFsuActionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuAction/DeleteFsuAction"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "DeleteFsuAction", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteFsuCollection Deletes a Exadata Fleet Update Collection resource by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/DeleteFsuCollection.go.html to see an example of how to use DeleteFsuCollection API.
// A default retry strategy applies to this operation DeleteFsuCollection()
func (client FleetSoftwareUpdateClient) DeleteFsuCollection(ctx context.Context, request DeleteFsuCollectionRequest) (response DeleteFsuCollectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteFsuCollection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteFsuCollectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteFsuCollectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteFsuCollectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteFsuCollectionResponse")
	}
	return
}

// deleteFsuCollection implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) deleteFsuCollection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/fsuCollections/{fsuCollectionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteFsuCollectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuCollection/DeleteFsuCollection"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "DeleteFsuCollection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteFsuCollectionTarget Removes a target from an existing Exadata Fleet Update Collection.
// This operation can only be performed on Collections that do not have an Action executing under an active Fleet Software Update Cycle.
// Additionally, during an active Fleet Software Update Cycle, a target can be removed only prior to executing an Apply Action.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/DeleteFsuCollectionTarget.go.html to see an example of how to use DeleteFsuCollectionTarget API.
// A default retry strategy applies to this operation DeleteFsuCollectionTarget()
func (client FleetSoftwareUpdateClient) DeleteFsuCollectionTarget(ctx context.Context, request DeleteFsuCollectionTargetRequest) (response DeleteFsuCollectionTargetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteFsuCollectionTarget, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteFsuCollectionTargetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteFsuCollectionTargetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteFsuCollectionTargetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteFsuCollectionTargetResponse")
	}
	return
}

// deleteFsuCollectionTarget implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) deleteFsuCollectionTarget(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/fsuCollections/{fsuCollectionId}/targets/{targetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteFsuCollectionTargetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuCollection/DeleteFsuCollectionTarget"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "DeleteFsuCollectionTarget", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteFsuCycle Deletes a Exadata Fleet Update Cycle resource by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/DeleteFsuCycle.go.html to see an example of how to use DeleteFsuCycle API.
// A default retry strategy applies to this operation DeleteFsuCycle()
func (client FleetSoftwareUpdateClient) DeleteFsuCycle(ctx context.Context, request DeleteFsuCycleRequest) (response DeleteFsuCycleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteFsuCycle, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteFsuCycleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteFsuCycleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteFsuCycleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteFsuCycleResponse")
	}
	return
}

// deleteFsuCycle implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) deleteFsuCycle(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/fsuCycles/{fsuCycleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteFsuCycleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuCycle/DeleteFsuCycle"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "DeleteFsuCycle", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteFsuDiscovery Deletes a Exadata Fleet Update Discovery resource by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/DeleteFsuDiscovery.go.html to see an example of how to use DeleteFsuDiscovery API.
// A default retry strategy applies to this operation DeleteFsuDiscovery()
func (client FleetSoftwareUpdateClient) DeleteFsuDiscovery(ctx context.Context, request DeleteFsuDiscoveryRequest) (response DeleteFsuDiscoveryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteFsuDiscovery, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteFsuDiscoveryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteFsuDiscoveryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteFsuDiscoveryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteFsuDiscoveryResponse")
	}
	return
}

// deleteFsuDiscovery implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) deleteFsuDiscovery(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/fsuDiscoveries/{fsuDiscoveryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteFsuDiscoveryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuDiscovery/DeleteFsuDiscovery"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "DeleteFsuDiscovery", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteFsuJob Deletes the Exadata Fleet Update Job resource by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/DeleteFsuJob.go.html to see an example of how to use DeleteFsuJob API.
// A default retry strategy applies to this operation DeleteFsuJob()
func (client FleetSoftwareUpdateClient) DeleteFsuJob(ctx context.Context, request DeleteFsuJobRequest) (response DeleteFsuJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteFsuJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteFsuJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteFsuJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteFsuJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteFsuJobResponse")
	}
	return
}

// deleteFsuJob implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) deleteFsuJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/fsuJobs/{fsuJobId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteFsuJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuJob/DeleteFsuJob"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "DeleteFsuJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetFsuAction Gets a Exadata Fleet Update Action by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/GetFsuAction.go.html to see an example of how to use GetFsuAction API.
// A default retry strategy applies to this operation GetFsuAction()
func (client FleetSoftwareUpdateClient) GetFsuAction(ctx context.Context, request GetFsuActionRequest) (response GetFsuActionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFsuAction, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFsuActionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFsuActionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFsuActionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFsuActionResponse")
	}
	return
}

// getFsuAction implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) getFsuAction(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fsuActions/{fsuActionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetFsuActionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuAction/GetFsuAction"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "GetFsuAction", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &fsuaction{})
	return response, err
}

// GetFsuActionOutputContent Gets the Exadata Fleet Update Action Output content as a binary file (string).
// This will only include the output from FAILED Exadata Fleet Update Jobs. No content in case there are no FAILED jobs.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/GetFsuActionOutputContent.go.html to see an example of how to use GetFsuActionOutputContent API.
// A default retry strategy applies to this operation GetFsuActionOutputContent()
func (client FleetSoftwareUpdateClient) GetFsuActionOutputContent(ctx context.Context, request GetFsuActionOutputContentRequest) (response GetFsuActionOutputContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFsuActionOutputContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFsuActionOutputContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFsuActionOutputContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFsuActionOutputContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFsuActionOutputContentResponse")
	}
	return
}

// getFsuActionOutputContent implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) getFsuActionOutputContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fsuActions/{fsuActionId}/output/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetFsuActionOutputContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuAction/GetFsuActionOutputContent"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "GetFsuActionOutputContent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetFsuCollection Gets a Exadata Fleet Update Collection by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/GetFsuCollection.go.html to see an example of how to use GetFsuCollection API.
// A default retry strategy applies to this operation GetFsuCollection()
func (client FleetSoftwareUpdateClient) GetFsuCollection(ctx context.Context, request GetFsuCollectionRequest) (response GetFsuCollectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFsuCollection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFsuCollectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFsuCollectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFsuCollectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFsuCollectionResponse")
	}
	return
}

// getFsuCollection implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) getFsuCollection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fsuCollections/{fsuCollectionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetFsuCollectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuCollection/GetFsuCollection"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "GetFsuCollection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &fsucollection{})
	return response, err
}

// GetFsuCollectionTarget Gets a Exadata Fleet Update Collection Target by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/GetFsuCollectionTarget.go.html to see an example of how to use GetFsuCollectionTarget API.
// A default retry strategy applies to this operation GetFsuCollectionTarget()
func (client FleetSoftwareUpdateClient) GetFsuCollectionTarget(ctx context.Context, request GetFsuCollectionTargetRequest) (response GetFsuCollectionTargetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFsuCollectionTarget, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFsuCollectionTargetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFsuCollectionTargetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFsuCollectionTargetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFsuCollectionTargetResponse")
	}
	return
}

// getFsuCollectionTarget implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) getFsuCollectionTarget(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fsuCollections/{fsuCollectionId}/targets/{targetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetFsuCollectionTargetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuCollectionTarget/GetFsuCollectionTarget"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "GetFsuCollectionTarget", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetFsuCycle Gets a Exadata Fleet Update Cycle by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/GetFsuCycle.go.html to see an example of how to use GetFsuCycle API.
// A default retry strategy applies to this operation GetFsuCycle()
func (client FleetSoftwareUpdateClient) GetFsuCycle(ctx context.Context, request GetFsuCycleRequest) (response GetFsuCycleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFsuCycle, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFsuCycleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFsuCycleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFsuCycleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFsuCycleResponse")
	}
	return
}

// getFsuCycle implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) getFsuCycle(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fsuCycles/{fsuCycleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetFsuCycleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuCycle/GetFsuCycle"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "GetFsuCycle", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &fsucycle{})
	return response, err
}

// GetFsuDiscovery Gets a Exadata Fleet Update Discovery by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/GetFsuDiscovery.go.html to see an example of how to use GetFsuDiscovery API.
// A default retry strategy applies to this operation GetFsuDiscovery()
func (client FleetSoftwareUpdateClient) GetFsuDiscovery(ctx context.Context, request GetFsuDiscoveryRequest) (response GetFsuDiscoveryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFsuDiscovery, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFsuDiscoveryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFsuDiscoveryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFsuDiscoveryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFsuDiscoveryResponse")
	}
	return
}

// getFsuDiscovery implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) getFsuDiscovery(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fsuDiscoveries/{fsuDiscoveryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetFsuDiscoveryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuDiscovery/GetFsuDiscovery"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "GetFsuDiscovery", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetFsuJob Gets a Exadata Fleet Update Job by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/GetFsuJob.go.html to see an example of how to use GetFsuJob API.
// A default retry strategy applies to this operation GetFsuJob()
func (client FleetSoftwareUpdateClient) GetFsuJob(ctx context.Context, request GetFsuJobRequest) (response GetFsuJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFsuJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFsuJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFsuJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFsuJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFsuJobResponse")
	}
	return
}

// getFsuJob implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) getFsuJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fsuJobs/{fsuJobId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetFsuJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuJob/GetFsuJob"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "GetFsuJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &fsujob{})
	return response, err
}

// GetFsuJobOutputContent Get the Exadata Fleet Update Job Output content as a binary file (string).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/GetFsuJobOutputContent.go.html to see an example of how to use GetFsuJobOutputContent API.
// A default retry strategy applies to this operation GetFsuJobOutputContent()
func (client FleetSoftwareUpdateClient) GetFsuJobOutputContent(ctx context.Context, request GetFsuJobOutputContentRequest) (response GetFsuJobOutputContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFsuJobOutputContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFsuJobOutputContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFsuJobOutputContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFsuJobOutputContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFsuJobOutputContentResponse")
	}
	return
}

// getFsuJobOutputContent implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) getFsuJobOutputContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fsuJobs/{fsuJobId}/output/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetFsuJobOutputContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuJob/GetFsuJobOutputContent"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "GetFsuJobOutputContent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets the status of the work request with the specified ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
// A default retry strategy applies to this operation GetWorkRequest()
func (client FleetSoftwareUpdateClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
func (client FleetSoftwareUpdateClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListFsuActions Gets a list of all Exadata Fleet Update Actions in a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/ListFsuActions.go.html to see an example of how to use ListFsuActions API.
// A default retry strategy applies to this operation ListFsuActions()
func (client FleetSoftwareUpdateClient) ListFsuActions(ctx context.Context, request ListFsuActionsRequest) (response ListFsuActionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listFsuActions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListFsuActionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListFsuActionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListFsuActionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListFsuActionsResponse")
	}
	return
}

// listFsuActions implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) listFsuActions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fsuActions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListFsuActionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuActionSummaryCollection/ListFsuActions"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "ListFsuActions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListFsuCollectionTargets Gets a list of all Targets that are members of a specific Exadata Fleet Update Collection.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/ListFsuCollectionTargets.go.html to see an example of how to use ListFsuCollectionTargets API.
// A default retry strategy applies to this operation ListFsuCollectionTargets()
func (client FleetSoftwareUpdateClient) ListFsuCollectionTargets(ctx context.Context, request ListFsuCollectionTargetsRequest) (response ListFsuCollectionTargetsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listFsuCollectionTargets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListFsuCollectionTargetsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListFsuCollectionTargetsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListFsuCollectionTargetsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListFsuCollectionTargetsResponse")
	}
	return
}

// listFsuCollectionTargets implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) listFsuCollectionTargets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fsuCollections/{fsuCollectionId}/targets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListFsuCollectionTargetsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/TargetSummaryCollection/ListFsuCollectionTargets"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "ListFsuCollectionTargets", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListFsuCollections Gets a list of all Exadata Fleet Update Collections in a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/ListFsuCollections.go.html to see an example of how to use ListFsuCollections API.
// A default retry strategy applies to this operation ListFsuCollections()
func (client FleetSoftwareUpdateClient) ListFsuCollections(ctx context.Context, request ListFsuCollectionsRequest) (response ListFsuCollectionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listFsuCollections, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListFsuCollectionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListFsuCollectionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListFsuCollectionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListFsuCollectionsResponse")
	}
	return
}

// listFsuCollections implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) listFsuCollections(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fsuCollections", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListFsuCollectionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuCollectionSummaryCollection/ListFsuCollections"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "ListFsuCollections", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListFsuCycles Gets a list of all Exadata Fleet Update Cycles in a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/ListFsuCycles.go.html to see an example of how to use ListFsuCycles API.
// A default retry strategy applies to this operation ListFsuCycles()
func (client FleetSoftwareUpdateClient) ListFsuCycles(ctx context.Context, request ListFsuCyclesRequest) (response ListFsuCyclesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listFsuCycles, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListFsuCyclesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListFsuCyclesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListFsuCyclesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListFsuCyclesResponse")
	}
	return
}

// listFsuCycles implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) listFsuCycles(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fsuCycles", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListFsuCyclesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuCycleSummary/ListFsuCycles"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "ListFsuCycles", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListFsuDiscoveries Returns a list of Exadata Fleet Update Discoveries resources in the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/ListFsuDiscoveries.go.html to see an example of how to use ListFsuDiscoveries API.
// A default retry strategy applies to this operation ListFsuDiscoveries()
func (client FleetSoftwareUpdateClient) ListFsuDiscoveries(ctx context.Context, request ListFsuDiscoveriesRequest) (response ListFsuDiscoveriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listFsuDiscoveries, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListFsuDiscoveriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListFsuDiscoveriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListFsuDiscoveriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListFsuDiscoveriesResponse")
	}
	return
}

// listFsuDiscoveries implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) listFsuDiscoveries(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fsuDiscoveries", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListFsuDiscoveriesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuDiscoverySummary/ListFsuDiscoveries"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "ListFsuDiscoveries", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListFsuDiscoveryTargets Gets a list of all Targets in the results of a Exadata Fleet Update Discovery.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/ListFsuDiscoveryTargets.go.html to see an example of how to use ListFsuDiscoveryTargets API.
// A default retry strategy applies to this operation ListFsuDiscoveryTargets()
func (client FleetSoftwareUpdateClient) ListFsuDiscoveryTargets(ctx context.Context, request ListFsuDiscoveryTargetsRequest) (response ListFsuDiscoveryTargetsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listFsuDiscoveryTargets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListFsuDiscoveryTargetsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListFsuDiscoveryTargetsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListFsuDiscoveryTargetsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListFsuDiscoveryTargetsResponse")
	}
	return
}

// listFsuDiscoveryTargets implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) listFsuDiscoveryTargets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fsuDiscoveries/{fsuDiscoveryId}/targets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListFsuDiscoveryTargetsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/TargetSummaryCollection/ListFsuDiscoveryTargets"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "ListFsuDiscoveryTargets", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListFsuJobOutputs Lists the Exadata Fleet Update Job Output messages, if any.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/ListFsuJobOutputs.go.html to see an example of how to use ListFsuJobOutputs API.
// A default retry strategy applies to this operation ListFsuJobOutputs()
func (client FleetSoftwareUpdateClient) ListFsuJobOutputs(ctx context.Context, request ListFsuJobOutputsRequest) (response ListFsuJobOutputsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listFsuJobOutputs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListFsuJobOutputsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListFsuJobOutputsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListFsuJobOutputsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListFsuJobOutputsResponse")
	}
	return
}

// listFsuJobOutputs implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) listFsuJobOutputs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fsuJobs/{fsuJobId}/output", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListFsuJobOutputsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuJobOutputSummary/ListFsuJobOutputs"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "ListFsuJobOutputs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListFsuJobs Lists all the Exadata Fleet Update Jobs associated to the specified Exadata Fleet Update Action.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/ListFsuJobs.go.html to see an example of how to use ListFsuJobs API.
// A default retry strategy applies to this operation ListFsuJobs()
func (client FleetSoftwareUpdateClient) ListFsuJobs(ctx context.Context, request ListFsuJobsRequest) (response ListFsuJobsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listFsuJobs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListFsuJobsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListFsuJobsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListFsuJobsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListFsuJobsResponse")
	}
	return
}

// listFsuJobs implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) listFsuJobs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/fsuJobs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListFsuJobsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuJobSummary/ListFsuJobs"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "ListFsuJobs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Returns a paginated list of errors for a specified Work Request..
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client FleetSoftwareUpdateClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
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
func (client FleetSoftwareUpdateClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/WorkRequestError/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Returns a paginated list of logs for a specified Work Request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client FleetSoftwareUpdateClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
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
func (client FleetSoftwareUpdateClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/WorkRequestLogEntry/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
// A default retry strategy applies to this operation ListWorkRequests()
func (client FleetSoftwareUpdateClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
func (client FleetSoftwareUpdateClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/WorkRequest/ListWorkRequests"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveFsuCollectionTargets Removes targets from an existing Exadata Fleet Update Collection.
// This operation can only be performed on Collections that do not have an Action executing under an active Fleet Software Update Cycle.
// Additionally, during an active Fleet Software Update Cycle, targets can be removed only prior to executing an Apply Action.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/RemoveFsuCollectionTargets.go.html to see an example of how to use RemoveFsuCollectionTargets API.
// A default retry strategy applies to this operation RemoveFsuCollectionTargets()
func (client FleetSoftwareUpdateClient) RemoveFsuCollectionTargets(ctx context.Context, request RemoveFsuCollectionTargetsRequest) (response RemoveFsuCollectionTargetsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.removeFsuCollectionTargets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveFsuCollectionTargetsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveFsuCollectionTargetsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveFsuCollectionTargetsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveFsuCollectionTargetsResponse")
	}
	return
}

// removeFsuCollectionTargets implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) removeFsuCollectionTargets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/fsuCollections/{fsuCollectionId}/targets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveFsuCollectionTargetsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuCollection/RemoveFsuCollectionTargets"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "RemoveFsuCollectionTargets", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ResumeFsuAction Resumes an Action that has batches of targets waiting to execute.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/ResumeFsuAction.go.html to see an example of how to use ResumeFsuAction API.
// A default retry strategy applies to this operation ResumeFsuAction()
func (client FleetSoftwareUpdateClient) ResumeFsuAction(ctx context.Context, request ResumeFsuActionRequest) (response ResumeFsuActionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.resumeFsuAction, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ResumeFsuActionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ResumeFsuActionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ResumeFsuActionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ResumeFsuActionResponse")
	}
	return
}

// resumeFsuAction implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) resumeFsuAction(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fsuActions/{fsuActionId}/actions/resume", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ResumeFsuActionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuAction/ResumeFsuAction"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "ResumeFsuAction", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RetryFsuJob Retry a failed Job, only while the current Action is being executed.
// After the Action reaches a terminal state, a new Action of the same kind is required to retry on failed targets.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/RetryFsuJob.go.html to see an example of how to use RetryFsuJob API.
// A default retry strategy applies to this operation RetryFsuJob()
func (client FleetSoftwareUpdateClient) RetryFsuJob(ctx context.Context, request RetryFsuJobRequest) (response RetryFsuJobResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.retryFsuJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RetryFsuJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RetryFsuJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RetryFsuJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RetryFsuJobResponse")
	}
	return
}

// retryFsuJob implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) retryFsuJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/fsuJobs/{fsuJobId}/actions/retry", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RetryFsuJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuJob/RetryFsuJob"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "RetryFsuJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateFsuAction Updates the Exadata Fleet Update Action identified by the ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/UpdateFsuAction.go.html to see an example of how to use UpdateFsuAction API.
// A default retry strategy applies to this operation UpdateFsuAction()
func (client FleetSoftwareUpdateClient) UpdateFsuAction(ctx context.Context, request UpdateFsuActionRequest) (response UpdateFsuActionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateFsuAction, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateFsuActionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateFsuActionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateFsuActionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateFsuActionResponse")
	}
	return
}

// updateFsuAction implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) updateFsuAction(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/fsuActions/{fsuActionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateFsuActionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuAction/UpdateFsuAction"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "UpdateFsuAction", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateFsuCollection Updates the Exadata Fleet Update Collection identified by the ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/UpdateFsuCollection.go.html to see an example of how to use UpdateFsuCollection API.
// A default retry strategy applies to this operation UpdateFsuCollection()
func (client FleetSoftwareUpdateClient) UpdateFsuCollection(ctx context.Context, request UpdateFsuCollectionRequest) (response UpdateFsuCollectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateFsuCollection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateFsuCollectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateFsuCollectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateFsuCollectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateFsuCollectionResponse")
	}
	return
}

// updateFsuCollection implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) updateFsuCollection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/fsuCollections/{fsuCollectionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateFsuCollectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuCollection/UpdateFsuCollection"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "UpdateFsuCollection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateFsuCycle Updates the Exadata Fleet Update Cycle identified by the ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/UpdateFsuCycle.go.html to see an example of how to use UpdateFsuCycle API.
// A default retry strategy applies to this operation UpdateFsuCycle()
func (client FleetSoftwareUpdateClient) UpdateFsuCycle(ctx context.Context, request UpdateFsuCycleRequest) (response UpdateFsuCycleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateFsuCycle, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateFsuCycleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateFsuCycleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateFsuCycleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateFsuCycleResponse")
	}
	return
}

// updateFsuCycle implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) updateFsuCycle(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/fsuCycles/{fsuCycleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateFsuCycleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuCycle/UpdateFsuCycle"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "UpdateFsuCycle", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateFsuDiscovery Updates the Exadata Fleet Update Discovery identified by the ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/UpdateFsuDiscovery.go.html to see an example of how to use UpdateFsuDiscovery API.
// A default retry strategy applies to this operation UpdateFsuDiscovery()
func (client FleetSoftwareUpdateClient) UpdateFsuDiscovery(ctx context.Context, request UpdateFsuDiscoveryRequest) (response UpdateFsuDiscoveryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateFsuDiscovery, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateFsuDiscoveryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateFsuDiscoveryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateFsuDiscoveryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateFsuDiscoveryResponse")
	}
	return
}

// updateFsuDiscovery implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) updateFsuDiscovery(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/fsuDiscoveries/{fsuDiscoveryId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateFsuDiscoveryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuDiscovery/UpdateFsuDiscovery"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "UpdateFsuDiscovery", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateFsuJob Updates Exadata Fleet Update Job resource details.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/UpdateFsuJob.go.html to see an example of how to use UpdateFsuJob API.
// A default retry strategy applies to this operation UpdateFsuJob()
func (client FleetSoftwareUpdateClient) UpdateFsuJob(ctx context.Context, request UpdateFsuJobRequest) (response UpdateFsuJobResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateFsuJob, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateFsuJobResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateFsuJobResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateFsuJobResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateFsuJobResponse")
	}
	return
}

// updateFsuJob implements the OCIOperation interface (enables retrying operations)
func (client FleetSoftwareUpdateClient) updateFsuJob(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/fsuJobs/{fsuJobId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateFsuJobResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/edsfu/20220528/FsuJob/UpdateFsuJob"
		err = common.PostProcessServiceError(err, "FleetSoftwareUpdate", "UpdateFsuJob", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &fsujob{})
	return response, err
}
