// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Email Delivery API
//
// Use the Email Delivery API to do the necessary set up to send high-volume and application-generated emails through the OCI Email Delivery service.
// For more information, see Overview of the Email Delivery Service (https://docs.oracle.com/iaas/Content/Email/Concepts/overview.htm).
//  **Note:** Write actions (POST, UPDATE, DELETE) may take several minutes to propagate and be reflected by the API.
//  If a subsequent read request fails to reflect your changes, wait a few minutes and try again.
//

package email

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// EmailClient a client for Email
type EmailClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewEmailClientWithConfigurationProvider Creates a new default Email client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewEmailClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client EmailClient, err error) {
	if enabled := common.CheckForEnabledServices("email"); !enabled {
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
	return newEmailClientFromBaseClient(baseClient, provider)
}

// NewEmailClientWithOboToken Creates a new default Email client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewEmailClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client EmailClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newEmailClientFromBaseClient(baseClient, configProvider)
}

func newEmailClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client EmailClient, err error) {
	// Email service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Email"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = EmailClient{BaseClient: baseClient}
	client.BasePath = "20170907"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *EmailClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("email", "https://ctrl.email.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *EmailClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *EmailClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AddEmailDomainLock Adds a lock to a resource.
func (client EmailClient) AddEmailDomainLock(ctx context.Context, request AddEmailDomainLockRequest) (response AddEmailDomainLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.addEmailDomainLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddEmailDomainLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddEmailDomainLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddEmailDomainLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddEmailDomainLockResponse")
	}
	return
}

// addEmailDomainLock implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) addEmailDomainLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/emailDomains/{emailDomainId}/actions/addLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddEmailDomainLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "AddEmailDomainLock")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailDomain/AddEmailDomainLock"
		err = common.PostProcessServiceError(err, "Email", "AddEmailDomainLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddEmailIpPoolLock Adds a lock to a resource.
func (client EmailClient) AddEmailIpPoolLock(ctx context.Context, request AddEmailIpPoolLockRequest) (response AddEmailIpPoolLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.addEmailIpPoolLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddEmailIpPoolLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddEmailIpPoolLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddEmailIpPoolLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddEmailIpPoolLockResponse")
	}
	return
}

// addEmailIpPoolLock implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) addEmailIpPoolLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/emailIpPools/{emailIpPoolId}/actions/addLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddEmailIpPoolLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "AddEmailIpPoolLock")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailIpPool/AddEmailIpPoolLock"
		err = common.PostProcessServiceError(err, "Email", "AddEmailIpPoolLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddEmailOutboundIp Add OutboundIps to EmailIpPool.
func (client EmailClient) AddEmailOutboundIp(ctx context.Context, request AddEmailOutboundIpRequest) (response AddEmailOutboundIpResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.addEmailOutboundIp, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddEmailOutboundIpResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddEmailOutboundIpResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddEmailOutboundIpResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddEmailOutboundIpResponse")
	}
	return
}

// addEmailOutboundIp implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) addEmailOutboundIp(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/emailIpPools/{emailIpPoolId}/actions/addEmailOutboundIp", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddEmailOutboundIpResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "AddEmailOutboundIp")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailIpPool/AddEmailOutboundIp"
		err = common.PostProcessServiceError(err, "Email", "AddEmailOutboundIp", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddEmailTrackConfigLock Adds a lock to a resource.
func (client EmailClient) AddEmailTrackConfigLock(ctx context.Context, request AddEmailTrackConfigLockRequest) (response AddEmailTrackConfigLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.addEmailTrackConfigLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddEmailTrackConfigLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddEmailTrackConfigLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddEmailTrackConfigLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddEmailTrackConfigLockResponse")
	}
	return
}

// addEmailTrackConfigLock implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) addEmailTrackConfigLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/emailTrackConfigs/{emailTrackConfigId}/actions/addLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddEmailTrackConfigLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "AddEmailTrackConfigLock")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailTrackConfig/AddEmailTrackConfigLock"
		err = common.PostProcessServiceError(err, "Email", "AddEmailTrackConfigLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddPrivateEndpointLock Adds a lock to a resource.
func (client EmailClient) AddPrivateEndpointLock(ctx context.Context, request AddPrivateEndpointLockRequest) (response AddPrivateEndpointLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.addPrivateEndpointLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddPrivateEndpointLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddPrivateEndpointLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddPrivateEndpointLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddPrivateEndpointLockResponse")
	}
	return
}

// addPrivateEndpointLock implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) addPrivateEndpointLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/emailPrivateEndpoints/{emailPrivateEndpointId}/actions/addLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddPrivateEndpointLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "AddPrivateEndpointLock")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailPrivateEndpoint/AddPrivateEndpointLock"
		err = common.PostProcessServiceError(err, "Email", "AddPrivateEndpointLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddReturnPathLock Adds a lock to a resource.
func (client EmailClient) AddReturnPathLock(ctx context.Context, request AddReturnPathLockRequest) (response AddReturnPathLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.addReturnPathLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddReturnPathLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddReturnPathLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddReturnPathLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddReturnPathLockResponse")
	}
	return
}

// addReturnPathLock implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) addReturnPathLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/emailReturnPaths/{emailReturnPathId}/actions/addLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddReturnPathLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "AddReturnPathLock")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailReturnPath/AddReturnPathLock"
		err = common.PostProcessServiceError(err, "Email", "AddReturnPathLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddSenderLock Adds a lock to a resource.
func (client EmailClient) AddSenderLock(ctx context.Context, request AddSenderLockRequest) (response AddSenderLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.addSenderLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddSenderLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddSenderLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddSenderLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddSenderLockResponse")
	}
	return
}

// addSenderLock implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) addSenderLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/senders/{senderId}/actions/addLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddSenderLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "AddSenderLock")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/Sender/AddSenderLock"
		err = common.PostProcessServiceError(err, "Email", "AddSenderLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeEmailDeliveryConfigCompartment Moves a resource into a different compartment. When provided, 'If-Match' is checked against 'ETag' values of the resource.
func (client EmailClient) ChangeEmailDeliveryConfigCompartment(ctx context.Context, request ChangeEmailDeliveryConfigCompartmentRequest) (response ChangeEmailDeliveryConfigCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeEmailDeliveryConfigCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeEmailDeliveryConfigCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeEmailDeliveryConfigCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeEmailDeliveryConfigCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeEmailDeliveryConfigCompartmentResponse")
	}
	return
}

// changeEmailDeliveryConfigCompartment implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) changeEmailDeliveryConfigCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/emailDeliveryConfigs/{emailDeliveryConfigId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeEmailDeliveryConfigCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "ChangeEmailDeliveryConfigCompartment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailDeliveryConfig/ChangeEmailDeliveryConfigCompartment"
		err = common.PostProcessServiceError(err, "Email", "ChangeEmailDeliveryConfigCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeEmailDeliveryConfigIpAssociationCompartment Moves a resource into a different compartment. When provided, 'If-Match' is checked against 'ETag' values of the resource.
func (client EmailClient) ChangeEmailDeliveryConfigIpAssociationCompartment(ctx context.Context, request ChangeEmailDeliveryConfigIpAssociationCompartmentRequest) (response ChangeEmailDeliveryConfigIpAssociationCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeEmailDeliveryConfigIpAssociationCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeEmailDeliveryConfigIpAssociationCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeEmailDeliveryConfigIpAssociationCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeEmailDeliveryConfigIpAssociationCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeEmailDeliveryConfigIpAssociationCompartmentResponse")
	}
	return
}

// changeEmailDeliveryConfigIpAssociationCompartment implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) changeEmailDeliveryConfigIpAssociationCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/emailDeliveryConfigIpAssociations/{emailDeliveryConfigIpAssociationId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeEmailDeliveryConfigIpAssociationCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "ChangeEmailDeliveryConfigIpAssociationCompartment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailDeliveryConfigIpAssociation/ChangeEmailDeliveryConfigIpAssociationCompartment"
		err = common.PostProcessServiceError(err, "Email", "ChangeEmailDeliveryConfigIpAssociationCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeEmailDomainCompartment Moves an email domain into a different compartment.
// When provided, If-Match is checked against ETag value of the resource.
// For information about moving resources between compartments, see
// Moving Resources to a Different Compartment (https://docs.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
// **Note:** All DKIM objects associated with this email domain will also be moved into the provided compartment.
func (client EmailClient) ChangeEmailDomainCompartment(ctx context.Context, request ChangeEmailDomainCompartmentRequest) (response ChangeEmailDomainCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeEmailDomainCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeEmailDomainCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeEmailDomainCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeEmailDomainCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeEmailDomainCompartmentResponse")
	}
	return
}

// changeEmailDomainCompartment implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) changeEmailDomainCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/emailDomains/{emailDomainId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeEmailDomainCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "ChangeEmailDomainCompartment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailDomain/ChangeEmailDomainCompartment"
		err = common.PostProcessServiceError(err, "Email", "ChangeEmailDomainCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeEmailIpPoolCompartment Moves a resource into a different compartment. When provided, If-Match is checked against ETag values of the resource.
func (client EmailClient) ChangeEmailIpPoolCompartment(ctx context.Context, request ChangeEmailIpPoolCompartmentRequest) (response ChangeEmailIpPoolCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeEmailIpPoolCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeEmailIpPoolCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeEmailIpPoolCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeEmailIpPoolCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeEmailIpPoolCompartmentResponse")
	}
	return
}

// changeEmailIpPoolCompartment implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) changeEmailIpPoolCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/emailIpPools/{emailIpPoolId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeEmailIpPoolCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "ChangeEmailIpPoolCompartment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailIpPool/ChangeEmailIpPoolCompartment"
		err = common.PostProcessServiceError(err, "Email", "ChangeEmailIpPoolCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeEmailPrivateEndpointCompartment Moves a resource into a different compartment. When provided, 'If-Match' is checked against 'ETag' values of the resource.
func (client EmailClient) ChangeEmailPrivateEndpointCompartment(ctx context.Context, request ChangeEmailPrivateEndpointCompartmentRequest) (response ChangeEmailPrivateEndpointCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeEmailPrivateEndpointCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeEmailPrivateEndpointCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeEmailPrivateEndpointCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeEmailPrivateEndpointCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeEmailPrivateEndpointCompartmentResponse")
	}
	return
}

// changeEmailPrivateEndpointCompartment implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) changeEmailPrivateEndpointCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/emailPrivateEndpoints/{emailPrivateEndpointId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeEmailPrivateEndpointCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "ChangeEmailPrivateEndpointCompartment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailPrivateEndpoint/ChangeEmailPrivateEndpointCompartment"
		err = common.PostProcessServiceError(err, "Email", "ChangeEmailPrivateEndpointCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeEmailRecipientDomainCompartment Moves a resource into a different compartment. When provided, 'If-Match' is checked against 'ETag' values of the resource.
func (client EmailClient) ChangeEmailRecipientDomainCompartment(ctx context.Context, request ChangeEmailRecipientDomainCompartmentRequest) (response ChangeEmailRecipientDomainCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeEmailRecipientDomainCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeEmailRecipientDomainCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeEmailRecipientDomainCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeEmailRecipientDomainCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeEmailRecipientDomainCompartmentResponse")
	}
	return
}

// changeEmailRecipientDomainCompartment implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) changeEmailRecipientDomainCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/emailRecipientDomains/{emailRecipientDomainId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeEmailRecipientDomainCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "ChangeEmailRecipientDomainCompartment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailRecipientDomain/ChangeEmailRecipientDomainCompartment"
		err = common.PostProcessServiceError(err, "Email", "ChangeEmailRecipientDomainCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeEmailTrackConfigCompartment Moves email tracking configuration resource into a different compartment.
func (client EmailClient) ChangeEmailTrackConfigCompartment(ctx context.Context, request ChangeEmailTrackConfigCompartmentRequest) (response ChangeEmailTrackConfigCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeEmailTrackConfigCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeEmailTrackConfigCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeEmailTrackConfigCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeEmailTrackConfigCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeEmailTrackConfigCompartmentResponse")
	}
	return
}

// changeEmailTrackConfigCompartment implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) changeEmailTrackConfigCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/emailTrackConfigs/{emailTrackConfigId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeEmailTrackConfigCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "ChangeEmailTrackConfigCompartment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailTrackConfig/ChangeEmailTrackConfigCompartment"
		err = common.PostProcessServiceError(err, "Email", "ChangeEmailTrackConfigCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeSenderCompartment Moves a sender into a different compartment. When provided, If-Match is checked against ETag values of the resource.
func (client EmailClient) ChangeSenderCompartment(ctx context.Context, request ChangeSenderCompartmentRequest) (response ChangeSenderCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeSenderCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeSenderCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeSenderCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeSenderCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeSenderCompartmentResponse")
	}
	return
}

// changeSenderCompartment implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) changeSenderCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/senders/{senderId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeSenderCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "ChangeSenderCompartment")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/Sender/ChangeSenderCompartment"
		err = common.PostProcessServiceError(err, "Email", "ChangeSenderCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDkim Creates a new DKIM for an email domain.
// This DKIM signs all approved senders in the tenancy that are in this email domain.
// Best security practices indicate to periodically rotate the DKIM that is doing the signing.
// When a second DKIM is applied, all senders seamlessly pick up the new key
// without interruption in signing.
func (client EmailClient) CreateDkim(ctx context.Context, request CreateDkimRequest) (response CreateDkimResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDkim, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDkimResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDkimResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDkimResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDkimResponse")
	}
	return
}

// createDkim implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) createDkim(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dkims", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDkimResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "CreateDkim")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/Dkim/CreateDkim"
		err = common.PostProcessServiceError(err, "Email", "CreateDkim", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateEmailDeliveryConfig Creates a delivery config.
func (client EmailClient) CreateEmailDeliveryConfig(ctx context.Context, request CreateEmailDeliveryConfigRequest) (response CreateEmailDeliveryConfigResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createEmailDeliveryConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateEmailDeliveryConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateEmailDeliveryConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateEmailDeliveryConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateEmailDeliveryConfigResponse")
	}
	return
}

// createEmailDeliveryConfig implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) createEmailDeliveryConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/emailDeliveryConfigs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateEmailDeliveryConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "CreateEmailDeliveryConfig")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Email", "CreateEmailDeliveryConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateEmailDeliveryConfigIpAssociation Creates an email delivery configuration ip association.
func (client EmailClient) CreateEmailDeliveryConfigIpAssociation(ctx context.Context, request CreateEmailDeliveryConfigIpAssociationRequest) (response CreateEmailDeliveryConfigIpAssociationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createEmailDeliveryConfigIpAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateEmailDeliveryConfigIpAssociationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateEmailDeliveryConfigIpAssociationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateEmailDeliveryConfigIpAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateEmailDeliveryConfigIpAssociationResponse")
	}
	return
}

