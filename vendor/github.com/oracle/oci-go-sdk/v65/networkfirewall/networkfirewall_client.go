// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Firewall API
//
// Use the Network Firewall API to create network firewalls and configure policies that regulates network traffic in and across VCNs.
//

package networkfirewall

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// NetworkFirewallClient a client for NetworkFirewall
type NetworkFirewallClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewNetworkFirewallClientWithConfigurationProvider Creates a new default NetworkFirewall client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewNetworkFirewallClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client NetworkFirewallClient, err error) {
	if enabled := common.CheckForEnabledServices("networkfirewall"); !enabled {
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
	return newNetworkFirewallClientFromBaseClient(baseClient, provider)
}

// NewNetworkFirewallClientWithOboToken Creates a new default NetworkFirewall client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewNetworkFirewallClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client NetworkFirewallClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newNetworkFirewallClientFromBaseClient(baseClient, configProvider)
}

func newNetworkFirewallClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client NetworkFirewallClient, err error) {
	// NetworkFirewall service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("NetworkFirewall"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = NetworkFirewallClient{BaseClient: baseClient}
	client.BasePath = "20230501"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *NetworkFirewallClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("networkfirewall", "https://network-firewall.{region}.ocs.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *NetworkFirewallClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *NetworkFirewallClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ApplyNetworkFirewallPolicy Applies the candidate version of the NetworkFirewallPolicy resource. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/ApplyNetworkFirewallPolicy.go.html to see an example of how to use ApplyNetworkFirewallPolicy API.
// A default retry strategy applies to this operation ApplyNetworkFirewallPolicy()
func (client NetworkFirewallClient) ApplyNetworkFirewallPolicy(ctx context.Context, request ApplyNetworkFirewallPolicyRequest) (response ApplyNetworkFirewallPolicyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.applyNetworkFirewallPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ApplyNetworkFirewallPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ApplyNetworkFirewallPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ApplyNetworkFirewallPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ApplyNetworkFirewallPolicyResponse")
	}
	return
}

// applyNetworkFirewallPolicy implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) applyNetworkFirewallPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkFirewallPolicies/{networkFirewallPolicyId}/actions/applyPolicy", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ApplyNetworkFirewallPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/NetworkFirewallPolicy/ApplyNetworkFirewallPolicy"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "ApplyNetworkFirewallPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// BulkUploadAddressLists Creates a new Address Lists at bulk for the Network Firewall Policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/BulkUploadAddressLists.go.html to see an example of how to use BulkUploadAddressLists API.
// A default retry strategy applies to this operation BulkUploadAddressLists()
func (client NetworkFirewallClient) BulkUploadAddressLists(ctx context.Context, request BulkUploadAddressListsRequest) (response BulkUploadAddressListsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.bulkUploadAddressLists, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BulkUploadAddressListsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BulkUploadAddressListsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BulkUploadAddressListsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BulkUploadAddressListsResponse")
	}
	return
}

// bulkUploadAddressLists implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) bulkUploadAddressLists(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkFirewallPolicies/{networkFirewallPolicyId}/addressLists/actions/bulkUpload", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BulkUploadAddressListsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/AddressList/BulkUploadAddressLists"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "BulkUploadAddressLists", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// BulkUploadApplicationGroups Creates a new Application Group at bulk for the Network Firewall Policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/BulkUploadApplicationGroups.go.html to see an example of how to use BulkUploadApplicationGroups API.
// A default retry strategy applies to this operation BulkUploadApplicationGroups()
func (client NetworkFirewallClient) BulkUploadApplicationGroups(ctx context.Context, request BulkUploadApplicationGroupsRequest) (response BulkUploadApplicationGroupsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.bulkUploadApplicationGroups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BulkUploadApplicationGroupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BulkUploadApplicationGroupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BulkUploadApplicationGroupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BulkUploadApplicationGroupsResponse")
	}
	return
}

// bulkUploadApplicationGroups implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) bulkUploadApplicationGroups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkFirewallPolicies/{networkFirewallPolicyId}/applicationGroups/actions/bulkUpload", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BulkUploadApplicationGroupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/ApplicationGroup/BulkUploadApplicationGroups"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "BulkUploadApplicationGroups", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// BulkUploadApplications Creates new Applications at bulk for the Network Firewall Policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/BulkUploadApplications.go.html to see an example of how to use BulkUploadApplications API.
// A default retry strategy applies to this operation BulkUploadApplications()
func (client NetworkFirewallClient) BulkUploadApplications(ctx context.Context, request BulkUploadApplicationsRequest) (response BulkUploadApplicationsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.bulkUploadApplications, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BulkUploadApplicationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BulkUploadApplicationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BulkUploadApplicationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BulkUploadApplicationsResponse")
	}
	return
}

// bulkUploadApplications implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) bulkUploadApplications(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkFirewallPolicies/{networkFirewallPolicyId}/applications/actions/bulkUpload", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BulkUploadApplicationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/Application/BulkUploadApplications"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "BulkUploadApplications", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// BulkUploadDecryptionProfiles Creates new Decryption Profiles at bulk for the Network Firewall Policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/BulkUploadDecryptionProfiles.go.html to see an example of how to use BulkUploadDecryptionProfiles API.
// A default retry strategy applies to this operation BulkUploadDecryptionProfiles()
func (client NetworkFirewallClient) BulkUploadDecryptionProfiles(ctx context.Context, request BulkUploadDecryptionProfilesRequest) (response BulkUploadDecryptionProfilesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.bulkUploadDecryptionProfiles, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BulkUploadDecryptionProfilesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BulkUploadDecryptionProfilesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BulkUploadDecryptionProfilesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BulkUploadDecryptionProfilesResponse")
	}
	return
}

// bulkUploadDecryptionProfiles implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) bulkUploadDecryptionProfiles(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkFirewallPolicies/{networkFirewallPolicyId}/decryptionProfiles/actions/bulkUpload", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BulkUploadDecryptionProfilesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/DecryptionProfile/BulkUploadDecryptionProfiles"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "BulkUploadDecryptionProfiles", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// BulkUploadDecryptionRules Creates Decryption Rules at bulk for the Network Firewall Policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/BulkUploadDecryptionRules.go.html to see an example of how to use BulkUploadDecryptionRules API.
// A default retry strategy applies to this operation BulkUploadDecryptionRules()
func (client NetworkFirewallClient) BulkUploadDecryptionRules(ctx context.Context, request BulkUploadDecryptionRulesRequest) (response BulkUploadDecryptionRulesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.bulkUploadDecryptionRules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BulkUploadDecryptionRulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BulkUploadDecryptionRulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BulkUploadDecryptionRulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BulkUploadDecryptionRulesResponse")
	}
	return
}

// bulkUploadDecryptionRules implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) bulkUploadDecryptionRules(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkFirewallPolicies/{networkFirewallPolicyId}/decryptionRules/actions/bulkUpload", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BulkUploadDecryptionRulesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/DecryptionRule/BulkUploadDecryptionRules"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "BulkUploadDecryptionRules", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// BulkUploadMappedSecrets Creates new Mapped Secrets at bulk for the Network Firewall Policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/BulkUploadMappedSecrets.go.html to see an example of how to use BulkUploadMappedSecrets API.
// A default retry strategy applies to this operation BulkUploadMappedSecrets()
func (client NetworkFirewallClient) BulkUploadMappedSecrets(ctx context.Context, request BulkUploadMappedSecretsRequest) (response BulkUploadMappedSecretsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.bulkUploadMappedSecrets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BulkUploadMappedSecretsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BulkUploadMappedSecretsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BulkUploadMappedSecretsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BulkUploadMappedSecretsResponse")
	}
	return
}

