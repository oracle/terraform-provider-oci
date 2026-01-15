// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// WebLogic Management Service API
//
// WebLogic Management Service is an OCI service that enables a unified view and management of WebLogic domains
// in Oracle Cloud Infrastructure. Features include on-demand patching of WebLogic domains, rollback of the
// last applied patch, discovery and management of WebLogic instances on a compute host.
//

package wlms

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// WeblogicManagementServiceClient a client for WeblogicManagementService
type WeblogicManagementServiceClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewWeblogicManagementServiceClientWithConfigurationProvider Creates a new default WeblogicManagementService client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewWeblogicManagementServiceClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client WeblogicManagementServiceClient, err error) {
	if enabled := common.CheckForEnabledServices("wlms"); !enabled {
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
	return newWeblogicManagementServiceClientFromBaseClient(baseClient, provider)
}

// NewWeblogicManagementServiceClientWithOboToken Creates a new default WeblogicManagementService client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewWeblogicManagementServiceClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client WeblogicManagementServiceClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newWeblogicManagementServiceClientFromBaseClient(baseClient, configProvider)
}

func newWeblogicManagementServiceClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client WeblogicManagementServiceClient, err error) {
	// WeblogicManagementService service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("WeblogicManagementService"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = WeblogicManagementServiceClient{BaseClient: baseClient}
	client.BasePath = "20241101"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *WeblogicManagementServiceClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("wlms", "https://api.weblogicmanagement.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *WeblogicManagementServiceClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *WeblogicManagementServiceClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeWlsDomainCompartment Moves a WebLogic domain into a different compartment within the same tenancy.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/ChangeWlsDomainCompartment.go.html to see an example of how to use ChangeWlsDomainCompartment API.
// A default retry strategy applies to this operation ChangeWlsDomainCompartment()
func (client WeblogicManagementServiceClient) ChangeWlsDomainCompartment(ctx context.Context, request ChangeWlsDomainCompartmentRequest) (response ChangeWlsDomainCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeWlsDomainCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeWlsDomainCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeWlsDomainCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeWlsDomainCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeWlsDomainCompartmentResponse")
	}
	return
}

// changeWlsDomainCompartment implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) changeWlsDomainCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/wlsDomains/{wlsDomainId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeWlsDomainCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/WlsDomain/ChangeWlsDomainCompartment"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "ChangeWlsDomainCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateAgreementRecord Creates a terms of use agreement record for a WebLogic domain.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/CreateAgreementRecord.go.html to see an example of how to use CreateAgreementRecord API.
// A default retry strategy applies to this operation CreateAgreementRecord()
func (client WeblogicManagementServiceClient) CreateAgreementRecord(ctx context.Context, request CreateAgreementRecordRequest) (response CreateAgreementRecordResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createAgreementRecord, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAgreementRecordResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAgreementRecordResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAgreementRecordResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAgreementRecordResponse")
	}
	return
}

// createAgreementRecord implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) createAgreementRecord(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/wlsDomains/{wlsDomainId}/agreementRecords", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateAgreementRecordResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/WlsDomain/CreateAgreementRecord"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "CreateAgreementRecord", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteWlsDomain Delete the WebLogic domain.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/DeleteWlsDomain.go.html to see an example of how to use DeleteWlsDomain API.
// A default retry strategy applies to this operation DeleteWlsDomain()
func (client WeblogicManagementServiceClient) DeleteWlsDomain(ctx context.Context, request DeleteWlsDomainRequest) (response DeleteWlsDomainResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteWlsDomain, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteWlsDomainResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteWlsDomainResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteWlsDomainResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteWlsDomainResponse")
	}
	return
}

// deleteWlsDomain implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) deleteWlsDomain(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/wlsDomains/{wlsDomainId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteWlsDomainResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/WlsDomain/DeleteWlsDomain"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "DeleteWlsDomain", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAgreement Returns the terms and conditions of use agreement.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/GetAgreement.go.html to see an example of how to use GetAgreement API.
// A default retry strategy applies to this operation GetAgreement()
func (client WeblogicManagementServiceClient) GetAgreement(ctx context.Context, request GetAgreementRequest) (response GetAgreementResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAgreement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAgreementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAgreementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAgreementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAgreementResponse")
	}
	return
}

