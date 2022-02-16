// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Management API
//
// API for managing certificates.
//

package certificatesmanagement

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"github.com/oracle/oci-go-sdk/v58/common/auth"
	"net/http"
)

//CertificatesManagementClient a client for CertificatesManagement
type CertificatesManagementClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewCertificatesManagementClientWithConfigurationProvider Creates a new default CertificatesManagement client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewCertificatesManagementClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client CertificatesManagementClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newCertificatesManagementClientFromBaseClient(baseClient, provider)
}

// NewCertificatesManagementClientWithOboToken Creates a new default CertificatesManagement client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewCertificatesManagementClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client CertificatesManagementClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newCertificatesManagementClientFromBaseClient(baseClient, configProvider)
}

func newCertificatesManagementClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client CertificatesManagementClient, err error) {
	// CertificatesManagement service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSetting())
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = CertificatesManagementClient{BaseClient: baseClient}
	client.BasePath = "20210224"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *CertificatesManagementClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("certificatesmanagement", "https://certificatesmanagement.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *CertificatesManagementClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *CertificatesManagementClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CancelCertificateAuthorityDeletion Cancels the scheduled deletion of the specified certificate authority (CA).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/CancelCertificateAuthorityDeletion.go.html to see an example of how to use CancelCertificateAuthorityDeletion API.
func (client CertificatesManagementClient) CancelCertificateAuthorityDeletion(ctx context.Context, request CancelCertificateAuthorityDeletionRequest) (response CancelCertificateAuthorityDeletionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.cancelCertificateAuthorityDeletion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelCertificateAuthorityDeletionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelCertificateAuthorityDeletionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelCertificateAuthorityDeletionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelCertificateAuthorityDeletionResponse")
	}
	return
}

// cancelCertificateAuthorityDeletion implements the OCIOperation interface (enables retrying operations)
func (client CertificatesManagementClient) cancelCertificateAuthorityDeletion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/certificateAuthorities/{certificateAuthorityId}/actions/cancelDeletion", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelCertificateAuthorityDeletionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CancelCertificateAuthorityVersionDeletion Cancels the scheduled deletion of the specified certificate authority (CA) version. Canceling
// a scheduled deletion restores the CA version's lifecycle state to what
// it was before its scheduled deletion.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/CancelCertificateAuthorityVersionDeletion.go.html to see an example of how to use CancelCertificateAuthorityVersionDeletion API.
func (client CertificatesManagementClient) CancelCertificateAuthorityVersionDeletion(ctx context.Context, request CancelCertificateAuthorityVersionDeletionRequest) (response CancelCertificateAuthorityVersionDeletionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.cancelCertificateAuthorityVersionDeletion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelCertificateAuthorityVersionDeletionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelCertificateAuthorityVersionDeletionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelCertificateAuthorityVersionDeletionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelCertificateAuthorityVersionDeletionResponse")
	}
	return
}

// cancelCertificateAuthorityVersionDeletion implements the OCIOperation interface (enables retrying operations)
func (client CertificatesManagementClient) cancelCertificateAuthorityVersionDeletion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/certificateAuthorities/{certificateAuthorityId}/version/{certificateAuthorityVersionNumber}/actions/cancelDeletion", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelCertificateAuthorityVersionDeletionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CancelCertificateDeletion Cancels the pending deletion of the specified certificate. Canceling
// a scheduled deletion restores the certificate's lifecycle state to what
// it was before you scheduled the certificate for deletion.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/CancelCertificateDeletion.go.html to see an example of how to use CancelCertificateDeletion API.
func (client CertificatesManagementClient) CancelCertificateDeletion(ctx context.Context, request CancelCertificateDeletionRequest) (response CancelCertificateDeletionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.cancelCertificateDeletion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelCertificateDeletionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelCertificateDeletionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelCertificateDeletionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelCertificateDeletionResponse")
	}
	return
}

// cancelCertificateDeletion implements the OCIOperation interface (enables retrying operations)
func (client CertificatesManagementClient) cancelCertificateDeletion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/certificates/{certificateId}/actions/cancelDeletion", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelCertificateDeletionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CancelCertificateVersionDeletion Cancels the scheduled deletion of the specified certificate version.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/CancelCertificateVersionDeletion.go.html to see an example of how to use CancelCertificateVersionDeletion API.
func (client CertificatesManagementClient) CancelCertificateVersionDeletion(ctx context.Context, request CancelCertificateVersionDeletionRequest) (response CancelCertificateVersionDeletionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.cancelCertificateVersionDeletion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelCertificateVersionDeletionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelCertificateVersionDeletionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelCertificateVersionDeletionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelCertificateVersionDeletionResponse")
	}
	return
}