// bulkUploadMappedSecrets implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) bulkUploadMappedSecrets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkFirewallPolicies/{networkFirewallPolicyId}/mappedSecrets/actions/bulkUpload", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BulkUploadMappedSecretsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/MappedSecret/BulkUploadMappedSecrets"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "BulkUploadMappedSecrets", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// BulkUploadSecurityRules Creates a new Security Rule at bulk for the Network Firewall Policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/BulkUploadSecurityRules.go.html to see an example of how to use BulkUploadSecurityRules API.
// A default retry strategy applies to this operation BulkUploadSecurityRules()
func (client NetworkFirewallClient) BulkUploadSecurityRules(ctx context.Context, request BulkUploadSecurityRulesRequest) (response BulkUploadSecurityRulesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.bulkUploadSecurityRules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BulkUploadSecurityRulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BulkUploadSecurityRulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BulkUploadSecurityRulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BulkUploadSecurityRulesResponse")
	}
	return
}

// bulkUploadSecurityRules implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) bulkUploadSecurityRules(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkFirewallPolicies/{networkFirewallPolicyId}/securityRules/actions/bulkUpload", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BulkUploadSecurityRulesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/SecurityRule/BulkUploadSecurityRules"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "BulkUploadSecurityRules", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// BulkUploadServiceLists Creates a new Service List at bulk for the Network Firewall Policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/BulkUploadServiceLists.go.html to see an example of how to use BulkUploadServiceLists API.
// A default retry strategy applies to this operation BulkUploadServiceLists()
func (client NetworkFirewallClient) BulkUploadServiceLists(ctx context.Context, request BulkUploadServiceListsRequest) (response BulkUploadServiceListsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.bulkUploadServiceLists, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BulkUploadServiceListsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BulkUploadServiceListsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BulkUploadServiceListsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BulkUploadServiceListsResponse")
	}
	return
}

// bulkUploadServiceLists implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) bulkUploadServiceLists(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkFirewallPolicies/{networkFirewallPolicyId}/serviceLists/actions/bulkUpload", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BulkUploadServiceListsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/ServiceList/BulkUploadServiceLists"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "BulkUploadServiceLists", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// BulkUploadServices Creates new Services at bulk for the Network Firewall Policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/BulkUploadServices.go.html to see an example of how to use BulkUploadServices API.
// A default retry strategy applies to this operation BulkUploadServices()
func (client NetworkFirewallClient) BulkUploadServices(ctx context.Context, request BulkUploadServicesRequest) (response BulkUploadServicesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.bulkUploadServices, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BulkUploadServicesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BulkUploadServicesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BulkUploadServicesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BulkUploadServicesResponse")
	}
	return
}

// bulkUploadServices implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) bulkUploadServices(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkFirewallPolicies/{networkFirewallPolicyId}/services/actions/bulkUpload", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BulkUploadServicesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/Service/BulkUploadServices"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "BulkUploadServices", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// BulkUploadUrlLists Creates a new Url Lists at bulk for the Network Firewall Policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/BulkUploadUrlLists.go.html to see an example of how to use BulkUploadUrlLists API.
// A default retry strategy applies to this operation BulkUploadUrlLists()
func (client NetworkFirewallClient) BulkUploadUrlLists(ctx context.Context, request BulkUploadUrlListsRequest) (response BulkUploadUrlListsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.bulkUploadUrlLists, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BulkUploadUrlListsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BulkUploadUrlListsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BulkUploadUrlListsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BulkUploadUrlListsResponse")
	}
	return
}

// bulkUploadUrlLists implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) bulkUploadUrlLists(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkFirewallPolicies/{networkFirewallPolicyId}/urlLists/actions/bulkUpload", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BulkUploadUrlListsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/UrlList/BulkUploadUrlLists"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "BulkUploadUrlLists", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CancelWorkRequest Cancel work request with the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/CancelWorkRequest.go.html to see an example of how to use CancelWorkRequest API.
// A default retry strategy applies to this operation CancelWorkRequest()
func (client NetworkFirewallClient) CancelWorkRequest(ctx context.Context, request CancelWorkRequestRequest) (response CancelWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.cancelWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelWorkRequestResponse")
	}
	return
}

// cancelWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) cancelWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/WorkRequest/CancelWorkRequest"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "CancelWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeNetworkFirewallCompartment Moves a NetworkFirewall resource from one compartment identifier to another. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/ChangeNetworkFirewallCompartment.go.html to see an example of how to use ChangeNetworkFirewallCompartment API.
// A default retry strategy applies to this operation ChangeNetworkFirewallCompartment()
func (client NetworkFirewallClient) ChangeNetworkFirewallCompartment(ctx context.Context, request ChangeNetworkFirewallCompartmentRequest) (response ChangeNetworkFirewallCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeNetworkFirewallCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeNetworkFirewallCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeNetworkFirewallCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeNetworkFirewallCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeNetworkFirewallCompartmentResponse")
	}
	return
}

// changeNetworkFirewallCompartment implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) changeNetworkFirewallCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkFirewalls/{networkFirewallId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeNetworkFirewallCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/NetworkFirewall/ChangeNetworkFirewallCompartment"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "ChangeNetworkFirewallCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeNetworkFirewallPolicyCompartment Moves a NetworkFirewallPolicy resource from one compartment identifier to another. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/ChangeNetworkFirewallPolicyCompartment.go.html to see an example of how to use ChangeNetworkFirewallPolicyCompartment API.
// A default retry strategy applies to this operation ChangeNetworkFirewallPolicyCompartment()
func (client NetworkFirewallClient) ChangeNetworkFirewallPolicyCompartment(ctx context.Context, request ChangeNetworkFirewallPolicyCompartmentRequest) (response ChangeNetworkFirewallPolicyCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeNetworkFirewallPolicyCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeNetworkFirewallPolicyCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeNetworkFirewallPolicyCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeNetworkFirewallPolicyCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeNetworkFirewallPolicyCompartmentResponse")
	}
	return
}

// changeNetworkFirewallPolicyCompartment implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) changeNetworkFirewallPolicyCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkFirewallPolicies/{networkFirewallPolicyId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeNetworkFirewallPolicyCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/NetworkFirewallPolicy/ChangeNetworkFirewallPolicyCompartment"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "ChangeNetworkFirewallPolicyCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CloneNetworkFirewallPolicy Moves a NetworkFirewallPolicy resource from one compartment identifier to another. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/CloneNetworkFirewallPolicy.go.html to see an example of how to use CloneNetworkFirewallPolicy API.
// A default retry strategy applies to this operation CloneNetworkFirewallPolicy()
func (client NetworkFirewallClient) CloneNetworkFirewallPolicy(ctx context.Context, request CloneNetworkFirewallPolicyRequest) (response CloneNetworkFirewallPolicyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.cloneNetworkFirewallPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CloneNetworkFirewallPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CloneNetworkFirewallPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CloneNetworkFirewallPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CloneNetworkFirewallPolicyResponse")
	}
	return
}

// cloneNetworkFirewallPolicy implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) cloneNetworkFirewallPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkFirewallPolicies/{networkFirewallPolicyId}/actions/clonePolicy", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CloneNetworkFirewallPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/NetworkFirewallPolicy/CloneNetworkFirewallPolicy"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "CloneNetworkFirewallPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateAddressList Creates a new Address List for the Network Firewall Policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/CreateAddressList.go.html to see an example of how to use CreateAddressList API.
// A default retry strategy applies to this operation CreateAddressList()
func (client NetworkFirewallClient) CreateAddressList(ctx context.Context, request CreateAddressListRequest) (response CreateAddressListResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createAddressList, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAddressListResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAddressListResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAddressListResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAddressListResponse")
	}
	return
}