// getAgreement implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) getAgreement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/agreement", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAgreementResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/Agreement/GetAgreement"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "GetAgreement", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetManagedInstance Gets information about the specified managed instance.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/GetManagedInstance.go.html to see an example of how to use GetManagedInstance API.
// A default retry strategy applies to this operation GetManagedInstance()
func (client WeblogicManagementServiceClient) GetManagedInstance(ctx context.Context, request GetManagedInstanceRequest) (response GetManagedInstanceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetManagedInstanceResponse")
	}
	return
}

// getManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) getManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances/{managedInstanceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/ManagedInstance/GetManagedInstance"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "GetManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetManagedInstanceServer Gets information about the specified server in a managed instance.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/GetManagedInstanceServer.go.html to see an example of how to use GetManagedInstanceServer API.
// A default retry strategy applies to this operation GetManagedInstanceServer()
func (client WeblogicManagementServiceClient) GetManagedInstanceServer(ctx context.Context, request GetManagedInstanceServerRequest) (response GetManagedInstanceServerResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getManagedInstanceServer, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetManagedInstanceServerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetManagedInstanceServerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetManagedInstanceServerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetManagedInstanceServerResponse")
	}
	return
}

// getManagedInstanceServer implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) getManagedInstanceServer(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances/{managedInstanceId}/servers/{serverId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetManagedInstanceServerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/ManagedInstance/GetManagedInstanceServer"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "GetManagedInstanceServer", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWlsDomain Gets a specific WebLogic domain.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/GetWlsDomain.go.html to see an example of how to use GetWlsDomain API.
// A default retry strategy applies to this operation GetWlsDomain()
func (client WeblogicManagementServiceClient) GetWlsDomain(ctx context.Context, request GetWlsDomainRequest) (response GetWlsDomainResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getWlsDomain, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetWlsDomainResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetWlsDomainResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetWlsDomainResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetWlsDomainResponse")
	}
	return
}

// getWlsDomain implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) getWlsDomain(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/wlsDomains/{wlsDomainId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetWlsDomainResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/WlsDomain/GetWlsDomain"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "GetWlsDomain", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWlsDomainCredential Gets WebLogic and Node Manager credentials of a specific WebLogic domain.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/GetWlsDomainCredential.go.html to see an example of how to use GetWlsDomainCredential API.
// A default retry strategy applies to this operation GetWlsDomainCredential()
func (client WeblogicManagementServiceClient) GetWlsDomainCredential(ctx context.Context, request GetWlsDomainCredentialRequest) (response GetWlsDomainCredentialResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getWlsDomainCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetWlsDomainCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetWlsDomainCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetWlsDomainCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetWlsDomainCredentialResponse")
	}
	return
}

// getWlsDomainCredential implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) getWlsDomainCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/wlsDomains/{wlsDomainId}/credentials/{credentialType}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetWlsDomainCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/WlsDomain/GetWlsDomainCredential"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "GetWlsDomainCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWlsDomainServer Gets information about the specified server in a WebLogic domain.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/GetWlsDomainServer.go.html to see an example of how to use GetWlsDomainServer API.
// A default retry strategy applies to this operation GetWlsDomainServer()
func (client WeblogicManagementServiceClient) GetWlsDomainServer(ctx context.Context, request GetWlsDomainServerRequest) (response GetWlsDomainServerResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getWlsDomainServer, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetWlsDomainServerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetWlsDomainServerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetWlsDomainServerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetWlsDomainServerResponse")
	}
	return
}