// cancelCertificateVersionDeletion implements the OCIOperation interface (enables retrying operations)
func (client CertificatesManagementClient) cancelCertificateVersionDeletion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/certificates/{certificateId}/version/{certificateVersionNumber}/actions/cancelDeletion", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelCertificateVersionDeletionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeCaBundleCompartment Moves a CA bundle to a different compartment in the same tenancy. For information about
// moving resources between compartments, see Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
// When provided, if-match is checked against the ETag values of the secret.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/ChangeCaBundleCompartment.go.html to see an example of how to use ChangeCaBundleCompartment API.
func (client CertificatesManagementClient) ChangeCaBundleCompartment(ctx context.Context, request ChangeCaBundleCompartmentRequest) (response ChangeCaBundleCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeCaBundleCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeCaBundleCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeCaBundleCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeCaBundleCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeCaBundleCompartmentResponse")
	}
	return
}

// changeCaBundleCompartment implements the OCIOperation interface (enables retrying operations)
func (client CertificatesManagementClient) changeCaBundleCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/caBundles/{caBundleId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeCaBundleCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeCertificateAuthorityCompartment Moves a certificate authority (CA) to a different compartment within the same tenancy. For information about
// moving resources between compartments, see Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
// When provided, If-Match is checked against the ETag values of the source.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/ChangeCertificateAuthorityCompartment.go.html to see an example of how to use ChangeCertificateAuthorityCompartment API.
func (client CertificatesManagementClient) ChangeCertificateAuthorityCompartment(ctx context.Context, request ChangeCertificateAuthorityCompartmentRequest) (response ChangeCertificateAuthorityCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeCertificateAuthorityCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeCertificateAuthorityCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeCertificateAuthorityCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeCertificateAuthorityCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeCertificateAuthorityCompartmentResponse")
	}
	return
}

// changeCertificateAuthorityCompartment implements the OCIOperation interface (enables retrying operations)
func (client CertificatesManagementClient) changeCertificateAuthorityCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/certificateAuthorities/{certificateAuthorityId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeCertificateAuthorityCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeCertificateCompartment Moves a certificate to a different compartment within the same tenancy. For information about
// moving resources between compartments, see Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
// When provided, if-match is checked against the ETag values of the secret.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/ChangeCertificateCompartment.go.html to see an example of how to use ChangeCertificateCompartment API.
func (client CertificatesManagementClient) ChangeCertificateCompartment(ctx context.Context, request ChangeCertificateCompartmentRequest) (response ChangeCertificateCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeCertificateCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeCertificateCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeCertificateCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeCertificateCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeCertificateCompartmentResponse")
	}
	return
}

// changeCertificateCompartment implements the OCIOperation interface (enables retrying operations)
func (client CertificatesManagementClient) changeCertificateCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/certificates/{certificateId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeCertificateCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateCaBundle Creates a new CA bundle according to the details of the request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/CreateCaBundle.go.html to see an example of how to use CreateCaBundle API.
func (client CertificatesManagementClient) CreateCaBundle(ctx context.Context, request CreateCaBundleRequest) (response CreateCaBundleResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createCaBundle, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateCaBundleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateCaBundleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateCaBundleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateCaBundleResponse")
	}
	return
}