// createAddressList implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) createAddressList(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkFirewallPolicies/{networkFirewallPolicyId}/addressLists", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateAddressListResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/AddressList/CreateAddressList"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "CreateAddressList", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateApplication Creates a new Application for the Network Firewall Policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/CreateApplication.go.html to see an example of how to use CreateApplication API.
// A default retry strategy applies to this operation CreateApplication()
func (client NetworkFirewallClient) CreateApplication(ctx context.Context, request CreateApplicationRequest) (response CreateApplicationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createApplication, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateApplicationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateApplicationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateApplicationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateApplicationResponse")
	}
	return
}

// createApplication implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) createApplication(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkFirewallPolicies/{networkFirewallPolicyId}/applications", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateApplicationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/Application/CreateApplication"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "CreateApplication", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &application{})
	return response, err
}

// CreateApplicationGroup Creates a new ApplicationGroup for the Network Firewall Policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/CreateApplicationGroup.go.html to see an example of how to use CreateApplicationGroup API.
// A default retry strategy applies to this operation CreateApplicationGroup()
func (client NetworkFirewallClient) CreateApplicationGroup(ctx context.Context, request CreateApplicationGroupRequest) (response CreateApplicationGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createApplicationGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateApplicationGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateApplicationGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateApplicationGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateApplicationGroupResponse")
	}
	return
}

// createApplicationGroup implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) createApplicationGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkFirewallPolicies/{networkFirewallPolicyId}/applicationGroups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateApplicationGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/ApplicationGroup/CreateApplicationGroup"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "CreateApplicationGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDecryptionProfile Creates a new Decryption Profile for the Network Firewall Policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/CreateDecryptionProfile.go.html to see an example of how to use CreateDecryptionProfile API.
// A default retry strategy applies to this operation CreateDecryptionProfile()
func (client NetworkFirewallClient) CreateDecryptionProfile(ctx context.Context, request CreateDecryptionProfileRequest) (response CreateDecryptionProfileResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDecryptionProfile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDecryptionProfileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDecryptionProfileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDecryptionProfileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDecryptionProfileResponse")
	}
	return
}

// createDecryptionProfile implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) createDecryptionProfile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkFirewallPolicies/{networkFirewallPolicyId}/decryptionProfiles", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDecryptionProfileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/DecryptionProfile/CreateDecryptionProfile"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "CreateDecryptionProfile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &decryptionprofile{})
	return response, err
}

// CreateDecryptionRule Creates a new Decryption Rule for the Network Firewall Policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/CreateDecryptionRule.go.html to see an example of how to use CreateDecryptionRule API.
// A default retry strategy applies to this operation CreateDecryptionRule()
func (client NetworkFirewallClient) CreateDecryptionRule(ctx context.Context, request CreateDecryptionRuleRequest) (response CreateDecryptionRuleResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDecryptionRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDecryptionRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDecryptionRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDecryptionRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDecryptionRuleResponse")
	}
	return
}

// createDecryptionRule implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) createDecryptionRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkFirewallPolicies/{networkFirewallPolicyId}/decryptionRules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDecryptionRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/DecryptionRule/CreateDecryptionRule"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "CreateDecryptionRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateMappedSecret Creates a new Mapped Secret for the Network Firewall Policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/CreateMappedSecret.go.html to see an example of how to use CreateMappedSecret API.
// A default retry strategy applies to this operation CreateMappedSecret()
func (client NetworkFirewallClient) CreateMappedSecret(ctx context.Context, request CreateMappedSecretRequest) (response CreateMappedSecretResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMappedSecret, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMappedSecretResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMappedSecretResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMappedSecretResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMappedSecretResponse")
	}
	return
}

// createMappedSecret implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) createMappedSecret(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkFirewallPolicies/{networkFirewallPolicyId}/mappedSecrets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMappedSecretResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/MappedSecret/CreateMappedSecret"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "CreateMappedSecret", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &mappedsecret{})
	return response, err
}

// CreateNetworkFirewall Creates a new NetworkFirewall.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/CreateNetworkFirewall.go.html to see an example of how to use CreateNetworkFirewall API.
// A default retry strategy applies to this operation CreateNetworkFirewall()
func (client NetworkFirewallClient) CreateNetworkFirewall(ctx context.Context, request CreateNetworkFirewallRequest) (response CreateNetworkFirewallResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createNetworkFirewall, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateNetworkFirewallResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateNetworkFirewallResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateNetworkFirewallResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateNetworkFirewallResponse")
	}
	return
}

// createNetworkFirewall implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) createNetworkFirewall(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkFirewalls", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateNetworkFirewallResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/NetworkFirewall/CreateNetworkFirewall"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "CreateNetworkFirewall", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateNetworkFirewallPolicy Creates a new Network Firewall Policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/CreateNetworkFirewallPolicy.go.html to see an example of how to use CreateNetworkFirewallPolicy API.
// A default retry strategy applies to this operation CreateNetworkFirewallPolicy()
func (client NetworkFirewallClient) CreateNetworkFirewallPolicy(ctx context.Context, request CreateNetworkFirewallPolicyRequest) (response CreateNetworkFirewallPolicyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createNetworkFirewallPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateNetworkFirewallPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateNetworkFirewallPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateNetworkFirewallPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateNetworkFirewallPolicyResponse")
	}
	return
}

// createNetworkFirewallPolicy implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) createNetworkFirewallPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkFirewallPolicies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateNetworkFirewallPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/NetworkFirewallPolicy/CreateNetworkFirewallPolicy"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "CreateNetworkFirewallPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSecurityRule Creates a new Security Rule for the Network Firewall Policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/CreateSecurityRule.go.html to see an example of how to use CreateSecurityRule API.
// A default retry strategy applies to this operation CreateSecurityRule()
func (client NetworkFirewallClient) CreateSecurityRule(ctx context.Context, request CreateSecurityRuleRequest) (response CreateSecurityRuleResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createSecurityRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSecurityRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSecurityRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSecurityRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSecurityRuleResponse")
	}
	return
}

// createSecurityRule implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) createSecurityRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkFirewallPolicies/{networkFirewallPolicyId}/securityRules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSecurityRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/SecurityRule/CreateSecurityRule"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "CreateSecurityRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateService Creates a new Service for the Network Firewall Policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/CreateService.go.html to see an example of how to use CreateService API.
// A default retry strategy applies to this operation CreateService()
func (client NetworkFirewallClient) CreateService(ctx context.Context, request CreateServiceRequest) (response CreateServiceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createService, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateServiceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateServiceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateServiceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateServiceResponse")
	}
	return
}

// createService implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) createService(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkFirewallPolicies/{networkFirewallPolicyId}/services", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateServiceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/Service/CreateService"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "CreateService", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &service{})
	return response, err
}

// CreateServiceList Creates a new ServiceList for the Network Firewall Policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/CreateServiceList.go.html to see an example of how to use CreateServiceList API.
// A default retry strategy applies to this operation CreateServiceList()
func (client NetworkFirewallClient) CreateServiceList(ctx context.Context, request CreateServiceListRequest) (response CreateServiceListResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createServiceList, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateServiceListResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateServiceListResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateServiceListResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateServiceListResponse")
	}
	return
}

