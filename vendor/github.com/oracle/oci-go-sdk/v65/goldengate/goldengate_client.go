// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// GoldenGateClient a client for GoldenGate
type GoldenGateClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewGoldenGateClientWithConfigurationProvider Creates a new default GoldenGate client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewGoldenGateClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client GoldenGateClient, err error) {
	if enabled := common.CheckForEnabledServices("goldengate"); !enabled {
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
	return newGoldenGateClientFromBaseClient(baseClient, provider)
}

// NewGoldenGateClientWithOboToken Creates a new default GoldenGate client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewGoldenGateClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client GoldenGateClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newGoldenGateClientFromBaseClient(baseClient, configProvider)
}

func newGoldenGateClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client GoldenGateClient, err error) {
	// GoldenGate service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("GoldenGate"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = GoldenGateClient{BaseClient: baseClient}
	client.BasePath = "20200407"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *GoldenGateClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("goldengate", "https://goldengate.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *GoldenGateClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *GoldenGateClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AddConnectionLock Adds a lock to a Connection resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/AddConnectionLock.go.html to see an example of how to use AddConnectionLock API.
// A default retry strategy applies to this operation AddConnectionLock()
func (client GoldenGateClient) AddConnectionLock(ctx context.Context, request AddConnectionLockRequest) (response AddConnectionLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.addConnectionLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddConnectionLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddConnectionLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddConnectionLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddConnectionLockResponse")
	}
	return
}

// addConnectionLock implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) addConnectionLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/connections/{connectionId}/actions/addLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddConnectionLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Connection/AddConnectionLock"
		err = common.PostProcessServiceError(err, "GoldenGate", "AddConnectionLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &connection{})
	return response, err
}

// AddDeploymentBackupLock Adds a lock to a DeploymentBackup resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/AddDeploymentBackupLock.go.html to see an example of how to use AddDeploymentBackupLock API.
// A default retry strategy applies to this operation AddDeploymentBackupLock()
func (client GoldenGateClient) AddDeploymentBackupLock(ctx context.Context, request AddDeploymentBackupLockRequest) (response AddDeploymentBackupLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.addDeploymentBackupLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddDeploymentBackupLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddDeploymentBackupLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddDeploymentBackupLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddDeploymentBackupLockResponse")
	}
	return
}

// addDeploymentBackupLock implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) addDeploymentBackupLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deploymentBackups/{deploymentBackupId}/actions/addLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddDeploymentBackupLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/DeploymentBackup/AddDeploymentBackupLock"
		err = common.PostProcessServiceError(err, "GoldenGate", "AddDeploymentBackupLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddDeploymentLocalPeer Adds a new local peer to the deployment, this will add the given placement to deployment placement attribute. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/AddDeploymentLocalPeer.go.html to see an example of how to use AddDeploymentLocalPeer API.
// A default retry strategy applies to this operation AddDeploymentLocalPeer()
func (client GoldenGateClient) AddDeploymentLocalPeer(ctx context.Context, request AddDeploymentLocalPeerRequest) (response AddDeploymentLocalPeerResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.addDeploymentLocalPeer, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddDeploymentLocalPeerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddDeploymentLocalPeerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddDeploymentLocalPeerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddDeploymentLocalPeerResponse")
	}
	return
}

// addDeploymentLocalPeer implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) addDeploymentLocalPeer(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deployments/{deploymentId}/actions/addLocalPeer", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddDeploymentLocalPeerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Deployment/AddDeploymentLocalPeer"
		err = common.PostProcessServiceError(err, "GoldenGate", "AddDeploymentLocalPeer", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// AddDeploymentLock Adds a lock to a Deployment resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/AddDeploymentLock.go.html to see an example of how to use AddDeploymentLock API.
// A default retry strategy applies to this operation AddDeploymentLock()
func (client GoldenGateClient) AddDeploymentLock(ctx context.Context, request AddDeploymentLockRequest) (response AddDeploymentLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.addDeploymentLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddDeploymentLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddDeploymentLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddDeploymentLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddDeploymentLockResponse")
	}
	return
}

// addDeploymentLock implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) addDeploymentLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deployments/{deploymentId}/actions/addLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddDeploymentLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Deployment/AddDeploymentLock"
		err = common.PostProcessServiceError(err, "GoldenGate", "AddDeploymentLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CancelDeploymentBackup Cancels a Deployment Backup creation process.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/CancelDeploymentBackup.go.html to see an example of how to use CancelDeploymentBackup API.
// A default retry strategy applies to this operation CancelDeploymentBackup()
func (client GoldenGateClient) CancelDeploymentBackup(ctx context.Context, request CancelDeploymentBackupRequest) (response CancelDeploymentBackupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.cancelDeploymentBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelDeploymentBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelDeploymentBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelDeploymentBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelDeploymentBackupResponse")
	}
	return
}

// cancelDeploymentBackup implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) cancelDeploymentBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deploymentBackups/{deploymentBackupId}/actions/cancel", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelDeploymentBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/DeploymentBackup/CancelDeploymentBackup"
		err = common.PostProcessServiceError(err, "GoldenGate", "CancelDeploymentBackup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CancelDeploymentUpgrade Cancels a DeploymentUpgrade, applicable only for DeploymentUpgrade in Waiting state. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/CancelDeploymentUpgrade.go.html to see an example of how to use CancelDeploymentUpgrade API.
// A default retry strategy applies to this operation CancelDeploymentUpgrade()
func (client GoldenGateClient) CancelDeploymentUpgrade(ctx context.Context, request CancelDeploymentUpgradeRequest) (response CancelDeploymentUpgradeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.cancelDeploymentUpgrade, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelDeploymentUpgradeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelDeploymentUpgradeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelDeploymentUpgradeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelDeploymentUpgradeResponse")
	}
	return
}

// cancelDeploymentUpgrade implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) cancelDeploymentUpgrade(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deploymentUpgrades/{deploymentUpgradeId}/actions/cancel", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelDeploymentUpgradeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/DeploymentUpgrade/CancelDeploymentUpgrade"
		err = common.PostProcessServiceError(err, "GoldenGate", "CancelDeploymentUpgrade", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CancelSnoozeDeploymentUpgrade Cancel snooze of a DeploymentUpgrade. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/CancelSnoozeDeploymentUpgrade.go.html to see an example of how to use CancelSnoozeDeploymentUpgrade API.
// A default retry strategy applies to this operation CancelSnoozeDeploymentUpgrade()
func (client GoldenGateClient) CancelSnoozeDeploymentUpgrade(ctx context.Context, request CancelSnoozeDeploymentUpgradeRequest) (response CancelSnoozeDeploymentUpgradeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.cancelSnoozeDeploymentUpgrade, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelSnoozeDeploymentUpgradeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelSnoozeDeploymentUpgradeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelSnoozeDeploymentUpgradeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelSnoozeDeploymentUpgradeResponse")
	}
	return
}

// cancelSnoozeDeploymentUpgrade implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) cancelSnoozeDeploymentUpgrade(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deploymentUpgrades/{deploymentUpgradeId}/actions/cancelSnooze", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelSnoozeDeploymentUpgradeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/DeploymentUpgrade/CancelSnoozeDeploymentUpgrade"
		err = common.PostProcessServiceError(err, "GoldenGate", "CancelSnoozeDeploymentUpgrade", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeConnectionCompartment Moves the Connection into a different compartment within the same tenancy. When
// provided, If-Match is checked against ETag values of the resource.  For information about
// moving resources between compartments, see Moving Resources Between
// Compartments (https://docs.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ChangeConnectionCompartment.go.html to see an example of how to use ChangeConnectionCompartment API.
// A default retry strategy applies to this operation ChangeConnectionCompartment()
func (client GoldenGateClient) ChangeConnectionCompartment(ctx context.Context, request ChangeConnectionCompartmentRequest) (response ChangeConnectionCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeConnectionCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeConnectionCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeConnectionCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeConnectionCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeConnectionCompartmentResponse")
	}
	return
}

// changeConnectionCompartment implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) changeConnectionCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/connections/{connectionId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeConnectionCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Connection/ChangeConnectionCompartment"
		err = common.PostProcessServiceError(err, "GoldenGate", "ChangeConnectionCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeDatabaseRegistrationCompartment Note: Deprecated. Use the /connections API instead.
// Moves the DatabaseRegistration into a different compartment within the same tenancy. When
// provided, If-Match is checked against ETag values of the resource.  For information about
// moving resources between compartments, see Moving Resources Between
// Compartments (https://docs.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ChangeDatabaseRegistrationCompartment.go.html to see an example of how to use ChangeDatabaseRegistrationCompartment API.
// A default retry strategy applies to this operation ChangeDatabaseRegistrationCompartment()
func (client GoldenGateClient) ChangeDatabaseRegistrationCompartment(ctx context.Context, request ChangeDatabaseRegistrationCompartmentRequest) (response ChangeDatabaseRegistrationCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeDatabaseRegistrationCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDatabaseRegistrationCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDatabaseRegistrationCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDatabaseRegistrationCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDatabaseRegistrationCompartmentResponse")
	}
	return
}