// createCaBundle implements the OCIOperation interface (enables retrying operations)
func (client CertificatesManagementClient) createCaBundle(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/caBundles", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateCaBundleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateCertificate Creates a new certificate according to the details of the request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/CreateCertificate.go.html to see an example of how to use CreateCertificate API.
func (client CertificatesManagementClient) CreateCertificate(ctx context.Context, request CreateCertificateRequest) (response CreateCertificateResponse, err error) {
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
func (client CertificatesManagementClient) createCertificate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/certificates", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateCertificateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateCertificateAuthority Creates a new certificate authority (CA) according to the details of the request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/CreateCertificateAuthority.go.html to see an example of how to use CreateCertificateAuthority API.
func (client CertificatesManagementClient) CreateCertificateAuthority(ctx context.Context, request CreateCertificateAuthorityRequest) (response CreateCertificateAuthorityResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createCertificateAuthority, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateCertificateAuthorityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateCertificateAuthorityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateCertificateAuthorityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateCertificateAuthorityResponse")
	}
	return
}

// createCertificateAuthority implements the OCIOperation interface (enables retrying operations)
func (client CertificatesManagementClient) createCertificateAuthority(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/certificateAuthorities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateCertificateAuthorityResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteCaBundle Deletes the specified CA bundle.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/DeleteCaBundle.go.html to see an example of how to use DeleteCaBundle API.
func (client CertificatesManagementClient) DeleteCaBundle(ctx context.Context, request DeleteCaBundleRequest) (response DeleteCaBundleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteCaBundle, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteCaBundleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteCaBundleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteCaBundleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteCaBundleResponse")
	}
	return
}

// deleteCaBundle implements the OCIOperation interface (enables retrying operations)
func (client CertificatesManagementClient) deleteCaBundle(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/caBundles/{caBundleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteCaBundleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAssociation Gets details about the specified association.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/GetAssociation.go.html to see an example of how to use GetAssociation API.
func (client CertificatesManagementClient) GetAssociation(ctx context.Context, request GetAssociationRequest) (response GetAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAssociationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAssociationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAssociationResponse")
	}
	return
}

// getAssociation implements the OCIOperation interface (enables retrying operations)
func (client CertificatesManagementClient) getAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/associations/{associationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAssociationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCaBundle Gets details about the specified CA bundle.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/GetCaBundle.go.html to see an example of how to use GetCaBundle API.
func (client CertificatesManagementClient) GetCaBundle(ctx context.Context, request GetCaBundleRequest) (response GetCaBundleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCaBundle, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCaBundleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCaBundleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCaBundleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCaBundleResponse")
	}
	return
}

// getCaBundle implements the OCIOperation interface (enables retrying operations)
func (client CertificatesManagementClient) getCaBundle(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/caBundles/{caBundleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCaBundleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCertificate Gets details about the specified certificate.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/GetCertificate.go.html to see an example of how to use GetCertificate API.
func (client CertificatesManagementClient) GetCertificate(ctx context.Context, request GetCertificateRequest) (response GetCertificateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client CertificatesManagementClient) getCertificate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/certificates/{certificateId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCertificateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCertificateAuthority Gets details about the specified certificate authority (CA).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/GetCertificateAuthority.go.html to see an example of how to use GetCertificateAuthority API.
func (client CertificatesManagementClient) GetCertificateAuthority(ctx context.Context, request GetCertificateAuthorityRequest) (response GetCertificateAuthorityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCertificateAuthority, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCertificateAuthorityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCertificateAuthorityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCertificateAuthorityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCertificateAuthorityResponse")
	}
	return
}

// getCertificateAuthority implements the OCIOperation interface (enables retrying operations)
func (client CertificatesManagementClient) getCertificateAuthority(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/certificateAuthorities/{certificateAuthorityId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCertificateAuthorityResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCertificateAuthorityVersion Gets details about the specified certificate authority (CA) version.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/GetCertificateAuthorityVersion.go.html to see an example of how to use GetCertificateAuthorityVersion API.
func (client CertificatesManagementClient) GetCertificateAuthorityVersion(ctx context.Context, request GetCertificateAuthorityVersionRequest) (response GetCertificateAuthorityVersionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCertificateAuthorityVersion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCertificateAuthorityVersionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCertificateAuthorityVersionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCertificateAuthorityVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCertificateAuthorityVersionResponse")
	}
	return
}

// getCertificateAuthorityVersion implements the OCIOperation interface (enables retrying operations)
func (client CertificatesManagementClient) getCertificateAuthorityVersion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/certificateAuthorities/{certificateAuthorityId}/version/{certificateAuthorityVersionNumber}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCertificateAuthorityVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetCertificateVersion Gets details about the specified version of a certificate.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/GetCertificateVersion.go.html to see an example of how to use GetCertificateVersion API.
func (client CertificatesManagementClient) GetCertificateVersion(ctx context.Context, request GetCertificateVersionRequest) (response GetCertificateVersionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCertificateVersion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCertificateVersionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCertificateVersionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCertificateVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCertificateVersionResponse")
	}
	return
}

// getCertificateVersion implements the OCIOperation interface (enables retrying operations)
func (client CertificatesManagementClient) getCertificateVersion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/certificates/{certificateId}/version/{certificateVersionNumber}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCertificateVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAssociations Lists all associations that match the query parameters.
// Optionally, you can use the parameter `FilterByAssociationIdQueryParam` to limit the result set to a single item that matches the specified association.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/ListAssociations.go.html to see an example of how to use ListAssociations API.
func (client CertificatesManagementClient) ListAssociations(ctx context.Context, request ListAssociationsRequest) (response ListAssociationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAssociations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAssociationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAssociationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAssociationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAssociationsResponse")
	}
	return
}

// listAssociations implements the OCIOperation interface (enables retrying operations)
func (client CertificatesManagementClient) listAssociations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/associations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAssociationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCaBundles Lists all CA bundles that match the query parameters.
// Optionally, you can use the parameter `FilterByCaBundleIdQueryParam` to limit the result set to a single item that matches the specified CA bundle.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/ListCaBundles.go.html to see an example of how to use ListCaBundles API.
func (client CertificatesManagementClient) ListCaBundles(ctx context.Context, request ListCaBundlesRequest) (response ListCaBundlesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCaBundles, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCaBundlesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCaBundlesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCaBundlesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCaBundlesResponse")
	}
	return
}

// listCaBundles implements the OCIOperation interface (enables retrying operations)
func (client CertificatesManagementClient) listCaBundles(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/caBundles", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCaBundlesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCertificateAuthorities Lists all certificate authorities (CAs) in the specified compartment.
// Optionally, you can use the parameter `FilterByCertificateAuthorityIdQueryParam` to limit the results to a single item that matches the specified CA.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/ListCertificateAuthorities.go.html to see an example of how to use ListCertificateAuthorities API.
func (client CertificatesManagementClient) ListCertificateAuthorities(ctx context.Context, request ListCertificateAuthoritiesRequest) (response ListCertificateAuthoritiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCertificateAuthorities, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCertificateAuthoritiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCertificateAuthoritiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCertificateAuthoritiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCertificateAuthoritiesResponse")
	}
	return
}

// listCertificateAuthorities implements the OCIOperation interface (enables retrying operations)
func (client CertificatesManagementClient) listCertificateAuthorities(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/certificateAuthorities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCertificateAuthoritiesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCertificateAuthorityVersions Lists all versions for the specified certificate authority (CA).
// Optionally, you can use the parameter `FilterByVersionNumberQueryParam` to limit the results to a single item that matches the specified version number.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/ListCertificateAuthorityVersions.go.html to see an example of how to use ListCertificateAuthorityVersions API.
func (client CertificatesManagementClient) ListCertificateAuthorityVersions(ctx context.Context, request ListCertificateAuthorityVersionsRequest) (response ListCertificateAuthorityVersionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCertificateAuthorityVersions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCertificateAuthorityVersionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCertificateAuthorityVersionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCertificateAuthorityVersionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCertificateAuthorityVersionsResponse")
	}
	return
}

// listCertificateAuthorityVersions implements the OCIOperation interface (enables retrying operations)
func (client CertificatesManagementClient) listCertificateAuthorityVersions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/certificateAuthorities/{certificateAuthorityId}/versions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCertificateAuthorityVersionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCertificateVersions Lists all certificate versions for the specified certificate.
// Optionally, you can use the parameter `FilterByVersionNumberQueryParam` to limit the result set to a single item that matches the specified version number.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/ListCertificateVersions.go.html to see an example of how to use ListCertificateVersions API.
func (client CertificatesManagementClient) ListCertificateVersions(ctx context.Context, request ListCertificateVersionsRequest) (response ListCertificateVersionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCertificateVersions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCertificateVersionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCertificateVersionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCertificateVersionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCertificateVersionsResponse")
	}
	return
}

// listCertificateVersions implements the OCIOperation interface (enables retrying operations)
func (client CertificatesManagementClient) listCertificateVersions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/certificates/{certificateId}/versions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCertificateVersionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCertificates Lists all certificates that match the query parameters.
// Optionally, you can use the parameter `FilterByCertificateIdQueryParam` to limit the result set to a single item that matches the specified certificate.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/ListCertificates.go.html to see an example of how to use ListCertificates API.
func (client CertificatesManagementClient) ListCertificates(ctx context.Context, request ListCertificatesRequest) (response ListCertificatesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client CertificatesManagementClient) listCertificates(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/certificates", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCertificatesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RevokeCertificateAuthorityVersion Revokes a certificate authority (CA) version.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/RevokeCertificateAuthorityVersion.go.html to see an example of how to use RevokeCertificateAuthorityVersion API.
func (client CertificatesManagementClient) RevokeCertificateAuthorityVersion(ctx context.Context, request RevokeCertificateAuthorityVersionRequest) (response RevokeCertificateAuthorityVersionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.revokeCertificateAuthorityVersion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RevokeCertificateAuthorityVersionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RevokeCertificateAuthorityVersionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RevokeCertificateAuthorityVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RevokeCertificateAuthorityVersionResponse")
	}
	return
}

// revokeCertificateAuthorityVersion implements the OCIOperation interface (enables retrying operations)
func (client CertificatesManagementClient) revokeCertificateAuthorityVersion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/certificateAuthorities/{certificateAuthorityId}/version/{certificateAuthorityVersionNumber}/actions/revoke", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RevokeCertificateAuthorityVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RevokeCertificateVersion Revokes the specified certificate version.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/RevokeCertificateVersion.go.html to see an example of how to use RevokeCertificateVersion API.
func (client CertificatesManagementClient) RevokeCertificateVersion(ctx context.Context, request RevokeCertificateVersionRequest) (response RevokeCertificateVersionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.revokeCertificateVersion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RevokeCertificateVersionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RevokeCertificateVersionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RevokeCertificateVersionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RevokeCertificateVersionResponse")
	}
	return
}