// createServiceList implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) createServiceList(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkFirewallPolicies/{networkFirewallPolicyId}/serviceLists", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateServiceListResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/ServiceList/CreateServiceList"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "CreateServiceList", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateUrlList Creates a new Url List for the Network Firewall Policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/CreateUrlList.go.html to see an example of how to use CreateUrlList API.
// A default retry strategy applies to this operation CreateUrlList()
func (client NetworkFirewallClient) CreateUrlList(ctx context.Context, request CreateUrlListRequest) (response CreateUrlListResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createUrlList, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateUrlListResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateUrlListResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateUrlListResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateUrlListResponse")
	}
	return
}

// createUrlList implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) createUrlList(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkFirewallPolicies/{networkFirewallPolicyId}/urlLists", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateUrlListResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/UrlList/CreateUrlList"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "CreateUrlList", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteAddressList Deletes a Address List resource with the given identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/DeleteAddressList.go.html to see an example of how to use DeleteAddressList API.
// A default retry strategy applies to this operation DeleteAddressList()
func (client NetworkFirewallClient) DeleteAddressList(ctx context.Context, request DeleteAddressListRequest) (response DeleteAddressListResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAddressList, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAddressListResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAddressListResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAddressListResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAddressListResponse")
	}
	return
}

// deleteAddressList implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) deleteAddressList(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/networkFirewallPolicies/{networkFirewallPolicyId}/addressLists/{addressListName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteAddressListResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/AddressList/DeleteAddressList"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "DeleteAddressList", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteApplication Deletes a Application resource with the given identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/DeleteApplication.go.html to see an example of how to use DeleteApplication API.
// A default retry strategy applies to this operation DeleteApplication()
func (client NetworkFirewallClient) DeleteApplication(ctx context.Context, request DeleteApplicationRequest) (response DeleteApplicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteApplication, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteApplicationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteApplicationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteApplicationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteApplicationResponse")
	}
	return
}

// deleteApplication implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) deleteApplication(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/networkFirewallPolicies/{networkFirewallPolicyId}/applications/{applicationName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteApplicationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/Application/DeleteApplication"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "DeleteApplication", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteApplicationGroup Deletes a ApplicationGroup resource with the given identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/DeleteApplicationGroup.go.html to see an example of how to use DeleteApplicationGroup API.
// A default retry strategy applies to this operation DeleteApplicationGroup()
func (client NetworkFirewallClient) DeleteApplicationGroup(ctx context.Context, request DeleteApplicationGroupRequest) (response DeleteApplicationGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteApplicationGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteApplicationGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteApplicationGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteApplicationGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteApplicationGroupResponse")
	}
	return
}

// deleteApplicationGroup implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) deleteApplicationGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/networkFirewallPolicies/{networkFirewallPolicyId}/applicationGroups/{applicationGroupName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteApplicationGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/ApplicationGroup/DeleteApplicationGroup"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "DeleteApplicationGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDecryptionProfile Deletes a Decryption Profile resource with the given identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/DeleteDecryptionProfile.go.html to see an example of how to use DeleteDecryptionProfile API.
// A default retry strategy applies to this operation DeleteDecryptionProfile()
func (client NetworkFirewallClient) DeleteDecryptionProfile(ctx context.Context, request DeleteDecryptionProfileRequest) (response DeleteDecryptionProfileResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDecryptionProfile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDecryptionProfileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDecryptionProfileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDecryptionProfileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDecryptionProfileResponse")
	}
	return
}

// deleteDecryptionProfile implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) deleteDecryptionProfile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/networkFirewallPolicies/{networkFirewallPolicyId}/decryptionProfiles/{decryptionProfileName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDecryptionProfileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/DecryptionProfile/DeleteDecryptionProfile"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "DeleteDecryptionProfile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDecryptionRule Deletes a Decryption Rule resource with the given identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/DeleteDecryptionRule.go.html to see an example of how to use DeleteDecryptionRule API.
// A default retry strategy applies to this operation DeleteDecryptionRule()
func (client NetworkFirewallClient) DeleteDecryptionRule(ctx context.Context, request DeleteDecryptionRuleRequest) (response DeleteDecryptionRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDecryptionRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDecryptionRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDecryptionRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDecryptionRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDecryptionRuleResponse")
	}
	return
}

// deleteDecryptionRule implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) deleteDecryptionRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/networkFirewallPolicies/{networkFirewallPolicyId}/decryptionRules/{decryptionRuleName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDecryptionRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/DecryptionRule/DeleteDecryptionRule"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "DeleteDecryptionRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteMappedSecret Deletes a Mapped Secret resource with the given identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/DeleteMappedSecret.go.html to see an example of how to use DeleteMappedSecret API.
// A default retry strategy applies to this operation DeleteMappedSecret()
func (client NetworkFirewallClient) DeleteMappedSecret(ctx context.Context, request DeleteMappedSecretRequest) (response DeleteMappedSecretResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteMappedSecret, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMappedSecretResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMappedSecretResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMappedSecretResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMappedSecretResponse")
	}
	return
}

// deleteMappedSecret implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) deleteMappedSecret(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/networkFirewallPolicies/{networkFirewallPolicyId}/mappedSecrets/{mappedSecretName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMappedSecretResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/MappedSecret/DeleteMappedSecret"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "DeleteMappedSecret", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteNetworkFirewall Deletes a NetworkFirewall resource by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/DeleteNetworkFirewall.go.html to see an example of how to use DeleteNetworkFirewall API.
// A default retry strategy applies to this operation DeleteNetworkFirewall()
func (client NetworkFirewallClient) DeleteNetworkFirewall(ctx context.Context, request DeleteNetworkFirewallRequest) (response DeleteNetworkFirewallResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteNetworkFirewall, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteNetworkFirewallResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteNetworkFirewallResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteNetworkFirewallResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteNetworkFirewallResponse")
	}
	return
}

// deleteNetworkFirewall implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) deleteNetworkFirewall(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/networkFirewalls/{networkFirewallId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteNetworkFirewallResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/NetworkFirewall/DeleteNetworkFirewall"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "DeleteNetworkFirewall", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteNetworkFirewallPolicy Deletes a NetworkFirewallPolicy resource with the given identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/DeleteNetworkFirewallPolicy.go.html to see an example of how to use DeleteNetworkFirewallPolicy API.
// A default retry strategy applies to this operation DeleteNetworkFirewallPolicy()
func (client NetworkFirewallClient) DeleteNetworkFirewallPolicy(ctx context.Context, request DeleteNetworkFirewallPolicyRequest) (response DeleteNetworkFirewallPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteNetworkFirewallPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteNetworkFirewallPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteNetworkFirewallPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteNetworkFirewallPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteNetworkFirewallPolicyResponse")
	}
	return
}

// deleteNetworkFirewallPolicy implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) deleteNetworkFirewallPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/networkFirewallPolicies/{networkFirewallPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteNetworkFirewallPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/NetworkFirewallPolicy/DeleteNetworkFirewallPolicy"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "DeleteNetworkFirewallPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSecurityRule Deletes a Security Rule resource with the given identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/DeleteSecurityRule.go.html to see an example of how to use DeleteSecurityRule API.
// A default retry strategy applies to this operation DeleteSecurityRule()
func (client NetworkFirewallClient) DeleteSecurityRule(ctx context.Context, request DeleteSecurityRuleRequest) (response DeleteSecurityRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSecurityRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSecurityRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSecurityRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSecurityRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSecurityRuleResponse")
	}
	return
}