// changeDatabaseRegistrationCompartment implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) changeDatabaseRegistrationCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseRegistrations/{databaseRegistrationId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDatabaseRegistrationCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/DatabaseRegistration/ChangeDatabaseRegistrationCompartment"
		err = common.PostProcessServiceError(err, "GoldenGate", "ChangeDatabaseRegistrationCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeDeploymentBackupCompartment Moves a DeploymentBackup into a different compartment within the same tenancy.  When provided,
// If-Match is checked against ETag values of the resource.  For information about moving
// resources between compartments, see Moving Resources Between
// Compartments (https://docs.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ChangeDeploymentBackupCompartment.go.html to see an example of how to use ChangeDeploymentBackupCompartment API.
// A default retry strategy applies to this operation ChangeDeploymentBackupCompartment()
func (client GoldenGateClient) ChangeDeploymentBackupCompartment(ctx context.Context, request ChangeDeploymentBackupCompartmentRequest) (response ChangeDeploymentBackupCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeDeploymentBackupCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDeploymentBackupCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDeploymentBackupCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDeploymentBackupCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDeploymentBackupCompartmentResponse")
	}
	return
}

// changeDeploymentBackupCompartment implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) changeDeploymentBackupCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deploymentBackups/{deploymentBackupId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDeploymentBackupCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/DeploymentBackup/ChangeDeploymentBackupCompartment"
		err = common.PostProcessServiceError(err, "GoldenGate", "ChangeDeploymentBackupCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeDeploymentCompartment Moves the Deployment into a different compartment within the same tenancy.  When provided,
// If-Match is checked against ETag values of the resource.  For information about moving
// resources between compartments, see Moving Resources Between
// Compartments (https://docs.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ChangeDeploymentCompartment.go.html to see an example of how to use ChangeDeploymentCompartment API.
// A default retry strategy applies to this operation ChangeDeploymentCompartment()
func (client GoldenGateClient) ChangeDeploymentCompartment(ctx context.Context, request ChangeDeploymentCompartmentRequest) (response ChangeDeploymentCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeDeploymentCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDeploymentCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDeploymentCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDeploymentCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDeploymentCompartmentResponse")
	}
	return
}

// changeDeploymentCompartment implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) changeDeploymentCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deployments/{deploymentId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDeploymentCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Deployment/ChangeDeploymentCompartment"
		err = common.PostProcessServiceError(err, "GoldenGate", "ChangeDeploymentCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangePipelineCompartment Moves the Pipeline into a different compartment within the same tenancy. When
// provided, If-Match is checked against ETag values of the resource.  For information about
// moving resources between compartments, see Moving Resources Between
// Compartments (https://docs.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ChangePipelineCompartment.go.html to see an example of how to use ChangePipelineCompartment API.
// A default retry strategy applies to this operation ChangePipelineCompartment()
func (client GoldenGateClient) ChangePipelineCompartment(ctx context.Context, request ChangePipelineCompartmentRequest) (response ChangePipelineCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changePipelineCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangePipelineCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangePipelineCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangePipelineCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangePipelineCompartmentResponse")
	}
	return
}

// changePipelineCompartment implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) changePipelineCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/pipelines/{pipelineId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangePipelineCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Pipeline/ChangePipelineCompartment"
		err = common.PostProcessServiceError(err, "GoldenGate", "ChangePipelineCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CollectDeploymentDiagnostic Collects the diagnostic of a Deployment. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/CollectDeploymentDiagnostic.go.html to see an example of how to use CollectDeploymentDiagnostic API.
// A default retry strategy applies to this operation CollectDeploymentDiagnostic()
func (client GoldenGateClient) CollectDeploymentDiagnostic(ctx context.Context, request CollectDeploymentDiagnosticRequest) (response CollectDeploymentDiagnosticResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.collectDeploymentDiagnostic, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CollectDeploymentDiagnosticResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CollectDeploymentDiagnosticResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CollectDeploymentDiagnosticResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CollectDeploymentDiagnosticResponse")
	}
	return
}

// collectDeploymentDiagnostic implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) collectDeploymentDiagnostic(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deployments/{deploymentId}/actions/collectDiagnostics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CollectDeploymentDiagnosticResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Deployment/CollectDeploymentDiagnostic"
		err = common.PostProcessServiceError(err, "GoldenGate", "CollectDeploymentDiagnostic", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CollectPipelineDiagnostic Collects diagnostics for the pipeline
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/CollectPipelineDiagnostic.go.html to see an example of how to use CollectPipelineDiagnostic API.
// A default retry strategy applies to this operation CollectPipelineDiagnostic()
func (client GoldenGateClient) CollectPipelineDiagnostic(ctx context.Context, request CollectPipelineDiagnosticRequest) (response CollectPipelineDiagnosticResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.collectPipelineDiagnostic, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CollectPipelineDiagnosticResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CollectPipelineDiagnosticResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CollectPipelineDiagnosticResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CollectPipelineDiagnosticResponse")
	}
	return
}

// collectPipelineDiagnostic implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) collectPipelineDiagnostic(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/pipelines/{pipelineId}/actions/collectDiagnostics", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CollectPipelineDiagnosticResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Pipeline/CollectPipelineDiagnostic"
		err = common.PostProcessServiceError(err, "GoldenGate", "CollectPipelineDiagnostic", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CopyDeploymentBackup Creates a copy of a Deployment Backup.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/CopyDeploymentBackup.go.html to see an example of how to use CopyDeploymentBackup API.
// A default retry strategy applies to this operation CopyDeploymentBackup()
func (client GoldenGateClient) CopyDeploymentBackup(ctx context.Context, request CopyDeploymentBackupRequest) (response CopyDeploymentBackupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.copyDeploymentBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CopyDeploymentBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CopyDeploymentBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CopyDeploymentBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CopyDeploymentBackupResponse")
	}
	return
}

// copyDeploymentBackup implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) copyDeploymentBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deploymentBackups/{deploymentBackupId}/actions/copyToBucket", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CopyDeploymentBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/DeploymentBackup/CopyDeploymentBackup"
		err = common.PostProcessServiceError(err, "GoldenGate", "CopyDeploymentBackup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateCertificate Creates a new certificate to truststore.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/CreateCertificate.go.html to see an example of how to use CreateCertificate API.
// A default retry strategy applies to this operation CreateCertificate()
func (client GoldenGateClient) CreateCertificate(ctx context.Context, request CreateCertificateRequest) (response CreateCertificateResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createCertificate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateCertificateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateCertificateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateCertificateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateCertificateResponse")
	}
	return
}

// createCertificate implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) createCertificate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deployments/{deploymentId}/certificates", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateCertificateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/CertificateCollection/CreateCertificate"
		err = common.PostProcessServiceError(err, "GoldenGate", "CreateCertificate", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateConnection Creates a new Connection.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/CreateConnection.go.html to see an example of how to use CreateConnection API.
// A default retry strategy applies to this operation CreateConnection()
func (client GoldenGateClient) CreateConnection(ctx context.Context, request CreateConnectionRequest) (response CreateConnectionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateConnectionResponse")
	}
	return
}

// createConnection implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) createConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/connections", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Connection/CreateConnection"
		err = common.PostProcessServiceError(err, "GoldenGate", "CreateConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &connection{})
	return response, err
}

// CreateConnectionAssignment Creates a new Connection Assignment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/CreateConnectionAssignment.go.html to see an example of how to use CreateConnectionAssignment API.
// A default retry strategy applies to this operation CreateConnectionAssignment()
func (client GoldenGateClient) CreateConnectionAssignment(ctx context.Context, request CreateConnectionAssignmentRequest) (response CreateConnectionAssignmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createConnectionAssignment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateConnectionAssignmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateConnectionAssignmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateConnectionAssignmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateConnectionAssignmentResponse")
	}
	return
}

