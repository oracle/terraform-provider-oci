// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Catalog API
//
// Manage solutions in Oracle Cloud Infrastructure Service Catalog.
//

package servicecatalog

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// ServiceCatalogClient a client for ServiceCatalog
type ServiceCatalogClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewServiceCatalogClientWithConfigurationProvider Creates a new default ServiceCatalog client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewServiceCatalogClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ServiceCatalogClient, err error) {
	if enabled := common.CheckForEnabledServices("servicecatalog"); !enabled {
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
	return newServiceCatalogClientFromBaseClient(baseClient, provider)
}

// NewServiceCatalogClientWithOboToken Creates a new default ServiceCatalog client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewServiceCatalogClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ServiceCatalogClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newServiceCatalogClientFromBaseClient(baseClient, configProvider)
}

func newServiceCatalogClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ServiceCatalogClient, err error) {
	// ServiceCatalog service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("ServiceCatalog"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ServiceCatalogClient{BaseClient: baseClient}
	client.BasePath = "20210527"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ServiceCatalogClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("servicecatalog", "https://service-catalog.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ServiceCatalogClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ServiceCatalogClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// BulkReplaceServiceCatalogAssociations Replace all associations of a given service catalog in one bulk transaction.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicecatalog/BulkReplaceServiceCatalogAssociations.go.html to see an example of how to use BulkReplaceServiceCatalogAssociations API.
func (client ServiceCatalogClient) BulkReplaceServiceCatalogAssociations(ctx context.Context, request BulkReplaceServiceCatalogAssociationsRequest) (response BulkReplaceServiceCatalogAssociationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.bulkReplaceServiceCatalogAssociations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BulkReplaceServiceCatalogAssociationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BulkReplaceServiceCatalogAssociationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BulkReplaceServiceCatalogAssociationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BulkReplaceServiceCatalogAssociationsResponse")
	}
	return
}

// bulkReplaceServiceCatalogAssociations implements the OCIOperation interface (enables retrying operations)
func (client ServiceCatalogClient) bulkReplaceServiceCatalogAssociations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/serviceCatalogs/{serviceCatalogId}/actions/bulkReplaceAssociations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BulkReplaceServiceCatalogAssociationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-catalog/20210527/ServiceCatalogAssociation/BulkReplaceServiceCatalogAssociations"
		err = common.PostProcessServiceError(err, "ServiceCatalog", "BulkReplaceServiceCatalogAssociations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangePrivateApplicationCompartment Moves the specified private application from one compartment to another.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicecatalog/ChangePrivateApplicationCompartment.go.html to see an example of how to use ChangePrivateApplicationCompartment API.
func (client ServiceCatalogClient) ChangePrivateApplicationCompartment(ctx context.Context, request ChangePrivateApplicationCompartmentRequest) (response ChangePrivateApplicationCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changePrivateApplicationCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangePrivateApplicationCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangePrivateApplicationCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangePrivateApplicationCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangePrivateApplicationCompartmentResponse")
	}
	return
}

// changePrivateApplicationCompartment implements the OCIOperation interface (enables retrying operations)
func (client ServiceCatalogClient) changePrivateApplicationCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/privateApplications/{privateApplicationId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangePrivateApplicationCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-catalog/20210527/PrivateApplication/ChangePrivateApplicationCompartment"
		err = common.PostProcessServiceError(err, "ServiceCatalog", "ChangePrivateApplicationCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeServiceCatalogCompartment Moves the specified service catalog from one compartment to another.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicecatalog/ChangeServiceCatalogCompartment.go.html to see an example of how to use ChangeServiceCatalogCompartment API.
func (client ServiceCatalogClient) ChangeServiceCatalogCompartment(ctx context.Context, request ChangeServiceCatalogCompartmentRequest) (response ChangeServiceCatalogCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeServiceCatalogCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeServiceCatalogCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeServiceCatalogCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeServiceCatalogCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeServiceCatalogCompartmentResponse")
	}
	return
}

// changeServiceCatalogCompartment implements the OCIOperation interface (enables retrying operations)
func (client ServiceCatalogClient) changeServiceCatalogCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/serviceCatalogs/{serviceCatalogId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeServiceCatalogCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-catalog/20210527/ServiceCatalog/ChangeServiceCatalogCompartment"
		err = common.PostProcessServiceError(err, "ServiceCatalog", "ChangeServiceCatalogCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreatePrivateApplication Creates a private application along with a single package to be hosted.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicecatalog/CreatePrivateApplication.go.html to see an example of how to use CreatePrivateApplication API.
func (client ServiceCatalogClient) CreatePrivateApplication(ctx context.Context, request CreatePrivateApplicationRequest) (response CreatePrivateApplicationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createPrivateApplication, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreatePrivateApplicationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreatePrivateApplicationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreatePrivateApplicationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreatePrivateApplicationResponse")
	}
	return
}

// createPrivateApplication implements the OCIOperation interface (enables retrying operations)
func (client ServiceCatalogClient) createPrivateApplication(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/privateApplications", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreatePrivateApplicationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-catalog/20210527/PrivateApplication/CreatePrivateApplication"
		err = common.PostProcessServiceError(err, "ServiceCatalog", "CreatePrivateApplication", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateServiceCatalog Creates a brand new service catalog in a given compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicecatalog/CreateServiceCatalog.go.html to see an example of how to use CreateServiceCatalog API.
func (client ServiceCatalogClient) CreateServiceCatalog(ctx context.Context, request CreateServiceCatalogRequest) (response CreateServiceCatalogResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createServiceCatalog, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateServiceCatalogResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateServiceCatalogResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateServiceCatalogResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateServiceCatalogResponse")
	}
	return
}

// createServiceCatalog implements the OCIOperation interface (enables retrying operations)
func (client ServiceCatalogClient) createServiceCatalog(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/serviceCatalogs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateServiceCatalogResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-catalog/20210527/ServiceCatalog/CreateServiceCatalog"
		err = common.PostProcessServiceError(err, "ServiceCatalog", "CreateServiceCatalog", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateServiceCatalogAssociation Creates an association between service catalog and a resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicecatalog/CreateServiceCatalogAssociation.go.html to see an example of how to use CreateServiceCatalogAssociation API.
func (client ServiceCatalogClient) CreateServiceCatalogAssociation(ctx context.Context, request CreateServiceCatalogAssociationRequest) (response CreateServiceCatalogAssociationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createServiceCatalogAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateServiceCatalogAssociationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateServiceCatalogAssociationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateServiceCatalogAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateServiceCatalogAssociationResponse")
	}
	return
}

// createServiceCatalogAssociation implements the OCIOperation interface (enables retrying operations)
func (client ServiceCatalogClient) createServiceCatalogAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/serviceCatalogAssociations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateServiceCatalogAssociationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-catalog/20210527/ServiceCatalogAssociation/CreateServiceCatalogAssociation"
		err = common.PostProcessServiceError(err, "ServiceCatalog", "CreateServiceCatalogAssociation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeletePrivateApplication Deletes an existing private application.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicecatalog/DeletePrivateApplication.go.html to see an example of how to use DeletePrivateApplication API.
func (client ServiceCatalogClient) DeletePrivateApplication(ctx context.Context, request DeletePrivateApplicationRequest) (response DeletePrivateApplicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deletePrivateApplication, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeletePrivateApplicationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeletePrivateApplicationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeletePrivateApplicationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeletePrivateApplicationResponse")
	}
	return
}

// deletePrivateApplication implements the OCIOperation interface (enables retrying operations)
func (client ServiceCatalogClient) deletePrivateApplication(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/privateApplications/{privateApplicationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeletePrivateApplicationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-catalog/20210527/PrivateApplication/DeletePrivateApplication"
		err = common.PostProcessServiceError(err, "ServiceCatalog", "DeletePrivateApplication", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteServiceCatalog Deletes the specified service catalog from the compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicecatalog/DeleteServiceCatalog.go.html to see an example of how to use DeleteServiceCatalog API.
func (client ServiceCatalogClient) DeleteServiceCatalog(ctx context.Context, request DeleteServiceCatalogRequest) (response DeleteServiceCatalogResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteServiceCatalog, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteServiceCatalogResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteServiceCatalogResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteServiceCatalogResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteServiceCatalogResponse")
	}
	return
}

// deleteServiceCatalog implements the OCIOperation interface (enables retrying operations)
func (client ServiceCatalogClient) deleteServiceCatalog(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/serviceCatalogs/{serviceCatalogId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteServiceCatalogResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-catalog/20210527/ServiceCatalog/DeleteServiceCatalog"
		err = common.PostProcessServiceError(err, "ServiceCatalog", "DeleteServiceCatalog", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteServiceCatalogAssociation Removes an association between service catalog and a resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicecatalog/DeleteServiceCatalogAssociation.go.html to see an example of how to use DeleteServiceCatalogAssociation API.
func (client ServiceCatalogClient) DeleteServiceCatalogAssociation(ctx context.Context, request DeleteServiceCatalogAssociationRequest) (response DeleteServiceCatalogAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteServiceCatalogAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteServiceCatalogAssociationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteServiceCatalogAssociationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteServiceCatalogAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteServiceCatalogAssociationResponse")
	}
	return
}

// deleteServiceCatalogAssociation implements the OCIOperation interface (enables retrying operations)
func (client ServiceCatalogClient) deleteServiceCatalogAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/serviceCatalogAssociations/{serviceCatalogAssociationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteServiceCatalogAssociationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-catalog/20210527/ServiceCatalogAssociation/DeleteServiceCatalogAssociation"
		err = common.PostProcessServiceError(err, "ServiceCatalog", "DeleteServiceCatalogAssociation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPrivateApplication Gets the details of the specified private application.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicecatalog/GetPrivateApplication.go.html to see an example of how to use GetPrivateApplication API.
func (client ServiceCatalogClient) GetPrivateApplication(ctx context.Context, request GetPrivateApplicationRequest) (response GetPrivateApplicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPrivateApplication, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPrivateApplicationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPrivateApplicationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPrivateApplicationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPrivateApplicationResponse")
	}
	return
}

// getPrivateApplication implements the OCIOperation interface (enables retrying operations)
func (client ServiceCatalogClient) getPrivateApplication(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/privateApplications/{privateApplicationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPrivateApplicationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-catalog/20210527/PrivateApplication/GetPrivateApplication"
		err = common.PostProcessServiceError(err, "ServiceCatalog", "GetPrivateApplication", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPrivateApplicationActionDownloadLogo Downloads the binary payload of the logo image of the private application.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicecatalog/GetPrivateApplicationActionDownloadLogo.go.html to see an example of how to use GetPrivateApplicationActionDownloadLogo API.
func (client ServiceCatalogClient) GetPrivateApplicationActionDownloadLogo(ctx context.Context, request GetPrivateApplicationActionDownloadLogoRequest) (response GetPrivateApplicationActionDownloadLogoResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPrivateApplicationActionDownloadLogo, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPrivateApplicationActionDownloadLogoResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPrivateApplicationActionDownloadLogoResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPrivateApplicationActionDownloadLogoResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPrivateApplicationActionDownloadLogoResponse")
	}
	return
}

// getPrivateApplicationActionDownloadLogo implements the OCIOperation interface (enables retrying operations)
func (client ServiceCatalogClient) getPrivateApplicationActionDownloadLogo(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/privateApplications/{privateApplicationId}/actions/downloadLogo", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPrivateApplicationActionDownloadLogoResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-catalog/20210527/PrivateApplication/GetPrivateApplicationActionDownloadLogo"
		err = common.PostProcessServiceError(err, "ServiceCatalog", "GetPrivateApplicationActionDownloadLogo", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPrivateApplicationPackage Gets the details of a specific package within a given private application.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicecatalog/GetPrivateApplicationPackage.go.html to see an example of how to use GetPrivateApplicationPackage API.
func (client ServiceCatalogClient) GetPrivateApplicationPackage(ctx context.Context, request GetPrivateApplicationPackageRequest) (response GetPrivateApplicationPackageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPrivateApplicationPackage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPrivateApplicationPackageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPrivateApplicationPackageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPrivateApplicationPackageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPrivateApplicationPackageResponse")
	}
	return
}

// getPrivateApplicationPackage implements the OCIOperation interface (enables retrying operations)
func (client ServiceCatalogClient) getPrivateApplicationPackage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/privateApplicationPackages/{privateApplicationPackageId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPrivateApplicationPackageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-catalog/20210527/PrivateApplicationPackage/GetPrivateApplicationPackage"
		err = common.PostProcessServiceError(err, "ServiceCatalog", "GetPrivateApplicationPackage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &privateapplicationpackage{})
	return response, err
}

// GetPrivateApplicationPackageActionDownloadConfig Downloads the configuration that was used to create the private application package.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicecatalog/GetPrivateApplicationPackageActionDownloadConfig.go.html to see an example of how to use GetPrivateApplicationPackageActionDownloadConfig API.
func (client ServiceCatalogClient) GetPrivateApplicationPackageActionDownloadConfig(ctx context.Context, request GetPrivateApplicationPackageActionDownloadConfigRequest) (response GetPrivateApplicationPackageActionDownloadConfigResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPrivateApplicationPackageActionDownloadConfig, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPrivateApplicationPackageActionDownloadConfigResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPrivateApplicationPackageActionDownloadConfigResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPrivateApplicationPackageActionDownloadConfigResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPrivateApplicationPackageActionDownloadConfigResponse")
	}
	return
}

// getPrivateApplicationPackageActionDownloadConfig implements the OCIOperation interface (enables retrying operations)
func (client ServiceCatalogClient) getPrivateApplicationPackageActionDownloadConfig(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/privateApplicationPackages/{privateApplicationPackageId}/actions/downloadConfig", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPrivateApplicationPackageActionDownloadConfigResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-catalog/20210527/PrivateApplicationPackage/GetPrivateApplicationPackageActionDownloadConfig"
		err = common.PostProcessServiceError(err, "ServiceCatalog", "GetPrivateApplicationPackageActionDownloadConfig", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetServiceCatalog Gets detailed information about the service catalog including name, compartmentId
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicecatalog/GetServiceCatalog.go.html to see an example of how to use GetServiceCatalog API.
func (client ServiceCatalogClient) GetServiceCatalog(ctx context.Context, request GetServiceCatalogRequest) (response GetServiceCatalogResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getServiceCatalog, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetServiceCatalogResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetServiceCatalogResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetServiceCatalogResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetServiceCatalogResponse")
	}
	return
}

// getServiceCatalog implements the OCIOperation interface (enables retrying operations)
func (client ServiceCatalogClient) getServiceCatalog(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/serviceCatalogs/{serviceCatalogId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetServiceCatalogResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-catalog/20210527/ServiceCatalog/GetServiceCatalog"
		err = common.PostProcessServiceError(err, "ServiceCatalog", "GetServiceCatalog", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetServiceCatalogAssociation Gets detailed information about specific service catalog association.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicecatalog/GetServiceCatalogAssociation.go.html to see an example of how to use GetServiceCatalogAssociation API.
func (client ServiceCatalogClient) GetServiceCatalogAssociation(ctx context.Context, request GetServiceCatalogAssociationRequest) (response GetServiceCatalogAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getServiceCatalogAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetServiceCatalogAssociationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetServiceCatalogAssociationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetServiceCatalogAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetServiceCatalogAssociationResponse")
	}
	return
}

// getServiceCatalogAssociation implements the OCIOperation interface (enables retrying operations)
func (client ServiceCatalogClient) getServiceCatalogAssociation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/serviceCatalogAssociations/{serviceCatalogAssociationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetServiceCatalogAssociationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-catalog/20210527/ServiceCatalogAssociation/GetServiceCatalogAssociation"
		err = common.PostProcessServiceError(err, "ServiceCatalog", "GetServiceCatalogAssociation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets the status of the work request with the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicecatalog/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
func (client ServiceCatalogClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
func (client ServiceCatalogClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-catalog/20210527/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "ServiceCatalog", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListApplications Lists all the applications in a service catalog or a tenancy.
// If no parameter is specified, all catalogs from all compartments in
// the tenancy will be scanned for any type of content.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicecatalog/ListApplications.go.html to see an example of how to use ListApplications API.
func (client ServiceCatalogClient) ListApplications(ctx context.Context, request ListApplicationsRequest) (response ListApplicationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listApplications, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListApplicationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListApplicationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListApplicationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListApplicationsResponse")
	}
	return
}

// listApplications implements the OCIOperation interface (enables retrying operations)
func (client ServiceCatalogClient) listApplications(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/applications", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListApplicationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-catalog/20210527/ApplicationSummary/ListApplications"
		err = common.PostProcessServiceError(err, "ServiceCatalog", "ListApplications", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPrivateApplicationPackages Lists the packages in the specified private application.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicecatalog/ListPrivateApplicationPackages.go.html to see an example of how to use ListPrivateApplicationPackages API.
func (client ServiceCatalogClient) ListPrivateApplicationPackages(ctx context.Context, request ListPrivateApplicationPackagesRequest) (response ListPrivateApplicationPackagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPrivateApplicationPackages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPrivateApplicationPackagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPrivateApplicationPackagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPrivateApplicationPackagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPrivateApplicationPackagesResponse")
	}
	return
}

// listPrivateApplicationPackages implements the OCIOperation interface (enables retrying operations)
func (client ServiceCatalogClient) listPrivateApplicationPackages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/privateApplicationPackages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPrivateApplicationPackagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-catalog/20210527/PrivateApplicationPackage/ListPrivateApplicationPackages"
		err = common.PostProcessServiceError(err, "ServiceCatalog", "ListPrivateApplicationPackages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPrivateApplications Lists all the private applications in a given compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicecatalog/ListPrivateApplications.go.html to see an example of how to use ListPrivateApplications API.
func (client ServiceCatalogClient) ListPrivateApplications(ctx context.Context, request ListPrivateApplicationsRequest) (response ListPrivateApplicationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPrivateApplications, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPrivateApplicationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPrivateApplicationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPrivateApplicationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPrivateApplicationsResponse")
	}
	return
}

// listPrivateApplications implements the OCIOperation interface (enables retrying operations)
func (client ServiceCatalogClient) listPrivateApplications(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/privateApplications", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPrivateApplicationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-catalog/20210527/PrivateApplication/ListPrivateApplications"
		err = common.PostProcessServiceError(err, "ServiceCatalog", "ListPrivateApplications", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListServiceCatalogAssociations Lists all the resource associations for a specific service catalog.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicecatalog/ListServiceCatalogAssociations.go.html to see an example of how to use ListServiceCatalogAssociations API.
func (client ServiceCatalogClient) ListServiceCatalogAssociations(ctx context.Context, request ListServiceCatalogAssociationsRequest) (response ListServiceCatalogAssociationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listServiceCatalogAssociations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListServiceCatalogAssociationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListServiceCatalogAssociationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListServiceCatalogAssociationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListServiceCatalogAssociationsResponse")
	}
	return
}

// listServiceCatalogAssociations implements the OCIOperation interface (enables retrying operations)
func (client ServiceCatalogClient) listServiceCatalogAssociations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/serviceCatalogAssociations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListServiceCatalogAssociationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-catalog/20210527/ServiceCatalogAssociation/ListServiceCatalogAssociations"
		err = common.PostProcessServiceError(err, "ServiceCatalog", "ListServiceCatalogAssociations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListServiceCatalogs Lists all the service catalogs in the given compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicecatalog/ListServiceCatalogs.go.html to see an example of how to use ListServiceCatalogs API.
func (client ServiceCatalogClient) ListServiceCatalogs(ctx context.Context, request ListServiceCatalogsRequest) (response ListServiceCatalogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listServiceCatalogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListServiceCatalogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListServiceCatalogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListServiceCatalogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListServiceCatalogsResponse")
	}
	return
}

// listServiceCatalogs implements the OCIOperation interface (enables retrying operations)
func (client ServiceCatalogClient) listServiceCatalogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/serviceCatalogs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListServiceCatalogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-catalog/20210527/ServiceCatalog/ListServiceCatalogs"
		err = common.PostProcessServiceError(err, "ServiceCatalog", "ListServiceCatalogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Return a (paginated) list of errors for a given work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicecatalog/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
func (client ServiceCatalogClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
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
func (client ServiceCatalogClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-catalog/20210527/WorkRequestError/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "ServiceCatalog", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Return a (paginated) list of logs for a given work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicecatalog/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
func (client ServiceCatalogClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
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
func (client ServiceCatalogClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-catalog/20210527/WorkRequestLogEntry/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "ServiceCatalog", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicecatalog/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
func (client ServiceCatalogClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
func (client ServiceCatalogClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-catalog/20210527/WorkRequest/ListWorkRequests"
		err = common.PostProcessServiceError(err, "ServiceCatalog", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdatePrivateApplication Updates the details of an existing private application.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicecatalog/UpdatePrivateApplication.go.html to see an example of how to use UpdatePrivateApplication API.
func (client ServiceCatalogClient) UpdatePrivateApplication(ctx context.Context, request UpdatePrivateApplicationRequest) (response UpdatePrivateApplicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updatePrivateApplication, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdatePrivateApplicationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdatePrivateApplicationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdatePrivateApplicationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdatePrivateApplicationResponse")
	}
	return
}

// updatePrivateApplication implements the OCIOperation interface (enables retrying operations)
func (client ServiceCatalogClient) updatePrivateApplication(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/privateApplications/{privateApplicationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdatePrivateApplicationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-catalog/20210527/PrivateApplication/UpdatePrivateApplication"
		err = common.PostProcessServiceError(err, "ServiceCatalog", "UpdatePrivateApplication", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateServiceCatalog Updates the details of a previously created service catalog.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicecatalog/UpdateServiceCatalog.go.html to see an example of how to use UpdateServiceCatalog API.
func (client ServiceCatalogClient) UpdateServiceCatalog(ctx context.Context, request UpdateServiceCatalogRequest) (response UpdateServiceCatalogResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateServiceCatalog, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateServiceCatalogResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateServiceCatalogResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateServiceCatalogResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateServiceCatalogResponse")
	}
	return
}

// updateServiceCatalog implements the OCIOperation interface (enables retrying operations)
func (client ServiceCatalogClient) updateServiceCatalog(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/serviceCatalogs/{serviceCatalogId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateServiceCatalogResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/service-catalog/20210527/ServiceCatalog/UpdateServiceCatalog"
		err = common.PostProcessServiceError(err, "ServiceCatalog", "UpdateServiceCatalog", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