// deleteSecurityRule implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) deleteSecurityRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/networkFirewallPolicies/{networkFirewallPolicyId}/securityRules/{securityRuleName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSecurityRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/SecurityRule/DeleteSecurityRule"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "DeleteSecurityRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteService Deletes a Service resource with the given identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/DeleteService.go.html to see an example of how to use DeleteService API.
// A default retry strategy applies to this operation DeleteService()
func (client NetworkFirewallClient) DeleteService(ctx context.Context, request DeleteServiceRequest) (response DeleteServiceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteService, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteServiceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteServiceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteServiceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteServiceResponse")
	}
	return
}

// deleteService implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) deleteService(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/networkFirewallPolicies/{networkFirewallPolicyId}/services/{serviceName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteServiceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/Service/DeleteService"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "DeleteService", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteServiceList Deletes a ServiceList resource with the given identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/DeleteServiceList.go.html to see an example of how to use DeleteServiceList API.
// A default retry strategy applies to this operation DeleteServiceList()
func (client NetworkFirewallClient) DeleteServiceList(ctx context.Context, request DeleteServiceListRequest) (response DeleteServiceListResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteServiceList, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteServiceListResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteServiceListResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteServiceListResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteServiceListResponse")
	}
	return
}

// deleteServiceList implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) deleteServiceList(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/networkFirewallPolicies/{networkFirewallPolicyId}/serviceLists/{serviceListName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteServiceListResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/ServiceList/DeleteServiceList"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "DeleteServiceList", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteUrlList Deletes a Url List resource with the given identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/DeleteUrlList.go.html to see an example of how to use DeleteUrlList API.
// A default retry strategy applies to this operation DeleteUrlList()
func (client NetworkFirewallClient) DeleteUrlList(ctx context.Context, request DeleteUrlListRequest) (response DeleteUrlListResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteUrlList, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteUrlListResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteUrlListResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteUrlListResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteUrlListResponse")
	}
	return
}

// deleteUrlList implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) deleteUrlList(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/networkFirewallPolicies/{networkFirewallPolicyId}/urlLists/{urlListName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteUrlListResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/UrlList/DeleteUrlList"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "DeleteUrlList", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAddressList Get Address List by the given name in the context of network firewall policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/GetAddressList.go.html to see an example of how to use GetAddressList API.
// A default retry strategy applies to this operation GetAddressList()
func (client NetworkFirewallClient) GetAddressList(ctx context.Context, request GetAddressListRequest) (response GetAddressListResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAddressList, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAddressListResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAddressListResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAddressListResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAddressListResponse")
	}
	return
}

// getAddressList implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) getAddressList(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkFirewallPolicies/{networkFirewallPolicyId}/addressLists/{addressListName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAddressListResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/AddressList/GetAddressList"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "GetAddressList", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetApplication Get Application by the given name in the context of network firewall policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/GetApplication.go.html to see an example of how to use GetApplication API.
// A default retry strategy applies to this operation GetApplication()
func (client NetworkFirewallClient) GetApplication(ctx context.Context, request GetApplicationRequest) (response GetApplicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getApplication, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetApplicationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetApplicationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetApplicationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetApplicationResponse")
	}
	return
}

// getApplication implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) getApplication(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkFirewallPolicies/{networkFirewallPolicyId}/applications/{applicationName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetApplicationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/Application/GetApplication"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "GetApplication", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &application{})
	return response, err
}

// GetApplicationGroup Get ApplicationGroup by the given name in the context of network firewall policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/GetApplicationGroup.go.html to see an example of how to use GetApplicationGroup API.
// A default retry strategy applies to this operation GetApplicationGroup()
func (client NetworkFirewallClient) GetApplicationGroup(ctx context.Context, request GetApplicationGroupRequest) (response GetApplicationGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getApplicationGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetApplicationGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetApplicationGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetApplicationGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetApplicationGroupResponse")
	}
	return
}

// getApplicationGroup implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) getApplicationGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkFirewallPolicies/{networkFirewallPolicyId}/applicationGroups/{applicationGroupName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetApplicationGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/ApplicationGroup/GetApplicationGroup"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "GetApplicationGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDecryptionProfile Get Decryption Profile by the given name in the context of network firewall policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/GetDecryptionProfile.go.html to see an example of how to use GetDecryptionProfile API.
// A default retry strategy applies to this operation GetDecryptionProfile()
func (client NetworkFirewallClient) GetDecryptionProfile(ctx context.Context, request GetDecryptionProfileRequest) (response GetDecryptionProfileResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDecryptionProfile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDecryptionProfileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDecryptionProfileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDecryptionProfileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDecryptionProfileResponse")
	}
	return
}

// getDecryptionProfile implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) getDecryptionProfile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkFirewallPolicies/{networkFirewallPolicyId}/decryptionProfiles/{decryptionProfileName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDecryptionProfileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/DecryptionProfile/GetDecryptionProfile"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "GetDecryptionProfile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &decryptionprofile{})
	return response, err
}

// GetDecryptionRule Get Decryption Rule by the given name in the context of network firewall policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/GetDecryptionRule.go.html to see an example of how to use GetDecryptionRule API.
// A default retry strategy applies to this operation GetDecryptionRule()
func (client NetworkFirewallClient) GetDecryptionRule(ctx context.Context, request GetDecryptionRuleRequest) (response GetDecryptionRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDecryptionRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDecryptionRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDecryptionRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDecryptionRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDecryptionRuleResponse")
	}
	return
}

// getDecryptionRule implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) getDecryptionRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkFirewallPolicies/{networkFirewallPolicyId}/decryptionRules/{decryptionRuleName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDecryptionRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/DecryptionRule/GetDecryptionRule"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "GetDecryptionRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetMappedSecret Get Mapped Secret by the given name in the context of network firewall policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/GetMappedSecret.go.html to see an example of how to use GetMappedSecret API.
// A default retry strategy applies to this operation GetMappedSecret()
func (client NetworkFirewallClient) GetMappedSecret(ctx context.Context, request GetMappedSecretRequest) (response GetMappedSecretResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMappedSecret, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMappedSecretResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMappedSecretResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMappedSecretResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMappedSecretResponse")
	}
	return
}

// getMappedSecret implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) getMappedSecret(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkFirewallPolicies/{networkFirewallPolicyId}/mappedSecrets/{mappedSecretName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMappedSecretResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/MappedSecret/GetMappedSecret"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "GetMappedSecret", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &mappedsecret{})
	return response, err
}

// GetNetworkFirewall Gets a NetworkFirewall by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/GetNetworkFirewall.go.html to see an example of how to use GetNetworkFirewall API.
// A default retry strategy applies to this operation GetNetworkFirewall()
func (client NetworkFirewallClient) GetNetworkFirewall(ctx context.Context, request GetNetworkFirewallRequest) (response GetNetworkFirewallResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getNetworkFirewall, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetNetworkFirewallResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetNetworkFirewallResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetNetworkFirewallResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetNetworkFirewallResponse")
	}
	return
}

