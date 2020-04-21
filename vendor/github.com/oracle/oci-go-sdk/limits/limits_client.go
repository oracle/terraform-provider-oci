// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service limits APIs
//
// APIs that interact with the resource limits of a specific resource type
//

package limits

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//LimitsClient a client for Limits
type LimitsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewLimitsClientWithConfigurationProvider Creates a new default Limits client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewLimitsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client LimitsClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newLimitsClientFromBaseClient(baseClient, configProvider)
}

// NewLimitsClientWithOboToken Creates a new default Limits client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewLimitsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client LimitsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newLimitsClientFromBaseClient(baseClient, configProvider)
}

func newLimitsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client LimitsClient, err error) {
	client = LimitsClient{BaseClient: baseClient}
	client.BasePath = ""
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *LimitsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("limits", "https://limits.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *LimitsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *LimitsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetResourceAvailability For a given compartmentId, resource limit name, and scope, returns the following:
//   - the number of available resources associated with the given limit
//   - the usage in the selected compartment for the given limit
//   Note: not all resource limits support this API. If the value is not available, the API will return 404.
func (client LimitsClient) GetResourceAvailability(ctx context.Context, request GetResourceAvailabilityRequest) (response GetResourceAvailabilityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getResourceAvailability, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetResourceAvailabilityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetResourceAvailabilityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetResourceAvailabilityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetResourceAvailabilityResponse")
	}
	return
}

// getResourceAvailability implements the OCIOperation interface (enables retrying operations)
func (client LimitsClient) getResourceAvailability(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/20190729/services/{serviceName}/limits/{limitName}/resourceAvailability")
	if err != nil {
		return nil, err
	}

	var response GetResourceAvailabilityResponse
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

// ListLimitDefinitions Includes a list of resource limits that are currently supported.
// If the 'areQuotasSupported' property is true, you can create quota policies on top of this limit at the
// compartment level.
func (client LimitsClient) ListLimitDefinitions(ctx context.Context, request ListLimitDefinitionsRequest) (response ListLimitDefinitionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLimitDefinitions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLimitDefinitionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLimitDefinitionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLimitDefinitionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLimitDefinitionsResponse")
	}
	return
}

// listLimitDefinitions implements the OCIOperation interface (enables retrying operations)
func (client LimitsClient) listLimitDefinitions(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/20190729/limitDefinitions")
	if err != nil {
		return nil, err
	}

	var response ListLimitDefinitionsResponse
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

// ListLimitValues Includes a full list of resource limits belonging to a given service.
func (client LimitsClient) ListLimitValues(ctx context.Context, request ListLimitValuesRequest) (response ListLimitValuesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLimitValues, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLimitValuesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLimitValuesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLimitValuesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLimitValuesResponse")
	}
	return
}

// listLimitValues implements the OCIOperation interface (enables retrying operations)
func (client LimitsClient) listLimitValues(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/20190729/limitValues")
	if err != nil {
		return nil, err
	}

	var response ListLimitValuesResponse
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

// ListServices Returns the list of supported services.
// This will include the programmatic service name, along with the friendly service name.
func (client LimitsClient) ListServices(ctx context.Context, request ListServicesRequest) (response ListServicesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listServices, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListServicesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListServicesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListServicesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListServicesResponse")
	}
	return
}

// listServices implements the OCIOperation interface (enables retrying operations)
func (client LimitsClient) listServices(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/20190729/services")
	if err != nil {
		return nil, err
	}

	var response ListServicesResponse
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