// createEmailDeliveryConfigIpAssociation implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) createEmailDeliveryConfigIpAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/emailDeliveryConfigIpAssociations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateEmailDeliveryConfigIpAssociationResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "CreateEmailDeliveryConfigIpAssociation")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Email", "CreateEmailDeliveryConfigIpAssociation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateEmailDomain Creates a new email domain. Avoid entering confidential information.
func (client EmailClient) CreateEmailDomain(ctx context.Context, request CreateEmailDomainRequest) (response CreateEmailDomainResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createEmailDomain, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateEmailDomainResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateEmailDomainResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateEmailDomainResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateEmailDomainResponse")
	}
	return
}

// createEmailDomain implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) createEmailDomain(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/emailDomains", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateEmailDomainResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "CreateEmailDomain")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailDomain/CreateEmailDomain"
		err = common.PostProcessServiceError(err, "Email", "CreateEmailDomain", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateEmailIpPool Creates a new EmailIpPool.
func (client EmailClient) CreateEmailIpPool(ctx context.Context, request CreateEmailIpPoolRequest) (response CreateEmailIpPoolResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createEmailIpPool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateEmailIpPoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateEmailIpPoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateEmailIpPoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateEmailIpPoolResponse")
	}
	return
}