// getNetworkFirewall implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) getNetworkFirewall(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkFirewalls/{networkFirewallId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetNetworkFirewallResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/NetworkFirewall/GetNetworkFirewall"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "GetNetworkFirewall", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetNetworkFirewallPolicy Gets a NetworkFirewallPolicy given the network firewall policy identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/GetNetworkFirewallPolicy.go.html to see an example of how to use GetNetworkFirewallPolicy API.
// A default retry strategy applies to this operation GetNetworkFirewallPolicy()
func (client NetworkFirewallClient) GetNetworkFirewallPolicy(ctx context.Context, request GetNetworkFirewallPolicyRequest) (response GetNetworkFirewallPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getNetworkFirewallPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetNetworkFirewallPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetNetworkFirewallPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetNetworkFirewallPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetNetworkFirewallPolicyResponse")
	}
	return
}

// getNetworkFirewallPolicy implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) getNetworkFirewallPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkFirewallPolicies/{networkFirewallPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetNetworkFirewallPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/NetworkFirewallPolicy/GetNetworkFirewallPolicy"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "GetNetworkFirewallPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSecurityRule Get Security Rule by the given name in the context of network firewall policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/GetSecurityRule.go.html to see an example of how to use GetSecurityRule API.
// A default retry strategy applies to this operation GetSecurityRule()
func (client NetworkFirewallClient) GetSecurityRule(ctx context.Context, request GetSecurityRuleRequest) (response GetSecurityRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSecurityRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSecurityRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSecurityRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSecurityRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSecurityRuleResponse")
	}
	return
}

// getSecurityRule implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) getSecurityRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkFirewallPolicies/{networkFirewallPolicyId}/securityRules/{securityRuleName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSecurityRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/SecurityRule/GetSecurityRule"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "GetSecurityRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetService Get Service by the given name in the context of network firewall policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/GetService.go.html to see an example of how to use GetService API.
// A default retry strategy applies to this operation GetService()
func (client NetworkFirewallClient) GetService(ctx context.Context, request GetServiceRequest) (response GetServiceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getService, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetServiceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetServiceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetServiceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetServiceResponse")
	}
	return
}

// getService implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) getService(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkFirewallPolicies/{networkFirewallPolicyId}/services/{serviceName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetServiceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/Service/GetService"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "GetService", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &service{})
	return response, err
}

// GetServiceList Get ServiceList by the given name in the context of network firewall policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/GetServiceList.go.html to see an example of how to use GetServiceList API.
// A default retry strategy applies to this operation GetServiceList()
func (client NetworkFirewallClient) GetServiceList(ctx context.Context, request GetServiceListRequest) (response GetServiceListResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getServiceList, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetServiceListResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetServiceListResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetServiceListResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetServiceListResponse")
	}
	return
}

// getServiceList implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) getServiceList(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkFirewallPolicies/{networkFirewallPolicyId}/serviceLists/{serviceListName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetServiceListResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/ServiceList/GetServiceList"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "GetServiceList", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetUrlList Get Url List by the given name in the context of network firewall policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/GetUrlList.go.html to see an example of how to use GetUrlList API.
// A default retry strategy applies to this operation GetUrlList()
func (client NetworkFirewallClient) GetUrlList(ctx context.Context, request GetUrlListRequest) (response GetUrlListResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getUrlList, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetUrlListResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetUrlListResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetUrlListResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetUrlListResponse")
	}
	return
}

// getUrlList implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) getUrlList(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkFirewallPolicies/{networkFirewallPolicyId}/urlLists/{urlListName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetUrlListResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/UrlList/GetUrlList"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "GetUrlList", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets the status of the work request with the given ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
// A default retry strategy applies to this operation GetWorkRequest()
func (client NetworkFirewallClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
func (client NetworkFirewallClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAddressLists Returns a list of Network Firewall Policies.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/ListAddressLists.go.html to see an example of how to use ListAddressLists API.
// A default retry strategy applies to this operation ListAddressLists()
func (client NetworkFirewallClient) ListAddressLists(ctx context.Context, request ListAddressListsRequest) (response ListAddressListsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAddressLists, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAddressListsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAddressListsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAddressListsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAddressListsResponse")
	}
	return
}

// listAddressLists implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) listAddressLists(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkFirewallPolicies/{networkFirewallPolicyId}/addressLists", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAddressListsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/AddressList/ListAddressLists"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "ListAddressLists", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListApplicationGroups Returns a list of ApplicationGroups for the policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/ListApplicationGroups.go.html to see an example of how to use ListApplicationGroups API.
// A default retry strategy applies to this operation ListApplicationGroups()
func (client NetworkFirewallClient) ListApplicationGroups(ctx context.Context, request ListApplicationGroupsRequest) (response ListApplicationGroupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listApplicationGroups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListApplicationGroupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListApplicationGroupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListApplicationGroupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListApplicationGroupsResponse")
	}
	return
}

// listApplicationGroups implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) listApplicationGroups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkFirewallPolicies/{networkFirewallPolicyId}/applicationGroups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListApplicationGroupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/ApplicationGroup/ListApplicationGroups"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "ListApplicationGroups", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListApplications Returns a list of Applications for the policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/ListApplications.go.html to see an example of how to use ListApplications API.
// A default retry strategy applies to this operation ListApplications()
func (client NetworkFirewallClient) ListApplications(ctx context.Context, request ListApplicationsRequest) (response ListApplicationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
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
func (client NetworkFirewallClient) listApplications(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkFirewallPolicies/{networkFirewallPolicyId}/applications", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListApplicationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/Application/ListApplications"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "ListApplications", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDecryptionProfiles Returns a list of Decryption Profile for the Network Firewall Policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/ListDecryptionProfiles.go.html to see an example of how to use ListDecryptionProfiles API.
// A default retry strategy applies to this operation ListDecryptionProfiles()
func (client NetworkFirewallClient) ListDecryptionProfiles(ctx context.Context, request ListDecryptionProfilesRequest) (response ListDecryptionProfilesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDecryptionProfiles, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDecryptionProfilesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDecryptionProfilesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDecryptionProfilesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDecryptionProfilesResponse")
	}
	return
}

// listDecryptionProfiles implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) listDecryptionProfiles(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkFirewallPolicies/{networkFirewallPolicyId}/decryptionProfiles", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDecryptionProfilesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/DecryptionProfile/ListDecryptionProfiles"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "ListDecryptionProfiles", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDecryptionRules Returns a list of Decryption Rule for the Network Firewall Policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/ListDecryptionRules.go.html to see an example of how to use ListDecryptionRules API.
// A default retry strategy applies to this operation ListDecryptionRules()
func (client NetworkFirewallClient) ListDecryptionRules(ctx context.Context, request ListDecryptionRulesRequest) (response ListDecryptionRulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDecryptionRules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDecryptionRulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDecryptionRulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDecryptionRulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDecryptionRulesResponse")
	}
	return
}

// listDecryptionRules implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) listDecryptionRules(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkFirewallPolicies/{networkFirewallPolicyId}/decryptionRules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDecryptionRulesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/DecryptionRule/ListDecryptionRules"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "ListDecryptionRules", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMappedSecrets Returns a list of Mapped Secret for the Network Firewall Policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/ListMappedSecrets.go.html to see an example of how to use ListMappedSecrets API.
// A default retry strategy applies to this operation ListMappedSecrets()
func (client NetworkFirewallClient) ListMappedSecrets(ctx context.Context, request ListMappedSecretsRequest) (response ListMappedSecretsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMappedSecrets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMappedSecretsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMappedSecretsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMappedSecretsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMappedSecretsResponse")
	}
	return
}