// createConnectionAssignment implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) createConnectionAssignment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/connectionAssignments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateConnectionAssignmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/ConnectionAssignment/CreateConnectionAssignment"
		err = common.PostProcessServiceError(err, "GoldenGate", "CreateConnectionAssignment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDatabaseRegistration Note: Deprecated. Use the /connections API instead.
// Creates a new DatabaseRegistration.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/CreateDatabaseRegistration.go.html to see an example of how to use CreateDatabaseRegistration API.
// A default retry strategy applies to this operation CreateDatabaseRegistration()
func (client GoldenGateClient) CreateDatabaseRegistration(ctx context.Context, request CreateDatabaseRegistrationRequest) (response CreateDatabaseRegistrationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDatabaseRegistration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDatabaseRegistrationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDatabaseRegistrationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDatabaseRegistrationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDatabaseRegistrationResponse")
	}
	return
}

// createDatabaseRegistration implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) createDatabaseRegistration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/databaseRegistrations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDatabaseRegistrationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/DatabaseRegistration/CreateDatabaseRegistration"
		err = common.PostProcessServiceError(err, "GoldenGate", "CreateDatabaseRegistration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDeployment Creates a new Deployment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/CreateDeployment.go.html to see an example of how to use CreateDeployment API.
// A default retry strategy applies to this operation CreateDeployment()
func (client GoldenGateClient) CreateDeployment(ctx context.Context, request CreateDeploymentRequest) (response CreateDeploymentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDeploymentResponse")
	}
	return
}

// createDeployment implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) createDeployment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deployments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDeploymentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Deployment/CreateDeployment"
		err = common.PostProcessServiceError(err, "GoldenGate", "CreateDeployment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDeploymentBackup Creates a new DeploymentBackup.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/CreateDeploymentBackup.go.html to see an example of how to use CreateDeploymentBackup API.
// A default retry strategy applies to this operation CreateDeploymentBackup()
func (client GoldenGateClient) CreateDeploymentBackup(ctx context.Context, request CreateDeploymentBackupRequest) (response CreateDeploymentBackupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDeploymentBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDeploymentBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDeploymentBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDeploymentBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDeploymentBackupResponse")
	}
	return
}

// createDeploymentBackup implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) createDeploymentBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deploymentBackups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDeploymentBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/DeploymentBackup/CreateDeploymentBackup"
		err = common.PostProcessServiceError(err, "GoldenGate", "CreateDeploymentBackup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreatePipeline Creates a new Pipeline.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/CreatePipeline.go.html to see an example of how to use CreatePipeline API.
// A default retry strategy applies to this operation CreatePipeline()
func (client GoldenGateClient) CreatePipeline(ctx context.Context, request CreatePipelineRequest) (response CreatePipelineResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createPipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreatePipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreatePipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreatePipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreatePipelineResponse")
	}
	return
}

// createPipeline implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) createPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/pipelines", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreatePipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Pipeline/CreatePipeline"
		err = common.PostProcessServiceError(err, "GoldenGate", "CreatePipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &pipeline{})
	return response, err
}

// DeleteCertificate Deletes the certificate from truststore.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/DeleteCertificate.go.html to see an example of how to use DeleteCertificate API.
// A default retry strategy applies to this operation DeleteCertificate()
func (client GoldenGateClient) DeleteCertificate(ctx context.Context, request DeleteCertificateRequest) (response DeleteCertificateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteCertificate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteCertificateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteCertificateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteCertificateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteCertificateResponse")
	}
	return
}

// deleteCertificate implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) deleteCertificate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/deployments/{deploymentId}/certificates/{certificateKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteCertificateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Certificate/DeleteCertificate"
		err = common.PostProcessServiceError(err, "GoldenGate", "DeleteCertificate", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteConnection Deletes a Connection.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/DeleteConnection.go.html to see an example of how to use DeleteConnection API.
// A default retry strategy applies to this operation DeleteConnection()
func (client GoldenGateClient) DeleteConnection(ctx context.Context, request DeleteConnectionRequest) (response DeleteConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteConnectionResponse")
	}
	return
}

// deleteConnection implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) deleteConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/connections/{connectionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Connection/DeleteConnection"
		err = common.PostProcessServiceError(err, "GoldenGate", "DeleteConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteConnectionAssignment Deletes a Connection Assignment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/DeleteConnectionAssignment.go.html to see an example of how to use DeleteConnectionAssignment API.
// A default retry strategy applies to this operation DeleteConnectionAssignment()
func (client GoldenGateClient) DeleteConnectionAssignment(ctx context.Context, request DeleteConnectionAssignmentRequest) (response DeleteConnectionAssignmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteConnectionAssignment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteConnectionAssignmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteConnectionAssignmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteConnectionAssignmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteConnectionAssignmentResponse")
	}
	return
}

// deleteConnectionAssignment implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) deleteConnectionAssignment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/connectionAssignments/{connectionAssignmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteConnectionAssignmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/ConnectionAssignment/DeleteConnectionAssignment"
		err = common.PostProcessServiceError(err, "GoldenGate", "DeleteConnectionAssignment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDatabaseRegistration Note: Deprecated. Use the /connections API instead.
// Deletes a DatabaseRegistration.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/DeleteDatabaseRegistration.go.html to see an example of how to use DeleteDatabaseRegistration API.
// A default retry strategy applies to this operation DeleteDatabaseRegistration()
func (client GoldenGateClient) DeleteDatabaseRegistration(ctx context.Context, request DeleteDatabaseRegistrationRequest) (response DeleteDatabaseRegistrationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDatabaseRegistration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDatabaseRegistrationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDatabaseRegistrationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDatabaseRegistrationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDatabaseRegistrationResponse")
	}
	return
}

// deleteDatabaseRegistration implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) deleteDatabaseRegistration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/databaseRegistrations/{databaseRegistrationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDatabaseRegistrationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/DatabaseRegistration/DeleteDatabaseRegistration"
		err = common.PostProcessServiceError(err, "GoldenGate", "DeleteDatabaseRegistration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDeployment Deletes the Deployment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/DeleteDeployment.go.html to see an example of how to use DeleteDeployment API.
// A default retry strategy applies to this operation DeleteDeployment()
func (client GoldenGateClient) DeleteDeployment(ctx context.Context, request DeleteDeploymentRequest) (response DeleteDeploymentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDeploymentResponse")
	}
	return
}

// deleteDeployment implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) deleteDeployment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/deployments/{deploymentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDeploymentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Deployment/DeleteDeployment"
		err = common.PostProcessServiceError(err, "GoldenGate", "DeleteDeployment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDeploymentBackup Deletes a DeploymentBackup.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/DeleteDeploymentBackup.go.html to see an example of how to use DeleteDeploymentBackup API.
// A default retry strategy applies to this operation DeleteDeploymentBackup()
func (client GoldenGateClient) DeleteDeploymentBackup(ctx context.Context, request DeleteDeploymentBackupRequest) (response DeleteDeploymentBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDeploymentBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDeploymentBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDeploymentBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDeploymentBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDeploymentBackupResponse")
	}
	return
}

// deleteDeploymentBackup implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) deleteDeploymentBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/deploymentBackups/{deploymentBackupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDeploymentBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/DeploymentBackup/DeleteDeploymentBackup"
		err = common.PostProcessServiceError(err, "GoldenGate", "DeleteDeploymentBackup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeletePipeline Deletes a Pipeline.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/DeletePipeline.go.html to see an example of how to use DeletePipeline API.
// A default retry strategy applies to this operation DeletePipeline()
func (client GoldenGateClient) DeletePipeline(ctx context.Context, request DeletePipelineRequest) (response DeletePipelineResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deletePipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeletePipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeletePipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeletePipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeletePipelineResponse")
	}
	return
}

// deletePipeline implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) deletePipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/pipelines/{pipelineId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeletePipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Pipeline/DeletePipeline"
		err = common.PostProcessServiceError(err, "GoldenGate", "DeletePipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeploymentWalletExists Checks if a wallet is already present in the deployment. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/DeploymentWalletExists.go.html to see an example of how to use DeploymentWalletExists API.
// A default retry strategy applies to this operation DeploymentWalletExists()
func (client GoldenGateClient) DeploymentWalletExists(ctx context.Context, request DeploymentWalletExistsRequest) (response DeploymentWalletExistsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deploymentWalletExists, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeploymentWalletExistsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeploymentWalletExistsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeploymentWalletExistsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeploymentWalletExistsResponse")
	}
	return
}

// deploymentWalletExists implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) deploymentWalletExists(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deployments/{deploymentId}/actions/walletExists", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeploymentWalletExistsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Deployment/DeploymentWalletExists"
		err = common.PostProcessServiceError(err, "GoldenGate", "DeploymentWalletExists", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ExportDeploymentWallet Export the OGG wallet from the deployment to OCI vault. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ExportDeploymentWallet.go.html to see an example of how to use ExportDeploymentWallet API.
// A default retry strategy applies to this operation ExportDeploymentWallet()
func (client GoldenGateClient) ExportDeploymentWallet(ctx context.Context, request ExportDeploymentWalletRequest) (response ExportDeploymentWalletResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.exportDeploymentWallet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ExportDeploymentWalletResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ExportDeploymentWalletResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ExportDeploymentWalletResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ExportDeploymentWalletResponse")
	}
	return
}