// createEmailIpPool implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) createEmailIpPool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/emailIpPools", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateEmailIpPoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "CreateEmailIpPool")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailIpPool/CreateEmailIpPool"
		err = common.PostProcessServiceError(err, "Email", "CreateEmailIpPool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateEmailPrivateEndpoint Create a new private connection endpoint.
func (client EmailClient) CreateEmailPrivateEndpoint(ctx context.Context, request CreateEmailPrivateEndpointRequest) (response CreateEmailPrivateEndpointResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createEmailPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateEmailPrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateEmailPrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateEmailPrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateEmailPrivateEndpointResponse")
	}
	return
}

// createEmailPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) createEmailPrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/emailPrivateEndpoints", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateEmailPrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "CreateEmailPrivateEndpoint")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailPrivateEndpoint/CreateEmailPrivateEndpoint"
		err = common.PostProcessServiceError(err, "Email", "CreateEmailPrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateEmailRecipientDomain Creates a recipient domain for a tenancy.
func (client EmailClient) CreateEmailRecipientDomain(ctx context.Context, request CreateEmailRecipientDomainRequest) (response CreateEmailRecipientDomainResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createEmailRecipientDomain, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateEmailRecipientDomainResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateEmailRecipientDomainResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateEmailRecipientDomainResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateEmailRecipientDomainResponse")
	}
	return
}

// createEmailRecipientDomain implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) createEmailRecipientDomain(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/emailRecipientDomains", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateEmailRecipientDomainResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "CreateEmailRecipientDomain")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Email", "CreateEmailRecipientDomain", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateEmailReturnPath Creates a new email return path. Avoid entering confidential information.
func (client EmailClient) CreateEmailReturnPath(ctx context.Context, request CreateEmailReturnPathRequest) (response CreateEmailReturnPathResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createEmailReturnPath, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateEmailReturnPathResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateEmailReturnPathResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateEmailReturnPathResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateEmailReturnPathResponse")
	}
	return
}

// createEmailReturnPath implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) createEmailReturnPath(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/emailReturnPaths", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateEmailReturnPathResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "CreateEmailReturnPath")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailReturnPath/CreateEmailReturnPath"
		err = common.PostProcessServiceError(err, "Email", "CreateEmailReturnPath", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateEmailTrackConfig Create email tracking configuration resource used to configure email open tracking, click tracking, and addition of List-Unsubscribe header.
func (client EmailClient) CreateEmailTrackConfig(ctx context.Context, request CreateEmailTrackConfigRequest) (response CreateEmailTrackConfigResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createEmailTrackConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateEmailTrackConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateEmailTrackConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateEmailTrackConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateEmailTrackConfigResponse")
	}
	return
}

// createEmailTrackConfig implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) createEmailTrackConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/emailTrackConfigs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateEmailTrackConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "CreateEmailTrackConfig")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailTrackConfig/CreateEmailTrackConfig"
		err = common.PostProcessServiceError(err, "Email", "CreateEmailTrackConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSender Creates a sender for a tenancy in a given compartment.
func (client EmailClient) CreateSender(ctx context.Context, request CreateSenderRequest) (response CreateSenderResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createSender, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSenderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSenderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSenderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSenderResponse")
	}
	return
}

// createSender implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) createSender(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/senders", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSenderResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "CreateSender")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/Sender/CreateSender"
		err = common.PostProcessServiceError(err, "Email", "CreateSender", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSuppression Adds recipient email addresses to the suppression list for a tenancy.
// Addresses added to the suppression list via the API are denoted as
// "MANUAL" in the `reason` field. *Note:* All email addresses added to the
// suppression list are normalized to include only lowercase letters.
func (client EmailClient) CreateSuppression(ctx context.Context, request CreateSuppressionRequest) (response CreateSuppressionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createSuppression, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSuppressionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSuppressionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSuppressionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSuppressionResponse")
	}
	return
}

// createSuppression implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) createSuppression(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/suppressions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSuppressionResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "CreateSuppression")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/Suppression/CreateSuppression"
		err = common.PostProcessServiceError(err, "Email", "CreateSuppression", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDkim Deletes a DKIM.
// If this key is currently the active key for the email domain, deleting the key
// will stop signing the domain's outgoing mail.
// DKIM keys are left in DELETING state for about a day to allow DKIM signatures on
// in-transit mail to be validated.
// Consider creating a new DKIM for this domain so the signing can be rotated to it instead of deletion.
func (client EmailClient) DeleteDkim(ctx context.Context, request DeleteDkimRequest) (response DeleteDkimResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDkim, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDkimResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDkimResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDkimResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDkimResponse")
	}
	return
}

// deleteDkim implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) deleteDkim(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/dkims/{dkimId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDkimResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "DeleteDkim")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/Dkim/DeleteDkim"
		err = common.PostProcessServiceError(err, "Email", "DeleteDkim", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteEmailDeliveryConfig Removes a delivery configuration for a tenancy in a given compartment for a provided `emailDeliveryConfigId`.
func (client EmailClient) DeleteEmailDeliveryConfig(ctx context.Context, request DeleteEmailDeliveryConfigRequest) (response DeleteEmailDeliveryConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteEmailDeliveryConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteEmailDeliveryConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteEmailDeliveryConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteEmailDeliveryConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteEmailDeliveryConfigResponse")
	}
	return
}

