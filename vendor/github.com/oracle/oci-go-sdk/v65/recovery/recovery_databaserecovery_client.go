// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database Autonomous Recovery Service API
//
// Use Oracle Database Autonomous Recovery Service API to manage Protected Databases.
//

package recovery

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// DatabaseRecoveryClient a client for DatabaseRecovery
type DatabaseRecoveryClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDatabaseRecoveryClientWithConfigurationProvider Creates a new default DatabaseRecovery client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDatabaseRecoveryClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DatabaseRecoveryClient, err error) {
	if enabled := common.CheckForEnabledServices("recovery"); !enabled {
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
	return newDatabaseRecoveryClientFromBaseClient(baseClient, provider)
}

// NewDatabaseRecoveryClientWithOboToken Creates a new default DatabaseRecovery client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewDatabaseRecoveryClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DatabaseRecoveryClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDatabaseRecoveryClientFromBaseClient(baseClient, configProvider)
}

func newDatabaseRecoveryClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DatabaseRecoveryClient, err error) {
	// DatabaseRecovery service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("DatabaseRecovery"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DatabaseRecoveryClient{BaseClient: baseClient}
	client.BasePath = "20210216"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DatabaseRecoveryClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("recovery", "https://recovery.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DatabaseRecoveryClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *DatabaseRecoveryClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeProtectedDatabaseCompartment Moves a protected database resource from the existing compartment to the specified compartment. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/recovery/ChangeProtectedDatabaseCompartment.go.html to see an example of how to use ChangeProtectedDatabaseCompartment API.
// A default retry strategy applies to this operation ChangeProtectedDatabaseCompartment()
func (client DatabaseRecoveryClient) ChangeProtectedDatabaseCompartment(ctx context.Context, request ChangeProtectedDatabaseCompartmentRequest) (response ChangeProtectedDatabaseCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeProtectedDatabaseCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeProtectedDatabaseCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeProtectedDatabaseCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeProtectedDatabaseCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeProtectedDatabaseCompartmentResponse")
	}
	return
}

// changeProtectedDatabaseCompartment implements the OCIOperation interface (enables retrying operations)
func (client DatabaseRecoveryClient) changeProtectedDatabaseCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/protectedDatabases/{protectedDatabaseId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeProtectedDatabaseCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/recovery-service/20210216/ProtectedDatabase/ChangeProtectedDatabaseCompartment"
		err = common.PostProcessServiceError(err, "DatabaseRecovery", "ChangeProtectedDatabaseCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeProtectionPolicyCompartment Moves a protection policy resource from the existing compartment to the specified compartment. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/recovery/ChangeProtectionPolicyCompartment.go.html to see an example of how to use ChangeProtectionPolicyCompartment API.
// A default retry strategy applies to this operation ChangeProtectionPolicyCompartment()
func (client DatabaseRecoveryClient) ChangeProtectionPolicyCompartment(ctx context.Context, request ChangeProtectionPolicyCompartmentRequest) (response ChangeProtectionPolicyCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeProtectionPolicyCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeProtectionPolicyCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeProtectionPolicyCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeProtectionPolicyCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeProtectionPolicyCompartmentResponse")
	}
	return
}

// changeProtectionPolicyCompartment implements the OCIOperation interface (enables retrying operations)
func (client DatabaseRecoveryClient) changeProtectionPolicyCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/protectionPolicies/{protectionPolicyId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeProtectionPolicyCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/recovery-service/20210216/ProtectionPolicy/ChangeProtectionPolicyCompartment"
		err = common.PostProcessServiceError(err, "DatabaseRecovery", "ChangeProtectionPolicyCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeRecoveryServiceSubnetCompartment Moves a recovery service subnet resource from the existing compartment to the specified compartment. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/recovery/ChangeRecoveryServiceSubnetCompartment.go.html to see an example of how to use ChangeRecoveryServiceSubnetCompartment API.
// A default retry strategy applies to this operation ChangeRecoveryServiceSubnetCompartment()
func (client DatabaseRecoveryClient) ChangeRecoveryServiceSubnetCompartment(ctx context.Context, request ChangeRecoveryServiceSubnetCompartmentRequest) (response ChangeRecoveryServiceSubnetCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeRecoveryServiceSubnetCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeRecoveryServiceSubnetCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeRecoveryServiceSubnetCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeRecoveryServiceSubnetCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeRecoveryServiceSubnetCompartmentResponse")
	}
	return
}

// changeRecoveryServiceSubnetCompartment implements the OCIOperation interface (enables retrying operations)
func (client DatabaseRecoveryClient) changeRecoveryServiceSubnetCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/recoveryServiceSubnets/{recoveryServiceSubnetId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeRecoveryServiceSubnetCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/recovery-service/20210216/RecoveryServiceSubnet/ChangeRecoveryServiceSubnetCompartment"
		err = common.PostProcessServiceError(err, "DatabaseRecovery", "ChangeRecoveryServiceSubnetCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateProtectedDatabase Creates a new Protected Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/recovery/CreateProtectedDatabase.go.html to see an example of how to use CreateProtectedDatabase API.
// A default retry strategy applies to this operation CreateProtectedDatabase()
func (client DatabaseRecoveryClient) CreateProtectedDatabase(ctx context.Context, request CreateProtectedDatabaseRequest) (response CreateProtectedDatabaseResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createProtectedDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateProtectedDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateProtectedDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateProtectedDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateProtectedDatabaseResponse")
	}
	return
}

// createProtectedDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseRecoveryClient) createProtectedDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/protectedDatabases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateProtectedDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/recovery-service/20210216/ProtectedDatabase/CreateProtectedDatabase"
		err = common.PostProcessServiceError(err, "DatabaseRecovery", "CreateProtectedDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateProtectionPolicy Creates a new Protection Policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/recovery/CreateProtectionPolicy.go.html to see an example of how to use CreateProtectionPolicy API.
// A default retry strategy applies to this operation CreateProtectionPolicy()
func (client DatabaseRecoveryClient) CreateProtectionPolicy(ctx context.Context, request CreateProtectionPolicyRequest) (response CreateProtectionPolicyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createProtectionPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateProtectionPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateProtectionPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateProtectionPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateProtectionPolicyResponse")
	}
	return
}

// createProtectionPolicy implements the OCIOperation interface (enables retrying operations)
func (client DatabaseRecoveryClient) createProtectionPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/protectionPolicies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateProtectionPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/recovery-service/20210216/ProtectionPolicy/CreateProtectionPolicy"
		err = common.PostProcessServiceError(err, "DatabaseRecovery", "CreateProtectionPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateRecoveryServiceSubnet Creates a new Recovery Service Subnet.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/recovery/CreateRecoveryServiceSubnet.go.html to see an example of how to use CreateRecoveryServiceSubnet API.
// A default retry strategy applies to this operation CreateRecoveryServiceSubnet()
func (client DatabaseRecoveryClient) CreateRecoveryServiceSubnet(ctx context.Context, request CreateRecoveryServiceSubnetRequest) (response CreateRecoveryServiceSubnetResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createRecoveryServiceSubnet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateRecoveryServiceSubnetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateRecoveryServiceSubnetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateRecoveryServiceSubnetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateRecoveryServiceSubnetResponse")
	}
	return
}

// createRecoveryServiceSubnet implements the OCIOperation interface (enables retrying operations)
func (client DatabaseRecoveryClient) createRecoveryServiceSubnet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/recoveryServiceSubnets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateRecoveryServiceSubnetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/recovery-service/20210216/RecoveryServiceSubnet/CreateRecoveryServiceSubnet"
		err = common.PostProcessServiceError(err, "DatabaseRecovery", "CreateRecoveryServiceSubnet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteProtectedDatabase Deletes a protected database based on the specified protected database ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/recovery/DeleteProtectedDatabase.go.html to see an example of how to use DeleteProtectedDatabase API.
// A default retry strategy applies to this operation DeleteProtectedDatabase()
func (client DatabaseRecoveryClient) DeleteProtectedDatabase(ctx context.Context, request DeleteProtectedDatabaseRequest) (response DeleteProtectedDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteProtectedDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteProtectedDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteProtectedDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteProtectedDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteProtectedDatabaseResponse")
	}
	return
}

// deleteProtectedDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseRecoveryClient) deleteProtectedDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/protectedDatabases/{protectedDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteProtectedDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/recovery-service/20210216/ProtectedDatabase/DeleteProtectedDatabase"
		err = common.PostProcessServiceError(err, "DatabaseRecovery", "DeleteProtectedDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteProtectionPolicy Deletes a specified protection policy. You can delete custom policies only.
// Deleting a Oracle predefined policies will result in status code 405 Method Not Allowed.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/recovery/DeleteProtectionPolicy.go.html to see an example of how to use DeleteProtectionPolicy API.
// A default retry strategy applies to this operation DeleteProtectionPolicy()
func (client DatabaseRecoveryClient) DeleteProtectionPolicy(ctx context.Context, request DeleteProtectionPolicyRequest) (response DeleteProtectionPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteProtectionPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteProtectionPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteProtectionPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteProtectionPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteProtectionPolicyResponse")
	}
	return
}

// deleteProtectionPolicy implements the OCIOperation interface (enables retrying operations)
func (client DatabaseRecoveryClient) deleteProtectionPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/protectionPolicies/{protectionPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteProtectionPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/recovery-service/20210216/ProtectionPolicy/DeleteProtectionPolicy"
		err = common.PostProcessServiceError(err, "DatabaseRecovery", "DeleteProtectionPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteRecoveryServiceSubnet Deletes a specified recovery service subnet.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/recovery/DeleteRecoveryServiceSubnet.go.html to see an example of how to use DeleteRecoveryServiceSubnet API.
// A default retry strategy applies to this operation DeleteRecoveryServiceSubnet()
func (client DatabaseRecoveryClient) DeleteRecoveryServiceSubnet(ctx context.Context, request DeleteRecoveryServiceSubnetRequest) (response DeleteRecoveryServiceSubnetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteRecoveryServiceSubnet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteRecoveryServiceSubnetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteRecoveryServiceSubnetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteRecoveryServiceSubnetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteRecoveryServiceSubnetResponse")
	}
	return
}

// deleteRecoveryServiceSubnet implements the OCIOperation interface (enables retrying operations)
func (client DatabaseRecoveryClient) deleteRecoveryServiceSubnet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/recoveryServiceSubnets/{recoveryServiceSubnetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteRecoveryServiceSubnetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/recovery-service/20210216/RecoveryServiceSubnet/DeleteRecoveryServiceSubnet"
		err = common.PostProcessServiceError(err, "DatabaseRecovery", "DeleteRecoveryServiceSubnet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// FetchProtectedDatabaseConfiguration Downloads the network service configuration file 'tnsnames.ora' for a specified protected database. Applies to user-defined recovery systems only.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/recovery/FetchProtectedDatabaseConfiguration.go.html to see an example of how to use FetchProtectedDatabaseConfiguration API.
// A default retry strategy applies to this operation FetchProtectedDatabaseConfiguration()
func (client DatabaseRecoveryClient) FetchProtectedDatabaseConfiguration(ctx context.Context, request FetchProtectedDatabaseConfigurationRequest) (response FetchProtectedDatabaseConfigurationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.fetchProtectedDatabaseConfiguration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = FetchProtectedDatabaseConfigurationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = FetchProtectedDatabaseConfigurationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(FetchProtectedDatabaseConfigurationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into FetchProtectedDatabaseConfigurationResponse")
	}
	return
}

// fetchProtectedDatabaseConfiguration implements the OCIOperation interface (enables retrying operations)
func (client DatabaseRecoveryClient) fetchProtectedDatabaseConfiguration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/protectedDatabases/{protectedDatabaseId}/actions/getConfiguration", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response FetchProtectedDatabaseConfigurationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/recovery-service/20210216/ProtectedDatabase/FetchProtectedDatabaseConfiguration"
		err = common.PostProcessServiceError(err, "DatabaseRecovery", "FetchProtectedDatabaseConfiguration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetProtectedDatabase Gets information about a specified protected database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/recovery/GetProtectedDatabase.go.html to see an example of how to use GetProtectedDatabase API.
// A default retry strategy applies to this operation GetProtectedDatabase()
func (client DatabaseRecoveryClient) GetProtectedDatabase(ctx context.Context, request GetProtectedDatabaseRequest) (response GetProtectedDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getProtectedDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetProtectedDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetProtectedDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetProtectedDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetProtectedDatabaseResponse")
	}
	return
}

// getProtectedDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseRecoveryClient) getProtectedDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/protectedDatabases/{protectedDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetProtectedDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/recovery-service/20210216/ProtectedDatabase/GetProtectedDatabase"
		err = common.PostProcessServiceError(err, "DatabaseRecovery", "GetProtectedDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetProtectionPolicy Gets information about a specified protection policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/recovery/GetProtectionPolicy.go.html to see an example of how to use GetProtectionPolicy API.
// A default retry strategy applies to this operation GetProtectionPolicy()
func (client DatabaseRecoveryClient) GetProtectionPolicy(ctx context.Context, request GetProtectionPolicyRequest) (response GetProtectionPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getProtectionPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetProtectionPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetProtectionPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetProtectionPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetProtectionPolicyResponse")
	}
	return
}

// getProtectionPolicy implements the OCIOperation interface (enables retrying operations)
func (client DatabaseRecoveryClient) getProtectionPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/protectionPolicies/{protectionPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetProtectionPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/recovery-service/20210216/ProtectionPolicy/GetProtectionPolicy"
		err = common.PostProcessServiceError(err, "DatabaseRecovery", "GetProtectionPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetRecoveryServiceSubnet Gets information about a specified recovery service subnet.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/recovery/GetRecoveryServiceSubnet.go.html to see an example of how to use GetRecoveryServiceSubnet API.
// A default retry strategy applies to this operation GetRecoveryServiceSubnet()
func (client DatabaseRecoveryClient) GetRecoveryServiceSubnet(ctx context.Context, request GetRecoveryServiceSubnetRequest) (response GetRecoveryServiceSubnetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getRecoveryServiceSubnet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetRecoveryServiceSubnetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetRecoveryServiceSubnetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetRecoveryServiceSubnetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetRecoveryServiceSubnetResponse")
	}
	return
}

// getRecoveryServiceSubnet implements the OCIOperation interface (enables retrying operations)
func (client DatabaseRecoveryClient) getRecoveryServiceSubnet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/recoveryServiceSubnets/{recoveryServiceSubnetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetRecoveryServiceSubnetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/recovery-service/20210216/RecoveryServiceSubnet/GetRecoveryServiceSubnet"
		err = common.PostProcessServiceError(err, "DatabaseRecovery", "GetRecoveryServiceSubnet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets the status of the work request based on the specified ID
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/recovery/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
// A default retry strategy applies to this operation GetWorkRequest()
func (client DatabaseRecoveryClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
func (client DatabaseRecoveryClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/recovery-service/20210216/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "DatabaseRecovery", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListProtectedDatabases Lists the protected databases based on the specified parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/recovery/ListProtectedDatabases.go.html to see an example of how to use ListProtectedDatabases API.
// A default retry strategy applies to this operation ListProtectedDatabases()
func (client DatabaseRecoveryClient) ListProtectedDatabases(ctx context.Context, request ListProtectedDatabasesRequest) (response ListProtectedDatabasesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listProtectedDatabases, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListProtectedDatabasesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListProtectedDatabasesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListProtectedDatabasesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListProtectedDatabasesResponse")
	}
	return
}

// listProtectedDatabases implements the OCIOperation interface (enables retrying operations)
func (client DatabaseRecoveryClient) listProtectedDatabases(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/protectedDatabases", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListProtectedDatabasesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/recovery-service/20210216/ProtectedDatabaseCollection/ListProtectedDatabases"
		err = common.PostProcessServiceError(err, "DatabaseRecovery", "ListProtectedDatabases", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListProtectionPolicies Gets a list of protection policies based on the specified parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/recovery/ListProtectionPolicies.go.html to see an example of how to use ListProtectionPolicies API.
// A default retry strategy applies to this operation ListProtectionPolicies()
func (client DatabaseRecoveryClient) ListProtectionPolicies(ctx context.Context, request ListProtectionPoliciesRequest) (response ListProtectionPoliciesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listProtectionPolicies, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListProtectionPoliciesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListProtectionPoliciesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListProtectionPoliciesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListProtectionPoliciesResponse")
	}
	return
}

// listProtectionPolicies implements the OCIOperation interface (enables retrying operations)
func (client DatabaseRecoveryClient) listProtectionPolicies(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/protectionPolicies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListProtectionPoliciesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/recovery-service/20210216/ProtectionPolicyCollection/ListProtectionPolicies"
		err = common.PostProcessServiceError(err, "DatabaseRecovery", "ListProtectionPolicies", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListRecoveryServiceSubnets Returns a list of Recovery Service Subnets.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/recovery/ListRecoveryServiceSubnets.go.html to see an example of how to use ListRecoveryServiceSubnets API.
// A default retry strategy applies to this operation ListRecoveryServiceSubnets()
func (client DatabaseRecoveryClient) ListRecoveryServiceSubnets(ctx context.Context, request ListRecoveryServiceSubnetsRequest) (response ListRecoveryServiceSubnetsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRecoveryServiceSubnets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRecoveryServiceSubnetsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRecoveryServiceSubnetsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRecoveryServiceSubnetsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRecoveryServiceSubnetsResponse")
	}
	return
}

// listRecoveryServiceSubnets implements the OCIOperation interface (enables retrying operations)
func (client DatabaseRecoveryClient) listRecoveryServiceSubnets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/recoveryServiceSubnets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListRecoveryServiceSubnetsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/recovery-service/20210216/RecoveryServiceSubnetCollection/ListRecoveryServiceSubnets"
		err = common.PostProcessServiceError(err, "DatabaseRecovery", "ListRecoveryServiceSubnets", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Return a (paginated) list of errors for a given work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/recovery/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client DatabaseRecoveryClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
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
func (client DatabaseRecoveryClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/recovery-service/20210216/WorkRequestErrorCollection/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "DatabaseRecovery", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Return a (paginated) list of logs for a given work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/recovery/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client DatabaseRecoveryClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
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
func (client DatabaseRecoveryClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/recovery-service/20210216/WorkRequestLogEntryCollection/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "DatabaseRecovery", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/recovery/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
// A default retry strategy applies to this operation ListWorkRequests()
func (client DatabaseRecoveryClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
func (client DatabaseRecoveryClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/recovery-service/20210216/WorkRequestSummaryCollection/ListWorkRequests"
		err = common.PostProcessServiceError(err, "DatabaseRecovery", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateProtectedDatabase Updates the Protected Database
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/recovery/UpdateProtectedDatabase.go.html to see an example of how to use UpdateProtectedDatabase API.
// A default retry strategy applies to this operation UpdateProtectedDatabase()
func (client DatabaseRecoveryClient) UpdateProtectedDatabase(ctx context.Context, request UpdateProtectedDatabaseRequest) (response UpdateProtectedDatabaseResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateProtectedDatabase, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateProtectedDatabaseResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateProtectedDatabaseResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateProtectedDatabaseResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateProtectedDatabaseResponse")
	}
	return
}

// updateProtectedDatabase implements the OCIOperation interface (enables retrying operations)
func (client DatabaseRecoveryClient) updateProtectedDatabase(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/protectedDatabases/{protectedDatabaseId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateProtectedDatabaseResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/recovery-service/20210216/ProtectedDatabase/UpdateProtectedDatabase"
		err = common.PostProcessServiceError(err, "DatabaseRecovery", "UpdateProtectedDatabase", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateProtectionPolicy Updates the specified protection policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/recovery/UpdateProtectionPolicy.go.html to see an example of how to use UpdateProtectionPolicy API.
// A default retry strategy applies to this operation UpdateProtectionPolicy()
func (client DatabaseRecoveryClient) UpdateProtectionPolicy(ctx context.Context, request UpdateProtectionPolicyRequest) (response UpdateProtectionPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateProtectionPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateProtectionPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateProtectionPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateProtectionPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateProtectionPolicyResponse")
	}
	return
}

// updateProtectionPolicy implements the OCIOperation interface (enables retrying operations)
func (client DatabaseRecoveryClient) updateProtectionPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/protectionPolicies/{protectionPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateProtectionPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/recovery-service/20210216/ProtectionPolicy/UpdateProtectionPolicy"
		err = common.PostProcessServiceError(err, "DatabaseRecovery", "UpdateProtectionPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateRecoveryServiceSubnet Updates the specified recovery service subnet.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/recovery/UpdateRecoveryServiceSubnet.go.html to see an example of how to use UpdateRecoveryServiceSubnet API.
// A default retry strategy applies to this operation UpdateRecoveryServiceSubnet()
func (client DatabaseRecoveryClient) UpdateRecoveryServiceSubnet(ctx context.Context, request UpdateRecoveryServiceSubnetRequest) (response UpdateRecoveryServiceSubnetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateRecoveryServiceSubnet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateRecoveryServiceSubnetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateRecoveryServiceSubnetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateRecoveryServiceSubnetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateRecoveryServiceSubnetResponse")
	}
	return
}

// updateRecoveryServiceSubnet implements the OCIOperation interface (enables retrying operations)
func (client DatabaseRecoveryClient) updateRecoveryServiceSubnet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/recoveryServiceSubnets/{recoveryServiceSubnetId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateRecoveryServiceSubnetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/recovery-service/20210216/RecoveryServiceSubnet/UpdateRecoveryServiceSubnet"
		err = common.PostProcessServiceError(err, "DatabaseRecovery", "UpdateRecoveryServiceSubnet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
