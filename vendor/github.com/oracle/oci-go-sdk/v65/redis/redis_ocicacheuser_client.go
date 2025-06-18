// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Cache API
//
// Use the OCI Cache API to create and manage clusters. A cluster is a memory-based storage solution. For more information, see OCI Cache (https://docs.oracle.com/iaas/Content/ocicache/home.htm).
//

package redis

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// OciCacheUserClient a client for OciCacheUser
type OciCacheUserClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOciCacheUserClientWithConfigurationProvider Creates a new default OciCacheUser client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOciCacheUserClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OciCacheUserClient, err error) {
	if enabled := common.CheckForEnabledServices("redis"); !enabled {
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
	return newOciCacheUserClientFromBaseClient(baseClient, provider)
}

// NewOciCacheUserClientWithOboToken Creates a new default OciCacheUser client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewOciCacheUserClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OciCacheUserClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newOciCacheUserClientFromBaseClient(baseClient, configProvider)
}

func newOciCacheUserClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OciCacheUserClient, err error) {
	// OciCacheUser service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("OciCacheUser"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = OciCacheUserClient{BaseClient: baseClient}
	client.BasePath = "20220315"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OciCacheUserClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("redis", "https://redis.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OciCacheUserClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OciCacheUserClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeOciCacheUserCompartment Moves an OCI Cache User from one compartment to another within the same tenancy.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/redis/ChangeOciCacheUserCompartment.go.html to see an example of how to use ChangeOciCacheUserCompartment API.
// A default retry strategy applies to this operation ChangeOciCacheUserCompartment()
func (client OciCacheUserClient) ChangeOciCacheUserCompartment(ctx context.Context, request ChangeOciCacheUserCompartmentRequest) (response ChangeOciCacheUserCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeOciCacheUserCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeOciCacheUserCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeOciCacheUserCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeOciCacheUserCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeOciCacheUserCompartmentResponse")
	}
	return
}

// changeOciCacheUserCompartment implements the OCIOperation interface (enables retrying operations)
func (client OciCacheUserClient) changeOciCacheUserCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/ociCacheUsers/{ociCacheUserId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeOciCacheUserCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/ocicache/20220315/OciCacheUser/ChangeOciCacheUserCompartment"
		err = common.PostProcessServiceError(err, "OciCacheUser", "ChangeOciCacheUserCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateOciCacheUser Creates a new OCI Cache user. OCI Cache user is required to authenticate to OCI Cache cluster.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/redis/CreateOciCacheUser.go.html to see an example of how to use CreateOciCacheUser API.
// A default retry strategy applies to this operation CreateOciCacheUser()
func (client OciCacheUserClient) CreateOciCacheUser(ctx context.Context, request CreateOciCacheUserRequest) (response CreateOciCacheUserResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOciCacheUser, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOciCacheUserResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOciCacheUserResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOciCacheUserResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOciCacheUserResponse")
	}
	return
}

// createOciCacheUser implements the OCIOperation interface (enables retrying operations)
func (client OciCacheUserClient) createOciCacheUser(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/ociCacheUsers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOciCacheUserResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/ocicache/20220315/CreateOciCacheUserDetails/CreateOciCacheUser"
		err = common.PostProcessServiceError(err, "OciCacheUser", "CreateOciCacheUser", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteOciCacheUser Deletes an existing OCI Cache User based on the OCI cache user unique ID (OCID).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/redis/DeleteOciCacheUser.go.html to see an example of how to use DeleteOciCacheUser API.
// A default retry strategy applies to this operation DeleteOciCacheUser()
func (client OciCacheUserClient) DeleteOciCacheUser(ctx context.Context, request DeleteOciCacheUserRequest) (response DeleteOciCacheUserResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOciCacheUser, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOciCacheUserResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOciCacheUserResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOciCacheUserResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOciCacheUserResponse")
	}
	return
}

// deleteOciCacheUser implements the OCIOperation interface (enables retrying operations)
func (client OciCacheUserClient) deleteOciCacheUser(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/ociCacheUsers/{ociCacheUserId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOciCacheUserResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/ocicache/20220315/OciCacheUser/DeleteOciCacheUser"
		err = common.PostProcessServiceError(err, "OciCacheUser", "DeleteOciCacheUser", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetOciCacheUser Get an existing OCI Cache users based on the ID (OCID).
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/redis/GetOciCacheUser.go.html to see an example of how to use GetOciCacheUser API.
// A default retry strategy applies to this operation GetOciCacheUser()
func (client OciCacheUserClient) GetOciCacheUser(ctx context.Context, request GetOciCacheUserRequest) (response GetOciCacheUserResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getOciCacheUser, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetOciCacheUserResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetOciCacheUserResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetOciCacheUserResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetOciCacheUserResponse")
	}
	return
}

// getOciCacheUser implements the OCIOperation interface (enables retrying operations)
func (client OciCacheUserClient) getOciCacheUser(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/ociCacheUsers/{ociCacheUserId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetOciCacheUserResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/ocicache/20220315/OciCacheUser/GetOciCacheUser"
		err = common.PostProcessServiceError(err, "OciCacheUser", "GetOciCacheUser", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAttachedRedisClusters Gets a list of associated redis cluster for an OCI cache user.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/redis/ListAttachedRedisClusters.go.html to see an example of how to use ListAttachedRedisClusters API.
// A default retry strategy applies to this operation ListAttachedRedisClusters()
func (client OciCacheUserClient) ListAttachedRedisClusters(ctx context.Context, request ListAttachedRedisClustersRequest) (response ListAttachedRedisClustersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAttachedRedisClusters, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAttachedRedisClustersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAttachedRedisClustersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAttachedRedisClustersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAttachedRedisClustersResponse")
	}
	return
}

// listAttachedRedisClusters implements the OCIOperation interface (enables retrying operations)
func (client OciCacheUserClient) listAttachedRedisClusters(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/ociCacheUsers/{ociCacheUserId}/actions/getRedisClusters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAttachedRedisClustersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/ocicache/20220315/AttachedOciCacheCluster/ListAttachedRedisClusters"
		err = common.PostProcessServiceError(err, "OciCacheUser", "ListAttachedRedisClusters", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListOciCacheUsers Lists the OCI Cache users based on the supplied parameters.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/redis/ListOciCacheUsers.go.html to see an example of how to use ListOciCacheUsers API.
// A default retry strategy applies to this operation ListOciCacheUsers()
func (client OciCacheUserClient) ListOciCacheUsers(ctx context.Context, request ListOciCacheUsersRequest) (response ListOciCacheUsersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOciCacheUsers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOciCacheUsersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOciCacheUsersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOciCacheUsersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOciCacheUsersResponse")
	}
	return
}

// listOciCacheUsers implements the OCIOperation interface (enables retrying operations)
func (client OciCacheUserClient) listOciCacheUsers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/ociCacheUsers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOciCacheUsersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/ocicache/20220315/OciCacheUserSummary/ListOciCacheUsers"
		err = common.PostProcessServiceError(err, "OciCacheUser", "ListOciCacheUsers", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateOciCacheUser Update an existing OCI Cache User with new details.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/redis/UpdateOciCacheUser.go.html to see an example of how to use UpdateOciCacheUser API.
// A default retry strategy applies to this operation UpdateOciCacheUser()
func (client OciCacheUserClient) UpdateOciCacheUser(ctx context.Context, request UpdateOciCacheUserRequest) (response UpdateOciCacheUserResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOciCacheUser, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOciCacheUserResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOciCacheUserResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOciCacheUserResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOciCacheUserResponse")
	}
	return
}

// updateOciCacheUser implements the OCIOperation interface (enables retrying operations)
func (client OciCacheUserClient) updateOciCacheUser(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/ociCacheUsers/{ociCacheUserId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOciCacheUserResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/ocicache/20220315/OciCacheUser/UpdateOciCacheUser"
		err = common.PostProcessServiceError(err, "OciCacheUser", "UpdateOciCacheUser", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