// revokeCertificateVersion implements the OCIOperation interface (enables retrying operations)
func (client CertificatesManagementClient) revokeCertificateVersion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/certificates/{certificateId}/version/{certificateVersionNumber}/actions/revoke", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RevokeCertificateVersionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ScheduleCertificateAuthorityDeletion Schedules the deletion of the specified certificate authority (CA). This sets the lifecycle state of the CA to `PENDING_DELETION` and then deletes it after the specified retention period ends. If needed, you can determine the status of the deletion by using `GetCertificateAuthority`.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/ScheduleCertificateAuthorityDeletion.go.html to see an example of how to use ScheduleCertificateAuthorityDeletion API.
func (client CertificatesManagementClient) ScheduleCertificateAuthorityDeletion(ctx context.Context, request ScheduleCertificateAuthorityDeletionRequest) (response ScheduleCertificateAuthorityDeletionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.scheduleCertificateAuthorityDeletion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ScheduleCertificateAuthorityDeletionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ScheduleCertificateAuthorityDeletionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ScheduleCertificateAuthorityDeletionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ScheduleCertificateAuthorityDeletionResponse")
	}
	return
}

// scheduleCertificateAuthorityDeletion implements the OCIOperation interface (enables retrying operations)
func (client CertificatesManagementClient) scheduleCertificateAuthorityDeletion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/certificateAuthorities/{certificateAuthorityId}/actions/scheduleDeletion", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ScheduleCertificateAuthorityDeletionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ScheduleCertificateAuthorityVersionDeletion Schedules the deletion of the specified certificate authority (CA) version.
// This sets the lifecycle state of the CA version to `PENDING_DELETION`
// and then deletes it after the specified retention period ends. If needed, you can determine the status of the deletion by using `GetCertificateAuthorityVersion`.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/ScheduleCertificateAuthorityVersionDeletion.go.html to see an example of how to use ScheduleCertificateAuthorityVersionDeletion API.
func (client CertificatesManagementClient) ScheduleCertificateAuthorityVersionDeletion(ctx context.Context, request ScheduleCertificateAuthorityVersionDeletionRequest) (response ScheduleCertificateAuthorityVersionDeletionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.scheduleCertificateAuthorityVersionDeletion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ScheduleCertificateAuthorityVersionDeletionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ScheduleCertificateAuthorityVersionDeletionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ScheduleCertificateAuthorityVersionDeletionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ScheduleCertificateAuthorityVersionDeletionResponse")
	}
	return
}

