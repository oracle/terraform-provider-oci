// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// BlueGreenDeploymentsClient a client for BlueGreenDeployments
type BlueGreenDeploymentsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewBlueGreenDeploymentsClientWithConfigurationProvider Creates a new default BlueGreenDeployments client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewBlueGreenDeploymentsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client BlueGreenDeploymentsClient, err error) {
	if enabled := common.CheckForEnabledServices("mysql"); !enabled {
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
	return newBlueGreenDeploymentsClientFromBaseClient(baseClient, provider)
}

// NewBlueGreenDeploymentsClientWithOboToken Creates a new default BlueGreenDeployments client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewBlueGreenDeploymentsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client BlueGreenDeploymentsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newBlueGreenDeploymentsClientFromBaseClient(baseClient, configProvider)
}

func newBlueGreenDeploymentsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client BlueGreenDeploymentsClient, err error) {
	// BlueGreenDeployments service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("BlueGreenDeployments"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = BlueGreenDeploymentsClient{BaseClient: baseClient}
	client.BasePath = "20190415"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *BlueGreenDeploymentsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("mysql", "https://mysql.{region}.ocp.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *BlueGreenDeploymentsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *BlueGreenDeploymentsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeBlueGreenDeploymentCompartment Moves only the blue/green deployment wrapper resource into a different compartment.
// This operation does not move the source DB system, target DB system, or replication channel referenced by the deployment.
// When provided, If-Match is checked against ETag values of the blue/green deployment.
// A default retry strategy applies to this operation ChangeBlueGreenDeploymentCompartment()
func (client BlueGreenDeploymentsClient) ChangeBlueGreenDeploymentCompartment(ctx context.Context, request ChangeBlueGreenDeploymentCompartmentRequest) (response ChangeBlueGreenDeploymentCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeBlueGreenDeploymentCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeBlueGreenDeploymentCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeBlueGreenDeploymentCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeBlueGreenDeploymentCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeBlueGreenDeploymentCompartmentResponse")
	}
	return
}

// changeBlueGreenDeploymentCompartment implements the OCIOperation interface (enables retrying operations)
func (client BlueGreenDeploymentsClient) changeBlueGreenDeploymentCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/blueGreenDeployments/{blueGreenDeploymentId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeBlueGreenDeploymentCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "blueGreenDeployments", "ChangeBlueGreenDeploymentCompartment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/mysql/20190415/BlueGreenDeployment/ChangeBlueGreenDeploymentCompartment"
		err = common.PostProcessServiceError(err, "BlueGreenDeployments", "ChangeBlueGreenDeploymentCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateBlueGreenDeployment Creates a blue/green deployment resource and its replication channel.
// A default retry strategy applies to this operation CreateBlueGreenDeployment()
func (client BlueGreenDeploymentsClient) CreateBlueGreenDeployment(ctx context.Context, request CreateBlueGreenDeploymentRequest) (response CreateBlueGreenDeploymentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createBlueGreenDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateBlueGreenDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateBlueGreenDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateBlueGreenDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateBlueGreenDeploymentResponse")
	}
	return
}

// createBlueGreenDeployment implements the OCIOperation interface (enables retrying operations)
func (client BlueGreenDeploymentsClient) createBlueGreenDeployment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/blueGreenDeployments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateBlueGreenDeploymentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "blueGreenDeployments", "CreateBlueGreenDeployment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "BlueGreenDeployments", "CreateBlueGreenDeployment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteBlueGreenDeployment Deletes only the blue/green deployment wrapper resource. This operation does not delete the source DB system, target DB system, or replication channel associated with the deployment.
// A default retry strategy applies to this operation DeleteBlueGreenDeployment()
func (client BlueGreenDeploymentsClient) DeleteBlueGreenDeployment(ctx context.Context, request DeleteBlueGreenDeploymentRequest) (response DeleteBlueGreenDeploymentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteBlueGreenDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteBlueGreenDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteBlueGreenDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteBlueGreenDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteBlueGreenDeploymentResponse")
	}
	return
}

// deleteBlueGreenDeployment implements the OCIOperation interface (enables retrying operations)
func (client BlueGreenDeploymentsClient) deleteBlueGreenDeployment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/blueGreenDeployments/{blueGreenDeploymentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteBlueGreenDeploymentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "blueGreenDeployments", "DeleteBlueGreenDeployment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/mysql/20190415/BlueGreenDeployment/DeleteBlueGreenDeployment"
		err = common.PostProcessServiceError(err, "BlueGreenDeployments", "DeleteBlueGreenDeployment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetBlueGreenDeployment Gets a blue/green deployment by identifier.
// A default retry strategy applies to this operation GetBlueGreenDeployment()
func (client BlueGreenDeploymentsClient) GetBlueGreenDeployment(ctx context.Context, request GetBlueGreenDeploymentRequest) (response GetBlueGreenDeploymentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getBlueGreenDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetBlueGreenDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetBlueGreenDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetBlueGreenDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetBlueGreenDeploymentResponse")
	}
	return
}

// getBlueGreenDeployment implements the OCIOperation interface (enables retrying operations)
func (client BlueGreenDeploymentsClient) getBlueGreenDeployment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/blueGreenDeployments/{blueGreenDeploymentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetBlueGreenDeploymentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "blueGreenDeployments", "GetBlueGreenDeployment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/mysql/20190415/BlueGreenDeployment/GetBlueGreenDeployment"
		err = common.PostProcessServiceError(err, "BlueGreenDeployments", "GetBlueGreenDeployment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListBlueGreenDeployments Lists blue/green deployments in a compartment.
// A default retry strategy applies to this operation ListBlueGreenDeployments()
func (client BlueGreenDeploymentsClient) ListBlueGreenDeployments(ctx context.Context, request ListBlueGreenDeploymentsRequest) (response ListBlueGreenDeploymentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listBlueGreenDeployments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListBlueGreenDeploymentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListBlueGreenDeploymentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListBlueGreenDeploymentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListBlueGreenDeploymentsResponse")
	}
	return
}

// listBlueGreenDeployments implements the OCIOperation interface (enables retrying operations)
func (client BlueGreenDeploymentsClient) listBlueGreenDeployments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/blueGreenDeployments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListBlueGreenDeploymentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "blueGreenDeployments", "ListBlueGreenDeployments")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/mysql/20190415/BlueGreenDeploymentCollection/ListBlueGreenDeployments"
		err = common.PostProcessServiceError(err, "BlueGreenDeployments", "ListBlueGreenDeployments", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SwitchoverBlueGreenDeployment Initiates switchover for a blue/green deployment.
// This action is asynchronous and supports idempotent retry only when the
// same `opc-retry-token` is reused for an equivalent switchover request on
// the same blue/green deployment. Equivalent request identity includes
// materially equivalent request parameters after applying defaults and
// preconditions, including the effective `waitTimeInSeconds`
// from the request body and the same `If-Match` value when provided. An
// empty switchover details body and a body that explicitly passes the
// default `waitTimeInSeconds=300` are equivalent.
// Requests that change `If-Match` or other material preconditions are
// treated as new requests and follow normal conflict (`409`) or
// precondition (`412`) handling instead of ordinary retry reuse.
// Switchover outcome and partial/failure conditions are surfaced through
// `lifecycleState`, `lifecycleDetails`, `switchoverStatus`,
// and `activeDbSystemId` in the
// blue/green deployment resource.
// When `waitTimeInSeconds` is omitted from the request body, the default
// wait time applies. The effective wait time provides upper bound guidance
// for waiting while the target DB System applies remaining replication
// changes during switchover processing.
// A default retry strategy applies to this operation SwitchoverBlueGreenDeployment()
func (client BlueGreenDeploymentsClient) SwitchoverBlueGreenDeployment(ctx context.Context, request SwitchoverBlueGreenDeploymentRequest) (response SwitchoverBlueGreenDeploymentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.switchoverBlueGreenDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SwitchoverBlueGreenDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SwitchoverBlueGreenDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SwitchoverBlueGreenDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SwitchoverBlueGreenDeploymentResponse")
	}
	return
}

// switchoverBlueGreenDeployment implements the OCIOperation interface (enables retrying operations)
func (client BlueGreenDeploymentsClient) switchoverBlueGreenDeployment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/blueGreenDeployments/{blueGreenDeploymentId}/actions/switchover", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SwitchoverBlueGreenDeploymentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "blueGreenDeployments", "SwitchoverBlueGreenDeployment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/mysql/20190415/BlueGreenDeployment/SwitchoverBlueGreenDeployment"
		err = common.PostProcessServiceError(err, "BlueGreenDeployments", "SwitchoverBlueGreenDeployment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateBlueGreenDeployment Updates mutable metadata for a blue/green deployment.
// Currently only `displayName`, `freeformTags`, and `definedTags` can be changed.
// A default retry strategy applies to this operation UpdateBlueGreenDeployment()
func (client BlueGreenDeploymentsClient) UpdateBlueGreenDeployment(ctx context.Context, request UpdateBlueGreenDeploymentRequest) (response UpdateBlueGreenDeploymentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateBlueGreenDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateBlueGreenDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateBlueGreenDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateBlueGreenDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateBlueGreenDeploymentResponse")
	}
	return
}

// updateBlueGreenDeployment implements the OCIOperation interface (enables retrying operations)
func (client BlueGreenDeploymentsClient) updateBlueGreenDeployment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/blueGreenDeployments/{blueGreenDeploymentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateBlueGreenDeploymentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "blueGreenDeployments", "UpdateBlueGreenDeployment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/mysql/20190415/BlueGreenDeployment/UpdateBlueGreenDeployment"
		err = common.PostProcessServiceError(err, "BlueGreenDeployments", "UpdateBlueGreenDeployment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
