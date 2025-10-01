// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Multicloud API
//
// Use the Oracle Multicloud API to retrieve resource anchors and network anchors, and the metadata mappings related a Cloud Service Provider. For more information, see <link to docs>.
//

package multicloud

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// MetadataClient a client for Metadata
type MetadataClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewMetadataClientWithConfigurationProvider Creates a new default Metadata client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewMetadataClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client MetadataClient, err error) {
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
	return newMetadataClientFromBaseClient(baseClient, provider)
}

// NewMetadataClientWithOboToken Creates a new default Metadata client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewMetadataClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client MetadataClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newMetadataClientFromBaseClient(baseClient, configProvider)
}

func newMetadataClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client MetadataClient, err error) {
	// Metadata service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Metadata"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = MetadataClient{BaseClient: baseClient}
	client.BasePath = "20180828"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *MetadataClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("multicloud", "https://multicloud.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *MetadataClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *MetadataClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ListExternalLocationDetailsMetadata List externalLocationDetail metadata from OCI to Cloud  Service Provider for regions, Availability Zones, and Cluster Placement Group ID.
// examples:
//
//	application-json: |
//	  [
//	    {
//	        "externalLocation": {
//	          "cspRegion": "East US",
//	          "cspPhysicalAz": "az1-xyz",
//	          "cspPhysicalAzDisplayName": "(US) East US 2",
//	          "cspLogicalAz": "az1",
//	          "serviceName": "ORACLEDBATAZURE",
//	          "cspZoneKeyReferenceId": {
//	            "keyName": "AzureSubscriptionId or AwsAccountId, GcpProjectName",
//	            "keyValue": "azure-subscriptionId-1 or aws-account-id-1, gcp-project-id-1"
//	          }
//	        },
//	        "ociPhysicalAd": "ad1-xyb",
//	        "ociLogicalAd": "ad2",
//	        "ociRegion": "us-ashburn-1",
//	        "cpgId": "cpg-1"
//	    },
//	      {
//	        "externalLocation": {
//	          "cspRegion": "East US",
//	          "cspPhysicalAz": "az2-abc",
//	          "cspPhysicalAzDisplayName": "(US) East US 2",
//	          "cspLogicalAz": "az2",
//	          "serviceName": "ORACLEDBATAZURE",
//	          "cspZoneKeyReferenceId": {
//	            "keyName": "AzureSubscriptionId or AwsAccountId, GcpProjectName",
//	            "keyValue": "azure-subscriptionId-2 or aws-account-id-2, gcp-project-id-2"
//	          }
//	        },
//	        "ociPhysicalAd": "ad2-xby",
//	        "ociLogicalAd": "ad1",
//	        "ociRegion": "us-ashburn-1",
//	        "cpgId": "cpg-2"
//	      },
//	      {
//	        "externalLocation": {
//	          "cspRegion": "East US",
//	          "cspPhysicalAz": "az3-abz",
//	          "cspPhysicalAzDisplayName": "(US) East US 2",
//	          "cspLogicalAz": "az3",
//	          "serviceName": "ORACLEDBATAZURE",
//	          "cspZoneKeyReferenceId": {
//	            "keyName": "AzureSubscriptionId or AwsAccountId, GcpProjectName",
//	            "keyValue": "azure-subscriptionId-3 or aws-account-id-3, gcp-project-id-3"
//	          }
//	        },
//	        "ociPhysicalAd": "ad3-cde",
//	        "ociLogicalAd": "ad3",
//	        "ociRegion": "us-ashburn-1",
//	        "cpgId": "cpg-3"
//	      },
//	      {
//	        "externalLocation": {
//	          "cspRegion": "East US 2",
//	          "cspPhysicalAz": "az1-def",
//	          "cspPhysicalAzDisplayName": "(US) East US 2",
//	          "cspLogicalAz": "az1",
//	          "serviceName": "ORACLEDBATAZURE",
//	          "cspZoneKeyReferenceId": {
//	            "keyName": "AzureSubscriptionId or AwsAccountId, GcpProjectName",
//	            "keyValue": "azure-subscriptionId-4 or aws-account-id-4, gcp-project-id-4"
//	          }
//	        },
//	        "ociPhysicalAd": "ad1-bce",
//	        "ociLogicalAd": "ad2",
//	        "ociRegion": "us-ashburn-1",
//	        "cpgId": "cpg-4"
//	      },
//	      {
//	        "externalLocation": {
//	          "cspRegion": "East US 2",
//	          "cspPhysicalAz": "az2-uvw",
//	          "cspPhysicalAzDisplayName": "(US) East US 2",
//	          "cspLogicalAz": "az2",
//	          "serviceName": "ORACLEDBATAZURE",
//	          "cspZoneKeyReferenceId": {
//	            "keyName": "AzureSubscriptionId or AwsAccountId, GcpProjectName",
//	            "keyValue": "azure-subscriptionId-3 or aws-account-id-3, gcp-project-id-3"
//	          }
//	        },
//	        "ociPhysicalAd": "ad2-ftc",
//	        "ociLogicalAd": "ad1",
//	        "ociRegion": "us-ashburn-1",
//	        "cpgId": "cpg-5"
//	      },
//	      {
//	        "externalLocation": {
//	          "cspRegion": "East US 2",
//	          "cspPhysicalAz": "az3-uvw",
//	          "cspPhysicalAzDisplayName": "(US) East US 2",
//	          "cspLogicalAz": "az3",
//	          "serviceName": "ORACLEDBATAZURE",
//	          "cspZoneKeyReferenceId": {
//	            "keyName": "AzureSubscriptionId or AwsAccountId, GcpProjectName",
//	            "keyValue": "azure-subscriptionId-3 or aws-account-id-3, gcp-project-id-3"
//	          }
//	        },
//	        "ociPhysicalAd": "ad3-stc",
//	        "ociLogicalAd": "ad3",
//	        "ociRegion": "us-ashburn-1",
//	        "cpgId": "cpg-6"
//	      }
//	    ]
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/multicloud/ListExternalLocationDetailsMetadata.go.html to see an example of how to use ListExternalLocationDetailsMetadata API.
// A default retry strategy applies to this operation ListExternalLocationDetailsMetadata()
func (client MetadataClient) ListExternalLocationDetailsMetadata(ctx context.Context, request ListExternalLocationDetailsMetadataRequest) (response ListExternalLocationDetailsMetadataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExternalLocationDetailsMetadata, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExternalLocationDetailsMetadataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExternalLocationDetailsMetadataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExternalLocationDetailsMetadataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExternalLocationDetailsMetadataResponse")
	}
	return
}

// listExternalLocationDetailsMetadata implements the OCIOperation interface (enables retrying operations)
func (client MetadataClient) listExternalLocationDetailsMetadata(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalLocationsMetadata", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExternalLocationDetailsMetadataResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/multicloud-omhub-cp/20180828/ExternalLocationsMetadatumCollection/ListExternalLocationDetailsMetadata"
		err = common.PostProcessServiceError(err, "Metadata", "ListExternalLocationDetailsMetadata", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListExternalLocationMappingMetadata List externalLocation metadata from OCI to the Cloud Service Provider for regions, Physical Availability Zones.
// examples:
//
//	application-json: |
//	  [
//	    {
//	      "externalLocation": {
//	        "cspRegion": "eastus",
//	        "cspPhysicalAz": "eastus-az1",
//	        "cspPhysicalAzDisplayName": "(US) East US 1",
//	        "serviceName": "ORACLEDBATAZURE"
//	      },
//	      "ociPhysicalAd": "iad-ad-1",
//	      "ociLogicalAd": "ad1",
//	      "ociRegion": "us-ashburn-1"
//	  },
//	    {
//	      "externalLocation": {
//	        "cspRegion": "eastus",
//	        "cspPhysicalAz": "eastus-az1",
//	        "cspPhysicalAzDisplayName": "(US) East US 1",
//	        "serviceName": "ORACLEDBATAZURE"
//	      },
//	      "ociPhysicalAd": "iad-ad-1",
//	      "ociLogicalAd": "ad1",
//	      "ociRegion": "us-ashburn-1"
//	    },
//	    {
//	      "externalLocation": {
//	        "cspRegion": "eastus2",
//	        "cspPhysicalAz": "eastus2-az3",
//	        "cspPhysicalAzDisplayName": "(US) East US 1",
//	        "serviceName": "ORACLEDBATAZURE"
//	      },
//	      "ociPhysicalAd": "iad-ad-2",
//	      "ociLogicalAd": "ad1",
//	      "ociRegion": "us-ashburn-1"
//	    },
//	    {
//	      "externalLocation": {
//	        "cspRegion": "eastus",
//	        "cspPhysicalAz": "eastus-az3"
//	        "cspPhysicalAzDisplayName": "(US) East US 1",
//	        "serviceName": "ORACLEDBATAZURE"
//	      },
//	      "ociPhysicalAd": "iad-ad-333",
//	      "ociLogicalAd": "ad1",
//	      "ociRegion": "us-ashburn-1"
//	    }
//	  ]
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/multicloud/ListExternalLocationMappingMetadata.go.html to see an example of how to use ListExternalLocationMappingMetadata API.
// A default retry strategy applies to this operation ListExternalLocationMappingMetadata()
func (client MetadataClient) ListExternalLocationMappingMetadata(ctx context.Context, request ListExternalLocationMappingMetadataRequest) (response ListExternalLocationMappingMetadataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExternalLocationMappingMetadata, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExternalLocationMappingMetadataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExternalLocationMappingMetadataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExternalLocationMappingMetadataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExternalLocationMappingMetadataResponse")
	}
	return
}

// listExternalLocationMappingMetadata implements the OCIOperation interface (enables retrying operations)
func (client MetadataClient) listExternalLocationMappingMetadata(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalLocationMappingMetadata", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExternalLocationMappingMetadataResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/multicloud-omhub-cp/20180828/ExternalLocationMappingMetadatumSummaryCollection/ListExternalLocationMappingMetadata"
		err = common.PostProcessServiceError(err, "Metadata", "ListExternalLocationMappingMetadata", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListExternalLocationSummariesMetadata List externalLocationSummary metadata from OCI Region to the Cloud Service Provider region across all regions.
// examples:
//
//	application-json: |
//	  [
//	    {
//	        "externalLocation": {
//	          "cspRegion": "East US"
//	        },
//	        "ociRegion": "us-ashburn-1"
//	    },
//	      {
//	        "externalLocation": {
//	          "cspRegion": "East US 2"
//	        },
//	        "ociRegion": "us-ashburn-1"
//	      },
//	      {
//	        "externalLocation": {
//	          "cspRegion": "Germany West Central"
//	        },
//	        "ociRegion": "eu-frankfurt-1",
//	      }
//	    ]
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/multicloud/ListExternalLocationSummariesMetadata.go.html to see an example of how to use ListExternalLocationSummariesMetadata API.
// A default retry strategy applies to this operation ListExternalLocationSummariesMetadata()
func (client MetadataClient) ListExternalLocationSummariesMetadata(ctx context.Context, request ListExternalLocationSummariesMetadataRequest) (response ListExternalLocationSummariesMetadataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listExternalLocationSummariesMetadata, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListExternalLocationSummariesMetadataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListExternalLocationSummariesMetadataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListExternalLocationSummariesMetadataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListExternalLocationSummariesMetadataResponse")
	}
	return
}

// listExternalLocationSummariesMetadata implements the OCIOperation interface (enables retrying operations)
func (client MetadataClient) listExternalLocationSummariesMetadata(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/externalLocationSummariesMetadata", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListExternalLocationSummariesMetadataResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/multicloud-omhub-cp/20180828/ExternalLocationSummariesMetadatumSummaryCollection/ListExternalLocationSummariesMetadata"
		err = common.PostProcessServiceError(err, "Metadata", "ListExternalLocationSummariesMetadata", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