// exportDeploymentWallet implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) exportDeploymentWallet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deployments/{deploymentId}/actions/exportWallet", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ExportDeploymentWalletResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Deployment/ExportDeploymentWallet"
		err = common.PostProcessServiceError(err, "GoldenGate", "ExportDeploymentWallet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GenerateLibraryUrl Generates a Pre-Authenticated Request Object URL to a DB2 for z/OS library that needs to be uploaded to your DB2 for z/OS server in order to establish GoldenGate connections to it. For licensing reasons, the URL is accessible for 10 minutes only.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/GenerateLibraryUrl.go.html to see an example of how to use GenerateLibraryUrl API.
// A default retry strategy applies to this operation GenerateLibraryUrl()
func (client GoldenGateClient) GenerateLibraryUrl(ctx context.Context, request GenerateLibraryUrlRequest) (response GenerateLibraryUrlResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.generateLibraryUrl, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateLibraryUrlResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateLibraryUrlResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateLibraryUrlResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateLibraryUrlResponse")
	}
	return
}

// generateLibraryUrl implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) generateLibraryUrl(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deployments/{deploymentId}/actions/generateLibraryUrl", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateLibraryUrlResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Deployment/GenerateLibraryUrl"
		err = common.PostProcessServiceError(err, "GoldenGate", "GenerateLibraryUrl", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCertificate Retrieves a Certificate.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/GetCertificate.go.html to see an example of how to use GetCertificate API.
// A default retry strategy applies to this operation GetCertificate()
func (client GoldenGateClient) GetCertificate(ctx context.Context, request GetCertificateRequest) (response GetCertificateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCertificate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCertificateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCertificateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCertificateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCertificateResponse")
	}
	return
}

// getCertificate implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) getCertificate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/deployments/{deploymentId}/certificates/{certificateKey}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCertificateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Certificate/GetCertificate"
		err = common.PostProcessServiceError(err, "GoldenGate", "GetCertificate", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetConnection Retrieves a Connection.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/GetConnection.go.html to see an example of how to use GetConnection API.
// A default retry strategy applies to this operation GetConnection()
func (client GoldenGateClient) GetConnection(ctx context.Context, request GetConnectionRequest) (response GetConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetConnectionResponse")
	}
	return
}

// getConnection implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) getConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/connections/{connectionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Connection/GetConnection"
		err = common.PostProcessServiceError(err, "GoldenGate", "GetConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &connection{})
	return response, err
}

// GetConnectionAssignment Retrieves a Connection Assignment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/GetConnectionAssignment.go.html to see an example of how to use GetConnectionAssignment API.
// A default retry strategy applies to this operation GetConnectionAssignment()
func (client GoldenGateClient) GetConnectionAssignment(ctx context.Context, request GetConnectionAssignmentRequest) (response GetConnectionAssignmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getConnectionAssignment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetConnectionAssignmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetConnectionAssignmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetConnectionAssignmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetConnectionAssignmentResponse")
	}
	return
}

// getConnectionAssignment implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) getConnectionAssignment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/connectionAssignments/{connectionAssignmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetConnectionAssignmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/ConnectionAssignment/GetConnectionAssignment"
		err = common.PostProcessServiceError(err, "GoldenGate", "GetConnectionAssignment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDatabaseRegistration Note: Deprecated. Use the /connections API instead.
// Retrieves a DatabaseRegistration.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/GetDatabaseRegistration.go.html to see an example of how to use GetDatabaseRegistration API.
// A default retry strategy applies to this operation GetDatabaseRegistration()
func (client GoldenGateClient) GetDatabaseRegistration(ctx context.Context, request GetDatabaseRegistrationRequest) (response GetDatabaseRegistrationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDatabaseRegistration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDatabaseRegistrationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDatabaseRegistrationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDatabaseRegistrationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDatabaseRegistrationResponse")
	}
	return
}

// getDatabaseRegistration implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) getDatabaseRegistration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseRegistrations/{databaseRegistrationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDatabaseRegistrationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/DatabaseRegistration/GetDatabaseRegistration"
		err = common.PostProcessServiceError(err, "GoldenGate", "GetDatabaseRegistration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDeployment Retrieves a deployment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/GetDeployment.go.html to see an example of how to use GetDeployment API.
// A default retry strategy applies to this operation GetDeployment()
func (client GoldenGateClient) GetDeployment(ctx context.Context, request GetDeploymentRequest) (response GetDeploymentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDeploymentResponse")
	}
	return
}

// getDeployment implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) getDeployment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/deployments/{deploymentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDeploymentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Deployment/GetDeployment"
		err = common.PostProcessServiceError(err, "GoldenGate", "GetDeployment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDeploymentBackup Retrieves a DeploymentBackup.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/GetDeploymentBackup.go.html to see an example of how to use GetDeploymentBackup API.
// A default retry strategy applies to this operation GetDeploymentBackup()
func (client GoldenGateClient) GetDeploymentBackup(ctx context.Context, request GetDeploymentBackupRequest) (response GetDeploymentBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDeploymentBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDeploymentBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDeploymentBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDeploymentBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDeploymentBackupResponse")
	}
	return
}

// getDeploymentBackup implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) getDeploymentBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/deploymentBackups/{deploymentBackupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDeploymentBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/DeploymentBackup/GetDeploymentBackup"
		err = common.PostProcessServiceError(err, "GoldenGate", "GetDeploymentBackup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDeploymentUpgrade Retrieves a deployment upgrade.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/GetDeploymentUpgrade.go.html to see an example of how to use GetDeploymentUpgrade API.
// A default retry strategy applies to this operation GetDeploymentUpgrade()
func (client GoldenGateClient) GetDeploymentUpgrade(ctx context.Context, request GetDeploymentUpgradeRequest) (response GetDeploymentUpgradeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDeploymentUpgrade, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDeploymentUpgradeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDeploymentUpgradeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDeploymentUpgradeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDeploymentUpgradeResponse")
	}
	return
}

// getDeploymentUpgrade implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) getDeploymentUpgrade(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/deploymentUpgrades/{deploymentUpgradeId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDeploymentUpgradeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/DeploymentUpgrade/GetDeploymentUpgrade"
		err = common.PostProcessServiceError(err, "GoldenGate", "GetDeploymentUpgrade", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPipeline Retrieves a Pipeline details.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/GetPipeline.go.html to see an example of how to use GetPipeline API.
// A default retry strategy applies to this operation GetPipeline()
func (client GoldenGateClient) GetPipeline(ctx context.Context, request GetPipelineRequest) (response GetPipelineResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPipelineResponse")
	}
	return
}

// getPipeline implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) getPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/pipelines/{pipelineId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Pipeline/GetPipeline"
		err = common.PostProcessServiceError(err, "GoldenGate", "GetPipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &pipeline{})
	return response, err
}