// deleteEmailDeliveryConfig implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) deleteEmailDeliveryConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/emailDeliveryConfigs/{emailDeliveryConfigId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteEmailDeliveryConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "DeleteEmailDeliveryConfig")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailDeliveryConfig/DeleteEmailDeliveryConfig"
		err = common.PostProcessServiceError(err, "Email", "DeleteEmailDeliveryConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteEmailDeliveryConfigIpAssociation Removes a delivery configuration ip association for a tenancy in a given compartment for a provided `emailDeliveryConfigIpAssociationId`.
func (client EmailClient) DeleteEmailDeliveryConfigIpAssociation(ctx context.Context, request DeleteEmailDeliveryConfigIpAssociationRequest) (response DeleteEmailDeliveryConfigIpAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteEmailDeliveryConfigIpAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteEmailDeliveryConfigIpAssociationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteEmailDeliveryConfigIpAssociationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteEmailDeliveryConfigIpAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteEmailDeliveryConfigIpAssociationResponse")
	}
	return
}

// deleteEmailDeliveryConfigIpAssociation implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) deleteEmailDeliveryConfigIpAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/emailDeliveryConfigIpAssociations/{emailDeliveryConfigIpAssociationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteEmailDeliveryConfigIpAssociationResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "DeleteEmailDeliveryConfigIpAssociation")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailDeliveryConfigIpAssociation/DeleteEmailDeliveryConfigIpAssociation"
		err = common.PostProcessServiceError(err, "Email", "DeleteEmailDeliveryConfigIpAssociation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteEmailDomain Deletes an email domain.
func (client EmailClient) DeleteEmailDomain(ctx context.Context, request DeleteEmailDomainRequest) (response DeleteEmailDomainResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteEmailDomain, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteEmailDomainResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteEmailDomainResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteEmailDomainResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteEmailDomainResponse")
	}
	return
}

// deleteEmailDomain implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) deleteEmailDomain(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/emailDomains/{emailDomainId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteEmailDomainResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "DeleteEmailDomain")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailDomain/DeleteEmailDomain"
		err = common.PostProcessServiceError(err, "Email", "DeleteEmailDomain", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteEmailIpPool Deletes an EmailIpPool resource by identifier
func (client EmailClient) DeleteEmailIpPool(ctx context.Context, request DeleteEmailIpPoolRequest) (response DeleteEmailIpPoolResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteEmailIpPool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteEmailIpPoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteEmailIpPoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteEmailIpPoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteEmailIpPoolResponse")
	}
	return
}

// deleteEmailIpPool implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) deleteEmailIpPool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/emailIpPools/{emailIpPoolId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteEmailIpPoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "DeleteEmailIpPool")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailIpPool/DeleteEmailIpPool"
		err = common.PostProcessServiceError(err, "Email", "DeleteEmailIpPool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteEmailPrivateEndpoint Deletes a private endpoint by identifier.
func (client EmailClient) DeleteEmailPrivateEndpoint(ctx context.Context, request DeleteEmailPrivateEndpointRequest) (response DeleteEmailPrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteEmailPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteEmailPrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteEmailPrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteEmailPrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteEmailPrivateEndpointResponse")
	}
	return
}

// deleteEmailPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) deleteEmailPrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/emailPrivateEndpoints/{emailPrivateEndpointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteEmailPrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "DeleteEmailPrivateEndpoint")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailPrivateEndpoint/DeleteEmailPrivateEndpoint"
		err = common.PostProcessServiceError(err, "Email", "DeleteEmailPrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteEmailRecipientDomain Removes a recipient domain for a tenancy in a given compartment for a provided `emailRecipientDomainId`.
func (client EmailClient) DeleteEmailRecipientDomain(ctx context.Context, request DeleteEmailRecipientDomainRequest) (response DeleteEmailRecipientDomainResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteEmailRecipientDomain, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteEmailRecipientDomainResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteEmailRecipientDomainResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteEmailRecipientDomainResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteEmailRecipientDomainResponse")
	}
	return
}

// deleteEmailRecipientDomain implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) deleteEmailRecipientDomain(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/emailRecipientDomains/{emailRecipientDomainId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteEmailRecipientDomainResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "DeleteEmailRecipientDomain")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailRecipientDomain/DeleteEmailRecipientDomain"
		err = common.PostProcessServiceError(err, "Email", "DeleteEmailRecipientDomain", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteEmailReturnPath Deletes an email return path.
func (client EmailClient) DeleteEmailReturnPath(ctx context.Context, request DeleteEmailReturnPathRequest) (response DeleteEmailReturnPathResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteEmailReturnPath, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteEmailReturnPathResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteEmailReturnPathResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteEmailReturnPathResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteEmailReturnPathResponse")
	}
	return
}

// deleteEmailReturnPath implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) deleteEmailReturnPath(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/emailReturnPaths/{emailReturnPathId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteEmailReturnPathResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "DeleteEmailReturnPath")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailReturnPath/DeleteEmailReturnPath"
		err = common.PostProcessServiceError(err, "Email", "DeleteEmailReturnPath", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteEmailTrackConfig Deletes email tracking configuration resource by identifier.
func (client EmailClient) DeleteEmailTrackConfig(ctx context.Context, request DeleteEmailTrackConfigRequest) (response DeleteEmailTrackConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteEmailTrackConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteEmailTrackConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteEmailTrackConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteEmailTrackConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteEmailTrackConfigResponse")
	}
	return
}

// deleteEmailTrackConfig implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) deleteEmailTrackConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/emailTrackConfigs/{emailTrackConfigId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteEmailTrackConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "DeleteEmailTrackConfig")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailTrackConfig/DeleteEmailTrackConfig"
		err = common.PostProcessServiceError(err, "Email", "DeleteEmailTrackConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSender Deletes an approved sender for a tenancy in a given compartment for a
// provided `senderId`.
func (client EmailClient) DeleteSender(ctx context.Context, request DeleteSenderRequest) (response DeleteSenderResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSender, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSenderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSenderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSenderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSenderResponse")
	}
	return
}

// deleteSender implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) deleteSender(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/senders/{senderId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSenderResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "DeleteSender")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/Sender/DeleteSender"
		err = common.PostProcessServiceError(err, "Email", "DeleteSender", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSuppression Removes a suppressed recipient email address from the suppression list
// for a tenancy in a given compartment for a provided `suppressionId`.
func (client EmailClient) DeleteSuppression(ctx context.Context, request DeleteSuppressionRequest) (response DeleteSuppressionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSuppression, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSuppressionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSuppressionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSuppressionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSuppressionResponse")
	}
	return
}

// deleteSuppression implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) deleteSuppression(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/suppressions/{suppressionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSuppressionResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "DeleteSuppression")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/Suppression/DeleteSuppression"
		err = common.PostProcessServiceError(err, "Email", "DeleteSuppression", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDkim Retrieves the specified DKIM.
func (client EmailClient) GetDkim(ctx context.Context, request GetDkimRequest) (response GetDkimResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDkim, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDkimResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDkimResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDkimResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDkimResponse")
	}
	return
}

// getDkim implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) getDkim(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dkims/{dkimId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDkimResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "GetDkim")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/Dkim/GetDkim"
		err = common.PostProcessServiceError(err, "Email", "GetDkim", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetEmailConfiguration Returns  email configuration associated with the specified compartment.
func (client EmailClient) GetEmailConfiguration(ctx context.Context, request GetEmailConfigurationRequest) (response GetEmailConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getEmailConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetEmailConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetEmailConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetEmailConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetEmailConfigurationResponse")
	}
	return
}