// getWlsDomainServer implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) getWlsDomainServer(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/wlsDomains/{wlsDomainId}/servers/{serverId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetWlsDomainServerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/WlsDomain/GetWlsDomainServer"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "GetWlsDomainServer", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWlsDomainServerBackup Get details of specific backup for the WebLogic Domain.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/GetWlsDomainServerBackup.go.html to see an example of how to use GetWlsDomainServerBackup API.
// A default retry strategy applies to this operation GetWlsDomainServerBackup()
func (client WeblogicManagementServiceClient) GetWlsDomainServerBackup(ctx context.Context, request GetWlsDomainServerBackupRequest) (response GetWlsDomainServerBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getWlsDomainServerBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetWlsDomainServerBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetWlsDomainServerBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetWlsDomainServerBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetWlsDomainServerBackupResponse")
	}
	return
}

// getWlsDomainServerBackup implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) getWlsDomainServerBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/wlsDomains/{wlsDomainId}/servers/{serverId}/backups/{backupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetWlsDomainServerBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/WlsDomain/GetWlsDomainServerBackup"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "GetWlsDomainServerBackup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWlsDomainServerBackupContent Get details of specific backup for the WebLogic Domain.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/GetWlsDomainServerBackupContent.go.html to see an example of how to use GetWlsDomainServerBackupContent API.
// A default retry strategy applies to this operation GetWlsDomainServerBackupContent()
func (client WeblogicManagementServiceClient) GetWlsDomainServerBackupContent(ctx context.Context, request GetWlsDomainServerBackupContentRequest) (response GetWlsDomainServerBackupContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getWlsDomainServerBackupContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetWlsDomainServerBackupContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetWlsDomainServerBackupContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetWlsDomainServerBackupContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetWlsDomainServerBackupContentResponse")
	}
	return
}

// getWlsDomainServerBackupContent implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) getWlsDomainServerBackupContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/wlsDomains/{wlsDomainId}/servers/{serverId}/backups/{backupId}/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetWlsDomainServerBackupContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/WlsDomain/GetWlsDomainServerBackupContent"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "GetWlsDomainServerBackupContent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &backupcontent{})
	return response, err
}

// GetWorkRequest Gets the details of a work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
// A default retry strategy applies to this operation GetWorkRequest()
func (client WeblogicManagementServiceClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
func (client WeblogicManagementServiceClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// InstallLatestPatchesOnWlsDomain Install the latest patches on a WebLogic domain.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/InstallLatestPatchesOnWlsDomain.go.html to see an example of how to use InstallLatestPatchesOnWlsDomain API.
// A default retry strategy applies to this operation InstallLatestPatchesOnWlsDomain()
func (client WeblogicManagementServiceClient) InstallLatestPatchesOnWlsDomain(ctx context.Context, request InstallLatestPatchesOnWlsDomainRequest) (response InstallLatestPatchesOnWlsDomainResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.installLatestPatchesOnWlsDomain, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = InstallLatestPatchesOnWlsDomainResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = InstallLatestPatchesOnWlsDomainResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(InstallLatestPatchesOnWlsDomainResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into InstallLatestPatchesOnWlsDomainResponse")
	}
	return
}

// installLatestPatchesOnWlsDomain implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) installLatestPatchesOnWlsDomain(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/wlsDomains/{wlsDomainId}/actions/installLatestPatches", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response InstallLatestPatchesOnWlsDomainResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/WlsDomain/InstallLatestPatchesOnWlsDomain"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "InstallLatestPatchesOnWlsDomain", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAgreementRecords List the terms of use agreement record for the WebLogic domain.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/ListAgreementRecords.go.html to see an example of how to use ListAgreementRecords API.
// A default retry strategy applies to this operation ListAgreementRecords()
func (client WeblogicManagementServiceClient) ListAgreementRecords(ctx context.Context, request ListAgreementRecordsRequest) (response ListAgreementRecordsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAgreementRecords, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAgreementRecordsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAgreementRecordsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAgreementRecordsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAgreementRecordsResponse")
	}
	return
}

// listAgreementRecords implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) listAgreementRecords(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/wlsDomains/{wlsDomainId}/agreementRecords", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAgreementRecordsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/WlsDomain/ListAgreementRecords"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "ListAgreementRecords", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListApplicablePatches Gets the latest patches that can be installed to the WebLogic domains.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/ListApplicablePatches.go.html to see an example of how to use ListApplicablePatches API.
// A default retry strategy applies to this operation ListApplicablePatches()
func (client WeblogicManagementServiceClient) ListApplicablePatches(ctx context.Context, request ListApplicablePatchesRequest) (response ListApplicablePatchesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listApplicablePatches, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListApplicablePatchesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListApplicablePatchesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListApplicablePatchesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListApplicablePatchesResponse")
	}
	return
}