// listMappedSecrets implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) listMappedSecrets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkFirewallPolicies/{networkFirewallPolicyId}/mappedSecrets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMappedSecretsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/MappedSecret/ListMappedSecrets"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "ListMappedSecrets", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListNetworkFirewallPolicies Returns a list of Network Firewall Policies.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/ListNetworkFirewallPolicies.go.html to see an example of how to use ListNetworkFirewallPolicies API.
// A default retry strategy applies to this operation ListNetworkFirewallPolicies()
func (client NetworkFirewallClient) ListNetworkFirewallPolicies(ctx context.Context, request ListNetworkFirewallPoliciesRequest) (response ListNetworkFirewallPoliciesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listNetworkFirewallPolicies, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListNetworkFirewallPoliciesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListNetworkFirewallPoliciesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListNetworkFirewallPoliciesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListNetworkFirewallPoliciesResponse")
	}
	return
}

// listNetworkFirewallPolicies implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) listNetworkFirewallPolicies(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkFirewallPolicies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListNetworkFirewallPoliciesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/NetworkFirewallPolicy/ListNetworkFirewallPolicies"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "ListNetworkFirewallPolicies", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListNetworkFirewalls Returns a list of NetworkFirewalls.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/ListNetworkFirewalls.go.html to see an example of how to use ListNetworkFirewalls API.
// A default retry strategy applies to this operation ListNetworkFirewalls()
func (client NetworkFirewallClient) ListNetworkFirewalls(ctx context.Context, request ListNetworkFirewallsRequest) (response ListNetworkFirewallsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listNetworkFirewalls, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListNetworkFirewallsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListNetworkFirewallsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListNetworkFirewallsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListNetworkFirewallsResponse")
	}
	return
}

// listNetworkFirewalls implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) listNetworkFirewalls(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkFirewalls", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListNetworkFirewallsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/NetworkFirewall/ListNetworkFirewalls"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "ListNetworkFirewalls", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSecurityRules Returns a list of Security Rule for the Network Firewall Policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/ListSecurityRules.go.html to see an example of how to use ListSecurityRules API.
// A default retry strategy applies to this operation ListSecurityRules()
func (client NetworkFirewallClient) ListSecurityRules(ctx context.Context, request ListSecurityRulesRequest) (response ListSecurityRulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSecurityRules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSecurityRulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSecurityRulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSecurityRulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSecurityRulesResponse")
	}
	return
}

// listSecurityRules implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) listSecurityRules(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkFirewallPolicies/{networkFirewallPolicyId}/securityRules", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSecurityRulesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/SecurityRule/ListSecurityRules"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "ListSecurityRules", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListServiceLists Returns a list of ServiceLists for the policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/ListServiceLists.go.html to see an example of how to use ListServiceLists API.
// A default retry strategy applies to this operation ListServiceLists()
func (client NetworkFirewallClient) ListServiceLists(ctx context.Context, request ListServiceListsRequest) (response ListServiceListsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listServiceLists, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListServiceListsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListServiceListsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListServiceListsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListServiceListsResponse")
	}
	return
}

// listServiceLists implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) listServiceLists(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkFirewallPolicies/{networkFirewallPolicyId}/serviceLists", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListServiceListsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/ServiceList/ListServiceLists"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "ListServiceLists", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListServices Returns a list of Services for the policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/ListServices.go.html to see an example of how to use ListServices API.
// A default retry strategy applies to this operation ListServices()
func (client NetworkFirewallClient) ListServices(ctx context.Context, request ListServicesRequest) (response ListServicesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
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
func (client NetworkFirewallClient) listServices(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkFirewallPolicies/{networkFirewallPolicyId}/services", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListServicesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/Service/ListServices"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "ListServices", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListUrlLists Returns a list of URL lists for the Network Firewall Policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/ListUrlLists.go.html to see an example of how to use ListUrlLists API.
// A default retry strategy applies to this operation ListUrlLists()
func (client NetworkFirewallClient) ListUrlLists(ctx context.Context, request ListUrlListsRequest) (response ListUrlListsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listUrlLists, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListUrlListsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListUrlListsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListUrlListsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListUrlListsResponse")
	}
	return
}

// listUrlLists implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) listUrlLists(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkFirewallPolicies/{networkFirewallPolicyId}/urlLists", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListUrlListsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/UrlList/ListUrlLists"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "ListUrlLists", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Return a (paginated) list of errors for a given work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrors API.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client NetworkFirewallClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
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
func (client NetworkFirewallClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/WorkRequestError/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Return a (paginated) list of logs for a given work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogs API.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client NetworkFirewallClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
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
func (client NetworkFirewallClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/WorkRequestLogEntry/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
// A default retry strategy applies to this operation ListWorkRequests()
func (client NetworkFirewallClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
func (client NetworkFirewallClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/WorkRequest/ListWorkRequests"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// MigrateNetworkFirewallPolicy Moves a NetworkFirewallPolicy resource from one version to latest version. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/MigrateNetworkFirewallPolicy.go.html to see an example of how to use MigrateNetworkFirewallPolicy API.
// A default retry strategy applies to this operation MigrateNetworkFirewallPolicy()
func (client NetworkFirewallClient) MigrateNetworkFirewallPolicy(ctx context.Context, request MigrateNetworkFirewallPolicyRequest) (response MigrateNetworkFirewallPolicyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.migrateNetworkFirewallPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = MigrateNetworkFirewallPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = MigrateNetworkFirewallPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(MigrateNetworkFirewallPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into MigrateNetworkFirewallPolicyResponse")
	}
	return
}

// migrateNetworkFirewallPolicy implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) migrateNetworkFirewallPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkFirewallPolicies/{networkFirewallPolicyId}/actions/migrate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response MigrateNetworkFirewallPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/NetworkFirewallPolicy/MigrateNetworkFirewallPolicy"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "MigrateNetworkFirewallPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAddressList Updates the Address list with the given name in the network firewall policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/UpdateAddressList.go.html to see an example of how to use UpdateAddressList API.
// A default retry strategy applies to this operation UpdateAddressList()
func (client NetworkFirewallClient) UpdateAddressList(ctx context.Context, request UpdateAddressListRequest) (response UpdateAddressListResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAddressList, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAddressListResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAddressListResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAddressListResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAddressListResponse")
	}
	return
}

// updateAddressList implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) updateAddressList(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/networkFirewallPolicies/{networkFirewallPolicyId}/addressLists/{addressListName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAddressListResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/AddressList/UpdateAddressList"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "UpdateAddressList", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateApplication Updates the Application with the given name in the network firewall policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/UpdateApplication.go.html to see an example of how to use UpdateApplication API.
// A default retry strategy applies to this operation UpdateApplication()
func (client NetworkFirewallClient) UpdateApplication(ctx context.Context, request UpdateApplicationRequest) (response UpdateApplicationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateApplication, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateApplicationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateApplicationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateApplicationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateApplicationResponse")
	}
	return
}

// updateApplication implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) updateApplication(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/networkFirewallPolicies/{networkFirewallPolicyId}/applications/{applicationName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateApplicationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/Application/UpdateApplication"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "UpdateApplication", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &application{})
	return response, err
}

// UpdateApplicationGroup Updates the ApplicationGroup with the given name in the network firewall policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/UpdateApplicationGroup.go.html to see an example of how to use UpdateApplicationGroup API.
// A default retry strategy applies to this operation UpdateApplicationGroup()
func (client NetworkFirewallClient) UpdateApplicationGroup(ctx context.Context, request UpdateApplicationGroupRequest) (response UpdateApplicationGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateApplicationGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateApplicationGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateApplicationGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateApplicationGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateApplicationGroupResponse")
	}
	return
}

// updateApplicationGroup implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) updateApplicationGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/networkFirewallPolicies/{networkFirewallPolicyId}/applicationGroups/{applicationGroupName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateApplicationGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/ApplicationGroup/UpdateApplicationGroup"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "UpdateApplicationGroup", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDecryptionProfile Updates the Decryption Profile with the given name in the network firewall policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/UpdateDecryptionProfile.go.html to see an example of how to use UpdateDecryptionProfile API.
// A default retry strategy applies to this operation UpdateDecryptionProfile()
func (client NetworkFirewallClient) UpdateDecryptionProfile(ctx context.Context, request UpdateDecryptionProfileRequest) (response UpdateDecryptionProfileResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDecryptionProfile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDecryptionProfileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDecryptionProfileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDecryptionProfileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDecryptionProfileResponse")
	}
	return
}