// getEmailConfiguration implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) getEmailConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/configuration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetEmailConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "GetEmailConfiguration")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/Configuration/GetEmailConfiguration"
		err = common.PostProcessServiceError(err, "Email", "GetEmailConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetEmailDeliveryConfig Gets the details of a delivery configuration for a given
// `emailDeliveryConfigId`. Each delivery configuration is given a unique OCID.
func (client EmailClient) GetEmailDeliveryConfig(ctx context.Context, request GetEmailDeliveryConfigRequest) (response GetEmailDeliveryConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getEmailDeliveryConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetEmailDeliveryConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetEmailDeliveryConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetEmailDeliveryConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetEmailDeliveryConfigResponse")
	}
	return
}

// getEmailDeliveryConfig implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) getEmailDeliveryConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/emailDeliveryConfigs/{emailDeliveryConfigId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetEmailDeliveryConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "GetEmailDeliveryConfig")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailDeliveryConfig/GetEmailDeliveryConfig"
		err = common.PostProcessServiceError(err, "Email", "GetEmailDeliveryConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetEmailDeliveryConfigIpAssociation Gets the details of a delivery configuration ip association resource for a given
// `emailDeliveryConfigIpAssociationId`.
func (client EmailClient) GetEmailDeliveryConfigIpAssociation(ctx context.Context, request GetEmailDeliveryConfigIpAssociationRequest) (response GetEmailDeliveryConfigIpAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getEmailDeliveryConfigIpAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetEmailDeliveryConfigIpAssociationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetEmailDeliveryConfigIpAssociationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetEmailDeliveryConfigIpAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetEmailDeliveryConfigIpAssociationResponse")
	}
	return
}

// getEmailDeliveryConfigIpAssociation implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) getEmailDeliveryConfigIpAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/emailDeliveryConfigIpAssociations/{emailDeliveryConfigIpAssociationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetEmailDeliveryConfigIpAssociationResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "GetEmailDeliveryConfigIpAssociation")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailDeliveryConfigIpAssociation/GetEmailDeliveryConfigIpAssociation"
		err = common.PostProcessServiceError(err, "Email", "GetEmailDeliveryConfigIpAssociation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetEmailDomain Retrieves the specified email domain.
func (client EmailClient) GetEmailDomain(ctx context.Context, request GetEmailDomainRequest) (response GetEmailDomainResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getEmailDomain, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetEmailDomainResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetEmailDomainResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetEmailDomainResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetEmailDomainResponse")
	}
	return
}

// getEmailDomain implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) getEmailDomain(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/emailDomains/{emailDomainId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetEmailDomainResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "GetEmailDomain")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailDomain/GetEmailDomain"
		err = common.PostProcessServiceError(err, "Email", "GetEmailDomain", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetEmailIpPool Retrieves the specified IpPool by identifier
func (client EmailClient) GetEmailIpPool(ctx context.Context, request GetEmailIpPoolRequest) (response GetEmailIpPoolResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getEmailIpPool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetEmailIpPoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetEmailIpPoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetEmailIpPoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetEmailIpPoolResponse")
	}
	return
}

// getEmailIpPool implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) getEmailIpPool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/emailIpPools/{emailIpPoolId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetEmailIpPoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "GetEmailIpPool")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailIpPool/GetEmailIpPool"
		err = common.PostProcessServiceError(err, "Email", "GetEmailIpPool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetEmailPrivateEndpoint Gets a specific private reverse connection by identifier.
func (client EmailClient) GetEmailPrivateEndpoint(ctx context.Context, request GetEmailPrivateEndpointRequest) (response GetEmailPrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getEmailPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetEmailPrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetEmailPrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetEmailPrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetEmailPrivateEndpointResponse")
	}
	return
}

// getEmailPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) getEmailPrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/emailPrivateEndpoints/{emailPrivateEndpointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetEmailPrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "GetEmailPrivateEndpoint")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailPrivateEndpoint/GetEmailPrivateEndpoint"
		err = common.PostProcessServiceError(err, "Email", "GetEmailPrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetEmailRecipientDomain Gets the details of a recipient domain for a given
// `emailRecipientDomainId`. Each recipient domain is given a unique OCID.
func (client EmailClient) GetEmailRecipientDomain(ctx context.Context, request GetEmailRecipientDomainRequest) (response GetEmailRecipientDomainResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getEmailRecipientDomain, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetEmailRecipientDomainResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetEmailRecipientDomainResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetEmailRecipientDomainResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetEmailRecipientDomainResponse")
	}
	return
}

// getEmailRecipientDomain implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) getEmailRecipientDomain(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/emailRecipientDomains/{emailRecipientDomainId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetEmailRecipientDomainResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "GetEmailRecipientDomain")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailRecipientDomain/GetEmailRecipientDomain"
		err = common.PostProcessServiceError(err, "Email", "GetEmailRecipientDomain", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetEmailReturnPath Retrieves the specified email return path.
func (client EmailClient) GetEmailReturnPath(ctx context.Context, request GetEmailReturnPathRequest) (response GetEmailReturnPathResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getEmailReturnPath, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetEmailReturnPathResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetEmailReturnPathResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetEmailReturnPathResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetEmailReturnPathResponse")
	}
	return
}

// getEmailReturnPath implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) getEmailReturnPath(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/emailReturnPaths/{emailReturnPathId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetEmailReturnPathResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "GetEmailReturnPath")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailReturnPath/GetEmailReturnPath"
		err = common.PostProcessServiceError(err, "Email", "GetEmailReturnPath", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetEmailTrackConfig Get email tracking configuration resource by identifier.
func (client EmailClient) GetEmailTrackConfig(ctx context.Context, request GetEmailTrackConfigRequest) (response GetEmailTrackConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getEmailTrackConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetEmailTrackConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetEmailTrackConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetEmailTrackConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetEmailTrackConfigResponse")
	}
	return
}

// getEmailTrackConfig implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) getEmailTrackConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/emailTrackConfigs/{emailTrackConfigId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetEmailTrackConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "GetEmailTrackConfig")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailTrackConfig/GetEmailTrackConfig"
		err = common.PostProcessServiceError(err, "Email", "GetEmailTrackConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSender Gets an approved sender for a given `senderId`.
func (client EmailClient) GetSender(ctx context.Context, request GetSenderRequest) (response GetSenderResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSender, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSenderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSenderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSenderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSenderResponse")
	}
	return
}

// getSender implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) getSender(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/senders/{senderId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSenderResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "GetSender")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/Sender/GetSender"
		err = common.PostProcessServiceError(err, "Email", "GetSender", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSuppression Gets the details of a suppressed recipient email address for a given
// `suppressionId`. Each suppression is given a unique OCID.
func (client EmailClient) GetSuppression(ctx context.Context, request GetSuppressionRequest) (response GetSuppressionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSuppression, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSuppressionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSuppressionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSuppressionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSuppressionResponse")
	}
	return
}

// getSuppression implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) getSuppression(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/suppressions/{suppressionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSuppressionResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "GetSuppression")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/Suppression/GetSuppression"
		err = common.PostProcessServiceError(err, "Email", "GetSuppression", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets the status of the work request with the given ID.
func (client EmailClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client EmailClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "GetWorkRequest")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "Email", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDkims Lists DKIMs for an email domain.
func (client EmailClient) ListDkims(ctx context.Context, request ListDkimsRequest) (response ListDkimsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDkims, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDkimsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDkimsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDkimsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDkimsResponse")
	}
	return
}