// listApplicablePatches implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) listApplicablePatches(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/wlsDomains/{wlsDomainId}/applicablePatches", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListApplicablePatchesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/WlsDomain/ListApplicablePatches"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "ListApplicablePatches", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListManagedInstanceScanResults Gets all the scan results for all WebLogic servers in the managed instance.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/ListManagedInstanceScanResults.go.html to see an example of how to use ListManagedInstanceScanResults API.
// A default retry strategy applies to this operation ListManagedInstanceScanResults()
func (client WeblogicManagementServiceClient) ListManagedInstanceScanResults(ctx context.Context, request ListManagedInstanceScanResultsRequest) (response ListManagedInstanceScanResultsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagedInstanceScanResults, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagedInstanceScanResultsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagedInstanceScanResultsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagedInstanceScanResultsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagedInstanceScanResultsResponse")
	}
	return
}

// listManagedInstanceScanResults implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) listManagedInstanceScanResults(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances/{managedInstanceId}/scanResults", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListManagedInstanceScanResultsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/ManagedInstance/ListManagedInstanceScanResults"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "ListManagedInstanceScanResults", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListManagedInstanceServerInstalledPatches Gets a list of installed patches on a server in a managed instance.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/ListManagedInstanceServerInstalledPatches.go.html to see an example of how to use ListManagedInstanceServerInstalledPatches API.
// A default retry strategy applies to this operation ListManagedInstanceServerInstalledPatches()
func (client WeblogicManagementServiceClient) ListManagedInstanceServerInstalledPatches(ctx context.Context, request ListManagedInstanceServerInstalledPatchesRequest) (response ListManagedInstanceServerInstalledPatchesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listManagedInstanceServerInstalledPatches, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagedInstanceServerInstalledPatchesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagedInstanceServerInstalledPatchesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagedInstanceServerInstalledPatchesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagedInstanceServerInstalledPatchesResponse")
	}
	return
}

// listManagedInstanceServerInstalledPatches implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) listManagedInstanceServerInstalledPatches(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances/{managedInstanceId}/servers/{serverId}/installedPatches", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListManagedInstanceServerInstalledPatchesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/ManagedInstance/ListManagedInstanceServerInstalledPatches"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "ListManagedInstanceServerInstalledPatches", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListManagedInstanceServers Gets list of servers in a specific managed instance.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/ListManagedInstanceServers.go.html to see an example of how to use ListManagedInstanceServers API.
// A default retry strategy applies to this operation ListManagedInstanceServers()
func (client WeblogicManagementServiceClient) ListManagedInstanceServers(ctx context.Context, request ListManagedInstanceServersRequest) (response ListManagedInstanceServersResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listManagedInstanceServers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagedInstanceServersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagedInstanceServersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagedInstanceServersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagedInstanceServersResponse")
	}
	return
}

// listManagedInstanceServers implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) listManagedInstanceServers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances/{managedInstanceId}/servers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListManagedInstanceServersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/ManagedInstance/ListManagedInstanceServers"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "ListManagedInstanceServers", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListManagedInstances Lists managed instances that match the specified compartment or managed instance OCID. Filter the list against a variety of criteria including but not limited to its name, status and compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/ListManagedInstances.go.html to see an example of how to use ListManagedInstances API.
// A default retry strategy applies to this operation ListManagedInstances()
func (client WeblogicManagementServiceClient) ListManagedInstances(ctx context.Context, request ListManagedInstancesRequest) (response ListManagedInstancesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listManagedInstances, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListManagedInstancesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListManagedInstancesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListManagedInstancesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListManagedInstancesResponse")
	}
	return
}

