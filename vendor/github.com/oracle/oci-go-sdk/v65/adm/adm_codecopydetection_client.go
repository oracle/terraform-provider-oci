// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Dependency Management API
//
// Use the Application Dependency Management API to create knowledge bases and vulnerability audits.  For more information, see ADM (https://docs.oracle.com/iaas/Content/application-dependency-management/home.htm).
//

package adm

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// CodeCopyDetectionClient a client for CodeCopyDetection
type CodeCopyDetectionClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewCodeCopyDetectionClientWithConfigurationProvider Creates a new default CodeCopyDetection client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewCodeCopyDetectionClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client CodeCopyDetectionClient, err error) {
	if enabled := common.CheckForEnabledServices("adm"); !enabled {
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
	return newCodeCopyDetectionClientFromBaseClient(baseClient, provider)
}

// NewCodeCopyDetectionClientWithOboToken Creates a new default CodeCopyDetection client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewCodeCopyDetectionClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client CodeCopyDetectionClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newCodeCopyDetectionClientFromBaseClient(baseClient, configProvider)
}

func newCodeCopyDetectionClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client CodeCopyDetectionClient, err error) {
	// CodeCopyDetection service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("CodeCopyDetection"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = CodeCopyDetectionClient{BaseClient: baseClient}
	client.BasePath = "20220421"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *CodeCopyDetectionClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("adm", "https://adm.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *CodeCopyDetectionClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *CodeCopyDetectionClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// DetectCodeCopies List source locations and licenses that are detected to match the the provided code snippet
// A default retry strategy applies to this operation DetectCodeCopies()
func (client CodeCopyDetectionClient) DetectCodeCopies(ctx context.Context, request DetectCodeCopiesRequest) (response DetectCodeCopiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.detectCodeCopies, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DetectCodeCopiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DetectCodeCopiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DetectCodeCopiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DetectCodeCopiesResponse")
	}
	return
}

// detectCodeCopies implements the OCIOperation interface (enables retrying operations)
func (client CodeCopyDetectionClient) detectCodeCopies(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/knowledgeBases/{knowledgeBaseId}/actions/detectCodeCopy", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DetectCodeCopiesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/adm/20220421/CodeCopyCollection/DetectCodeCopies"
		err = common.PostProcessServiceError(err, "CodeCopyDetection", "DetectCodeCopies", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
