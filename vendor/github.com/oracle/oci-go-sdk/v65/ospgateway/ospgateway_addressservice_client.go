// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OSP Gateway API
//
// This site describes all the Rest endpoints of OSP Gateway.
//

package ospgateway

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// AddressServiceClient a client for AddressService
type AddressServiceClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewAddressServiceClientWithConfigurationProvider Creates a new default AddressService client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewAddressServiceClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client AddressServiceClient, err error) {
	if enabled := common.CheckForEnabledServices("ospgateway"); !enabled {
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
	return newAddressServiceClientFromBaseClient(baseClient, provider)
}

// NewAddressServiceClientWithOboToken Creates a new default AddressService client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewAddressServiceClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client AddressServiceClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newAddressServiceClientFromBaseClient(baseClient, configProvider)
}

func newAddressServiceClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client AddressServiceClient, err error) {
	// AddressService service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("AddressService"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = AddressServiceClient{BaseClient: baseClient}
	client.BasePath = "20191001"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *AddressServiceClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("ospgateway", "https://ospap.oracle.com")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *AddressServiceClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *AddressServiceClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetAddress Get the address by id for the compartment
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ospgateway/GetAddress.go.html to see an example of how to use GetAddress API.
// A default retry strategy applies to this operation GetAddress()
func (client AddressServiceClient) GetAddress(ctx context.Context, request GetAddressRequest) (response GetAddressResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAddress, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAddressResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAddressResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAddressResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAddressResponse")
	}
	return
}

// getAddress implements the OCIOperation interface (enables retrying operations)
func (client AddressServiceClient) getAddress(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/addresses/{addressId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAddressResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "AddressService", "GetAddress", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// VerifyAddress Verify address
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ospgateway/VerifyAddress.go.html to see an example of how to use VerifyAddress API.
// A default retry strategy applies to this operation VerifyAddress()
func (client AddressServiceClient) VerifyAddress(ctx context.Context, request VerifyAddressRequest) (response VerifyAddressResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.verifyAddress, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = VerifyAddressResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = VerifyAddressResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(VerifyAddressResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into VerifyAddressResponse")
	}
	return
}

// verifyAddress implements the OCIOperation interface (enables retrying operations)
func (client AddressServiceClient) verifyAddress(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/addresses/action/verification", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response VerifyAddressResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "AddressService", "VerifyAddress", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