// listManagedInstances implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) listManagedInstances(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/managedInstances", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListManagedInstancesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/ManagedInstance/ListManagedInstances"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "ListManagedInstances", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListRequiredPolicies Gets all the required policies for the WebLogic Management Service.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/ListRequiredPolicies.go.html to see an example of how to use ListRequiredPolicies API.
// A default retry strategy applies to this operation ListRequiredPolicies()
func (client WeblogicManagementServiceClient) ListRequiredPolicies(ctx context.Context, request ListRequiredPoliciesRequest) (response ListRequiredPoliciesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRequiredPolicies, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRequiredPoliciesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRequiredPoliciesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRequiredPoliciesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRequiredPoliciesResponse")
	}
	return
}

// listRequiredPolicies implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) listRequiredPolicies(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/requiredPolicies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListRequiredPoliciesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/RequiredPolicyCollection/ListRequiredPolicies"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "ListRequiredPolicies", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWlsDomainCredentials Gets domain credentials of a specific domain.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/ListWlsDomainCredentials.go.html to see an example of how to use ListWlsDomainCredentials API.
// A default retry strategy applies to this operation ListWlsDomainCredentials()
func (client WeblogicManagementServiceClient) ListWlsDomainCredentials(ctx context.Context, request ListWlsDomainCredentialsRequest) (response ListWlsDomainCredentialsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listWlsDomainCredentials, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWlsDomainCredentialsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWlsDomainCredentialsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWlsDomainCredentialsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWlsDomainCredentialsResponse")
	}
	return
}

// listWlsDomainCredentials implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) listWlsDomainCredentials(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/wlsDomains/{wlsDomainId}/credentials", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWlsDomainCredentialsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/WlsDomain/ListWlsDomainCredentials"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "ListWlsDomainCredentials", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWlsDomainScanResults Get all scan results for a server in a specific WebLogic domain.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/ListWlsDomainScanResults.go.html to see an example of how to use ListWlsDomainScanResults API.
// A default retry strategy applies to this operation ListWlsDomainScanResults()
func (client WeblogicManagementServiceClient) ListWlsDomainScanResults(ctx context.Context, request ListWlsDomainScanResultsRequest) (response ListWlsDomainScanResultsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWlsDomainScanResults, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWlsDomainScanResultsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWlsDomainScanResultsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWlsDomainScanResultsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWlsDomainScanResultsResponse")
	}
	return
}

// listWlsDomainScanResults implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) listWlsDomainScanResults(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/wlsDomains/{wlsDomainId}/scanResults", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWlsDomainScanResultsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/WlsDomain/ListWlsDomainScanResults"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "ListWlsDomainScanResults", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWlsDomainServerBackups Gets a list of backups for the server of a specific WebLogic Domain.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/ListWlsDomainServerBackups.go.html to see an example of how to use ListWlsDomainServerBackups API.
// A default retry strategy applies to this operation ListWlsDomainServerBackups()
func (client WeblogicManagementServiceClient) ListWlsDomainServerBackups(ctx context.Context, request ListWlsDomainServerBackupsRequest) (response ListWlsDomainServerBackupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWlsDomainServerBackups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWlsDomainServerBackupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWlsDomainServerBackupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWlsDomainServerBackupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWlsDomainServerBackupsResponse")
	}
	return
}

// listWlsDomainServerBackups implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) listWlsDomainServerBackups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/wlsDomains/{wlsDomainId}/servers/{serverId}/backups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWlsDomainServerBackupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/WlsDomain/ListWlsDomainServerBackups"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "ListWlsDomainServerBackups", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWlsDomainServerInstalledPatches Gets a list of installed patches on a server for a domain.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/ListWlsDomainServerInstalledPatches.go.html to see an example of how to use ListWlsDomainServerInstalledPatches API.
// A default retry strategy applies to this operation ListWlsDomainServerInstalledPatches()
func (client WeblogicManagementServiceClient) ListWlsDomainServerInstalledPatches(ctx context.Context, request ListWlsDomainServerInstalledPatchesRequest) (response ListWlsDomainServerInstalledPatchesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWlsDomainServerInstalledPatches, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWlsDomainServerInstalledPatchesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWlsDomainServerInstalledPatchesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWlsDomainServerInstalledPatchesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWlsDomainServerInstalledPatchesResponse")
	}
	return
}