// listDkims implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) listDkims(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dkims", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDkimsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "ListDkims")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/Dkim/ListDkims"
		err = common.PostProcessServiceError(err, "Email", "ListDkims", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListEmailDeliveryConfigIpAssociations Gets a list of email delivery configurations ip associations. The returned list
// is sorted by creation time in descending order.
func (client EmailClient) ListEmailDeliveryConfigIpAssociations(ctx context.Context, request ListEmailDeliveryConfigIpAssociationsRequest) (response ListEmailDeliveryConfigIpAssociationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listEmailDeliveryConfigIpAssociations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListEmailDeliveryConfigIpAssociationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListEmailDeliveryConfigIpAssociationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListEmailDeliveryConfigIpAssociationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListEmailDeliveryConfigIpAssociationsResponse")
	}
	return
}

// listEmailDeliveryConfigIpAssociations implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) listEmailDeliveryConfigIpAssociations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/emailDeliveryConfigIpAssociations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListEmailDeliveryConfigIpAssociationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "ListEmailDeliveryConfigIpAssociations")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailDeliveryConfigIpAssociation/ListEmailDeliveryConfigIpAssociations"
		err = common.PostProcessServiceError(err, "Email", "ListEmailDeliveryConfigIpAssociations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListEmailDeliveryConfigs Gets a list of delivery configurations. The returned list
// is sorted by creation time in descending order.
func (client EmailClient) ListEmailDeliveryConfigs(ctx context.Context, request ListEmailDeliveryConfigsRequest) (response ListEmailDeliveryConfigsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listEmailDeliveryConfigs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListEmailDeliveryConfigsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListEmailDeliveryConfigsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListEmailDeliveryConfigsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListEmailDeliveryConfigsResponse")
	}
	return
}

// listEmailDeliveryConfigs implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) listEmailDeliveryConfigs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/emailDeliveryConfigs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListEmailDeliveryConfigsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "ListEmailDeliveryConfigs")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailDeliveryConfig/ListEmailDeliveryConfigs"
		err = common.PostProcessServiceError(err, "Email", "ListEmailDeliveryConfigs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListEmailDomains Lists email domains in the specified compartment.
func (client EmailClient) ListEmailDomains(ctx context.Context, request ListEmailDomainsRequest) (response ListEmailDomainsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listEmailDomains, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListEmailDomainsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListEmailDomainsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListEmailDomainsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListEmailDomainsResponse")
	}
	return
}

// listEmailDomains implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) listEmailDomains(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/emailDomains", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListEmailDomainsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "ListEmailDomains")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailDomain/ListEmailDomains"
		err = common.PostProcessServiceError(err, "Email", "ListEmailDomains", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListEmailIpPools Returns a list of EmailIpPools.
func (client EmailClient) ListEmailIpPools(ctx context.Context, request ListEmailIpPoolsRequest) (response ListEmailIpPoolsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listEmailIpPools, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListEmailIpPoolsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListEmailIpPoolsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListEmailIpPoolsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListEmailIpPoolsResponse")
	}
	return
}

// listEmailIpPools implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) listEmailIpPools(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/emailIpPools", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListEmailIpPoolsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "ListEmailIpPools")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailIpPoolCollection/ListEmailIpPools"
		err = common.PostProcessServiceError(err, "Email", "ListEmailIpPools", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListEmailOutboundIps Returns a list of all Outbound Public IPs assigned for a given tenant.
func (client EmailClient) ListEmailOutboundIps(ctx context.Context, request ListEmailOutboundIpsRequest) (response ListEmailOutboundIpsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listEmailOutboundIps, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListEmailOutboundIpsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListEmailOutboundIpsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListEmailOutboundIpsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListEmailOutboundIpsResponse")
	}
	return
}

// listEmailOutboundIps implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) listEmailOutboundIps(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/emailOutboundIps", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListEmailOutboundIpsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "ListEmailOutboundIps")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailOutboundIpCollection/ListEmailOutboundIps"
		err = common.PostProcessServiceError(err, "Email", "ListEmailOutboundIps", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListEmailPrivateEndpoints Returns a list of all the email private endpoints in the specified compartment.
func (client EmailClient) ListEmailPrivateEndpoints(ctx context.Context, request ListEmailPrivateEndpointsRequest) (response ListEmailPrivateEndpointsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listEmailPrivateEndpoints, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListEmailPrivateEndpointsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListEmailPrivateEndpointsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListEmailPrivateEndpointsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListEmailPrivateEndpointsResponse")
	}
	return
}

// listEmailPrivateEndpoints implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) listEmailPrivateEndpoints(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/emailPrivateEndpoints", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListEmailPrivateEndpointsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "ListEmailPrivateEndpoints")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailPrivateEndpoint/ListEmailPrivateEndpoints"
		err = common.PostProcessServiceError(err, "Email", "ListEmailPrivateEndpoints", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListEmailRecipientDomains Gets a list of email domains. The
// `compartmentId` for email domains must be a tenancy OCID. The returned list
// is sorted by creation time in descending order.
func (client EmailClient) ListEmailRecipientDomains(ctx context.Context, request ListEmailRecipientDomainsRequest) (response ListEmailRecipientDomainsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listEmailRecipientDomains, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListEmailRecipientDomainsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListEmailRecipientDomainsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListEmailRecipientDomainsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListEmailRecipientDomainsResponse")
	}
	return
}

// listEmailRecipientDomains implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) listEmailRecipientDomains(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/emailRecipientDomains", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListEmailRecipientDomainsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "ListEmailRecipientDomains")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailRecipientDomain/ListEmailRecipientDomains"
		err = common.PostProcessServiceError(err, "Email", "ListEmailRecipientDomains", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListEmailReturnPaths Lists email return paths in the specified compartment or emaildomain.
func (client EmailClient) ListEmailReturnPaths(ctx context.Context, request ListEmailReturnPathsRequest) (response ListEmailReturnPathsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listEmailReturnPaths, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListEmailReturnPathsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListEmailReturnPathsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListEmailReturnPathsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListEmailReturnPathsResponse")
	}
	return
}

// listEmailReturnPaths implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) listEmailReturnPaths(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/emailReturnPaths", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListEmailReturnPathsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "ListEmailReturnPaths")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailReturnPath/ListEmailReturnPaths"
		err = common.PostProcessServiceError(err, "Email", "ListEmailReturnPaths", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListEmailTrackConfigs Returns a list of email tracking configuration resources in the specified compartment.
func (client EmailClient) ListEmailTrackConfigs(ctx context.Context, request ListEmailTrackConfigsRequest) (response ListEmailTrackConfigsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listEmailTrackConfigs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListEmailTrackConfigsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListEmailTrackConfigsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListEmailTrackConfigsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListEmailTrackConfigsResponse")
	}
	return
}

// listEmailTrackConfigs implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) listEmailTrackConfigs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/emailTrackConfigs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListEmailTrackConfigsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "ListEmailTrackConfigs")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailTrackConfig/ListEmailTrackConfigs"
		err = common.PostProcessServiceError(err, "Email", "ListEmailTrackConfigs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSenders Gets a collection of approved sender email addresses and sender IDs.
func (client EmailClient) ListSenders(ctx context.Context, request ListSendersRequest) (response ListSendersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSenders, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSendersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSendersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSendersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSendersResponse")
	}
	return
}