// GetWorkRequest Retrieve the WorkRequest identified by the given OCID.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
// A default retry strategy applies to this operation GetWorkRequest()
func (client GoldenGateClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
func (client GoldenGateClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "GoldenGate", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ImportDeploymentWallet Imports an OGG wallet from the OCI Vault to the Deployment. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ImportDeploymentWallet.go.html to see an example of how to use ImportDeploymentWallet API.
// A default retry strategy applies to this operation ImportDeploymentWallet()
func (client GoldenGateClient) ImportDeploymentWallet(ctx context.Context, request ImportDeploymentWalletRequest) (response ImportDeploymentWalletResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.importDeploymentWallet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ImportDeploymentWalletResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ImportDeploymentWalletResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ImportDeploymentWalletResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ImportDeploymentWalletResponse")
	}
	return
}

// importDeploymentWallet implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) importDeploymentWallet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deployments/{deploymentId}/actions/importWallet", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ImportDeploymentWalletResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Deployment/ImportDeploymentWallet"
		err = common.PostProcessServiceError(err, "GoldenGate", "ImportDeploymentWallet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCertificates Returns a list of certificates from truststore.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListCertificates.go.html to see an example of how to use ListCertificates API.
// A default retry strategy applies to this operation ListCertificates()
func (client GoldenGateClient) ListCertificates(ctx context.Context, request ListCertificatesRequest) (response ListCertificatesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCertificates, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCertificatesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCertificatesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCertificatesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCertificatesResponse")
	}
	return
}

// listCertificates implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) listCertificates(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/deployments/{deploymentId}/certificates", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCertificatesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/CertificateCollection/ListCertificates"
		err = common.PostProcessServiceError(err, "GoldenGate", "ListCertificates", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListConnectionAssignments Lists the Connection Assignments in the compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListConnectionAssignments.go.html to see an example of how to use ListConnectionAssignments API.
// A default retry strategy applies to this operation ListConnectionAssignments()
func (client GoldenGateClient) ListConnectionAssignments(ctx context.Context, request ListConnectionAssignmentsRequest) (response ListConnectionAssignmentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listConnectionAssignments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListConnectionAssignmentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListConnectionAssignmentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListConnectionAssignmentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListConnectionAssignmentsResponse")
	}
	return
}

// listConnectionAssignments implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) listConnectionAssignments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/connectionAssignments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListConnectionAssignmentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/ConnectionAssignment/ListConnectionAssignments"
		err = common.PostProcessServiceError(err, "GoldenGate", "ListConnectionAssignments", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListConnections Lists the Connections in the compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListConnections.go.html to see an example of how to use ListConnections API.
// A default retry strategy applies to this operation ListConnections()
func (client GoldenGateClient) ListConnections(ctx context.Context, request ListConnectionsRequest) (response ListConnectionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listConnections, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListConnectionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListConnectionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListConnectionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListConnectionsResponse")
	}
	return
}

// listConnections implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) listConnections(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/connections", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListConnectionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Connection/ListConnections"
		err = common.PostProcessServiceError(err, "GoldenGate", "ListConnections", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDatabaseRegistrations Note: Deprecated. Use the /connections API instead.
// Lists the DatabaseRegistrations in the compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListDatabaseRegistrations.go.html to see an example of how to use ListDatabaseRegistrations API.
// A default retry strategy applies to this operation ListDatabaseRegistrations()
func (client GoldenGateClient) ListDatabaseRegistrations(ctx context.Context, request ListDatabaseRegistrationsRequest) (response ListDatabaseRegistrationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDatabaseRegistrations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDatabaseRegistrationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDatabaseRegistrationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDatabaseRegistrationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDatabaseRegistrationsResponse")
	}
	return
}

// listDatabaseRegistrations implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) listDatabaseRegistrations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/databaseRegistrations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDatabaseRegistrationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/DatabaseRegistration/ListDatabaseRegistrations"
		err = common.PostProcessServiceError(err, "GoldenGate", "ListDatabaseRegistrations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDeploymentBackups Lists the Backups in a compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListDeploymentBackups.go.html to see an example of how to use ListDeploymentBackups API.
// A default retry strategy applies to this operation ListDeploymentBackups()
func (client GoldenGateClient) ListDeploymentBackups(ctx context.Context, request ListDeploymentBackupsRequest) (response ListDeploymentBackupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDeploymentBackups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDeploymentBackupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDeploymentBackupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDeploymentBackupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDeploymentBackupsResponse")
	}
	return
}

// listDeploymentBackups implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) listDeploymentBackups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/deploymentBackups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDeploymentBackupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/DeploymentBackup/ListDeploymentBackups"
		err = common.PostProcessServiceError(err, "GoldenGate", "ListDeploymentBackups", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDeploymentEnvironments Returns an array of DeploymentEnvironmentDescriptor
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListDeploymentEnvironments.go.html to see an example of how to use ListDeploymentEnvironments API.
// A default retry strategy applies to this operation ListDeploymentEnvironments()
func (client GoldenGateClient) ListDeploymentEnvironments(ctx context.Context, request ListDeploymentEnvironmentsRequest) (response ListDeploymentEnvironmentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDeploymentEnvironments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDeploymentEnvironmentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDeploymentEnvironmentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDeploymentEnvironmentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDeploymentEnvironmentsResponse")
	}
	return
}

// listDeploymentEnvironments implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) listDeploymentEnvironments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/deploymentEnvironments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDeploymentEnvironmentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/DeploymentEnvironmentCollection/ListDeploymentEnvironments"
		err = common.PostProcessServiceError(err, "GoldenGate", "ListDeploymentEnvironments", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDeploymentPeers Lists the local and remote peers in a deployment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListDeploymentPeers.go.html to see an example of how to use ListDeploymentPeers API.
// A default retry strategy applies to this operation ListDeploymentPeers()
func (client GoldenGateClient) ListDeploymentPeers(ctx context.Context, request ListDeploymentPeersRequest) (response ListDeploymentPeersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDeploymentPeers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDeploymentPeersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDeploymentPeersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDeploymentPeersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDeploymentPeersResponse")
	}
	return
}

// listDeploymentPeers implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) listDeploymentPeers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/deployments/{deploymentId}/peers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDeploymentPeersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/DeploymentPeerSummary/ListDeploymentPeers"
		err = common.PostProcessServiceError(err, "GoldenGate", "ListDeploymentPeers", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDeploymentTypes Returns an array of DeploymentTypeDescriptor
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListDeploymentTypes.go.html to see an example of how to use ListDeploymentTypes API.
// A default retry strategy applies to this operation ListDeploymentTypes()
func (client GoldenGateClient) ListDeploymentTypes(ctx context.Context, request ListDeploymentTypesRequest) (response ListDeploymentTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDeploymentTypes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDeploymentTypesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDeploymentTypesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDeploymentTypesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDeploymentTypesResponse")
	}
	return
}

// listDeploymentTypes implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) listDeploymentTypes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/deploymentTypes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDeploymentTypesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/DeploymentTypeCollection/ListDeploymentTypes"
		err = common.PostProcessServiceError(err, "GoldenGate", "ListDeploymentTypes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDeploymentUpgrades Lists the Deployment Upgrades in a compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListDeploymentUpgrades.go.html to see an example of how to use ListDeploymentUpgrades API.
// A default retry strategy applies to this operation ListDeploymentUpgrades()
func (client GoldenGateClient) ListDeploymentUpgrades(ctx context.Context, request ListDeploymentUpgradesRequest) (response ListDeploymentUpgradesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDeploymentUpgrades, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDeploymentUpgradesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDeploymentUpgradesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDeploymentUpgradesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDeploymentUpgradesResponse")
	}
	return
}

// listDeploymentUpgrades implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) listDeploymentUpgrades(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/deploymentUpgrades", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDeploymentUpgradesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/DeploymentUpgrade/ListDeploymentUpgrades"
		err = common.PostProcessServiceError(err, "GoldenGate", "ListDeploymentUpgrades", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDeploymentVersions Returns the list of available deployment versions.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListDeploymentVersions.go.html to see an example of how to use ListDeploymentVersions API.
// A default retry strategy applies to this operation ListDeploymentVersions()
func (client GoldenGateClient) ListDeploymentVersions(ctx context.Context, request ListDeploymentVersionsRequest) (response ListDeploymentVersionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDeploymentVersions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDeploymentVersionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDeploymentVersionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDeploymentVersionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDeploymentVersionsResponse")
	}
	return
}

// listDeploymentVersions implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) listDeploymentVersions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/deploymentVersions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDeploymentVersionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/DeploymentVersionCollection/ListDeploymentVersions"
		err = common.PostProcessServiceError(err, "GoldenGate", "ListDeploymentVersions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDeploymentWalletsOperations Lists the wallets export/import operations to/from a deployment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListDeploymentWalletsOperations.go.html to see an example of how to use ListDeploymentWalletsOperations API.
// A default retry strategy applies to this operation ListDeploymentWalletsOperations()
func (client GoldenGateClient) ListDeploymentWalletsOperations(ctx context.Context, request ListDeploymentWalletsOperationsRequest) (response ListDeploymentWalletsOperationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDeploymentWalletsOperations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDeploymentWalletsOperationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDeploymentWalletsOperationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDeploymentWalletsOperationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDeploymentWalletsOperationsResponse")
	}
	return
}

// listDeploymentWalletsOperations implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) listDeploymentWalletsOperations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/deployments/{deploymentId}/deploymentWalletsOperations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDeploymentWalletsOperationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/DeploymentWalletsOperationSummary/ListDeploymentWalletsOperations"
		err = common.PostProcessServiceError(err, "GoldenGate", "ListDeploymentWalletsOperations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDeployments Lists the Deployments in a compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListDeployments.go.html to see an example of how to use ListDeployments API.
// A default retry strategy applies to this operation ListDeployments()
func (client GoldenGateClient) ListDeployments(ctx context.Context, request ListDeploymentsRequest) (response ListDeploymentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDeployments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDeploymentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDeploymentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDeploymentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDeploymentsResponse")
	}
	return
}

// listDeployments implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) listDeployments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/deployments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDeploymentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Deployment/ListDeployments"
		err = common.PostProcessServiceError(err, "GoldenGate", "ListDeployments", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMessages Lists the DeploymentMessages for a deployment. The sorting order is not important. By default first will be Upgrade message, next Exception message and then Storage Utilization message.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListMessages.go.html to see an example of how to use ListMessages API.
// A default retry strategy applies to this operation ListMessages()
func (client GoldenGateClient) ListMessages(ctx context.Context, request ListMessagesRequest) (response ListMessagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMessages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMessagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMessagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMessagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMessagesResponse")
	}
	return
}

// listMessages implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) listMessages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/messages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMessagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/MessageSummary/ListMessages"
		err = common.PostProcessServiceError(err, "GoldenGate", "ListMessages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPipelineInitializationSteps Retrieves a Pipeline recipe steps and its progress details.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListPipelineInitializationSteps.go.html to see an example of how to use ListPipelineInitializationSteps API.
// A default retry strategy applies to this operation ListPipelineInitializationSteps()
func (client GoldenGateClient) ListPipelineInitializationSteps(ctx context.Context, request ListPipelineInitializationStepsRequest) (response ListPipelineInitializationStepsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPipelineInitializationSteps, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPipelineInitializationStepsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPipelineInitializationStepsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPipelineInitializationStepsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPipelineInitializationStepsResponse")
	}
	return
}

// listPipelineInitializationSteps implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) listPipelineInitializationSteps(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/pipelines/{pipelineId}/initializationSteps", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPipelineInitializationStepsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Pipeline/ListPipelineInitializationSteps"
		err = common.PostProcessServiceError(err, "GoldenGate", "ListPipelineInitializationSteps", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPipelineRunningProcesses Retrieves a Pipeline's running replication process's status like Capture/Apply.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListPipelineRunningProcesses.go.html to see an example of how to use ListPipelineRunningProcesses API.
// A default retry strategy applies to this operation ListPipelineRunningProcesses()
func (client GoldenGateClient) ListPipelineRunningProcesses(ctx context.Context, request ListPipelineRunningProcessesRequest) (response ListPipelineRunningProcessesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPipelineRunningProcesses, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPipelineRunningProcessesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPipelineRunningProcessesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPipelineRunningProcessesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPipelineRunningProcessesResponse")
	}
	return
}

// listPipelineRunningProcesses implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) listPipelineRunningProcesses(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/pipelines/{pipelineId}/runningProcesses", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPipelineRunningProcessesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Pipeline/ListPipelineRunningProcesses"
		err = common.PostProcessServiceError(err, "GoldenGate", "ListPipelineRunningProcesses", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPipelineSchemaTables Returns an array of tables under the given schemas of the pipeline for given source and target schemas passed as query params.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListPipelineSchemaTables.go.html to see an example of how to use ListPipelineSchemaTables API.
// A default retry strategy applies to this operation ListPipelineSchemaTables()
func (client GoldenGateClient) ListPipelineSchemaTables(ctx context.Context, request ListPipelineSchemaTablesRequest) (response ListPipelineSchemaTablesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPipelineSchemaTables, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPipelineSchemaTablesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPipelineSchemaTablesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPipelineSchemaTablesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPipelineSchemaTablesResponse")
	}
	return
}

// listPipelineSchemaTables implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) listPipelineSchemaTables(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/pipelines/{pipelineId}/schemaTables", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPipelineSchemaTablesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Pipeline/ListPipelineSchemaTables"
		err = common.PostProcessServiceError(err, "GoldenGate", "ListPipelineSchemaTables", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPipelineSchemas Returns an array of schemas based on mapping rules for a pipeline.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListPipelineSchemas.go.html to see an example of how to use ListPipelineSchemas API.
// A default retry strategy applies to this operation ListPipelineSchemas()
func (client GoldenGateClient) ListPipelineSchemas(ctx context.Context, request ListPipelineSchemasRequest) (response ListPipelineSchemasResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPipelineSchemas, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPipelineSchemasResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPipelineSchemasResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPipelineSchemasResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPipelineSchemasResponse")
	}
	return
}

// listPipelineSchemas implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) listPipelineSchemas(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/pipelines/{pipelineId}/schemas", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPipelineSchemasResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Pipeline/ListPipelineSchemas"
		err = common.PostProcessServiceError(err, "GoldenGate", "ListPipelineSchemas", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPipelines Lists the Pipelines in the compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListPipelines.go.html to see an example of how to use ListPipelines API.
// A default retry strategy applies to this operation ListPipelines()
func (client GoldenGateClient) ListPipelines(ctx context.Context, request ListPipelinesRequest) (response ListPipelinesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPipelines, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPipelinesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPipelinesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPipelinesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPipelinesResponse")
	}
	return
}

// listPipelines implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) listPipelines(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/pipelines", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPipelinesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Pipeline/ListPipelines"
		err = common.PostProcessServiceError(err, "GoldenGate", "ListPipelines", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListRecipes Returns an array of Recipe Summary.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListRecipes.go.html to see an example of how to use ListRecipes API.
// A default retry strategy applies to this operation ListRecipes()
func (client GoldenGateClient) ListRecipes(ctx context.Context, request ListRecipesRequest) (response ListRecipesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRecipes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRecipesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRecipesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRecipesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRecipesResponse")
	}
	return
}

// listRecipes implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) listRecipes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/recipes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListRecipesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/RecipeSummaryCollection/ListRecipes"
		err = common.PostProcessServiceError(err, "GoldenGate", "ListRecipes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTrailFiles Lists the TrailFiles for a deployment.
// Deprecated: Please access trail file management functions directly on OGG console which are available since version Oracle GoldenGate 23c.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListTrailFiles.go.html to see an example of how to use ListTrailFiles API.
// A default retry strategy applies to this operation ListTrailFiles()
func (client GoldenGateClient) ListTrailFiles(ctx context.Context, request ListTrailFilesRequest) (response ListTrailFilesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTrailFiles, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTrailFilesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTrailFilesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTrailFilesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTrailFilesResponse")
	}
	return
}

// listTrailFiles implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) listTrailFiles(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/trailFiles", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTrailFilesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/TrailFileSummary/ListTrailFiles"
		err = common.PostProcessServiceError(err, "GoldenGate", "ListTrailFiles", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTrailSequences Lists the Trail Sequences for a TrailFile in a given deployment.
// Deprecated: Please access trail file management functions directly on OGG console which are available since version Oracle GoldenGate 23c.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListTrailSequences.go.html to see an example of how to use ListTrailSequences API.
// A default retry strategy applies to this operation ListTrailSequences()
func (client GoldenGateClient) ListTrailSequences(ctx context.Context, request ListTrailSequencesRequest) (response ListTrailSequencesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTrailSequences, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTrailSequencesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTrailSequencesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTrailSequencesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTrailSequencesResponse")
	}
	return
}

// listTrailSequences implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) listTrailSequences(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/trailSequences", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTrailSequencesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/TrailSequenceSummary/ListTrailSequences"
		err = common.PostProcessServiceError(err, "GoldenGate", "ListTrailSequences", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Lists work request errors.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client GoldenGateClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
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
func (client GoldenGateClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/WorkRequestError/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "GoldenGate", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Lists work request logs.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client GoldenGateClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
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
func (client GoldenGateClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/WorkRequestLogEntry/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "GoldenGate", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in the compartment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
// A default retry strategy applies to this operation ListWorkRequests()
func (client GoldenGateClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
func (client GoldenGateClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/WorkRequest/ListWorkRequests"
		err = common.PostProcessServiceError(err, "GoldenGate", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RefreshConnection Refresh the external Connection attributes.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/RefreshConnection.go.html to see an example of how to use RefreshConnection API.
// A default retry strategy applies to this operation RefreshConnection()
func (client GoldenGateClient) RefreshConnection(ctx context.Context, request RefreshConnectionRequest) (response RefreshConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.refreshConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RefreshConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RefreshConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RefreshConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RefreshConnectionResponse")
	}
	return
}

// refreshConnection implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) refreshConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/connections/{connectionId}/actions/refresh", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RefreshConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Connection/RefreshConnection"
		err = common.PostProcessServiceError(err, "GoldenGate", "RefreshConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveConnectionLock Removes a lock from a Connection resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/RemoveConnectionLock.go.html to see an example of how to use RemoveConnectionLock API.
// A default retry strategy applies to this operation RemoveConnectionLock()
func (client GoldenGateClient) RemoveConnectionLock(ctx context.Context, request RemoveConnectionLockRequest) (response RemoveConnectionLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.removeConnectionLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveConnectionLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveConnectionLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveConnectionLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveConnectionLockResponse")
	}
	return
}

// removeConnectionLock implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) removeConnectionLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/connections/{connectionId}/actions/removeLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveConnectionLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Connection/RemoveConnectionLock"
		err = common.PostProcessServiceError(err, "GoldenGate", "RemoveConnectionLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &connection{})
	return response, err
}

// RemoveDeploymentBackupLock Removes a lock from a DeploymentBackup resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/RemoveDeploymentBackupLock.go.html to see an example of how to use RemoveDeploymentBackupLock API.
// A default retry strategy applies to this operation RemoveDeploymentBackupLock()
func (client GoldenGateClient) RemoveDeploymentBackupLock(ctx context.Context, request RemoveDeploymentBackupLockRequest) (response RemoveDeploymentBackupLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.removeDeploymentBackupLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveDeploymentBackupLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveDeploymentBackupLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveDeploymentBackupLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveDeploymentBackupLockResponse")
	}
	return
}

// removeDeploymentBackupLock implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) removeDeploymentBackupLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deploymentBackups/{deploymentBackupId}/actions/removeLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveDeploymentBackupLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/DeploymentBackup/RemoveDeploymentBackupLock"
		err = common.PostProcessServiceError(err, "GoldenGate", "RemoveDeploymentBackupLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveDeploymentLocalPeer Removes a local peer of the deployment, this will remove the given placement entry from the deployment placement attribute. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/RemoveDeploymentLocalPeer.go.html to see an example of how to use RemoveDeploymentLocalPeer API.
// A default retry strategy applies to this operation RemoveDeploymentLocalPeer()
func (client GoldenGateClient) RemoveDeploymentLocalPeer(ctx context.Context, request RemoveDeploymentLocalPeerRequest) (response RemoveDeploymentLocalPeerResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.removeDeploymentLocalPeer, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveDeploymentLocalPeerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveDeploymentLocalPeerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveDeploymentLocalPeerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveDeploymentLocalPeerResponse")
	}
	return
}

// removeDeploymentLocalPeer implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) removeDeploymentLocalPeer(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deployments/{deploymentId}/actions/removeLocalPeer", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveDeploymentLocalPeerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Deployment/RemoveDeploymentLocalPeer"
		err = common.PostProcessServiceError(err, "GoldenGate", "RemoveDeploymentLocalPeer", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveDeploymentLock Removes a lock from a Deployment resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/RemoveDeploymentLock.go.html to see an example of how to use RemoveDeploymentLock API.
// A default retry strategy applies to this operation RemoveDeploymentLock()
func (client GoldenGateClient) RemoveDeploymentLock(ctx context.Context, request RemoveDeploymentLockRequest) (response RemoveDeploymentLockResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.removeDeploymentLock, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveDeploymentLockResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveDeploymentLockResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveDeploymentLockResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveDeploymentLockResponse")
	}
	return
}

// removeDeploymentLock implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) removeDeploymentLock(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deployments/{deploymentId}/actions/removeLock", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveDeploymentLockResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Deployment/RemoveDeploymentLock"
		err = common.PostProcessServiceError(err, "GoldenGate", "RemoveDeploymentLock", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RescheduleDeploymentUpgrade Reschedules a DeploymentUpgrade, applicable only for DeploymentUpgrade in Waiting state. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/RescheduleDeploymentUpgrade.go.html to see an example of how to use RescheduleDeploymentUpgrade API.
// A default retry strategy applies to this operation RescheduleDeploymentUpgrade()
func (client GoldenGateClient) RescheduleDeploymentUpgrade(ctx context.Context, request RescheduleDeploymentUpgradeRequest) (response RescheduleDeploymentUpgradeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.rescheduleDeploymentUpgrade, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RescheduleDeploymentUpgradeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RescheduleDeploymentUpgradeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RescheduleDeploymentUpgradeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RescheduleDeploymentUpgradeResponse")
	}
	return
}

// rescheduleDeploymentUpgrade implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) rescheduleDeploymentUpgrade(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deploymentUpgrades/{deploymentUpgradeId}/actions/reschedule", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RescheduleDeploymentUpgradeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/DeploymentUpgrade/RescheduleDeploymentUpgrade"
		err = common.PostProcessServiceError(err, "GoldenGate", "RescheduleDeploymentUpgrade", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RestoreDeployment Restores a Deployment from a Deployment Backup created from the same Deployment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/RestoreDeployment.go.html to see an example of how to use RestoreDeployment API.
// A default retry strategy applies to this operation RestoreDeployment()
func (client GoldenGateClient) RestoreDeployment(ctx context.Context, request RestoreDeploymentRequest) (response RestoreDeploymentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.restoreDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RestoreDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RestoreDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RestoreDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RestoreDeploymentResponse")
	}
	return
}

// restoreDeployment implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) restoreDeployment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deploymentBackups/{deploymentBackupId}/actions/restore", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RestoreDeploymentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/DeploymentBackup/RestoreDeployment"
		err = common.PostProcessServiceError(err, "GoldenGate", "RestoreDeployment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RollbackDeploymentUpgrade Rollback a deployment to it's previous version. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/RollbackDeploymentUpgrade.go.html to see an example of how to use RollbackDeploymentUpgrade API.
// A default retry strategy applies to this operation RollbackDeploymentUpgrade()
func (client GoldenGateClient) RollbackDeploymentUpgrade(ctx context.Context, request RollbackDeploymentUpgradeRequest) (response RollbackDeploymentUpgradeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.rollbackDeploymentUpgrade, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RollbackDeploymentUpgradeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RollbackDeploymentUpgradeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RollbackDeploymentUpgradeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RollbackDeploymentUpgradeResponse")
	}
	return
}

// rollbackDeploymentUpgrade implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) rollbackDeploymentUpgrade(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deploymentUpgrades/{deploymentUpgradeId}/actions/rollback", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RollbackDeploymentUpgradeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/DeploymentUpgrade/RollbackDeploymentUpgrade"
		err = common.PostProcessServiceError(err, "GoldenGate", "RollbackDeploymentUpgrade", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SnoozeDeploymentUpgrade Snooze a DeploymentUpgrade. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/SnoozeDeploymentUpgrade.go.html to see an example of how to use SnoozeDeploymentUpgrade API.
// A default retry strategy applies to this operation SnoozeDeploymentUpgrade()
func (client GoldenGateClient) SnoozeDeploymentUpgrade(ctx context.Context, request SnoozeDeploymentUpgradeRequest) (response SnoozeDeploymentUpgradeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.snoozeDeploymentUpgrade, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SnoozeDeploymentUpgradeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SnoozeDeploymentUpgradeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SnoozeDeploymentUpgradeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SnoozeDeploymentUpgradeResponse")
	}
	return
}

// snoozeDeploymentUpgrade implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) snoozeDeploymentUpgrade(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deploymentUpgrades/{deploymentUpgradeId}/actions/snooze", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SnoozeDeploymentUpgradeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/DeploymentUpgrade/SnoozeDeploymentUpgrade"
		err = common.PostProcessServiceError(err, "GoldenGate", "SnoozeDeploymentUpgrade", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StartDeployment Starts a Deployment. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/StartDeployment.go.html to see an example of how to use StartDeployment API.
// A default retry strategy applies to this operation StartDeployment()
func (client GoldenGateClient) StartDeployment(ctx context.Context, request StartDeploymentRequest) (response StartDeploymentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.startDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StartDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StartDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartDeploymentResponse")
	}
	return
}

// startDeployment implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) startDeployment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deployments/{deploymentId}/actions/start", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StartDeploymentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Deployment/StartDeployment"
		err = common.PostProcessServiceError(err, "GoldenGate", "StartDeployment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StartPipeline Starts the pipeline for data replication.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/StartPipeline.go.html to see an example of how to use StartPipeline API.
// A default retry strategy applies to this operation StartPipeline()
func (client GoldenGateClient) StartPipeline(ctx context.Context, request StartPipelineRequest) (response StartPipelineResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.startPipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StartPipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StartPipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartPipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartPipelineResponse")
	}
	return
}

// startPipeline implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) startPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/pipelines/{pipelineId}/actions/start", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StartPipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Pipeline/StartPipeline"
		err = common.PostProcessServiceError(err, "GoldenGate", "StartPipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StopDeployment Stops a Deployment. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/StopDeployment.go.html to see an example of how to use StopDeployment API.
// A default retry strategy applies to this operation StopDeployment()
func (client GoldenGateClient) StopDeployment(ctx context.Context, request StopDeploymentRequest) (response StopDeploymentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.stopDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StopDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StopDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopDeploymentResponse")
	}
	return
}

// stopDeployment implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) stopDeployment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deployments/{deploymentId}/actions/stop", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StopDeploymentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Deployment/StopDeployment"
		err = common.PostProcessServiceError(err, "GoldenGate", "StopDeployment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// StopPipeline Stops the pipeline for data replication.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/StopPipeline.go.html to see an example of how to use StopPipeline API.
// A default retry strategy applies to this operation StopPipeline()
func (client GoldenGateClient) StopPipeline(ctx context.Context, request StopPipelineRequest) (response StopPipelineResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.stopPipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StopPipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StopPipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopPipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopPipelineResponse")
	}
	return
}

// stopPipeline implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) stopPipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/pipelines/{pipelineId}/actions/stop", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StopPipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Pipeline/StopPipeline"
		err = common.PostProcessServiceError(err, "GoldenGate", "StopPipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SwitchoverDeploymentPeer Switchover to the selected standby peer, which can be a local placement or a standby deployment in different region. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/SwitchoverDeploymentPeer.go.html to see an example of how to use SwitchoverDeploymentPeer API.
// A default retry strategy applies to this operation SwitchoverDeploymentPeer()
func (client GoldenGateClient) SwitchoverDeploymentPeer(ctx context.Context, request SwitchoverDeploymentPeerRequest) (response SwitchoverDeploymentPeerResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.switchoverDeploymentPeer, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SwitchoverDeploymentPeerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SwitchoverDeploymentPeerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SwitchoverDeploymentPeerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SwitchoverDeploymentPeerResponse")
	}
	return
}

// switchoverDeploymentPeer implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) switchoverDeploymentPeer(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deployments/{deploymentId}/actions/switchover", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SwitchoverDeploymentPeerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Deployment/SwitchoverDeploymentPeer"
		err = common.PostProcessServiceError(err, "GoldenGate", "SwitchoverDeploymentPeer", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// TestConnectionAssignment Tests the connectivity between given GoldenGate deployment and one of the associated database / service.
// When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/TestConnectionAssignment.go.html to see an example of how to use TestConnectionAssignment API.
// A default retry strategy applies to this operation TestConnectionAssignment()
func (client GoldenGateClient) TestConnectionAssignment(ctx context.Context, request TestConnectionAssignmentRequest) (response TestConnectionAssignmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.testConnectionAssignment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = TestConnectionAssignmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = TestConnectionAssignmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(TestConnectionAssignmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into TestConnectionAssignmentResponse")
	}
	return
}

// testConnectionAssignment implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) testConnectionAssignment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/connectionAssignments/{connectionAssignmentId}/actions/test", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response TestConnectionAssignmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/ConnectionAssignment/TestConnectionAssignment"
		err = common.PostProcessServiceError(err, "GoldenGate", "TestConnectionAssignment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// TestPipelineConnection Tests pipeline connections against pipeline to verify the connectivity.
// When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/TestPipelineConnection.go.html to see an example of how to use TestPipelineConnection API.
// A default retry strategy applies to this operation TestPipelineConnection()
func (client GoldenGateClient) TestPipelineConnection(ctx context.Context, request TestPipelineConnectionRequest) (response TestPipelineConnectionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.testPipelineConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = TestPipelineConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = TestPipelineConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(TestPipelineConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into TestPipelineConnectionResponse")
	}
	return
}

// testPipelineConnection implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) testPipelineConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/pipelines/{pipelineId}/actions/testConnection", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response TestPipelineConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Pipeline/TestPipelineConnection"
		err = common.PostProcessServiceError(err, "GoldenGate", "TestPipelineConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateConnection Updates the Connection.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/UpdateConnection.go.html to see an example of how to use UpdateConnection API.
// A default retry strategy applies to this operation UpdateConnection()
func (client GoldenGateClient) UpdateConnection(ctx context.Context, request UpdateConnectionRequest) (response UpdateConnectionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateConnection, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateConnectionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateConnectionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateConnectionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateConnectionResponse")
	}
	return
}

// updateConnection implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) updateConnection(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/connections/{connectionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateConnectionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Connection/UpdateConnection"
		err = common.PostProcessServiceError(err, "GoldenGate", "UpdateConnection", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDatabaseRegistration Note: Deprecated. Use the /connections API instead.
// Updates the DatabaseRegistration.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/UpdateDatabaseRegistration.go.html to see an example of how to use UpdateDatabaseRegistration API.
// A default retry strategy applies to this operation UpdateDatabaseRegistration()
func (client GoldenGateClient) UpdateDatabaseRegistration(ctx context.Context, request UpdateDatabaseRegistrationRequest) (response UpdateDatabaseRegistrationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDatabaseRegistration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDatabaseRegistrationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDatabaseRegistrationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDatabaseRegistrationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDatabaseRegistrationResponse")
	}
	return
}

// updateDatabaseRegistration implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) updateDatabaseRegistration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/databaseRegistrations/{databaseRegistrationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDatabaseRegistrationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/DatabaseRegistration/UpdateDatabaseRegistration"
		err = common.PostProcessServiceError(err, "GoldenGate", "UpdateDatabaseRegistration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDeployment Modifies a Deployment.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/UpdateDeployment.go.html to see an example of how to use UpdateDeployment API.
// A default retry strategy applies to this operation UpdateDeployment()
func (client GoldenGateClient) UpdateDeployment(ctx context.Context, request UpdateDeploymentRequest) (response UpdateDeploymentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDeploymentResponse")
	}
	return
}

// updateDeployment implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) updateDeployment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/deployments/{deploymentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDeploymentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Deployment/UpdateDeployment"
		err = common.PostProcessServiceError(err, "GoldenGate", "UpdateDeployment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDeploymentBackup Modifies a Deployment Backup.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/UpdateDeploymentBackup.go.html to see an example of how to use UpdateDeploymentBackup API.
// A default retry strategy applies to this operation UpdateDeploymentBackup()
func (client GoldenGateClient) UpdateDeploymentBackup(ctx context.Context, request UpdateDeploymentBackupRequest) (response UpdateDeploymentBackupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDeploymentBackup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDeploymentBackupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDeploymentBackupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDeploymentBackupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDeploymentBackupResponse")
	}
	return
}

// updateDeploymentBackup implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) updateDeploymentBackup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/deploymentBackups/{deploymentBackupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDeploymentBackupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/DeploymentBackup/UpdateDeploymentBackup"
		err = common.PostProcessServiceError(err, "GoldenGate", "UpdateDeploymentBackup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdatePipeline Updates the Pipeline.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/UpdatePipeline.go.html to see an example of how to use UpdatePipeline API.
// A default retry strategy applies to this operation UpdatePipeline()
func (client GoldenGateClient) UpdatePipeline(ctx context.Context, request UpdatePipelineRequest) (response UpdatePipelineResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updatePipeline, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdatePipelineResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdatePipelineResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdatePipelineResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdatePipelineResponse")
	}
	return
}

// updatePipeline implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) updatePipeline(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/pipelines/{pipelineId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdatePipelineResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Pipeline/UpdatePipeline"
		err = common.PostProcessServiceError(err, "GoldenGate", "UpdatePipeline", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpgradeDeployment Upgrade a Deployment. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/UpgradeDeployment.go.html to see an example of how to use UpgradeDeployment API.
// A default retry strategy applies to this operation UpgradeDeployment()
func (client GoldenGateClient) UpgradeDeployment(ctx context.Context, request UpgradeDeploymentRequest) (response UpgradeDeploymentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.upgradeDeployment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpgradeDeploymentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpgradeDeploymentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpgradeDeploymentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpgradeDeploymentResponse")
	}
	return
}

// upgradeDeployment implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) upgradeDeployment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deployments/{deploymentId}/actions/upgrade", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpgradeDeploymentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/Deployment/UpgradeDeployment"
		err = common.PostProcessServiceError(err, "GoldenGate", "UpgradeDeployment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpgradeDeploymentUpgrade Upgrade a deployment. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/UpgradeDeploymentUpgrade.go.html to see an example of how to use UpgradeDeploymentUpgrade API.
// A default retry strategy applies to this operation UpgradeDeploymentUpgrade()
func (client GoldenGateClient) UpgradeDeploymentUpgrade(ctx context.Context, request UpgradeDeploymentUpgradeRequest) (response UpgradeDeploymentUpgradeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.upgradeDeploymentUpgrade, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpgradeDeploymentUpgradeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpgradeDeploymentUpgradeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpgradeDeploymentUpgradeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpgradeDeploymentUpgradeResponse")
	}
	return
}

// upgradeDeploymentUpgrade implements the OCIOperation interface (enables retrying operations)
func (client GoldenGateClient) upgradeDeploymentUpgrade(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/deploymentUpgrades/{deploymentUpgradeId}/actions/upgrade", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpgradeDeploymentUpgradeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/goldengate/20200407/DeploymentUpgrade/UpgradeDeploymentUpgrade"
		err = common.PostProcessServiceError(err, "GoldenGate", "UpgradeDeploymentUpgrade", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
