// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//MarketplaceClient a client for Marketplace
type MarketplaceClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewMarketplaceClientWithConfigurationProvider Creates a new default Marketplace client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewMarketplaceClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client MarketplaceClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newMarketplaceClientFromBaseClient(baseClient, configProvider)
}

// NewMarketplaceClientWithOboToken Creates a new default Marketplace client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewMarketplaceClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client MarketplaceClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newMarketplaceClientFromBaseClient(baseClient, configProvider)
}

func newMarketplaceClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client MarketplaceClient, err error) {
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
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *MarketplaceClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateAcceptedAgreement Accepts a terms of use agreement for a specific package version of a listing. You must accept all
// terms of use for a package before you can deploy the package.
func (client MarketplaceClient) CreateAcceptedAgreement(ctx context.Context, request CreateAcceptedAgreementRequest) (response CreateAcceptedAgreementResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client MarketplaceClient) createAcceptedAgreement(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/acceptedAgreements")
	if err != nil {
		return nil, err
	}

	var response CreateAcceptedAgreementResponse
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

// DeleteAcceptedAgreement Removes a previously accepted terms of use agreement from the list of agreements that Marketplace checks
// before initiating a deployment. Listings in the Marketplace that require acceptance of the specified terms
// of use can no longer be deployed, but existing deployments aren't affected.
func (client MarketplaceClient) DeleteAcceptedAgreement(ctx context.Context, request DeleteAcceptedAgreementRequest) (response DeleteAcceptedAgreementResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client MarketplaceClient) deleteAcceptedAgreement(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/acceptedAgreements/{acceptedAgreementId}")
	if err != nil {
		return nil, err
	}

	var response DeleteAcceptedAgreementResponse
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

// GetAcceptedAgreement Gets the details of a specific, previously accepted terms of use agreement.
func (client MarketplaceClient) GetAcceptedAgreement(ctx context.Context, request GetAcceptedAgreementRequest) (response GetAcceptedAgreementResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client MarketplaceClient) getAcceptedAgreement(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/acceptedAgreements/{acceptedAgreementId}")
	if err != nil {
		return nil, err
	}

	var response GetAcceptedAgreementResponse
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

// GetAgreement Returns a terms of use agreement for a package with a time-based signature that can be used to
// accept the agreement.
func (client MarketplaceClient) GetAgreement(ctx context.Context, request GetAgreementRequest) (response GetAgreementResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client MarketplaceClient) getAgreement(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/listings/{listingId}/packages/{packageVersion}/agreements/{agreementId}")
	if err != nil {
		return nil, err
	}

	var response GetAgreementResponse
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
func (client MarketplaceClient) GetListing(ctx context.Context, request GetListingRequest) (response GetListingResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client MarketplaceClient) getListing(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/listings/{listingId}")
	if err != nil {
		return nil, err
	}

	var response GetListingResponse
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
func (client MarketplaceClient) GetPackage(ctx context.Context, request GetPackageRequest) (response GetPackageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client MarketplaceClient) getPackage(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/listings/{listingId}/packages/{packageVersion}")
	if err != nil {
		return nil, err
	}

	var response GetPackageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &listingpackage{})
	return response, err
}

// ListAcceptedAgreements Lists the terms of use agreements that have been accepted in the specified compartment.
// You can filter results by specifying query parameters.
func (client MarketplaceClient) ListAcceptedAgreements(ctx context.Context, request ListAcceptedAgreementsRequest) (response ListAcceptedAgreementsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client MarketplaceClient) listAcceptedAgreements(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/acceptedAgreements")
	if err != nil {
		return nil, err
	}

	var response ListAcceptedAgreementsResponse
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

// ListAgreements Returns the terms of use agreements that must be accepted before you can deploy the specified version of a package.
func (client MarketplaceClient) ListAgreements(ctx context.Context, request ListAgreementsRequest) (response ListAgreementsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client MarketplaceClient) listAgreements(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/listings/{listingId}/packages/{packageVersion}/agreements")
	if err != nil {
		return nil, err
	}

	var response ListAgreementsResponse
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

// ListCategories Gets the list of all the categories for listings published to Oracle Cloud Infrastructure Marketplace. Categories apply
// to the software product provided by the listing.
func (client MarketplaceClient) ListCategories(ctx context.Context, request ListCategoriesRequest) (response ListCategoriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client MarketplaceClient) listCategories(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/categories")
	if err != nil {
		return nil, err
	}

	var response ListCategoriesResponse
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
func (client MarketplaceClient) ListListings(ctx context.Context, request ListListingsRequest) (response ListListingsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client MarketplaceClient) listListings(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/listings")
	if err != nil {
		return nil, err
	}

	var response ListListingsResponse
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
func (client MarketplaceClient) ListPackages(ctx context.Context, request ListPackagesRequest) (response ListPackagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client MarketplaceClient) listPackages(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/listings/{listingId}/packages")
	if err != nil {
		return nil, err
	}

	var response ListPackagesResponse
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

// ListPublishers Gets the list of all the publishers of listings available in Oracle Cloud Infrastructure Marketplace.
func (client MarketplaceClient) ListPublishers(ctx context.Context, request ListPublishersRequest) (response ListPublishersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client MarketplaceClient) listPublishers(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/publishers")
	if err != nil {
		return nil, err
	}

	var response ListPublishersResponse
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

// ListReportTypes Lists available types of reports for the compartment.
func (client MarketplaceClient) ListReportTypes(ctx context.Context, request ListReportTypesRequest) (response ListReportTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client MarketplaceClient) listReportTypes(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/reportTypes")
	if err != nil {
		return nil, err
	}

	var response ListReportTypesResponse
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

// ListReports Lists reports in the compartment that match the specified report type and date.
func (client MarketplaceClient) ListReports(ctx context.Context, request ListReportsRequest) (response ListReportsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client MarketplaceClient) listReports(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/reports")
	if err != nil {
		return nil, err
	}

	var response ListReportsResponse
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

// UpdateAcceptedAgreement Updates the display name or tags associated with a listing's previously accepted terms of use agreement.
func (client MarketplaceClient) UpdateAcceptedAgreement(ctx context.Context, request UpdateAcceptedAgreementRequest) (response UpdateAcceptedAgreementResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
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
func (client MarketplaceClient) updateAcceptedAgreement(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/acceptedAgreements/{acceptedAgreementId}")
	if err != nil {
		return nil, err
	}

	var response UpdateAcceptedAgreementResponse
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