// listWlsDomainServerInstalledPatches implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) listWlsDomainServerInstalledPatches(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/wlsDomains/{wlsDomainId}/servers/{serverId}/installedPatches", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWlsDomainServerInstalledPatchesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/WlsDomain/ListWlsDomainServerInstalledPatches"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "ListWlsDomainServerInstalledPatches", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWlsDomainServers Gets list of servers in a specific WebLogic domain.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/ListWlsDomainServers.go.html to see an example of how to use ListWlsDomainServers API.
// A default retry strategy applies to this operation ListWlsDomainServers()
func (client WeblogicManagementServiceClient) ListWlsDomainServers(ctx context.Context, request ListWlsDomainServersRequest) (response ListWlsDomainServersResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listWlsDomainServers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWlsDomainServersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWlsDomainServersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWlsDomainServersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWlsDomainServersResponse")
	}
	return
}

// listWlsDomainServers implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) listWlsDomainServers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/wlsDomains/{wlsDomainId}/servers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWlsDomainServersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/WlsDomain/ListWlsDomainServers"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "ListWlsDomainServers", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWlsDomains Gets all WebLogic domains in a given compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/ListWlsDomains.go.html to see an example of how to use ListWlsDomains API.
// A default retry strategy applies to this operation ListWlsDomains()
func (client WeblogicManagementServiceClient) ListWlsDomains(ctx context.Context, request ListWlsDomainsRequest) (response ListWlsDomainsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.listWlsDomains, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWlsDomainsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWlsDomainsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWlsDomainsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWlsDomainsResponse")
	}
	return
}

// listWlsDomains implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) listWlsDomains(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/wlsDomains", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWlsDomainsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/WlsDomain/ListWlsDomains"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "ListWlsDomains", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWlsDomainsSharingMiddlewares Gets a list of WebLogic domains that share middleware with a specific domain.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/ListWlsDomainsSharingMiddlewares.go.html to see an example of how to use ListWlsDomainsSharingMiddlewares API.
// A default retry strategy applies to this operation ListWlsDomainsSharingMiddlewares()
func (client WeblogicManagementServiceClient) ListWlsDomainsSharingMiddlewares(ctx context.Context, request ListWlsDomainsSharingMiddlewaresRequest) (response ListWlsDomainsSharingMiddlewaresResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWlsDomainsSharingMiddlewares, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWlsDomainsSharingMiddlewaresResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWlsDomainsSharingMiddlewaresResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWlsDomainsSharingMiddlewaresResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWlsDomainsSharingMiddlewaresResponse")
	}
	return
}

