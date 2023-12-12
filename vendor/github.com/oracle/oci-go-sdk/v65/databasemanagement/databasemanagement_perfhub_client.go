// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// PerfhubClient a client for Perfhub
type PerfhubClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewPerfhubClientWithConfigurationProvider Creates a new default Perfhub client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewPerfhubClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client PerfhubClient, err error) {
	if enabled := common.CheckForEnabledServices("databasemanagement"); !enabled {
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
	return newPerfhubClientFromBaseClient(baseClient, provider)
}

// NewPerfhubClientWithOboToken Creates a new default Perfhub client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewPerfhubClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client PerfhubClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newPerfhubClientFromBaseClient(baseClient, configProvider)
}

func newPerfhubClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client PerfhubClient, err error) {
	// Perfhub service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Perfhub"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = PerfhubClient{BaseClient: baseClient}
	client.BasePath = "20201101"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *PerfhubClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("databasemanagement", "https://dbmgmt.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *PerfhubClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *PerfhubClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ModifySnapshotSettings Modifies the snapshot settings for the specified Database.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ModifySnapshotSettings.go.html to see an example of how to use ModifySnapshotSettings API.
func (client PerfhubClient) ModifySnapshotSettings(ctx context.Context, request ModifySnapshotSettingsRequest) (response ModifySnapshotSettingsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.modifySnapshotSettings, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ModifySnapshotSettingsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ModifySnapshotSettingsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ModifySnapshotSettingsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ModifySnapshotSettingsResponse")
	}
	return
}

// modifySnapshotSettings implements the OCIOperation interface (enables retrying operations)
func (client PerfhubClient) modifySnapshotSettings(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/managedDatabases/{managedDatabaseId}/actions/modifySnapshotSettings", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ModifySnapshotSettingsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/database-management/20201101/ManagedDatabase/ModifySnapshotSettings"
		err = common.PostProcessServiceError(err, "Perfhub", "ModifySnapshotSettings", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