// updateDecryptionProfile implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) updateDecryptionProfile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/networkFirewallPolicies/{networkFirewallPolicyId}/decryptionProfiles/{decryptionProfileName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDecryptionProfileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/DecryptionProfile/UpdateDecryptionProfile"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "UpdateDecryptionProfile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &decryptionprofile{})
	return response, err
}

// UpdateDecryptionRule Updates the Decryption Rule with the given name in the network firewall policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/UpdateDecryptionRule.go.html to see an example of how to use UpdateDecryptionRule API.
// A default retry strategy applies to this operation UpdateDecryptionRule()
func (client NetworkFirewallClient) UpdateDecryptionRule(ctx context.Context, request UpdateDecryptionRuleRequest) (response UpdateDecryptionRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDecryptionRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDecryptionRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDecryptionRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDecryptionRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDecryptionRuleResponse")
	}
	return
}

// updateDecryptionRule implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) updateDecryptionRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/networkFirewallPolicies/{networkFirewallPolicyId}/decryptionRules/{decryptionRuleName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDecryptionRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/DecryptionRule/UpdateDecryptionRule"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "UpdateDecryptionRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateMappedSecret Updates the Mapped Secret with the given name in the network firewall policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/UpdateMappedSecret.go.html to see an example of how to use UpdateMappedSecret API.
// A default retry strategy applies to this operation UpdateMappedSecret()
func (client NetworkFirewallClient) UpdateMappedSecret(ctx context.Context, request UpdateMappedSecretRequest) (response UpdateMappedSecretResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateMappedSecret, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateMappedSecretResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateMappedSecretResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateMappedSecretResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateMappedSecretResponse")
	}
	return
}

// updateMappedSecret implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) updateMappedSecret(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/networkFirewallPolicies/{networkFirewallPolicyId}/mappedSecrets/{mappedSecretName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateMappedSecretResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/MappedSecret/UpdateMappedSecret"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "UpdateMappedSecret", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &mappedsecret{})
	return response, err
}

// UpdateNetworkFirewall Updates the NetworkFirewall
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/UpdateNetworkFirewall.go.html to see an example of how to use UpdateNetworkFirewall API.
// A default retry strategy applies to this operation UpdateNetworkFirewall()
func (client NetworkFirewallClient) UpdateNetworkFirewall(ctx context.Context, request UpdateNetworkFirewallRequest) (response UpdateNetworkFirewallResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateNetworkFirewall, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateNetworkFirewallResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateNetworkFirewallResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateNetworkFirewallResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateNetworkFirewallResponse")
	}
	return
}

// updateNetworkFirewall implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) updateNetworkFirewall(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/networkFirewalls/{networkFirewallId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateNetworkFirewallResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/NetworkFirewall/UpdateNetworkFirewall"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "UpdateNetworkFirewall", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateNetworkFirewallPolicy Updates the NetworkFirewallPolicy
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/UpdateNetworkFirewallPolicy.go.html to see an example of how to use UpdateNetworkFirewallPolicy API.
// A default retry strategy applies to this operation UpdateNetworkFirewallPolicy()
func (client NetworkFirewallClient) UpdateNetworkFirewallPolicy(ctx context.Context, request UpdateNetworkFirewallPolicyRequest) (response UpdateNetworkFirewallPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateNetworkFirewallPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateNetworkFirewallPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateNetworkFirewallPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateNetworkFirewallPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateNetworkFirewallPolicyResponse")
	}
	return
}

// updateNetworkFirewallPolicy implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) updateNetworkFirewallPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/networkFirewallPolicies/{networkFirewallPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateNetworkFirewallPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/NetworkFirewallPolicy/UpdateNetworkFirewallPolicy"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "UpdateNetworkFirewallPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSecurityRule Updates the Security Rule with the given name in the network firewall policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/UpdateSecurityRule.go.html to see an example of how to use UpdateSecurityRule API.
// A default retry strategy applies to this operation UpdateSecurityRule()
func (client NetworkFirewallClient) UpdateSecurityRule(ctx context.Context, request UpdateSecurityRuleRequest) (response UpdateSecurityRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSecurityRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSecurityRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSecurityRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSecurityRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSecurityRuleResponse")
	}
	return
}

// updateSecurityRule implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) updateSecurityRule(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/networkFirewallPolicies/{networkFirewallPolicyId}/securityRules/{securityRuleName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSecurityRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/SecurityRule/UpdateSecurityRule"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "UpdateSecurityRule", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateService Updates the Service with the given name in the network firewall policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/UpdateService.go.html to see an example of how to use UpdateService API.
// A default retry strategy applies to this operation UpdateService()
func (client NetworkFirewallClient) UpdateService(ctx context.Context, request UpdateServiceRequest) (response UpdateServiceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateService, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateServiceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateServiceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateServiceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateServiceResponse")
	}
	return
}

// updateService implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) updateService(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/networkFirewallPolicies/{networkFirewallPolicyId}/services/{serviceName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateServiceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/Service/UpdateService"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "UpdateService", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &service{})
	return response, err
}

// UpdateServiceList Updates the ServiceList with the given name in the network firewall policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/UpdateServiceList.go.html to see an example of how to use UpdateServiceList API.
// A default retry strategy applies to this operation UpdateServiceList()
func (client NetworkFirewallClient) UpdateServiceList(ctx context.Context, request UpdateServiceListRequest) (response UpdateServiceListResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateServiceList, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateServiceListResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateServiceListResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateServiceListResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateServiceListResponse")
	}
	return
}

// updateServiceList implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) updateServiceList(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/networkFirewallPolicies/{networkFirewallPolicyId}/serviceLists/{serviceListName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateServiceListResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/ServiceList/UpdateServiceList"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "UpdateServiceList", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateUrlList Updates the Url list with the given name in the network firewall policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/UpdateUrlList.go.html to see an example of how to use UpdateUrlList API.
// A default retry strategy applies to this operation UpdateUrlList()
func (client NetworkFirewallClient) UpdateUrlList(ctx context.Context, request UpdateUrlListRequest) (response UpdateUrlListResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateUrlList, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateUrlListResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateUrlListResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateUrlListResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateUrlListResponse")
	}
	return
}

// updateUrlList implements the OCIOperation interface (enables retrying operations)
func (client NetworkFirewallClient) updateUrlList(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/networkFirewallPolicies/{networkFirewallPolicyId}/urlLists/{urlListName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateUrlListResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/network-firewall/20230501/UrlList/UpdateUrlList"
		err = common.PostProcessServiceError(err, "NetworkFirewall", "UpdateUrlList", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
