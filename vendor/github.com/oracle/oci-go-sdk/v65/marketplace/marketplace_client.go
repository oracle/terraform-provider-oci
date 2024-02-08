// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Use the Marketplace API to manage applications in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.cloud.oracle.com/Content/Marketplace/Concepts/marketoverview.htm)
//

package marketplace

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// MarketplaceClient a client for Marketplace
type MarketplaceClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewMarketplaceClientWithConfigurationProvider Creates a new default Marketplace client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewMarketplaceClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client MarketplaceClient, err error) {
	if enabled := common.CheckForEnabledServices("marketplace"); !enabled {
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
	return newMarketplaceClientFromBaseClient(baseClient, provider)
}

// NewMarketplaceClientWithOboToken Creates a new default Marketplace client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewMarketplaceClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client MarketplaceClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newMarketplaceClientFromBaseClient(baseClient, configProvider)
}

func newMarketplaceClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client MarketplaceClient, err error) {
	// Marketplace service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Marketplace"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = MarketplaceClient{BaseClient: baseClient}
	client.BasePath = "20181001"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *MarketplaceClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("marketplace", "https://marketplace.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *MarketplaceClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *MarketplaceClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangePublicationCompartment Moves the specified publication from one compartment to another.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/ChangePublicationCompartment.go.html to see an example of how to use ChangePublicationCompartment API.
// A default retry strategy applies to this operation ChangePublicationCompartment()
func (client MarketplaceClient) ChangePublicationCompartment(ctx context.Context, request ChangePublicationCompartmentRequest) (response ChangePublicationCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changePublicationCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangePublicationCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangePublicationCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangePublicationCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangePublicationCompartmentResponse")
	}
	return
}

// changePublicationCompartment implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceClient) changePublicationCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/publications/{publicationId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangePublicationCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/marketplace/20181001/Publication/ChangePublicationCompartment"
		err = common.PostProcessServiceError(err, "Marketplace", "ChangePublicationCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateAcceptedAgreement Accepts a terms of use agreement for a specific package version of a listing. You must accept all
// terms of use for a package before you can deploy the package.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/CreateAcceptedAgreement.go.html to see an example of how to use CreateAcceptedAgreement API.
// A default retry strategy applies to this operation CreateAcceptedAgreement()
func (client MarketplaceClient) CreateAcceptedAgreement(ctx context.Context, request CreateAcceptedAgreementRequest) (response CreateAcceptedAgreementResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createAcceptedAgreement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAcceptedAgreementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAcceptedAgreementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAcceptedAgreementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAcceptedAgreementResponse")
	}
	return
}

// createAcceptedAgreement implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceClient) createAcceptedAgreement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/acceptedAgreements", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateAcceptedAgreementResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/marketplace/20181001/AcceptedAgreement/CreateAcceptedAgreement"
		err = common.PostProcessServiceError(err, "Marketplace", "CreateAcceptedAgreement", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreatePublication Creates a publication of the specified listing type with an optional default package.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/CreatePublication.go.html to see an example of how to use CreatePublication API.
// A default retry strategy applies to this operation CreatePublication()
func (client MarketplaceClient) CreatePublication(ctx context.Context, request CreatePublicationRequest) (response CreatePublicationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createPublication, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreatePublicationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreatePublicationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreatePublicationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreatePublicationResponse")
	}
	return
}

// createPublication implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceClient) createPublication(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/publications", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreatePublicationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/marketplace/20181001/Publication/CreatePublication"
		err = common.PostProcessServiceError(err, "Marketplace", "CreatePublication", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteAcceptedAgreement Removes a previously accepted terms of use agreement from the list of agreements that Marketplace checks
// before initiating a deployment. Listings in Marketplace that require acceptance of the specified terms
// of use can no longer be deployed, but existing deployments aren't affected.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/DeleteAcceptedAgreement.go.html to see an example of how to use DeleteAcceptedAgreement API.
// A default retry strategy applies to this operation DeleteAcceptedAgreement()
func (client MarketplaceClient) DeleteAcceptedAgreement(ctx context.Context, request DeleteAcceptedAgreementRequest) (response DeleteAcceptedAgreementResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAcceptedAgreement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAcceptedAgreementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAcceptedAgreementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAcceptedAgreementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAcceptedAgreementResponse")
	}
	return
}

// deleteAcceptedAgreement implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceClient) deleteAcceptedAgreement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/acceptedAgreements/{acceptedAgreementId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteAcceptedAgreementResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/marketplace/20181001/AcceptedAgreement/DeleteAcceptedAgreement"
		err = common.PostProcessServiceError(err, "Marketplace", "DeleteAcceptedAgreement", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeletePublication Deletes a publication, which also removes the associated listing from anywhere it was published, such as Marketplace or Compute.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/DeletePublication.go.html to see an example of how to use DeletePublication API.
// A default retry strategy applies to this operation DeletePublication()
func (client MarketplaceClient) DeletePublication(ctx context.Context, request DeletePublicationRequest) (response DeletePublicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deletePublication, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeletePublicationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeletePublicationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeletePublicationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeletePublicationResponse")
	}
	return
}

// deletePublication implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceClient) deletePublication(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/publications/{publicationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeletePublicationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/marketplace/20181001/Publication/DeletePublication"
		err = common.PostProcessServiceError(err, "Marketplace", "DeletePublication", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ExportListing Exports container images or helm chart from marketplace to customer's registry.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/ExportListing.go.html to see an example of how to use ExportListing API.
// A default retry strategy applies to this operation ExportListing()
func (client MarketplaceClient) ExportListing(ctx context.Context, request ExportListingRequest) (response ExportListingResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.exportListing, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ExportListingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ExportListingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ExportListingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ExportListingResponse")
	}
	return
}

// exportListing implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceClient) exportListing(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/listings/{listingId}/packages/{packageVersion}/actions/export", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ExportListingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/marketplace/20181001/Listing/ExportListing"
		err = common.PostProcessServiceError(err, "Marketplace", "ExportListing", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAcceptedAgreement Gets the details of a specific, previously accepted terms of use agreement.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/GetAcceptedAgreement.go.html to see an example of how to use GetAcceptedAgreement API.
// A default retry strategy applies to this operation GetAcceptedAgreement()
func (client MarketplaceClient) GetAcceptedAgreement(ctx context.Context, request GetAcceptedAgreementRequest) (response GetAcceptedAgreementResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAcceptedAgreement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAcceptedAgreementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAcceptedAgreementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAcceptedAgreementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAcceptedAgreementResponse")
	}
	return
}

// getAcceptedAgreement implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceClient) getAcceptedAgreement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/acceptedAgreements/{acceptedAgreementId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAcceptedAgreementResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/marketplace/20181001/AcceptedAgreement/GetAcceptedAgreement"
		err = common.PostProcessServiceError(err, "Marketplace", "GetAcceptedAgreement", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAgreement Returns a terms of use agreement for a package with a time-based signature that can be used to
// accept the agreement.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/GetAgreement.go.html to see an example of how to use GetAgreement API.
// A default retry strategy applies to this operation GetAgreement()
func (client MarketplaceClient) GetAgreement(ctx context.Context, request GetAgreementRequest) (response GetAgreementResponse, err error) {
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
func (client MarketplaceClient) getAgreement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/listings/{listingId}/packages/{packageVersion}/agreements/{agreementId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAgreementResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/marketplace/20181001/Agreement/GetAgreement"
		err = common.PostProcessServiceError(err, "Marketplace", "GetAgreement", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetListing Gets detailed information about a listing, including the listing's name, version, description, and
// resources.
// If you plan to launch an instance from an image listing, you must first subscribe to the listing. When
// you launch the instance, you also need to provide the image ID of the listing resource version that you want.
// Subscribing to the listing requires you to first get a signature from the terms of use agreement for the
// listing resource version. To get the signature, issue a GetAppCatalogListingAgreements (https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/latest/AppCatalogListingResourceVersionAgreements/GetAppCatalogListingAgreements) API call.
// The AppCatalogListingResourceVersionAgreements (https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/latest/AppCatalogListingResourceVersionAgreements) object, including
// its signature, is returned in the response. With the signature for the terms of use agreement for the desired
// listing resource version, create a subscription by issuing a
// CreateAppCatalogSubscription (https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/latest/AppCatalogSubscription/CreateAppCatalogSubscription) API call.
// To get the image ID to launch an instance, issue a GetAppCatalogListingResourceVersion (https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/latest/AppCatalogListingResourceVersion/GetAppCatalogListingResourceVersion) API call.
// Lastly, to launch the instance, use the image ID of the listing resource version to issue a LaunchInstance (https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/latest/Instance/LaunchInstance) API call.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/GetListing.go.html to see an example of how to use GetListing API.
// A default retry strategy applies to this operation GetListing()
func (client MarketplaceClient) GetListing(ctx context.Context, request GetListingRequest) (response GetListingResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getListing, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetListingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetListingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetListingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetListingResponse")
	}
	return
}

// getListing implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceClient) getListing(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/listings/{listingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetListingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/marketplace/20181001/Listing/GetListing"
		err = common.PostProcessServiceError(err, "Marketplace", "GetListing", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPackage Get the details of the specified version of a package, including information needed to launch the package.
// If you plan to launch an instance from an image listing, you must first subscribe to the listing. When
// you launch the instance, you also need to provide the image ID of the listing resource version that you want.
// Subscribing to the listing requires you to first get a signature from the terms of use agreement for the
// listing resource version. To get the signature, issue a GetAppCatalogListingAgreements (https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/latest/AppCatalogListingResourceVersionAgreements/GetAppCatalogListingAgreements) API call.
// The AppCatalogListingResourceVersionAgreements (https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/latest/AppCatalogListingResourceVersionAgreements) object, including
// its signature, is returned in the response. With the signature for the terms of use agreement for the desired
// listing resource version, create a subscription by issuing a
// CreateAppCatalogSubscription (https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/latest/AppCatalogSubscription/CreateAppCatalogSubscription) API call.
// To get the image ID to launch an instance, issue a GetAppCatalogListingResourceVersion (https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/latest/AppCatalogListingResourceVersion/GetAppCatalogListingResourceVersion) API call.
// Lastly, to launch the instance, use the image ID of the listing resource version to issue a LaunchInstance (https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/latest/Instance/LaunchInstance) API call.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/GetPackage.go.html to see an example of how to use GetPackage API.
// A default retry strategy applies to this operation GetPackage()
func (client MarketplaceClient) GetPackage(ctx context.Context, request GetPackageRequest) (response GetPackageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPackage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPackageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPackageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPackageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPackageResponse")
	}
	return
}

// getPackage implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceClient) getPackage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/listings/{listingId}/packages/{packageVersion}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPackageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/marketplace/20181001/ListingPackage/GetPackage"
		err = common.PostProcessServiceError(err, "Marketplace", "GetPackage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &listingpackage{})
	return response, err
}

// GetPublication Gets the details of the specified publication.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/GetPublication.go.html to see an example of how to use GetPublication API.
// A default retry strategy applies to this operation GetPublication()
func (client MarketplaceClient) GetPublication(ctx context.Context, request GetPublicationRequest) (response GetPublicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPublication, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPublicationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPublicationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPublicationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPublicationResponse")
	}
	return
}

// getPublication implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceClient) getPublication(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/publications/{publicationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPublicationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/marketplace/20181001/Publication/GetPublication"
		err = common.PostProcessServiceError(err, "Marketplace", "GetPublication", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPublicationPackage Gets the details of a specific package version within a given publication.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/GetPublicationPackage.go.html to see an example of how to use GetPublicationPackage API.
// A default retry strategy applies to this operation GetPublicationPackage()
func (client MarketplaceClient) GetPublicationPackage(ctx context.Context, request GetPublicationPackageRequest) (response GetPublicationPackageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPublicationPackage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPublicationPackageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPublicationPackageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPublicationPackageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPublicationPackageResponse")
	}
	return
}

// getPublicationPackage implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceClient) getPublicationPackage(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/publications/{publicationId}/packages/{packageVersion}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPublicationPackageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/marketplace/20181001/PublicationPackage/GetPublicationPackage"
		err = common.PostProcessServiceError(err, "Marketplace", "GetPublicationPackage", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &publicationpackage{})
	return response, err
}

// GetWorkRequest Gets the details of the specified work request
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
// A default retry strategy applies to this operation GetWorkRequest()
func (client MarketplaceClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
func (client MarketplaceClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/marketplace/20181001/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "Marketplace", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAcceptedAgreements Lists the terms of use agreements that have been accepted in the specified compartment.
// You can filter results by specifying query parameters.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/ListAcceptedAgreements.go.html to see an example of how to use ListAcceptedAgreements API.
// A default retry strategy applies to this operation ListAcceptedAgreements()
func (client MarketplaceClient) ListAcceptedAgreements(ctx context.Context, request ListAcceptedAgreementsRequest) (response ListAcceptedAgreementsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAcceptedAgreements, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAcceptedAgreementsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAcceptedAgreementsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAcceptedAgreementsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAcceptedAgreementsResponse")
	}
	return
}

// listAcceptedAgreements implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceClient) listAcceptedAgreements(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/acceptedAgreements", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAcceptedAgreementsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/marketplace/20181001/AcceptedAgreementSummary/ListAcceptedAgreements"
		err = common.PostProcessServiceError(err, "Marketplace", "ListAcceptedAgreements", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAgreements Returns the terms of use agreements that must be accepted before you can deploy the specified version of a package.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/ListAgreements.go.html to see an example of how to use ListAgreements API.
// A default retry strategy applies to this operation ListAgreements()
func (client MarketplaceClient) ListAgreements(ctx context.Context, request ListAgreementsRequest) (response ListAgreementsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAgreements, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAgreementsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAgreementsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAgreementsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAgreementsResponse")
	}
	return
}

// listAgreements implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceClient) listAgreements(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/listings/{listingId}/packages/{packageVersion}/agreements", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAgreementsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/marketplace/20181001/AgreementSummary/ListAgreements"
		err = common.PostProcessServiceError(err, "Marketplace", "ListAgreements", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListCategories Gets the list of all the categories for listings published to Oracle Cloud Infrastructure Marketplace. Categories apply
// to the software product provided by the listing.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/ListCategories.go.html to see an example of how to use ListCategories API.
// A default retry strategy applies to this operation ListCategories()
func (client MarketplaceClient) ListCategories(ctx context.Context, request ListCategoriesRequest) (response ListCategoriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCategories, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCategoriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCategoriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCategoriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCategoriesResponse")
	}
	return
}

// listCategories implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceClient) listCategories(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/categories", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCategoriesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/marketplace/20181001/CategorySummary/ListCategories"
		err = common.PostProcessServiceError(err, "Marketplace", "ListCategories", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListListings Gets a list of listings from Oracle Cloud Infrastructure Marketplace by searching keywords and
// filtering according to listing attributes.
// If you plan to launch an instance from an image listing, you must first subscribe to the listing. When
// you launch the instance, you also need to provide the image ID of the listing resource version that you want.
// Subscribing to the listing requires you to first get a signature from the terms of use agreement for the
// listing resource version. To get the signature, issue a GetAppCatalogListingAgreements (https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/latest/AppCatalogListingResourceVersionAgreements/GetAppCatalogListingAgreements) API call.
// The AppCatalogListingResourceVersionAgreements (https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/latest/AppCatalogListingResourceVersionAgreements) object, including
// its signature, is returned in the response. With the signature for the terms of use agreement for the desired
// listing resource version, create a subscription by issuing a
// CreateAppCatalogSubscription (https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/latest/AppCatalogSubscription/CreateAppCatalogSubscription) API call.
// To get the image ID to launch an instance, issue a GetAppCatalogListingResourceVersion (https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/latest/AppCatalogListingResourceVersion/GetAppCatalogListingResourceVersion) API call.
// Lastly, to launch the instance, use the image ID of the listing resource version to issue a LaunchInstance (https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/latest/Instance/LaunchInstance) API call.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/ListListings.go.html to see an example of how to use ListListings API.
// A default retry strategy applies to this operation ListListings()
func (client MarketplaceClient) ListListings(ctx context.Context, request ListListingsRequest) (response ListListingsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listListings, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListListingsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListListingsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListListingsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListListingsResponse")
	}
	return
}

// listListings implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceClient) listListings(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/listings", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListListingsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/marketplace/20181001/ListingSummary/ListListings"
		err = common.PostProcessServiceError(err, "Marketplace", "ListListings", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPackages Gets the list of packages for a listing.
// If you plan to launch an instance from an image listing, you must first subscribe to the listing. When
// you launch the instance, you also need to provide the image ID of the listing resource version that you want.
// Subscribing to the listing requires you to first get a signature from the terms of use agreement for the
// listing resource version. To get the signature, issue a GetAppCatalogListingAgreements (https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/latest/AppCatalogListingResourceVersionAgreements/GetAppCatalogListingAgreements) API call.
// The AppCatalogListingResourceVersionAgreements (https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/latest/AppCatalogListingResourceVersionAgreements) object, including
// its signature, is returned in the response. With the signature for the terms of use agreement for the desired
// listing resource version, create a subscription by issuing a
// CreateAppCatalogSubscription (https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/latest/AppCatalogSubscription/CreateAppCatalogSubscription) API call.
// To get the image ID to launch an instance, issue a GetAppCatalogListingResourceVersion (https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/latest/AppCatalogListingResourceVersion/GetAppCatalogListingResourceVersion) API call.
// Lastly, to launch the instance, use the image ID of the listing resource version to issue a LaunchInstance (https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/latest/Instance/LaunchInstance) API call.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/ListPackages.go.html to see an example of how to use ListPackages API.
// A default retry strategy applies to this operation ListPackages()
func (client MarketplaceClient) ListPackages(ctx context.Context, request ListPackagesRequest) (response ListPackagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPackages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPackagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPackagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPackagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPackagesResponse")
	}
	return
}

// listPackages implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceClient) listPackages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/listings/{listingId}/packages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPackagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/marketplace/20181001/ListingPackageSummary/ListPackages"
		err = common.PostProcessServiceError(err, "Marketplace", "ListPackages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPublicationPackages Lists the packages in the specified publication.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/ListPublicationPackages.go.html to see an example of how to use ListPublicationPackages API.
// A default retry strategy applies to this operation ListPublicationPackages()
func (client MarketplaceClient) ListPublicationPackages(ctx context.Context, request ListPublicationPackagesRequest) (response ListPublicationPackagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPublicationPackages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPublicationPackagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPublicationPackagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPublicationPackagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPublicationPackagesResponse")
	}
	return
}

// listPublicationPackages implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceClient) listPublicationPackages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/publications/{publicationId}/packages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPublicationPackagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/marketplace/20181001/PublicationPackageSummary/ListPublicationPackages"
		err = common.PostProcessServiceError(err, "Marketplace", "ListPublicationPackages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPublications Lists the publications in the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/ListPublications.go.html to see an example of how to use ListPublications API.
// A default retry strategy applies to this operation ListPublications()
func (client MarketplaceClient) ListPublications(ctx context.Context, request ListPublicationsRequest) (response ListPublicationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPublications, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPublicationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPublicationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPublicationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPublicationsResponse")
	}
	return
}

// listPublications implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceClient) listPublications(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/publications", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPublicationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/marketplace/20181001/PublicationSummary/ListPublications"
		err = common.PostProcessServiceError(err, "Marketplace", "ListPublications", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListPublishers Gets the list of all the publishers of listings available in Oracle Cloud Infrastructure Marketplace.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/ListPublishers.go.html to see an example of how to use ListPublishers API.
// A default retry strategy applies to this operation ListPublishers()
func (client MarketplaceClient) ListPublishers(ctx context.Context, request ListPublishersRequest) (response ListPublishersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPublishers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPublishersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPublishersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPublishersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPublishersResponse")
	}
	return
}

// listPublishers implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceClient) listPublishers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/publishers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPublishersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/marketplace/20181001/PublisherSummary/ListPublishers"
		err = common.PostProcessServiceError(err, "Marketplace", "ListPublishers", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListReportTypes Lists available types of reports for the compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/ListReportTypes.go.html to see an example of how to use ListReportTypes API.
// A default retry strategy applies to this operation ListReportTypes()
func (client MarketplaceClient) ListReportTypes(ctx context.Context, request ListReportTypesRequest) (response ListReportTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listReportTypes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListReportTypesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListReportTypesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListReportTypesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListReportTypesResponse")
	}
	return
}

// listReportTypes implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceClient) listReportTypes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/reportTypes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListReportTypesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/marketplace/20181001/ReportTypeCollection/ListReportTypes"
		err = common.PostProcessServiceError(err, "Marketplace", "ListReportTypes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListReports Lists reports in the compartment that match the specified report type and date.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/ListReports.go.html to see an example of how to use ListReports API.
// A default retry strategy applies to this operation ListReports()
func (client MarketplaceClient) ListReports(ctx context.Context, request ListReportsRequest) (response ListReportsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listReports, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListReportsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListReportsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListReportsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListReportsResponse")
	}
	return
}

// listReports implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceClient) listReports(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/reports", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListReportsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/marketplace/20181001/ReportCollection/ListReports"
		err = common.PostProcessServiceError(err, "Marketplace", "ListReports", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTaxes Returns list of all tax implications that current tenant may be liable to once they launch the listing.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/ListTaxes.go.html to see an example of how to use ListTaxes API.
// A default retry strategy applies to this operation ListTaxes()
func (client MarketplaceClient) ListTaxes(ctx context.Context, request ListTaxesRequest) (response ListTaxesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTaxes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTaxesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTaxesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTaxesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTaxesResponse")
	}
	return
}

// listTaxes implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceClient) listTaxes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/listings/{listingId}/taxes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTaxesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/marketplace/20181001/TaxSummary/ListTaxes"
		err = common.PostProcessServiceError(err, "Marketplace", "ListTaxes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors List all errors for a work request
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client MarketplaceClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
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
func (client MarketplaceClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/marketplace/20181001/WorkRequest/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "Marketplace", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs List all logs for a work request
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client MarketplaceClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
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
func (client MarketplaceClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/marketplace/20181001/WorkRequest/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "Marketplace", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests List all work requests in a compartment
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
// A default retry strategy applies to this operation ListWorkRequests()
func (client MarketplaceClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
func (client MarketplaceClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/marketplace/20181001/WorkRequest/ListWorkRequests"
		err = common.PostProcessServiceError(err, "Marketplace", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// SearchListings Queries all Marketplace Applications to find listings that match the specified criteria. To search
// for a listing, you can use a free text or structured search.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/SearchListings.go.html to see an example of how to use SearchListings API.
// A default retry strategy applies to this operation SearchListings()
func (client MarketplaceClient) SearchListings(ctx context.Context, request SearchListingsRequest) (response SearchListingsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.searchListings, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SearchListingsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SearchListingsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SearchListingsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SearchListingsResponse")
	}
	return
}

// searchListings implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceClient) searchListings(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/searchListings", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response SearchListingsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/marketplace/20181001/ListingSummary/SearchListings"
		err = common.PostProcessServiceError(err, "Marketplace", "SearchListings", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAcceptedAgreement Updates the display name or tags associated with a listing's previously accepted terms of use agreement.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/UpdateAcceptedAgreement.go.html to see an example of how to use UpdateAcceptedAgreement API.
// A default retry strategy applies to this operation UpdateAcceptedAgreement()
func (client MarketplaceClient) UpdateAcceptedAgreement(ctx context.Context, request UpdateAcceptedAgreementRequest) (response UpdateAcceptedAgreementResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateAcceptedAgreement, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAcceptedAgreementResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAcceptedAgreementResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAcceptedAgreementResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAcceptedAgreementResponse")
	}
	return
}

// updateAcceptedAgreement implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceClient) updateAcceptedAgreement(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/acceptedAgreements/{acceptedAgreementId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAcceptedAgreementResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/marketplace/20181001/AcceptedAgreement/UpdateAcceptedAgreement"
		err = common.PostProcessServiceError(err, "Marketplace", "UpdateAcceptedAgreement", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdatePublication Updates the details of an existing publication.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/UpdatePublication.go.html to see an example of how to use UpdatePublication API.
// A default retry strategy applies to this operation UpdatePublication()
func (client MarketplaceClient) UpdatePublication(ctx context.Context, request UpdatePublicationRequest) (response UpdatePublicationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updatePublication, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdatePublicationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdatePublicationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdatePublicationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdatePublicationResponse")
	}
	return
}

// updatePublication implements the OCIOperation interface (enables retrying operations)
func (client MarketplaceClient) updatePublication(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/publications/{publicationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdatePublicationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/marketplace/20181001/Publication/UpdatePublication"
		err = common.PostProcessServiceError(err, "Marketplace", "UpdatePublication", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
