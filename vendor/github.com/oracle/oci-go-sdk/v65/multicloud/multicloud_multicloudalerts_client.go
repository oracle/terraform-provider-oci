// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Multicloud API
//
// Use the Oracle Multicloud API to retrieve resource anchors and network anchors, and the metadata mappings related a Cloud Service Provider. For more information, see Oracle Multicloud Hub (https://docs.oracle.com/iaas/Content/multicloud-hub/home.htm).
//

package multicloud

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// MulticloudAlertsClient a client for MulticloudAlerts
type MulticloudAlertsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewMulticloudAlertsClientWithConfigurationProvider Creates a new default MulticloudAlerts client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewMulticloudAlertsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client MulticloudAlertsClient, err error) {
	if enabled := common.CheckForEnabledServices("multicloud"); !enabled {
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
	return newMulticloudAlertsClientFromBaseClient(baseClient, provider)
}

// NewMulticloudAlertsClientWithOboToken Creates a new default MulticloudAlerts client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewMulticloudAlertsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client MulticloudAlertsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newMulticloudAlertsClientFromBaseClient(baseClient, configProvider)
}

func newMulticloudAlertsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client MulticloudAlertsClient, err error) {
	// MulticloudAlerts service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("MulticloudAlerts"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = MulticloudAlertsClient{BaseClient: baseClient}
	client.BasePath = "20180828"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *MulticloudAlertsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("multicloud", "https://multicloud.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *MulticloudAlertsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *MulticloudAlertsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ListMulticloudAlerts Gets a list of Multicloud Alerts for a given root compartment.
// Optional query parameters can be used to filter alerts by resource,
// subscription, severity, lifecycle state, and alert status.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/multicloud/ListMulticloudAlerts.go.html to see an example of how to use ListMulticloudAlerts API.
// A default retry strategy applies to this operation ListMulticloudAlerts()
func (client MulticloudAlertsClient) ListMulticloudAlerts(ctx context.Context, request ListMulticloudAlertsRequest) (response ListMulticloudAlertsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMulticloudAlerts, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMulticloudAlertsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMulticloudAlertsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMulticloudAlertsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMulticloudAlertsResponse")
	}
	return
}

// listMulticloudAlerts implements the OCIOperation interface (enables retrying operations)
func (client MulticloudAlertsClient) listMulticloudAlerts(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/multicloudalerts", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMulticloudAlertsResponse
	var httpResponse *http.Response
	httpResponse, err = client.CallWithServiceAndOperationName(ctx, &httpRequest, "multicloudAlerts", "ListMulticloudAlerts")
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/multicloud-omhub-cp/20180828/MulticloudAlertCollection/ListMulticloudAlerts"
		err = common.PostProcessServiceError(err, "MulticloudAlerts", "ListMulticloudAlerts", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