// listWlsDomainsSharingMiddlewares implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) listWlsDomainsSharingMiddlewares(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/wlsDomains/{wlsDomainId}/wlsDomainsSharingMiddleware", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWlsDomainsSharingMiddlewaresResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/WlsDomain/ListWlsDomainsSharingMiddlewares"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "ListWlsDomainsSharingMiddlewares", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Lists the errors for a work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client WeblogicManagementServiceClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
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
func (client WeblogicManagementServiceClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/WorkRequest/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Lists the logs for a work request.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client WeblogicManagementServiceClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
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
func (client WeblogicManagementServiceClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/WorkRequest/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
// A default retry strategy applies to this operation ListWorkRequests()
func (client WeblogicManagementServiceClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
func (client WeblogicManagementServiceClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/WorkRequest/ListWorkRequests"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RestartWlsDomain Restarts all the servers in the WebLogic domains. Servers that are already stopped are ignored.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/RestartWlsDomain.go.html to see an example of how to use RestartWlsDomain API.
// A default retry strategy applies to this operation RestartWlsDomain()
func (client WeblogicManagementServiceClient) RestartWlsDomain(ctx context.Context, request RestartWlsDomainRequest) (response RestartWlsDomainResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.restartWlsDomain, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RestartWlsDomainResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RestartWlsDomainResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RestartWlsDomainResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RestartWlsDomainResponse")
	}
	return
}

// restartWlsDomain implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) restartWlsDomain(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/wlsDomains/{wlsDomainId}/actions/restart", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RestartWlsDomainResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/WlsDomain/RestartWlsDomain"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "RestartWlsDomain", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RestoreWlsDomain Restore a domain from backup. If the backup contains a MIDDLEWARE asset, then the middleware of the domain, including patches, will be restored.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/RestoreWlsDomain.go.html to see an example of how to use RestoreWlsDomain API.
// A default retry strategy applies to this operation RestoreWlsDomain()
func (client WeblogicManagementServiceClient) RestoreWlsDomain(ctx context.Context, request RestoreWlsDomainRequest) (response RestoreWlsDomainResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.restoreWlsDomain, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RestoreWlsDomainResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RestoreWlsDomainResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RestoreWlsDomainResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RestoreWlsDomainResponse")
	}
	return
}

// restoreWlsDomain implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) restoreWlsDomain(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/wlsDomains/{wlsDomainId}/actions/restore", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RestoreWlsDomainResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/WlsDomain/RestoreWlsDomain"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "RestoreWlsDomain", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ScanManagedInstance Scans a managed instance for WebLogic domains.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/ScanManagedInstance.go.html to see an example of how to use ScanManagedInstance API.
// A default retry strategy applies to this operation ScanManagedInstance()
func (client WeblogicManagementServiceClient) ScanManagedInstance(ctx context.Context, request ScanManagedInstanceRequest) (response ScanManagedInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.scanManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ScanManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ScanManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ScanManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ScanManagedInstanceResponse")
	}
	return
}

// scanManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) scanManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedInstances/{managedInstanceId}/actions/scan", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ScanManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/ManagedInstance/ScanManagedInstance"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "ScanManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ScanWlsDomain Runs a series of checks in the WebLogic domain.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/ScanWlsDomain.go.html to see an example of how to use ScanWlsDomain API.
// A default retry strategy applies to this operation ScanWlsDomain()
func (client WeblogicManagementServiceClient) ScanWlsDomain(ctx context.Context, request ScanWlsDomainRequest) (response ScanWlsDomainResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.scanWlsDomain, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ScanWlsDomainResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ScanWlsDomainResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ScanWlsDomainResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ScanWlsDomainResponse")
	}
	return
}

// scanWlsDomain implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) scanWlsDomain(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/wlsDomains/{wlsDomainId}/actions/scan", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ScanWlsDomainResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/WlsDomain/ScanWlsDomain"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "ScanWlsDomain", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SetRestartOrder Sets restart order of servers in specific WebLogic domain.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/SetRestartOrder.go.html to see an example of how to use SetRestartOrder API.
// A default retry strategy applies to this operation SetRestartOrder()
func (client WeblogicManagementServiceClient) SetRestartOrder(ctx context.Context, request SetRestartOrderRequest) (response SetRestartOrderResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.setRestartOrder, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SetRestartOrderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SetRestartOrderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SetRestartOrderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SetRestartOrderResponse")
	}
	return
}

// setRestartOrder implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) setRestartOrder(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/wlsDomains/{wlsDomainId}/actions/setRestartOrder", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SetRestartOrderResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/WlsDomain/SetRestartOrder"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "SetRestartOrder", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StartWlsDomain Starts all the servers in the WebLogic domain. Servers that are already started are ignored.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/StartWlsDomain.go.html to see an example of how to use StartWlsDomain API.
// A default retry strategy applies to this operation StartWlsDomain()
func (client WeblogicManagementServiceClient) StartWlsDomain(ctx context.Context, request StartWlsDomainRequest) (response StartWlsDomainResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.startWlsDomain, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StartWlsDomainResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StartWlsDomainResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartWlsDomainResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartWlsDomainResponse")
	}
	return
}