// listSenders implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) listSenders(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/senders", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSendersResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "ListSenders")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/Sender/ListSenders"
		err = common.PostProcessServiceError(err, "Email", "ListSenders", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSuppressions Gets a list of suppressed recipient email addresses for a user. The
// `compartmentId` for suppressions must be a tenancy OCID. The returned list
// is sorted by creation time in descending order.
func (client EmailClient) ListSuppressions(ctx context.Context, request ListSuppressionsRequest) (response ListSuppressionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSuppressions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSuppressionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSuppressionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSuppressionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSuppressionsResponse")
	}
	return
}

// listSuppressions implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) listSuppressions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/suppressions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSuppressionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "ListSuppressions")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/Suppression/ListSuppressions"
		err = common.PostProcessServiceError(err, "Email", "ListSuppressions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Return a (paginated) list of errors for a given work request.
func (client EmailClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client EmailClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/errors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestErrorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "ListWorkRequestErrors")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/WorkRequestErrorCollection/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "Email", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Return a (paginated) list of logs for a given work request.
func (client EmailClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client EmailClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/logs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "ListWorkRequestLogs")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/WorkRequestLogEntryCollection/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "Email", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment.
func (client EmailClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client EmailClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "ListWorkRequests")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/WorkRequestSummaryCollection/ListWorkRequests"
		err = common.PostProcessServiceError(err, "Email", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchEmailDeliveryConfig Patches delivery config identified by id.
func (client EmailClient) PatchEmailDeliveryConfig(ctx context.Context, request PatchEmailDeliveryConfigRequest) (response PatchEmailDeliveryConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.patchEmailDeliveryConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchEmailDeliveryConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchEmailDeliveryConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchEmailDeliveryConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchEmailDeliveryConfigResponse")
	}
	return
}

// patchEmailDeliveryConfig implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) patchEmailDeliveryConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/emailDeliveryConfigs/{emailDeliveryConfigId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchEmailDeliveryConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "PatchEmailDeliveryConfig")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailDeliveryConfig/PatchEmailDeliveryConfig"
		err = common.PostProcessServiceError(err, "Email", "PatchEmailDeliveryConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveEmailDomainLock Remove a lock to a resource.
func (client EmailClient) RemoveEmailDomainLock(ctx context.Context, request RemoveEmailDomainLockRequest) (response RemoveEmailDomainLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.removeEmailDomainLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveEmailDomainLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveEmailDomainLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveEmailDomainLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveEmailDomainLockResponse")
	}
	return
}

// removeEmailDomainLock implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) removeEmailDomainLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/emailDomains/{emailDomainId}/actions/removeLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveEmailDomainLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "RemoveEmailDomainLock")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailDomain/RemoveEmailDomainLock"
		err = common.PostProcessServiceError(err, "Email", "RemoveEmailDomainLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveEmailIpPoolLock Remove a lock to a resource.
func (client EmailClient) RemoveEmailIpPoolLock(ctx context.Context, request RemoveEmailIpPoolLockRequest) (response RemoveEmailIpPoolLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.removeEmailIpPoolLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveEmailIpPoolLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveEmailIpPoolLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveEmailIpPoolLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveEmailIpPoolLockResponse")
	}
	return
}

// removeEmailIpPoolLock implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) removeEmailIpPoolLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/emailIpPools/{emailIpPoolId}/actions/removeLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveEmailIpPoolLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "RemoveEmailIpPoolLock")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailIpPool/RemoveEmailIpPoolLock"
		err = common.PostProcessServiceError(err, "Email", "RemoveEmailIpPoolLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveEmailOutboundIp Remove OutboundIps from EmailIpPool.
func (client EmailClient) RemoveEmailOutboundIp(ctx context.Context, request RemoveEmailOutboundIpRequest) (response RemoveEmailOutboundIpResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.removeEmailOutboundIp, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveEmailOutboundIpResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveEmailOutboundIpResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveEmailOutboundIpResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveEmailOutboundIpResponse")
	}
	return
}

// removeEmailOutboundIp implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) removeEmailOutboundIp(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/emailIpPools/{emailIpPoolId}/actions/removeEmailOutboundIp", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveEmailOutboundIpResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "RemoveEmailOutboundIp")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailIpPool/RemoveEmailOutboundIp"
		err = common.PostProcessServiceError(err, "Email", "RemoveEmailOutboundIp", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveEmailTrackConfigLock Remove a lock to a resource.
func (client EmailClient) RemoveEmailTrackConfigLock(ctx context.Context, request RemoveEmailTrackConfigLockRequest) (response RemoveEmailTrackConfigLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.removeEmailTrackConfigLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveEmailTrackConfigLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveEmailTrackConfigLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveEmailTrackConfigLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveEmailTrackConfigLockResponse")
	}
	return
}

// removeEmailTrackConfigLock implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) removeEmailTrackConfigLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/emailTrackConfigs/{emailTrackConfigId}/actions/removeLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveEmailTrackConfigLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "RemoveEmailTrackConfigLock")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailTrackConfig/RemoveEmailTrackConfigLock"
		err = common.PostProcessServiceError(err, "Email", "RemoveEmailTrackConfigLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemovePrivateEndpointLock Removes a lock to a resource.
func (client EmailClient) RemovePrivateEndpointLock(ctx context.Context, request RemovePrivateEndpointLockRequest) (response RemovePrivateEndpointLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.removePrivateEndpointLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemovePrivateEndpointLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemovePrivateEndpointLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemovePrivateEndpointLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemovePrivateEndpointLockResponse")
	}
	return
}

// removePrivateEndpointLock implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) removePrivateEndpointLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/emailPrivateEndpoints/{emailPrivateEndpointId}/actions/removeLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemovePrivateEndpointLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "RemovePrivateEndpointLock")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailPrivateEndpoint/RemovePrivateEndpointLock"
		err = common.PostProcessServiceError(err, "Email", "RemovePrivateEndpointLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveReturnPathLock Remove a lock to a resource.
func (client EmailClient) RemoveReturnPathLock(ctx context.Context, request RemoveReturnPathLockRequest) (response RemoveReturnPathLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.removeReturnPathLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveReturnPathLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveReturnPathLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveReturnPathLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveReturnPathLockResponse")
	}
	return
}

// removeReturnPathLock implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) removeReturnPathLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/emailReturnPaths/{emailReturnPathId}/actions/removeLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveReturnPathLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "RemoveReturnPathLock")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailReturnPath/RemoveReturnPathLock"
		err = common.PostProcessServiceError(err, "Email", "RemoveReturnPathLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveSenderLock Remove a lock to a resource.
func (client EmailClient) RemoveSenderLock(ctx context.Context, request RemoveSenderLockRequest) (response RemoveSenderLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.removeSenderLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveSenderLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveSenderLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveSenderLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveSenderLockResponse")
	}
	return
}

// removeSenderLock implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) removeSenderLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/senders/{senderId}/actions/removeLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveSenderLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "RemoveSenderLock")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/Sender/RemoveSenderLock"
		err = common.PostProcessServiceError(err, "Email", "RemoveSenderLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateConfiguration Updates  email configuration associated with the specified compartment. The only valid value for compartmentId query parameter is a tenancy OCID.
func (client EmailClient) UpdateConfiguration(ctx context.Context, request UpdateConfigurationRequest) (response UpdateConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateConfigurationResponse")
	}
	return
}

// updateConfiguration implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) updateConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/configuration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "UpdateConfiguration")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/Configuration/UpdateConfiguration"
		err = common.PostProcessServiceError(err, "Email", "UpdateConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDkim Modifies a DKIM.
func (client EmailClient) UpdateDkim(ctx context.Context, request UpdateDkimRequest) (response UpdateDkimResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDkim, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDkimResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDkimResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDkimResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDkimResponse")
	}
	return
}

// updateDkim implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) updateDkim(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/dkims/{dkimId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDkimResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "UpdateDkim")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/Dkim/UpdateDkim"
		err = common.PostProcessServiceError(err, "Email", "UpdateDkim", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateEmailDeliveryConfig Update delivery config identified by id.
func (client EmailClient) UpdateEmailDeliveryConfig(ctx context.Context, request UpdateEmailDeliveryConfigRequest) (response UpdateEmailDeliveryConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateEmailDeliveryConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateEmailDeliveryConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateEmailDeliveryConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateEmailDeliveryConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateEmailDeliveryConfigResponse")
	}
	return
}

// updateEmailDeliveryConfig implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) updateEmailDeliveryConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/emailDeliveryConfigs/{emailDeliveryConfigId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateEmailDeliveryConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "UpdateEmailDeliveryConfig")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailDeliveryConfig/UpdateEmailDeliveryConfig"
		err = common.PostProcessServiceError(err, "Email", "UpdateEmailDeliveryConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateEmailDeliveryConfigIpAssociation Update delivery config ip association identified by emailDeliveryConfigIpAssociationId.
func (client EmailClient) UpdateEmailDeliveryConfigIpAssociation(ctx context.Context, request UpdateEmailDeliveryConfigIpAssociationRequest) (response UpdateEmailDeliveryConfigIpAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateEmailDeliveryConfigIpAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateEmailDeliveryConfigIpAssociationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateEmailDeliveryConfigIpAssociationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateEmailDeliveryConfigIpAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateEmailDeliveryConfigIpAssociationResponse")
	}
	return
}

// updateEmailDeliveryConfigIpAssociation implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) updateEmailDeliveryConfigIpAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/emailDeliveryConfigIpAssociations/{emailDeliveryConfigIpAssociationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateEmailDeliveryConfigIpAssociationResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "UpdateEmailDeliveryConfigIpAssociation")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailDeliveryConfig/UpdateEmailDeliveryConfigIpAssociation"
		err = common.PostProcessServiceError(err, "Email", "UpdateEmailDeliveryConfigIpAssociation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateEmailDomain Modifies an email domain.
func (client EmailClient) UpdateEmailDomain(ctx context.Context, request UpdateEmailDomainRequest) (response UpdateEmailDomainResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateEmailDomain, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateEmailDomainResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateEmailDomainResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateEmailDomainResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateEmailDomainResponse")
	}
	return
}

// updateEmailDomain implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) updateEmailDomain(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/emailDomains/{emailDomainId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateEmailDomainResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "UpdateEmailDomain")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailDomain/UpdateEmailDomain"
		err = common.PostProcessServiceError(err, "Email", "UpdateEmailDomain", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateEmailIpPool Updates the EmailIpPool
func (client EmailClient) UpdateEmailIpPool(ctx context.Context, request UpdateEmailIpPoolRequest) (response UpdateEmailIpPoolResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateEmailIpPool, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateEmailIpPoolResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateEmailIpPoolResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateEmailIpPoolResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateEmailIpPoolResponse")
	}
	return
}

// updateEmailIpPool implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) updateEmailIpPool(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/emailIpPools/{emailIpPoolId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateEmailIpPoolResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "UpdateEmailIpPool")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailIpPool/UpdateEmailIpPool"
		err = common.PostProcessServiceError(err, "Email", "UpdateEmailIpPool", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateEmailPrivateEndpoint Updates the private reverse connection endpoint.
func (client EmailClient) UpdateEmailPrivateEndpoint(ctx context.Context, request UpdateEmailPrivateEndpointRequest) (response UpdateEmailPrivateEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateEmailPrivateEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateEmailPrivateEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateEmailPrivateEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateEmailPrivateEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateEmailPrivateEndpointResponse")
	}
	return
}

// updateEmailPrivateEndpoint implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) updateEmailPrivateEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/emailPrivateEndpoints/{emailPrivateEndpointId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateEmailPrivateEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "UpdateEmailPrivateEndpoint")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailPrivateEndpoint/UpdateEmailPrivateEndpoint"
		err = common.PostProcessServiceError(err, "Email", "UpdateEmailPrivateEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateEmailRecipientDomain Update recipient domain identified by id.
func (client EmailClient) UpdateEmailRecipientDomain(ctx context.Context, request UpdateEmailRecipientDomainRequest) (response UpdateEmailRecipientDomainResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateEmailRecipientDomain, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateEmailRecipientDomainResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateEmailRecipientDomainResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateEmailRecipientDomainResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateEmailRecipientDomainResponse")
	}
	return
}

// updateEmailRecipientDomain implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) updateEmailRecipientDomain(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/emailRecipientDomains/{emailRecipientDomainId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateEmailRecipientDomainResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "UpdateEmailRecipientDomain")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailRecipientDomain/UpdateEmailRecipientDomain"
		err = common.PostProcessServiceError(err, "Email", "UpdateEmailRecipientDomain", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateEmailReturnPath Modifies an email return path.
func (client EmailClient) UpdateEmailReturnPath(ctx context.Context, request UpdateEmailReturnPathRequest) (response UpdateEmailReturnPathResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateEmailReturnPath, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateEmailReturnPathResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateEmailReturnPathResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateEmailReturnPathResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateEmailReturnPathResponse")
	}
	return
}

// updateEmailReturnPath implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) updateEmailReturnPath(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/emailReturnPaths/{emailReturnPathId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateEmailReturnPathResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "UpdateEmailReturnPath")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailReturnPath/UpdateEmailReturnPath"
		err = common.PostProcessServiceError(err, "Email", "UpdateEmailReturnPath", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateEmailTrackConfig Update email tracking configuration identified by id.
func (client EmailClient) UpdateEmailTrackConfig(ctx context.Context, request UpdateEmailTrackConfigRequest) (response UpdateEmailTrackConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateEmailTrackConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateEmailTrackConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateEmailTrackConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateEmailTrackConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateEmailTrackConfigResponse")
	}
	return
}

// updateEmailTrackConfig implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) updateEmailTrackConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/emailTrackConfigs/{emailTrackConfigId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateEmailTrackConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "UpdateEmailTrackConfig")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/EmailTrackConfig/UpdateEmailTrackConfig"
		err = common.PostProcessServiceError(err, "Email", "UpdateEmailTrackConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSender Replaces the set of tags for a sender with the tags provided. If either freeform
// or defined tags are omitted, the tags for that set remain the same. Each set must
// include the full set of tags for the sender, partial updates are not permitted.
// For more information about tagging, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
func (client EmailClient) UpdateSender(ctx context.Context, request UpdateSenderRequest) (response UpdateSenderResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSender, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSenderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSenderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSenderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSenderResponse")
	}
	return
}

// updateSender implements the OCIOperation interface (enables retrying operations)
func (client EmailClient) updateSender(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/senders/{senderId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSenderResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "email", "UpdateSender")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/emaildelivery/20170907/Sender/UpdateSender"
		err = common.PostProcessServiceError(err, "Email", "UpdateSender", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
