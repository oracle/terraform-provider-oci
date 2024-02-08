// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Key Management API
//
// Use the Key Management API to manage vaults and keys. For more information, see Managing Vaults (https://docs.cloud.oracle.com/Content/KeyManagement/Tasks/managingvaults.htm) and Managing Keys (https://docs.cloud.oracle.com/Content/KeyManagement/Tasks/managingkeys.htm).
//

package keymanagement

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// KmsHsmClusterClient a client for KmsHsmCluster
type KmsHsmClusterClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewKmsHsmClusterClientWithConfigurationProvider Creates a new default KmsHsmCluster client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewKmsHsmClusterClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client KmsHsmClusterClient, err error) {
	if enabled := common.CheckForEnabledServices("keymanagement"); !enabled {
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
	return newKmsHsmClusterClientFromBaseClient(baseClient, provider)
}

// NewKmsHsmClusterClientWithOboToken Creates a new default KmsHsmCluster client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewKmsHsmClusterClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client KmsHsmClusterClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newKmsHsmClusterClientFromBaseClient(baseClient, configProvider)
}

func newKmsHsmClusterClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client KmsHsmClusterClient, err error) {
	// KmsHsmCluster service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("KmsHsmCluster"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = KmsHsmClusterClient{BaseClient: baseClient}
	client.BasePath = ""
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *KmsHsmClusterClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("kms", "https://kms.{region}.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *KmsHsmClusterClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *KmsHsmClusterClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CancelHsmClusterDeletion Cancels deletion of specified HSM Cluster, restores it and associated HSM partitions to pre-deletion states.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/keymanagement/CancelHsmClusterDeletion.go.html to see an example of how to use CancelHsmClusterDeletion API.
func (client KmsHsmClusterClient) CancelHsmClusterDeletion(ctx context.Context, request CancelHsmClusterDeletionRequest) (response CancelHsmClusterDeletionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.cancelHsmClusterDeletion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelHsmClusterDeletionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelHsmClusterDeletionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelHsmClusterDeletionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelHsmClusterDeletionResponse")
	}
	return
}

// cancelHsmClusterDeletion implements the OCIOperation interface (enables retrying operations)
func (client KmsHsmClusterClient) cancelHsmClusterDeletion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/20180608/hsmClusters/{hsmClusterId}/actions/cancelDeletion", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelHsmClusterDeletionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/key/release/HsmCluster/CancelHsmClusterDeletion"
		err = common.PostProcessServiceError(err, "KmsHsmCluster", "CancelHsmClusterDeletion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeHsmClusterCompartment Moves a HSM Cluster resource to a different compartment within the same tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/keymanagement/ChangeHsmClusterCompartment.go.html to see an example of how to use ChangeHsmClusterCompartment API.
func (client KmsHsmClusterClient) ChangeHsmClusterCompartment(ctx context.Context, request ChangeHsmClusterCompartmentRequest) (response ChangeHsmClusterCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeHsmClusterCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeHsmClusterCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeHsmClusterCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeHsmClusterCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeHsmClusterCompartmentResponse")
	}
	return
}

// changeHsmClusterCompartment implements the OCIOperation interface (enables retrying operations)
func (client KmsHsmClusterClient) changeHsmClusterCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/20180608/hsmClusters/{hsmClusterId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeHsmClusterCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/key/release/HsmCluster/ChangeHsmClusterCompartment"
		err = common.PostProcessServiceError(err, "KmsHsmCluster", "ChangeHsmClusterCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateHsmCluster Creates a new HSM cluster resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/keymanagement/CreateHsmCluster.go.html to see an example of how to use CreateHsmCluster API.
func (client KmsHsmClusterClient) CreateHsmCluster(ctx context.Context, request CreateHsmClusterRequest) (response CreateHsmClusterResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createHsmCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateHsmClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateHsmClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateHsmClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateHsmClusterResponse")
	}
	return
}

// createHsmCluster implements the OCIOperation interface (enables retrying operations)
func (client KmsHsmClusterClient) createHsmCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/20180608/hsmClusters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateHsmClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/key/release/HsmCluster/CreateHsmCluster"
		err = common.PostProcessServiceError(err, "KmsHsmCluster", "CreateHsmCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DownloadCertificateSigningRequest Retrieves the certificate signing request for the designated HSM Cluster resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/keymanagement/DownloadCertificateSigningRequest.go.html to see an example of how to use DownloadCertificateSigningRequest API.
func (client KmsHsmClusterClient) DownloadCertificateSigningRequest(ctx context.Context, request DownloadCertificateSigningRequestRequest) (response DownloadCertificateSigningRequestResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.downloadCertificateSigningRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DownloadCertificateSigningRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DownloadCertificateSigningRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DownloadCertificateSigningRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DownloadCertificateSigningRequestResponse")
	}
	return
}

// downloadCertificateSigningRequest implements the OCIOperation interface (enables retrying operations)
func (client KmsHsmClusterClient) downloadCertificateSigningRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/20180608/hsmClusters/{hsmClusterId}/actions/downloadCertificateSigningRequest", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DownloadCertificateSigningRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/key/release/HsmCluster/DownloadCertificateSigningRequest"
		err = common.PostProcessServiceError(err, "KmsHsmCluster", "DownloadCertificateSigningRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetHsmCluster Retrieves configuration details for the specified HSM Cluster resource.
// As a provisioning operation, this call is subject to a Key Management limit that applies to
// the total number of requests across all provisioning read operations. Key Management might
// throttle this call to reject an otherwise valid request when the total rate of provisioning
// read operations exceeds 10 requests per second for a given tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/keymanagement/GetHsmCluster.go.html to see an example of how to use GetHsmCluster API.
func (client KmsHsmClusterClient) GetHsmCluster(ctx context.Context, request GetHsmClusterRequest) (response GetHsmClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getHsmCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetHsmClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetHsmClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetHsmClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetHsmClusterResponse")
	}
	return
}

// getHsmCluster implements the OCIOperation interface (enables retrying operations)
func (client KmsHsmClusterClient) getHsmCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/20180608/hsmClusters/{hsmClusterId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetHsmClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/key/release/HsmCluster/GetHsmCluster"
		err = common.PostProcessServiceError(err, "KmsHsmCluster", "GetHsmCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetHsmPartition Retrieves HSM partition details for the specified HSM cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/keymanagement/GetHsmPartition.go.html to see an example of how to use GetHsmPartition API.
func (client KmsHsmClusterClient) GetHsmPartition(ctx context.Context, request GetHsmPartitionRequest) (response GetHsmPartitionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getHsmPartition, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetHsmPartitionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetHsmPartitionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetHsmPartitionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetHsmPartitionResponse")
	}
	return
}

// getHsmPartition implements the OCIOperation interface (enables retrying operations)
func (client KmsHsmClusterClient) getHsmPartition(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/20180608/hsmClusters/{hsmClusterId}/hsmPartitions/{hsmPartitionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetHsmPartitionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/key/release/HsmPartition/GetHsmPartition"
		err = common.PostProcessServiceError(err, "KmsHsmCluster", "GetHsmPartition", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPreCoUserCredentials Retrieves Pre Crypto Officer user credentials for the specified HSM cluster.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/keymanagement/GetPreCoUserCredentials.go.html to see an example of how to use GetPreCoUserCredentials API.
func (client KmsHsmClusterClient) GetPreCoUserCredentials(ctx context.Context, request GetPreCoUserCredentialsRequest) (response GetPreCoUserCredentialsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.getPreCoUserCredentials, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPreCoUserCredentialsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPreCoUserCredentialsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPreCoUserCredentialsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPreCoUserCredentialsResponse")
	}
	return
}

// getPreCoUserCredentials implements the OCIOperation interface (enables retrying operations)
func (client KmsHsmClusterClient) getPreCoUserCredentials(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/20180608/hsmClusters/{hsmClusterId}/preCoUserCredentials", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPreCoUserCredentialsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/key/release/HsmCluster/GetPreCoUserCredentials"
		err = common.PostProcessServiceError(err, "KmsHsmCluster", "GetPreCoUserCredentials", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListHsmClusters Lists all HSM cluster resources contained within the specified compartment.
// As a provisioning operation, this call is subject to a Key Management limit that applies to
// the total number of requests across all provisioning read operations. Key Management might
// throttle this call to reject an otherwise valid request when the total rate of provisioning
// read operations exceeds 10 requests per second for a given tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/keymanagement/ListHsmClusters.go.html to see an example of how to use ListHsmClusters API.
func (client KmsHsmClusterClient) ListHsmClusters(ctx context.Context, request ListHsmClustersRequest) (response ListHsmClustersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listHsmClusters, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListHsmClustersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListHsmClustersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListHsmClustersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListHsmClustersResponse")
	}
	return
}

// listHsmClusters implements the OCIOperation interface (enables retrying operations)
func (client KmsHsmClusterClient) listHsmClusters(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/20180608/hsmClusters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListHsmClustersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/key/release/HsmCluster/ListHsmClusters"
		err = common.PostProcessServiceError(err, "KmsHsmCluster", "ListHsmClusters", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListHsmPartitions Lists all HSM partitions within the specified HSM Cluster resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/keymanagement/ListHsmPartitions.go.html to see an example of how to use ListHsmPartitions API.
func (client KmsHsmClusterClient) ListHsmPartitions(ctx context.Context, request ListHsmPartitionsRequest) (response ListHsmPartitionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listHsmPartitions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListHsmPartitionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListHsmPartitionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListHsmPartitionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListHsmPartitionsResponse")
	}
	return
}

// listHsmPartitions implements the OCIOperation interface (enables retrying operations)
func (client KmsHsmClusterClient) listHsmPartitions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/20180608/hsmClusters/{hsmClusterId}/hsmPartitions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListHsmPartitionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/key/release/HsmPartition/ListHsmPartitions"
		err = common.PostProcessServiceError(err, "KmsHsmCluster", "ListHsmPartitions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ScheduleHsmClusterDeletion Schedules HSM cluster for deletion, update its lifecycle state to 'PENDING_DELETION'
// and deletes it after the retention period.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/keymanagement/ScheduleHsmClusterDeletion.go.html to see an example of how to use ScheduleHsmClusterDeletion API.
func (client KmsHsmClusterClient) ScheduleHsmClusterDeletion(ctx context.Context, request ScheduleHsmClusterDeletionRequest) (response ScheduleHsmClusterDeletionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.scheduleHsmClusterDeletion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ScheduleHsmClusterDeletionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ScheduleHsmClusterDeletionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ScheduleHsmClusterDeletionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ScheduleHsmClusterDeletionResponse")
	}
	return
}

// scheduleHsmClusterDeletion implements the OCIOperation interface (enables retrying operations)
func (client KmsHsmClusterClient) scheduleHsmClusterDeletion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/20180608/hsmClusters/{hsmClusterId}/actions/scheduleDeletion", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ScheduleHsmClusterDeletionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/key/release/HsmCluster/ScheduleHsmClusterDeletion"
		err = common.PostProcessServiceError(err, "KmsHsmCluster", "ScheduleHsmClusterDeletion", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateHsmCluster Modifies properties of an HSM cluster resource, including `displayName`, `freeformTags` and `definedTags`.
// As a provisioning operation, this call is subject to a Key Management limit that applies to
// the total number of requests across all provisioning write operations. Key Management might
// throttle this call to reject an otherwise valid request when the total rate of provisioning
// write operations exceeds 10 requests per second for a given tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/keymanagement/UpdateHsmCluster.go.html to see an example of how to use UpdateHsmCluster API.
func (client KmsHsmClusterClient) UpdateHsmCluster(ctx context.Context, request UpdateHsmClusterRequest) (response UpdateHsmClusterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateHsmCluster, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateHsmClusterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateHsmClusterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateHsmClusterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateHsmClusterResponse")
	}
	return
}

// updateHsmCluster implements the OCIOperation interface (enables retrying operations)
func (client KmsHsmClusterClient) updateHsmCluster(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/20180608/hsmClusters/{hsmClusterId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateHsmClusterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/key/release/HsmCluster/UpdateHsmCluster"
		err = common.PostProcessServiceError(err, "KmsHsmCluster", "UpdateHsmCluster", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UploadPartitionCertificates Uploads the partition owner certificates to the HSM Cluster resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/keymanagement/UploadPartitionCertificates.go.html to see an example of how to use UploadPartitionCertificates API.
func (client KmsHsmClusterClient) UploadPartitionCertificates(ctx context.Context, request UploadPartitionCertificatesRequest) (response UploadPartitionCertificatesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.uploadPartitionCertificates, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UploadPartitionCertificatesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UploadPartitionCertificatesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UploadPartitionCertificatesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UploadPartitionCertificatesResponse")
	}
	return
}

// uploadPartitionCertificates implements the OCIOperation interface (enables retrying operations)
func (client KmsHsmClusterClient) uploadPartitionCertificates(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/20180608/hsmClusters/{hsmClusterId}/actions/uploadPartitionCertificates", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UploadPartitionCertificatesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/key/release/HsmCluster/UploadPartitionCertificates"
		err = common.PostProcessServiceError(err, "KmsHsmCluster", "UploadPartitionCertificates", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