// startWlsDomain implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) startWlsDomain(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/wlsDomains/{wlsDomainId}/actions/start", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StartWlsDomainResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/WlsDomain/StartWlsDomain"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "StartWlsDomain", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StopWlsDomain Stops all the servers in the WebLogic domain. Servers that are already stopped are ignored.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/StopWlsDomain.go.html to see an example of how to use StopWlsDomain API.
// A default retry strategy applies to this operation StopWlsDomain()
func (client WeblogicManagementServiceClient) StopWlsDomain(ctx context.Context, request StopWlsDomainRequest) (response StopWlsDomainResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.stopWlsDomain, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StopWlsDomainResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StopWlsDomainResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopWlsDomainResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopWlsDomainResponse")
	}
	return
}

// stopWlsDomain implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) stopWlsDomain(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/wlsDomains/{wlsDomainId}/actions/stop", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StopWlsDomainResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/WlsDomain/StopWlsDomain"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "StopWlsDomain", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SummarizeResourceInventory Gets the data to be shown in the Overview page of the service in a given compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/SummarizeResourceInventory.go.html to see an example of how to use SummarizeResourceInventory API.
// A default retry strategy applies to this operation SummarizeResourceInventory()
func (client WeblogicManagementServiceClient) SummarizeResourceInventory(ctx context.Context, request SummarizeResourceInventoryRequest) (response SummarizeResourceInventoryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.summarizeResourceInventory, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SummarizeResourceInventoryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SummarizeResourceInventoryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SummarizeResourceInventoryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SummarizeResourceInventoryResponse")
	}
	return
}

// summarizeResourceInventory implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) summarizeResourceInventory(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/resourceInventory", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SummarizeResourceInventoryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/ResourceInventory/SummarizeResourceInventory"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "SummarizeResourceInventory", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateManagedInstance Updates the specified managed instance information, such as discovery interval and domain search path.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/UpdateManagedInstance.go.html to see an example of how to use UpdateManagedInstance API.
// A default retry strategy applies to this operation UpdateManagedInstance()
func (client WeblogicManagementServiceClient) UpdateManagedInstance(ctx context.Context, request UpdateManagedInstanceRequest) (response UpdateManagedInstanceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateManagedInstance, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateManagedInstanceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateManagedInstanceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateManagedInstanceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateManagedInstanceResponse")
	}
	return
}

// updateManagedInstance implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) updateManagedInstance(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/managedInstances/{managedInstanceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateManagedInstanceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/ManagedInstance/UpdateManagedInstance"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "UpdateManagedInstance", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateWlsDomain Updates a specific WebLogic domain.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/UpdateWlsDomain.go.html to see an example of how to use UpdateWlsDomain API.
// A default retry strategy applies to this operation UpdateWlsDomain()
func (client WeblogicManagementServiceClient) UpdateWlsDomain(ctx context.Context, request UpdateWlsDomainRequest) (response UpdateWlsDomainResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateWlsDomain, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateWlsDomainResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateWlsDomainResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateWlsDomainResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateWlsDomainResponse")
	}
	return
}

// updateWlsDomain implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) updateWlsDomain(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/wlsDomains/{wlsDomainId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateWlsDomainResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/WlsDomain/UpdateWlsDomain"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "UpdateWlsDomain", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateWlsDomainCredential Updates WebLogic domain credentials of specific WebLogic domain.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/UpdateWlsDomainCredential.go.html to see an example of how to use UpdateWlsDomainCredential API.
// A default retry strategy applies to this operation UpdateWlsDomainCredential()
func (client WeblogicManagementServiceClient) UpdateWlsDomainCredential(ctx context.Context, request UpdateWlsDomainCredentialRequest) (response UpdateWlsDomainCredentialResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateWlsDomainCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateWlsDomainCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateWlsDomainCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateWlsDomainCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateWlsDomainCredentialResponse")
	}
	return
}

// updateWlsDomainCredential implements the OCIOperation interface (enables retrying operations)
func (client WeblogicManagementServiceClient) updateWlsDomainCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/wlsDomains/{wlsDomainId}/credentials/{credentialType}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateWlsDomainCredentialResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/wlms/20241101/WlsDomain/UpdateWlsDomainCredential"
		err = common.PostProcessServiceError(err, "WeblogicManagementService", "UpdateWlsDomainCredential", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