// scheduleCertificateAuthorityVersionDeletion implements the OCIOperation interface (enables retrying operations)
func (client CertificatesManagementClient) scheduleCertificateAuthorityVersionDeletion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/certificateAuthorities/{certificateAuthorityId}/version/{certificateAuthorityVersionNumber}/actions/scheduleDeletion", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ScheduleCertificateAuthorityVersionDeletionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ScheduleCertificateDeletion Schedules the deletion of the specified certificate. This sets the lifecycle state of the certificate
// to `PENDING_DELETION` and then deletes it after the specified retention period ends.
// You can subsequently use `GetCertificate` to determine the current deletion status.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/ScheduleCertificateDeletion.go.html to see an example of how to use ScheduleCertificateDeletion API.
func (client CertificatesManagementClient) ScheduleCertificateDeletion(ctx context.Context, request ScheduleCertificateDeletionRequest) (response ScheduleCertificateDeletionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.scheduleCertificateDeletion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ScheduleCertificateDeletionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ScheduleCertificateDeletionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ScheduleCertificateDeletionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ScheduleCertificateDeletionResponse")
	}
	return
}

// scheduleCertificateDeletion implements the OCIOperation interface (enables retrying operations)
func (client CertificatesManagementClient) scheduleCertificateDeletion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/certificates/{certificateId}/actions/scheduleDeletion", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ScheduleCertificateDeletionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ScheduleCertificateVersionDeletion Schedules the deletion of the specified certificate version. This sets the lifecycle state of the certificate version to `PENDING_DELETION` and then deletes it after the specified retention period ends. You can only
// delete a certificate version if the certificate version rotation state is marked as `DEPRECATED`.
// You can subsequently use `GetCertificateVersion` to determine the current certificate version deletion status.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/ScheduleCertificateVersionDeletion.go.html to see an example of how to use ScheduleCertificateVersionDeletion API.
func (client CertificatesManagementClient) ScheduleCertificateVersionDeletion(ctx context.Context, request ScheduleCertificateVersionDeletionRequest) (response ScheduleCertificateVersionDeletionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.scheduleCertificateVersionDeletion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ScheduleCertificateVersionDeletionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ScheduleCertificateVersionDeletionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ScheduleCertificateVersionDeletionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ScheduleCertificateVersionDeletionResponse")
	}
	return
}

// scheduleCertificateVersionDeletion implements the OCIOperation interface (enables retrying operations)
func (client CertificatesManagementClient) scheduleCertificateVersionDeletion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/certificates/{certificateId}/version/{certificateVersionNumber}/actions/scheduleDeletion", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ScheduleCertificateVersionDeletionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateCaBundle Updates the properties of a CA bundle.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/UpdateCaBundle.go.html to see an example of how to use UpdateCaBundle API.
func (client CertificatesManagementClient) UpdateCaBundle(ctx context.Context, request UpdateCaBundleRequest) (response UpdateCaBundleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateCaBundle, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateCaBundleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateCaBundleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateCaBundleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateCaBundleResponse")
	}
	return
}

// updateCaBundle implements the OCIOperation interface (enables retrying operations)
func (client CertificatesManagementClient) updateCaBundle(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/caBundles/{caBundleId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateCaBundleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateCertificate Updates the properties of a certificate.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/UpdateCertificate.go.html to see an example of how to use UpdateCertificate API.
func (client CertificatesManagementClient) UpdateCertificate(ctx context.Context, request UpdateCertificateRequest) (response UpdateCertificateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateCertificate, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateCertificateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateCertificateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateCertificateResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateCertificateResponse")
	}
	return
}

// updateCertificate implements the OCIOperation interface (enables retrying operations)
func (client CertificatesManagementClient) updateCertificate(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/certificates/{certificateId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateCertificateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateCertificateAuthority Updates the properties of the specified certificate authority (CA).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/UpdateCertificateAuthority.go.html to see an example of how to use UpdateCertificateAuthority API.
func (client CertificatesManagementClient) UpdateCertificateAuthority(ctx context.Context, request UpdateCertificateAuthorityRequest) (response UpdateCertificateAuthorityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateCertificateAuthority, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateCertificateAuthorityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateCertificateAuthorityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateCertificateAuthorityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateCertificateAuthorityResponse")
	}
	return
}

// updateCertificateAuthority implements the OCIOperation interface (enables retrying operations)
func (client CertificatesManagementClient) updateCertificateAuthority(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/certificateAuthorities/{certificateAuthorityId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateCertificateAuthorityResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
