// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v50/common"
	"github.com/oracle/oci-go-sdk/v50/common/auth"
	"net/http"
	"time"
)

//IdentityClient a client for Identity
type IdentityClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewIdentityClientWithConfigurationProvider Creates a new default Identity client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewIdentityClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client IdentityClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newIdentityClientFromBaseClient(baseClient, provider)
}

// NewIdentityClientWithOboToken Creates a new default Identity client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewIdentityClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client IdentityClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newIdentityClientFromBaseClient(baseClient, configProvider)
}

func newIdentityClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client IdentityClient, err error) {
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = IdentityClient{BaseClient: baseClient}
	client.BasePath = "20160918"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *IdentityClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("identity", "https://identity.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *IdentityClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *IdentityClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ActivateDomain If the domain's {@code lifecycleState} is INACTIVE,
// 1. Set the {@code lifecycleDetails} to ACTIVATING and asynchronously starts enabling
//    the domain and return 202 ACCEPTED.
//     1.1 Sets the domain status to ENABLED and set specified domain's
//         {@code lifecycleState} to ACTIVE and set the {@code lifecycleDetails} to null.
// To track progress, HTTP GET on /iamWorkRequests/{iamWorkRequestsId} endpoint will provide
// the async operation's status. Deactivate a domain can be done using HTTP POST
// /domains/{domainId}/actions/deactivate.
// - If the domain's {@code lifecycleState} is ACTIVE, returns 202 ACCEPTED with no action
//   taken on service side.
// - If domain is of {@code type} DEFAULT or DEFAULT_LIGHTWEIGHT or domain's {@code lifecycleState} is not INACTIVE,
//   returns 400 BAD REQUEST.
// - If the domain doesn't exists, returns 404 NOT FOUND.
// - If the authenticated user is part of the domain to be activated, returns 400 BAD REQUEST
// - If error occurs while activating domain, returns 500 INTERNAL SERVER ERROR.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ActivateDomain.go.html to see an example of how to use ActivateDomain API.
func (client IdentityClient) ActivateDomain(ctx context.Context, request ActivateDomainRequest) (response ActivateDomainResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.activateDomain, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ActivateDomainResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ActivateDomainResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ActivateDomainResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ActivateDomainResponse")
	}
	return
}

// activateDomain implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) activateDomain(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/domains/{domainId}/actions/activate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ActivateDomainResponse
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

// ActivateMfaTotpDevice Activates the specified MFA TOTP device for the user. Activation requires manual interaction with the Console.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ActivateMfaTotpDevice.go.html to see an example of how to use ActivateMfaTotpDevice API.
func (client IdentityClient) ActivateMfaTotpDevice(ctx context.Context, request ActivateMfaTotpDeviceRequest) (response ActivateMfaTotpDeviceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.activateMfaTotpDevice, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ActivateMfaTotpDeviceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ActivateMfaTotpDeviceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ActivateMfaTotpDeviceResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ActivateMfaTotpDeviceResponse")
	}
	return
}

// activateMfaTotpDevice implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) activateMfaTotpDevice(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/users/{userId}/mfaTotpDevices/{mfaTotpDeviceId}/actions/activate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ActivateMfaTotpDeviceResponse
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

// AddUserToGroup Adds the specified user to the specified group and returns a `UserGroupMembership` object with its own OCID.
// After you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before using the
// object, first make sure its `lifecycleState` has changed to ACTIVE.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/AddUserToGroup.go.html to see an example of how to use AddUserToGroup API.
func (client IdentityClient) AddUserToGroup(ctx context.Context, request AddUserToGroupRequest) (response AddUserToGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.addUserToGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddUserToGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddUserToGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddUserToGroupResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddUserToGroupResponse")
	}
	return
}

// addUserToGroup implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) addUserToGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/userGroupMemberships", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddUserToGroupResponse
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

// AssembleEffectiveTagSet Assembles tag defaults in the specified compartment and any parent compartments to determine
// the tags to apply. Tag defaults from parent compartments do not override tag defaults
// referencing the same tag in a compartment lower down the hierarchy. This set of tag defaults
// includes all tag defaults from the current compartment back to the root compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/AssembleEffectiveTagSet.go.html to see an example of how to use AssembleEffectiveTagSet API.
func (client IdentityClient) AssembleEffectiveTagSet(ctx context.Context, request AssembleEffectiveTagSetRequest) (response AssembleEffectiveTagSetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.assembleEffectiveTagSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AssembleEffectiveTagSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AssembleEffectiveTagSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AssembleEffectiveTagSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AssembleEffectiveTagSetResponse")
	}
	return
}

// assembleEffectiveTagSet implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) assembleEffectiveTagSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/tagDefaults/actions/assembleEffectiveTagSet", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AssembleEffectiveTagSetResponse
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

// BulkDeleteResources Deletes multiple resources in the compartment. All resources must be in the same compartment. You must have the appropriate
// permissions to delete the resources in the request. This API can only be invoked from the tenancy's
// home region (https://docs.cloud.oracle.com/Content/Identity/Tasks/managingregions.htm#Home). This operation creates a
// WorkRequest. Use the GetWorkRequest
// API to monitor the status of the bulk action.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/BulkDeleteResources.go.html to see an example of how to use BulkDeleteResources API.
func (client IdentityClient) BulkDeleteResources(ctx context.Context, request BulkDeleteResourcesRequest) (response BulkDeleteResourcesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.bulkDeleteResources, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BulkDeleteResourcesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BulkDeleteResourcesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BulkDeleteResourcesResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BulkDeleteResourcesResponse")
	}
	return
}

// bulkDeleteResources implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) bulkDeleteResources(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/compartments/{compartmentId}/actions/bulkDeleteResources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BulkDeleteResourcesResponse
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

// BulkDeleteTags Deletes the specified tag key definitions. This operation triggers a process that removes the
// tags from all resources in your tenancy. The tag key definitions must be within the same tag namespace.
// The following actions happen immediately:
//
//   * If the tag is a cost-tracking tag, the tag no longer counts against your
//   10 cost-tracking tags limit, even if you do not disable the tag before running this operation.
//   * If the tag is used with dynamic groups, the rules that contain the tag are no longer
//   evaluated against the tag.
// After you start this operation, the state of the tag changes to DELETING, and tag removal
// from resources begins. This process can take up to 48 hours depending on the number of resources that
// are tagged and the regions in which those resources reside.
// When all tags have been removed, the state changes to DELETED. You cannot restore a deleted tag. After the tag state
// changes to DELETED, you can use the same tag name again.
// After you start this operation, you cannot start either the DeleteTag or the CascadeDeleteTagNamespace operation until this process completes.
// In order to delete tags, you must first retire the tags. Use UpdateTag
// to retire a tag.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/BulkDeleteTags.go.html to see an example of how to use BulkDeleteTags API.
func (client IdentityClient) BulkDeleteTags(ctx context.Context, request BulkDeleteTagsRequest) (response BulkDeleteTagsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.bulkDeleteTags, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BulkDeleteTagsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BulkDeleteTagsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BulkDeleteTagsResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BulkDeleteTagsResponse")
	}
	return
}

// bulkDeleteTags implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) bulkDeleteTags(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/tags/actions/bulkDelete", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BulkDeleteTagsResponse
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

// BulkEditTags Edits the specified list of tag key definitions for the selected resources.
// This operation triggers a process that edits the tags on all selected resources. The possible actions are:
//   * Add a defined tag when the tag does not already exist on the resource.
//   * Update the value for a defined tag when the tag is present on the resource.
//   * Add a defined tag when it does not already exist on the resource or update the value for a defined tag when the tag is present on the resource.
//   * Remove a defined tag from a resource. The tag is removed from the resource regardless of the tag value.
// See BulkEditOperationDetails for more information.
// The edits can include a combination of operations and tag sets.
// However, multiple operations cannot apply to one key definition in the same request.
// For example, if one request adds `tag set-1` to a resource and sets a tag value to `tag set-2`,
// `tag set-1` and `tag set-2` cannot have any common tag definitions.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/BulkEditTags.go.html to see an example of how to use BulkEditTags API.
func (client IdentityClient) BulkEditTags(ctx context.Context, request BulkEditTagsRequest) (response BulkEditTagsResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.bulkEditTags, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BulkEditTagsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BulkEditTagsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BulkEditTagsResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BulkEditTagsResponse")
	}
	return
}

// bulkEditTags implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) bulkEditTags(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/tags/actions/bulkEdit", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BulkEditTagsResponse
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

// BulkMoveResources Moves multiple resources from one compartment to another. All resources must be in the same compartment.
// This API can only be invoked from the tenancy's home region (https://docs.cloud.oracle.com/Content/Identity/Tasks/managingregions.htm#Home).
// To move resources, you must have the appropriate permissions to move the resource in both the source and target
// compartments. This operation creates a WorkRequest.
// Use the GetWorkRequest API to monitor the status of the bulk action.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/BulkMoveResources.go.html to see an example of how to use BulkMoveResources API.
func (client IdentityClient) BulkMoveResources(ctx context.Context, request BulkMoveResourcesRequest) (response BulkMoveResourcesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.bulkMoveResources, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BulkMoveResourcesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BulkMoveResourcesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BulkMoveResourcesResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BulkMoveResourcesResponse")
	}
	return
}

// bulkMoveResources implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) bulkMoveResources(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/compartments/{compartmentId}/actions/bulkMoveResources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BulkMoveResourcesResponse
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

// CascadeDeleteTagNamespace Deletes the specified tag namespace. This operation triggers a process that removes all of the tags
// defined in the specified tag namespace from all resources in your tenancy and then deletes the tag namespace.
// After you start the delete operation:
//   * New tag key definitions cannot be created under the namespace.
//   * The state of the tag namespace changes to DELETING.
//   * Tag removal from the resources begins.
// This process can take up to 48 hours depending on the number of tag definitions in the namespace, the number of resources
// that are tagged, and the locations of the regions in which those resources reside.
// After all tags are removed, the state changes to DELETED. You cannot restore a deleted tag namespace. After the deleted tag namespace
// changes its state to DELETED, you can use the name of the deleted tag namespace again.
// After you start this operation, you cannot start either the DeleteTag or the BulkDeleteTags operation until this process completes.
// To delete a tag namespace, you must first retire it. Use UpdateTagNamespace
// to retire a tag namespace.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/CascadeDeleteTagNamespace.go.html to see an example of how to use CascadeDeleteTagNamespace API.
func (client IdentityClient) CascadeDeleteTagNamespace(ctx context.Context, request CascadeDeleteTagNamespaceRequest) (response CascadeDeleteTagNamespaceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.cascadeDeleteTagNamespace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CascadeDeleteTagNamespaceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CascadeDeleteTagNamespaceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CascadeDeleteTagNamespaceResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CascadeDeleteTagNamespaceResponse")
	}
	return
}

// cascadeDeleteTagNamespace implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) cascadeDeleteTagNamespace(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/tagNamespaces/{tagNamespaceId}/actions/cascadeDelete", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CascadeDeleteTagNamespaceResponse
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

// ChangeDomainCompartment Change the containing compartment for a domain.
// This is an asynchronous call where the Domain's compartment is changed and is updated with the new compartment information.
// To track progress, HTTP GET on /iamWorkRequests/{iamWorkRequestsId} endpoint will provide
// the async operation's status.
// The compartment change is complete when accessed via domain URL and
// also returns new compartment OCID.
// - If the domain doesn't exists, returns 404 NOT FOUND.
// - If Domain {@code type} is DEFAULT or DEFAULT_LIGHTWEIGHT, return 400 BAD Request
// - If Domain is not active or being updated, returns 400 BAD REQUEST.
// - If error occurs while changing compartment for domain, return 500 INTERNAL SERVER ERROR.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ChangeDomainCompartment.go.html to see an example of how to use ChangeDomainCompartment API.
func (client IdentityClient) ChangeDomainCompartment(ctx context.Context, request ChangeDomainCompartmentRequest) (response ChangeDomainCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeDomainCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDomainCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDomainCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDomainCompartmentResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDomainCompartmentResponse")
	}
	return
}

// changeDomainCompartment implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) changeDomainCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/domains/{domainId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDomainCompartmentResponse
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

// ChangeDomainLicenseType If the domain's {@code lifecycleState} is ACTIVE, validates the requested {@code licenseType} update
// is allowed and
// 1. Set the {@code lifecycleDetails} to UPDATING
// 2. Asynchronously starts updating the domain and return 202 ACCEPTED.
//     2.1 Successfully updates specified domain's {@code licenseType}.
// 3. On completion set the {@code lifecycleDetails} to null.
// To track progress, HTTP GET on /iamWorkRequests/{iamWorkRequestsId} endpoint will provide
// the async operation's status.
// - If license type update is successful, return 202 ACCEPTED
// - If requested {@code licenseType} validation fails, returns 400 Bad request.
// - If Domain is not active or being updated, returns 400 BAD REQUEST.
// - If Domain {@code type} is DEFAULT or DEFAULT_LIGHTWEIGHT, return 400 BAD Request
// - If the domain doesn't exists, returns 404 NOT FOUND
// - If any internal error occurs, returns 500 INTERNAL SERVER ERROR.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ChangeDomainLicenseType.go.html to see an example of how to use ChangeDomainLicenseType API.
func (client IdentityClient) ChangeDomainLicenseType(ctx context.Context, request ChangeDomainLicenseTypeRequest) (response ChangeDomainLicenseTypeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeDomainLicenseType, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDomainLicenseTypeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDomainLicenseTypeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDomainLicenseTypeResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDomainLicenseTypeResponse")
	}
	return
}

// changeDomainLicenseType implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) changeDomainLicenseType(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/domains/{domainId}/actions/changeLicenseType", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDomainLicenseTypeResponse
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

// ChangeTagNamespaceCompartment Moves the specified tag namespace to the specified compartment within the same tenancy.
// To move the tag namespace, you must have the manage tag-namespaces permission on both compartments.
// For more information about IAM policies, see Details for IAM (https://docs.cloud.oracle.com/Content/Identity/Reference/iampolicyreference.htm).
// Moving a tag namespace moves all the tag key definitions contained in the tag namespace.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ChangeTagNamespaceCompartment.go.html to see an example of how to use ChangeTagNamespaceCompartment API.
func (client IdentityClient) ChangeTagNamespaceCompartment(ctx context.Context, request ChangeTagNamespaceCompartmentRequest) (response ChangeTagNamespaceCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeTagNamespaceCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeTagNamespaceCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeTagNamespaceCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeTagNamespaceCompartmentResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeTagNamespaceCompartmentResponse")
	}
	return
}

// changeTagNamespaceCompartment implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) changeTagNamespaceCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/tagNamespaces/{tagNamespaceId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeTagNamespaceCompartmentResponse
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

// CreateAuthToken Creates a new auth token for the specified user. For information about what auth tokens are for, see
// Managing User Credentials (https://docs.cloud.oracle.com/Content/Identity/Tasks/managingcredentials.htm).
// You must specify a *description* for the auth token (although it can be an empty string). It does not
// have to be unique, and you can change it anytime with
// UpdateAuthToken.
// Every user has permission to create an auth token for *their own user ID*. An administrator in your organization
// does not need to write a policy to give users this ability. To compare, administrators who have permission to the
// tenancy can use this operation to create an auth token for any user, including themselves.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/CreateAuthToken.go.html to see an example of how to use CreateAuthToken API.
func (client IdentityClient) CreateAuthToken(ctx context.Context, request CreateAuthTokenRequest) (response CreateAuthTokenResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createAuthToken, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAuthTokenResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAuthTokenResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAuthTokenResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAuthTokenResponse")
	}
	return
}

// createAuthToken implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) createAuthToken(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/users/{userId}/authTokens", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateAuthTokenResponse
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

// CreateCompartment Creates a new compartment in the specified compartment.
// **Important:** Compartments cannot be deleted.
// Specify the parent compartment's OCID as the compartment ID in the request object. Remember that the tenancy
// is simply the root compartment. For information about OCIDs, see
// Resource Identifiers (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
// You must also specify a *name* for the compartment, which must be unique across all compartments in
// your tenancy. You can use this name or the OCID when writing policies that apply
// to the compartment. For more information about policies, see
// How Policies Work (https://docs.cloud.oracle.com/Content/Identity/Concepts/policies.htm).
// You must also specify a *description* for the compartment (although it can be an empty string). It does
// not have to be unique, and you can change it anytime with
// UpdateCompartment.
// After you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before using the
// object, first make sure its `lifecycleState` has changed to ACTIVE.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/CreateCompartment.go.html to see an example of how to use CreateCompartment API.
func (client IdentityClient) CreateCompartment(ctx context.Context, request CreateCompartmentRequest) (response CreateCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateCompartmentResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateCompartmentResponse")
	}
	return
}

// createCompartment implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) createCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/compartments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateCompartmentResponse
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

// CreateCustomerSecretKey Creates a new secret key for the specified user. Secret keys are used for authentication with the Object Storage Service's Amazon S3
// compatible API. The secret key consists of an Access Key/Secret Key pair. For information, see
// Managing User Credentials (https://docs.cloud.oracle.com/Content/Identity/Tasks/managingcredentials.htm).
// You must specify a *description* for the secret key (although it can be an empty string). It does not
// have to be unique, and you can change it anytime with
// UpdateCustomerSecretKey.
// Every user has permission to create a secret key for *their own user ID*. An administrator in your organization
// does not need to write a policy to give users this ability. To compare, administrators who have permission to the
// tenancy can use this operation to create a secret key for any user, including themselves.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/CreateCustomerSecretKey.go.html to see an example of how to use CreateCustomerSecretKey API.
func (client IdentityClient) CreateCustomerSecretKey(ctx context.Context, request CreateCustomerSecretKeyRequest) (response CreateCustomerSecretKeyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createCustomerSecretKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateCustomerSecretKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateCustomerSecretKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateCustomerSecretKeyResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateCustomerSecretKeyResponse")
	}
	return
}

// createCustomerSecretKey implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) createCustomerSecretKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/users/{userId}/customerSecretKeys", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateCustomerSecretKeyResponse
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

// CreateDomain Creates a new domain in the tenancy with domain home in {@code homeRegion}. This is an asynchronous call - where, at start,
// {@code lifecycleState} of this domain is set to CREATING and {@code lifecycleDetails} to UPDATING. On domain creation completion
// this Domain's {@code lifecycleState} will be set to ACTIVE and {@code lifecycleDetails} to null.
// To track progress, HTTP GET on /iamWorkRequests/{iamWorkRequestsId} endpoint will provide
// the async operation's status.
// After creating a `Domain`, make sure its `lifecycleState` changes from CREATING to ACTIVE
// before using it.
// If the domain's {@code displayName} already exists, returns 400 BAD REQUEST.
// If any one of admin related fields are provided and one of the following 3 fields
// - {@code adminEmail}, {@code adminLastName} and {@code adminUserName} - is not provided,
// returns 400 BAD REQUEST.
// - If {@code isNotificationBypassed} is NOT provided when admin information is provided,
// returns 400 BAD REQUEST.
// - If any internal error occurs, return 500 INTERNAL SERVER ERROR.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/CreateDomain.go.html to see an example of how to use CreateDomain API.
func (client IdentityClient) CreateDomain(ctx context.Context, request CreateDomainRequest) (response CreateDomainResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDomain, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDomainResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDomainResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDomainResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDomainResponse")
	}
	return
}

// createDomain implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) createDomain(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/domains", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDomainResponse
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

// CreateDynamicGroup Creates a new dynamic group in your tenancy.
// You must specify your tenancy's OCID as the compartment ID in the request object (remember that the tenancy
// is simply the root compartment). Notice that IAM resources (users, groups, compartments, and some policies)
// reside within the tenancy itself, unlike cloud resources such as compute instances, which typically
// reside within compartments inside the tenancy. For information about OCIDs, see
// Resource Identifiers (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
// You must also specify a *name* for the dynamic group, which must be unique across all dynamic groups in your
// tenancy, and cannot be changed. Note that this name has to be also unique across all groups in your tenancy.
// You can use this name or the OCID when writing policies that apply to the dynamic group. For more information
// about policies, see How Policies Work (https://docs.cloud.oracle.com/Content/Identity/Concepts/policies.htm).
// You must also specify a *description* for the dynamic group (although it can be an empty string). It does not
// have to be unique, and you can change it anytime with UpdateDynamicGroup.
// After you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before using the
// object, first make sure its `lifecycleState` has changed to ACTIVE.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/CreateDynamicGroup.go.html to see an example of how to use CreateDynamicGroup API.
func (client IdentityClient) CreateDynamicGroup(ctx context.Context, request CreateDynamicGroupRequest) (response CreateDynamicGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createDynamicGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDynamicGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDynamicGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDynamicGroupResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDynamicGroupResponse")
	}
	return
}

// createDynamicGroup implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) createDynamicGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/dynamicGroups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDynamicGroupResponse
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

// CreateGroup Creates a new group in your tenancy.
// You must specify your tenancy's OCID as the compartment ID in the request object (remember that the tenancy
// is simply the root compartment). Notice that IAM resources (users, groups, compartments, and some policies)
// reside within the tenancy itself, unlike cloud resources such as compute instances, which typically
// reside within compartments inside the tenancy. For information about OCIDs, see
// Resource Identifiers (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
// You must also specify a *name* for the group, which must be unique across all groups in your tenancy and
// cannot be changed. You can use this name or the OCID when writing policies that apply to the group. For more
// information about policies, see How Policies Work (https://docs.cloud.oracle.com/Content/Identity/Concepts/policies.htm).
// You must also specify a *description* for the group (although it can be an empty string). It does not
// have to be unique, and you can change it anytime with UpdateGroup.
// After you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before using the
// object, first make sure its `lifecycleState` has changed to ACTIVE.
// After creating the group, you need to put users in it and write policies for it.
// See AddUserToGroup and
// CreatePolicy.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/CreateGroup.go.html to see an example of how to use CreateGroup API.
func (client IdentityClient) CreateGroup(ctx context.Context, request CreateGroupRequest) (response CreateGroupResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateGroupResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateGroupResponse")
	}
	return
}

// createGroup implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) createGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/groups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateGroupResponse
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

// CreateIdentityProvider **Deprecated.** For more information, see Deprecated IAM Service APIs (https://docs.cloud.oracle.com/Content/Identity/Reference/deprecatediamapis.htm).
// Creates a new identity provider in your tenancy. For more information, see
// Identity Providers and Federation (https://docs.cloud.oracle.com/Content/Identity/Concepts/federation.htm).
// You must specify your tenancy's OCID as the compartment ID in the request object.
// Remember that the tenancy is simply the root compartment. For information about
// OCIDs, see Resource Identifiers (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
// You must also specify a *name* for the `IdentityProvider`, which must be unique
// across all `IdentityProvider` objects in your tenancy and cannot be changed.
// You must also specify a *description* for the `IdentityProvider` (although
// it can be an empty string). It does not have to be unique, and you can change
// it anytime with
// UpdateIdentityProvider.
// After you send your request, the new object's `lifecycleState` will temporarily
// be CREATING. Before using the object, first make sure its `lifecycleState` has
// changed to ACTIVE.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/CreateIdentityProvider.go.html to see an example of how to use CreateIdentityProvider API.
func (client IdentityClient) CreateIdentityProvider(ctx context.Context, request CreateIdentityProviderRequest) (response CreateIdentityProviderResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createIdentityProvider, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateIdentityProviderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateIdentityProviderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateIdentityProviderResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateIdentityProviderResponse")
	}
	return
}

// createIdentityProvider implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) createIdentityProvider(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/identityProviders", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateIdentityProviderResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &identityprovider{})
	return response, err
}

// CreateIdpGroupMapping **Deprecated.** For more information, see Deprecated IAM Service APIs (https://docs.cloud.oracle.com/Content/Identity/Reference/deprecatediamapis.htm).
// Creates a single mapping between an IdP group and an IAM Service
// Group.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/CreateIdpGroupMapping.go.html to see an example of how to use CreateIdpGroupMapping API.
func (client IdentityClient) CreateIdpGroupMapping(ctx context.Context, request CreateIdpGroupMappingRequest) (response CreateIdpGroupMappingResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createIdpGroupMapping, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateIdpGroupMappingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateIdpGroupMappingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateIdpGroupMappingResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateIdpGroupMappingResponse")
	}
	return
}

// createIdpGroupMapping implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) createIdpGroupMapping(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/identityProviders/{identityProviderId}/groupMappings", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateIdpGroupMappingResponse
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

// CreateMfaTotpDevice Creates a new MFA TOTP device for the user. A user can have one MFA TOTP device.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/CreateMfaTotpDevice.go.html to see an example of how to use CreateMfaTotpDevice API.
func (client IdentityClient) CreateMfaTotpDevice(ctx context.Context, request CreateMfaTotpDeviceRequest) (response CreateMfaTotpDeviceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createMfaTotpDevice, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateMfaTotpDeviceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateMfaTotpDeviceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateMfaTotpDeviceResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateMfaTotpDeviceResponse")
	}
	return
}

// createMfaTotpDevice implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) createMfaTotpDevice(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/users/{userId}/mfaTotpDevices", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateMfaTotpDeviceResponse
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

// CreateNetworkSource Creates a new network source in your tenancy.
// You must specify your tenancy's OCID as the compartment ID in the request object (remember that the tenancy
// is simply the root compartment). Notice that IAM resources (users, groups, compartments, and some policies)
// reside within the tenancy itself, unlike cloud resources such as compute instances, which typically
// reside within compartments inside the tenancy. For information about OCIDs, see
// Resource Identifiers (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
// You must also specify a *name* for the network source, which must be unique across all network sources in your
// tenancy, and cannot be changed.
// You can use this name or the OCID when writing policies that apply to the network source. For more information
// about policies, see How Policies Work (https://docs.cloud.oracle.com/Content/Identity/Concepts/policies.htm).
// You must also specify a *description* for the network source (although it can be an empty string). It does not
// have to be unique, and you can change it anytime with UpdateNetworkSource.
// After you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before using the
// object, first make sure its `lifecycleState` has changed to ACTIVE.
// After your network resource is created, you can use it in policy to restrict access to only requests made from an allowed
// IP address specified in your network source. For more information, see Managing Network Sources (https://docs.cloud.oracle.com/Content/Identity/Tasks/managingnetworksources.htm).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/CreateNetworkSource.go.html to see an example of how to use CreateNetworkSource API.
func (client IdentityClient) CreateNetworkSource(ctx context.Context, request CreateNetworkSourceRequest) (response CreateNetworkSourceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createNetworkSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateNetworkSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateNetworkSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateNetworkSourceResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateNetworkSourceResponse")
	}
	return
}

// createNetworkSource implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) createNetworkSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkSources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateNetworkSourceResponse
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

// CreateOAuthClientCredential Creates Oauth token for the user
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/CreateOAuthClientCredential.go.html to see an example of how to use CreateOAuthClientCredential API.
func (client IdentityClient) CreateOAuthClientCredential(ctx context.Context, request CreateOAuthClientCredentialRequest) (response CreateOAuthClientCredentialResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOAuthClientCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOAuthClientCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOAuthClientCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOAuthClientCredentialResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOAuthClientCredentialResponse")
	}
	return
}

// createOAuthClientCredential implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) createOAuthClientCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/users/{userId}/oauth2ClientCredentials", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOAuthClientCredentialResponse
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

// CreateOrResetUIPassword Creates a new Console one-time password for the specified user. For more information about user
// credentials, see User Credentials (https://docs.cloud.oracle.com/Content/Identity/Concepts/usercredentials.htm).
// Use this operation after creating a new user, or if a user forgets their password. The new one-time
// password is returned to you in the response, and you must securely deliver it to the user. They'll
// be prompted to change this password the next time they sign in to the Console. If they don't change
// it within 7 days, the password will expire and you'll need to create a new one-time password for the
// user.
// **Note:** The user's Console login is the unique name you specified when you created the user
// (see CreateUser).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/CreateOrResetUIPassword.go.html to see an example of how to use CreateOrResetUIPassword API.
func (client IdentityClient) CreateOrResetUIPassword(ctx context.Context, request CreateOrResetUIPasswordRequest) (response CreateOrResetUIPasswordResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createOrResetUIPassword, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateOrResetUIPasswordResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateOrResetUIPasswordResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateOrResetUIPasswordResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateOrResetUIPasswordResponse")
	}
	return
}

// createOrResetUIPassword implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) createOrResetUIPassword(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/users/{userId}/uiPassword", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateOrResetUIPasswordResponse
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

// CreatePolicy Creates a new policy in the specified compartment (either the tenancy or another of your compartments).
// If you're new to policies, see Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
// You must specify a *name* for the policy, which must be unique across all policies in your tenancy
// and cannot be changed.
// You must also specify a *description* for the policy (although it can be an empty string). It does not
// have to be unique, and you can change it anytime with UpdatePolicy.
// You must specify one or more policy statements in the statements array. For information about writing
// policies, see How Policies Work (https://docs.cloud.oracle.com/Content/Identity/Concepts/policies.htm) and
// Common Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/commonpolicies.htm).
// After you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before using the
// object, first make sure its `lifecycleState` has changed to ACTIVE.
// New policies take effect typically within 10 seconds.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/CreatePolicy.go.html to see an example of how to use CreatePolicy API.
func (client IdentityClient) CreatePolicy(ctx context.Context, request CreatePolicyRequest) (response CreatePolicyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreatePolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreatePolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreatePolicyResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreatePolicyResponse")
	}
	return
}

// createPolicy implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) createPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/policies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreatePolicyResponse
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

// CreateRegionSubscription Creates a subscription to a region for a tenancy.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/CreateRegionSubscription.go.html to see an example of how to use CreateRegionSubscription API.
func (client IdentityClient) CreateRegionSubscription(ctx context.Context, request CreateRegionSubscriptionRequest) (response CreateRegionSubscriptionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createRegionSubscription, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateRegionSubscriptionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateRegionSubscriptionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateRegionSubscriptionResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateRegionSubscriptionResponse")
	}
	return
}

// createRegionSubscription implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) createRegionSubscription(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/tenancies/{tenancyId}/regionSubscriptions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateRegionSubscriptionResponse
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

// CreateSmtpCredential Creates a new SMTP credential for the specified user. An SMTP credential has an SMTP user name and an SMTP password.
// You must specify a *description* for the SMTP credential (although it can be an empty string). It does not
// have to be unique, and you can change it anytime with
// UpdateSmtpCredential.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/CreateSmtpCredential.go.html to see an example of how to use CreateSmtpCredential API.
func (client IdentityClient) CreateSmtpCredential(ctx context.Context, request CreateSmtpCredentialRequest) (response CreateSmtpCredentialResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createSmtpCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSmtpCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSmtpCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSmtpCredentialResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSmtpCredentialResponse")
	}
	return
}

// createSmtpCredential implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) createSmtpCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/users/{userId}/smtpCredentials", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSmtpCredentialResponse
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

// CreateSwiftPassword **Deprecated. Use CreateAuthToken instead.**
// Creates a new Swift password for the specified user. For information about what Swift passwords are for, see
// Managing User Credentials (https://docs.cloud.oracle.com/Content/Identity/Tasks/managingcredentials.htm).
// You must specify a *description* for the Swift password (although it can be an empty string). It does not
// have to be unique, and you can change it anytime with
// UpdateSwiftPassword.
// Every user has permission to create a Swift password for *their own user ID*. An administrator in your organization
// does not need to write a policy to give users this ability. To compare, administrators who have permission to the
// tenancy can use this operation to create a Swift password for any user, including themselves.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/CreateSwiftPassword.go.html to see an example of how to use CreateSwiftPassword API.
func (client IdentityClient) CreateSwiftPassword(ctx context.Context, request CreateSwiftPasswordRequest) (response CreateSwiftPasswordResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createSwiftPassword, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSwiftPasswordResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSwiftPasswordResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSwiftPasswordResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSwiftPasswordResponse")
	}
	return
}

// createSwiftPassword implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) createSwiftPassword(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/users/{userId}/swiftPasswords", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSwiftPasswordResponse
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

// CreateTag Creates a new tag in the specified tag namespace.
// The tag requires either the OCID or the name of the tag namespace that will contain this
// tag definition.
// You must specify a *name* for the tag, which must be unique across all tags in the tag namespace
// and cannot be changed. The name can contain any ASCII character except the space (_) or period (.) characters.
// Names are case insensitive. That means, for example, "myTag" and "mytag" are not allowed in the same namespace.
// If you specify a name that's already in use in the tag namespace, a 409 error is returned.
// The tag must have a *description*. It does not have to be unique, and you can change it with
// UpdateTag.
// The tag must have a value type, which is specified with a validator. Tags can use either a
// static value or a list of possible values. Static values are entered by a user applying the tag
// to a resource. Lists are created by you and the user must apply a value from the list. Lists
// are validiated.
// * If no `validator` is set, the user applying the tag to a resource can type in a static
// value or leave the tag value empty.
// * If a `validator` is set, the user applying the tag to a resource must select from a list
// of values that you supply with EnumTagDefinitionValidator.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/CreateTag.go.html to see an example of how to use CreateTag API.
func (client IdentityClient) CreateTag(ctx context.Context, request CreateTagRequest) (response CreateTagResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createTag, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTagResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTagResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTagResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTagResponse")
	}
	return
}

// createTag implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) createTag(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/tagNamespaces/{tagNamespaceId}/tags", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateTagResponse
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

// CreateTagDefault Creates a new tag default in the specified compartment for the specified tag definition.
// If you specify that a value is required, a value is set during resource creation (either by
// the user creating the resource or another tag defualt). If no value is set, resource creation
// is blocked.
// * If the `isRequired` flag is set to "true", the value is set during resource creation.
// * If the `isRequired` flag is set to "false", the value you enter is set during resource creation.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/CreateTagDefault.go.html to see an example of how to use CreateTagDefault API.
func (client IdentityClient) CreateTagDefault(ctx context.Context, request CreateTagDefaultRequest) (response CreateTagDefaultResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createTagDefault, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTagDefaultResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTagDefaultResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTagDefaultResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTagDefaultResponse")
	}
	return
}

// createTagDefault implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) createTagDefault(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/tagDefaults", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateTagDefaultResponse
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

// CreateTagNamespace Creates a new tag namespace in the specified compartment.
// You must specify the compartment ID in the request object (remember that the tenancy is simply the root
// compartment).
// You must also specify a *name* for the namespace, which must be unique across all namespaces in your tenancy
// and cannot be changed. The name can contain any ASCII character except the space (_) or period (.).
// Names are case insensitive. That means, for example, "myNamespace" and "mynamespace" are not allowed
// in the same tenancy. Once you created a namespace, you cannot change the name.
// If you specify a name that's already in use in the tenancy, a 409 error is returned.
// You must also specify a *description* for the namespace.
// It does not have to be unique, and you can change it with
// UpdateTagNamespace.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/CreateTagNamespace.go.html to see an example of how to use CreateTagNamespace API.
func (client IdentityClient) CreateTagNamespace(ctx context.Context, request CreateTagNamespaceRequest) (response CreateTagNamespaceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createTagNamespace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTagNamespaceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTagNamespaceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTagNamespaceResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTagNamespaceResponse")
	}
	return
}

// createTagNamespace implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) createTagNamespace(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/tagNamespaces", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateTagNamespaceResponse
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

// CreateUser Creates a new user in your tenancy. For conceptual information about users, your tenancy, and other
// IAM Service components, see Overview of the IAM Service (https://docs.cloud.oracle.com/Content/Identity/Concepts/overview.htm).
// You must specify your tenancy's OCID as the compartment ID in the request object (remember that the
// tenancy is simply the root compartment). Notice that IAM resources (users, groups, compartments, and
// some policies) reside within the tenancy itself, unlike cloud resources such as compute instances,
// which typically reside within compartments inside the tenancy. For information about OCIDs, see
// Resource Identifiers (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
// You must also specify a *name* for the user, which must be unique across all users in your tenancy
// and cannot be changed. Allowed characters: No spaces. Only letters, numerals, hyphens, periods,
// underscores, +, and @. If you specify a name that's already in use, you'll get a 409 error.
// This name will be the user's login to the Console. You might want to pick a
// name that your company's own identity system (e.g., Active Directory, LDAP, etc.) already uses.
// If you delete a user and then create a new user with the same name, they'll be considered different
// users because they have different OCIDs.
// You must also specify a *description* for the user (although it can be an empty string).
// It does not have to be unique, and you can change it anytime with
// UpdateUser. You can use the field to provide the user's
// full name, a description, a nickname, or other information to generally identify the user.
// After you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before
// using the object, first make sure its `lifecycleState` has changed to ACTIVE.
// A new user has no permissions until you place the user in one or more groups (see
// AddUserToGroup). If the user needs to
// access the Console, you need to provide the user a password (see
// CreateOrResetUIPassword).
// If the user needs to access the Oracle Cloud Infrastructure REST API, you need to upload a
// public API signing key for that user (see
// Required Keys and OCIDs (https://docs.cloud.oracle.com/Content/API/Concepts/apisigningkey.htm) and also
// UploadApiKey).
// **Important:** Make sure to inform the new user which compartment(s) they have access to.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/CreateUser.go.html to see an example of how to use CreateUser API.
func (client IdentityClient) CreateUser(ctx context.Context, request CreateUserRequest) (response CreateUserResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createUser, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateUserResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateUserResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateUserResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateUserResponse")
	}
	return
}

// createUser implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) createUser(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/users", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateUserResponse
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

// DeactivateDomain If the domain's {@code lifecycleState} is ACTIVE and no active Apps are present in domain,
// 1. Set the {@code lifecycleDetails} to DEACTIVATING and asynchronously starts disabling
//    the domain and return 202 ACCEPTED.
//     1.1 Sets the domain status to DISABLED and set specified domain's
//         {@code lifecycleState} to INACTIVE and set the {@code lifecycleDetails} to null.
// To track progress, HTTP GET on /iamWorkRequests/{iamWorkRequestsId} endpoint will provide
// the async operation's status. Activate a domain can be done using HTTP POST
// /domains/{domainId}/actions/activate.
// - If the domain's {@code lifecycleState} is INACTIVE, returns 202 ACCEPTED with no action
//   taken on service side.
// - If domain is of {@code type} DEFAULT or DEFAULT_LIGHTWEIGHT or domain's {@code lifecycleState}
//   is not ACTIVE, returns 400 BAD REQUEST.
// - If the domain doesn't exists, returns 404 NOT FOUND.
// - If any active Apps in domain, returns 400 BAD REQUEST.
// - If the authenticated user is part of the domain to be activated, returns 400 BAD REQUEST
// - If error occurs while deactivating domain, returns 500 INTERNAL SERVER ERROR.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/DeactivateDomain.go.html to see an example of how to use DeactivateDomain API.
func (client IdentityClient) DeactivateDomain(ctx context.Context, request DeactivateDomainRequest) (response DeactivateDomainResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deactivateDomain, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeactivateDomainResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeactivateDomainResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeactivateDomainResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeactivateDomainResponse")
	}
	return
}

// deactivateDomain implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) deactivateDomain(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/domains/{domainId}/actions/deactivate", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeactivateDomainResponse
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

// DeleteApiKey Deletes the specified API signing key for the specified user.
// Every user has permission to use this operation to delete a key for *their own user ID*. An
// administrator in your organization does not need to write a policy to give users this ability.
// To compare, administrators who have permission to the tenancy can use this operation to delete
// a key for any user, including themselves.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/DeleteApiKey.go.html to see an example of how to use DeleteApiKey API.
func (client IdentityClient) DeleteApiKey(ctx context.Context, request DeleteApiKeyRequest) (response DeleteApiKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteApiKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteApiKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteApiKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteApiKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteApiKeyResponse")
	}
	return
}

// deleteApiKey implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) deleteApiKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/users/{userId}/apiKeys/{fingerprint}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteApiKeyResponse
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

// DeleteAuthToken Deletes the specified auth token for the specified user.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/DeleteAuthToken.go.html to see an example of how to use DeleteAuthToken API.
func (client IdentityClient) DeleteAuthToken(ctx context.Context, request DeleteAuthTokenRequest) (response DeleteAuthTokenResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAuthToken, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAuthTokenResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAuthTokenResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAuthTokenResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAuthTokenResponse")
	}
	return
}

// deleteAuthToken implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) deleteAuthToken(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/users/{userId}/authTokens/{authTokenId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteAuthTokenResponse
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

// DeleteCompartment Deletes the specified compartment. The compartment must be empty.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/DeleteCompartment.go.html to see an example of how to use DeleteCompartment API.
func (client IdentityClient) DeleteCompartment(ctx context.Context, request DeleteCompartmentRequest) (response DeleteCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteCompartmentResponse")
	}
	return
}

// deleteCompartment implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) deleteCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/compartments/{compartmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteCompartmentResponse
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

// DeleteCustomerSecretKey Deletes the specified secret key for the specified user.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/DeleteCustomerSecretKey.go.html to see an example of how to use DeleteCustomerSecretKey API.
func (client IdentityClient) DeleteCustomerSecretKey(ctx context.Context, request DeleteCustomerSecretKeyRequest) (response DeleteCustomerSecretKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteCustomerSecretKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteCustomerSecretKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteCustomerSecretKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteCustomerSecretKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteCustomerSecretKeyResponse")
	}
	return
}

// deleteCustomerSecretKey implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) deleteCustomerSecretKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/users/{userId}/customerSecretKeys/{customerSecretKeyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteCustomerSecretKeyResponse
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

// DeleteDomain Soft Deletes a domain.
// This is an asynchronous API, where, if the domain's {@code lifecycleState} is INACTIVE and
// no active Apps are present in underlying stripe,
//   1. Sets the specified domain's {@code lifecycleState} to DELETING.
//   2. Domains marked as DELETING will be cleaned up by a periodic task unless customer request it to be undo via ticket.
//   3. Work request is created and returned as opc-work-request-id along with 202 ACCEPTED.
// To track progress, HTTP GET on /iamWorkRequests/{iamWorkRequestsId} endpoint will provide
// the async operation's status.
// - If the domain's {@code lifecycleState} is DELETING, returns 202 Accepted as a deletion
//   is already in progress for this domain.
// - If the domain doesn't exists, returns 404 NOT FOUND.
// - If domain is of {@code type} DEFAULT or DEFAULT_LIGHTWEIGHT, returns 400 BAD REQUEST.
// - If any active Apps in domain, returns 400 BAD REQUEST.
// - If the authenticated user is part of the domain to be deleted, returns 400 BAD REQUEST.
// - If any internal error occurs, return 500 INTERNAL SERVER ERROR.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/DeleteDomain.go.html to see an example of how to use DeleteDomain API.
func (client IdentityClient) DeleteDomain(ctx context.Context, request DeleteDomainRequest) (response DeleteDomainResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDomain, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDomainResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDomainResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDomainResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDomainResponse")
	}
	return
}

// deleteDomain implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) deleteDomain(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/domains/{domainId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDomainResponse
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

// DeleteDynamicGroup Deletes the specified dynamic group.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/DeleteDynamicGroup.go.html to see an example of how to use DeleteDynamicGroup API.
func (client IdentityClient) DeleteDynamicGroup(ctx context.Context, request DeleteDynamicGroupRequest) (response DeleteDynamicGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDynamicGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDynamicGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDynamicGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDynamicGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDynamicGroupResponse")
	}
	return
}

// deleteDynamicGroup implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) deleteDynamicGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/dynamicGroups/{dynamicGroupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDynamicGroupResponse
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

// DeleteGroup Deletes the specified group. The group must be empty.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/DeleteGroup.go.html to see an example of how to use DeleteGroup API.
func (client IdentityClient) DeleteGroup(ctx context.Context, request DeleteGroupRequest) (response DeleteGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteGroupResponse")
	}
	return
}

// deleteGroup implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) deleteGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/groups/{groupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteGroupResponse
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

// DeleteIdentityProvider **Deprecated.** For more information, see Deprecated IAM Service APIs (https://docs.cloud.oracle.com/Content/Identity/Reference/deprecatediamapis.htm).
// Deletes the specified identity provider. The identity provider must not have
// any group mappings (see IdpGroupMapping).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/DeleteIdentityProvider.go.html to see an example of how to use DeleteIdentityProvider API.
func (client IdentityClient) DeleteIdentityProvider(ctx context.Context, request DeleteIdentityProviderRequest) (response DeleteIdentityProviderResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteIdentityProvider, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteIdentityProviderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteIdentityProviderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteIdentityProviderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteIdentityProviderResponse")
	}
	return
}

// deleteIdentityProvider implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) deleteIdentityProvider(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/identityProviders/{identityProviderId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteIdentityProviderResponse
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

// DeleteIdpGroupMapping **Deprecated.** For more information, see Deprecated IAM Service APIs (https://docs.cloud.oracle.com/Content/Identity/Reference/deprecatediamapis.htm).
// Deletes the specified group mapping.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/DeleteIdpGroupMapping.go.html to see an example of how to use DeleteIdpGroupMapping API.
func (client IdentityClient) DeleteIdpGroupMapping(ctx context.Context, request DeleteIdpGroupMappingRequest) (response DeleteIdpGroupMappingResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteIdpGroupMapping, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteIdpGroupMappingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteIdpGroupMappingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteIdpGroupMappingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteIdpGroupMappingResponse")
	}
	return
}

// deleteIdpGroupMapping implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) deleteIdpGroupMapping(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/identityProviders/{identityProviderId}/groupMappings/{mappingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteIdpGroupMappingResponse
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

// DeleteMfaTotpDevice Deletes the specified MFA TOTP device for the specified user.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/DeleteMfaTotpDevice.go.html to see an example of how to use DeleteMfaTotpDevice API.
func (client IdentityClient) DeleteMfaTotpDevice(ctx context.Context, request DeleteMfaTotpDeviceRequest) (response DeleteMfaTotpDeviceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteMfaTotpDevice, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteMfaTotpDeviceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteMfaTotpDeviceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteMfaTotpDeviceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteMfaTotpDeviceResponse")
	}
	return
}

// deleteMfaTotpDevice implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) deleteMfaTotpDevice(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/users/{userId}/mfaTotpDevices/{mfaTotpDeviceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteMfaTotpDeviceResponse
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

// DeleteNetworkSource Deletes the specified network source
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/DeleteNetworkSource.go.html to see an example of how to use DeleteNetworkSource API.
func (client IdentityClient) DeleteNetworkSource(ctx context.Context, request DeleteNetworkSourceRequest) (response DeleteNetworkSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteNetworkSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteNetworkSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteNetworkSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteNetworkSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteNetworkSourceResponse")
	}
	return
}

// deleteNetworkSource implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) deleteNetworkSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/networkSources/{networkSourceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteNetworkSourceResponse
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

// DeleteOAuthClientCredential Delete Oauth token for the user
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/DeleteOAuthClientCredential.go.html to see an example of how to use DeleteOAuthClientCredential API.
func (client IdentityClient) DeleteOAuthClientCredential(ctx context.Context, request DeleteOAuthClientCredentialRequest) (response DeleteOAuthClientCredentialResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteOAuthClientCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteOAuthClientCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteOAuthClientCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteOAuthClientCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteOAuthClientCredentialResponse")
	}
	return
}

// deleteOAuthClientCredential implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) deleteOAuthClientCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/users/{userId}/oauth2ClientCredentials/{oauth2ClientCredentialId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteOAuthClientCredentialResponse
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

// DeletePolicy Deletes the specified policy. The deletion takes effect typically within 10 seconds.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/DeletePolicy.go.html to see an example of how to use DeletePolicy API.
func (client IdentityClient) DeletePolicy(ctx context.Context, request DeletePolicyRequest) (response DeletePolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deletePolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeletePolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeletePolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeletePolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeletePolicyResponse")
	}
	return
}

// deletePolicy implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) deletePolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/policies/{policyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeletePolicyResponse
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

// DeleteSmtpCredential Deletes the specified SMTP credential for the specified user.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/DeleteSmtpCredential.go.html to see an example of how to use DeleteSmtpCredential API.
func (client IdentityClient) DeleteSmtpCredential(ctx context.Context, request DeleteSmtpCredentialRequest) (response DeleteSmtpCredentialResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSmtpCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSmtpCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSmtpCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSmtpCredentialResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSmtpCredentialResponse")
	}
	return
}

// deleteSmtpCredential implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) deleteSmtpCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/users/{userId}/smtpCredentials/{smtpCredentialId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSmtpCredentialResponse
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

// DeleteSwiftPassword **Deprecated. Use DeleteAuthToken instead.**
// Deletes the specified Swift password for the specified user.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/DeleteSwiftPassword.go.html to see an example of how to use DeleteSwiftPassword API.
func (client IdentityClient) DeleteSwiftPassword(ctx context.Context, request DeleteSwiftPasswordRequest) (response DeleteSwiftPasswordResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSwiftPassword, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSwiftPasswordResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSwiftPasswordResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSwiftPasswordResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSwiftPasswordResponse")
	}
	return
}

// deleteSwiftPassword implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) deleteSwiftPassword(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/users/{userId}/swiftPasswords/{swiftPasswordId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSwiftPasswordResponse
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

// DeleteTag Deletes the specified tag definition. This operation triggers a process that removes the
// tag from all resources in your tenancy.
// These things happen immediately:
//
//   * If the tag was a cost-tracking tag, it no longer counts against your 10 cost-tracking
//   tags limit, whether you first disabled it or not.
//   * If the tag was used with dynamic groups, none of the rules that contain the tag will
//   be evaluated against the tag.
// When you start the delete operation, the state of the tag changes to DELETING and tag removal
// from resources begins. This can take up to 48 hours depending on the number of resources that
// were tagged as well as the regions in which those resources reside.
// When all tags have been removed, the state changes to DELETED. You cannot restore a deleted tag. Once the deleted tag
// changes its state to DELETED, you can use the same tag name again.
// After you start this operation, you cannot start either the BulkDeleteTags or the CascadeDeleteTagNamespace operation until this process completes.
// To delete a tag, you must first retire it. Use UpdateTag
// to retire a tag.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/DeleteTag.go.html to see an example of how to use DeleteTag API.
func (client IdentityClient) DeleteTag(ctx context.Context, request DeleteTagRequest) (response DeleteTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTag, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteTagResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteTagResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTagResponse")
	}
	return
}

// deleteTag implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) deleteTag(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/tagNamespaces/{tagNamespaceId}/tags/{tagName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteTagResponse
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

// DeleteTagDefault Deletes the the specified tag default.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/DeleteTagDefault.go.html to see an example of how to use DeleteTagDefault API.
func (client IdentityClient) DeleteTagDefault(ctx context.Context, request DeleteTagDefaultRequest) (response DeleteTagDefaultResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTagDefault, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteTagDefaultResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteTagDefaultResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTagDefaultResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTagDefaultResponse")
	}
	return
}

// deleteTagDefault implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) deleteTagDefault(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/tagDefaults/{tagDefaultId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteTagDefaultResponse
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

// DeleteTagNamespace Deletes the specified tag namespace. Only an empty tag namespace can be deleted with this operation. To use this operation
// to delete a tag namespace that contains tag definitions, first delete all of its tag definitions.
// Use CascadeDeleteTagNamespace to delete a tag namespace along with all of
// the tag definitions contained within that namespace.
// Use DeleteTag to delete a tag definition.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/DeleteTagNamespace.go.html to see an example of how to use DeleteTagNamespace API.
func (client IdentityClient) DeleteTagNamespace(ctx context.Context, request DeleteTagNamespaceRequest) (response DeleteTagNamespaceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTagNamespace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteTagNamespaceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteTagNamespaceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTagNamespaceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTagNamespaceResponse")
	}
	return
}

// deleteTagNamespace implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) deleteTagNamespace(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/tagNamespaces/{tagNamespaceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteTagNamespaceResponse
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

// DeleteUser Deletes the specified user. The user must not be in any groups.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/DeleteUser.go.html to see an example of how to use DeleteUser API.
func (client IdentityClient) DeleteUser(ctx context.Context, request DeleteUserRequest) (response DeleteUserResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteUser, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteUserResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteUserResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteUserResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteUserResponse")
	}
	return
}

// deleteUser implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) deleteUser(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/users/{userId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteUserResponse
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

// EnableReplicationToRegion Replicate domain to a new region. This is an asynchronous call - where, at start,
// {@code state} of this domain in replica region is set to ENABLING_REPLICATION.
// On domain replication completion the {@code state} will be set to REPLICATION_ENABLED.
// To track progress, HTTP GET on /iamWorkRequests/{iamWorkRequestsId} endpoint will provide
// the async operation's status.
// If the replica region's {@code state} is already ENABLING_REPLICATION or REPLICATION_ENABLED,
// returns 409 CONFLICT.
// - If the domain doesn't exists, returns 404 NOT FOUND.
// - If home region is same as replication region, return 400 BAD REQUEST.
// - If Domain is not active or being updated, returns 400 BAD REQUEST.
// - If any internal error occurs, return 500 INTERNAL SERVER ERROR.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/EnableReplicationToRegion.go.html to see an example of how to use EnableReplicationToRegion API.
func (client IdentityClient) EnableReplicationToRegion(ctx context.Context, request EnableReplicationToRegionRequest) (response EnableReplicationToRegionResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.enableReplicationToRegion, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableReplicationToRegionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableReplicationToRegionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableReplicationToRegionResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableReplicationToRegionResponse")
	}
	return
}

// enableReplicationToRegion implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) enableReplicationToRegion(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/domains/{domainId}/actions/enableReplicationToRegion", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response EnableReplicationToRegionResponse
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

// GenerateTotpSeed Generate seed for the MFA TOTP device.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/GenerateTotpSeed.go.html to see an example of how to use GenerateTotpSeed API.
func (client IdentityClient) GenerateTotpSeed(ctx context.Context, request GenerateTotpSeedRequest) (response GenerateTotpSeedResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.generateTotpSeed, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GenerateTotpSeedResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GenerateTotpSeedResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GenerateTotpSeedResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GenerateTotpSeedResponse")
	}
	return
}

// generateTotpSeed implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) generateTotpSeed(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/users/{userId}/mfaTotpDevices/{mfaTotpDeviceId}/actions/generateSeed", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GenerateTotpSeedResponse
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

// GetAuthenticationPolicy Gets the authentication policy for the given tenancy. You must specify your tenant’s OCID as the value for
// the compartment ID (remember that the tenancy is simply the root compartment).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/GetAuthenticationPolicy.go.html to see an example of how to use GetAuthenticationPolicy API.
func (client IdentityClient) GetAuthenticationPolicy(ctx context.Context, request GetAuthenticationPolicyRequest) (response GetAuthenticationPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAuthenticationPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAuthenticationPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAuthenticationPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAuthenticationPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAuthenticationPolicyResponse")
	}
	return
}

// getAuthenticationPolicy implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) getAuthenticationPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/authenticationPolicies/{compartmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAuthenticationPolicyResponse
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

// GetCompartment Gets the specified compartment's information.
// This operation does not return a list of all the resources inside the compartment. There is no single
// API operation that does that. Compartments can contain multiple types of resources (instances, block
// storage volumes, etc.). To find out what's in a compartment, you must call the "List" operation for
// each resource type and specify the compartment's OCID as a query parameter in the request. For example,
// call the ListInstances operation in the Cloud Compute
// Service or the ListVolumes operation in Cloud Block Storage.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/GetCompartment.go.html to see an example of how to use GetCompartment API.
func (client IdentityClient) GetCompartment(ctx context.Context, request GetCompartmentRequest) (response GetCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetCompartmentResponse")
	}
	return
}

// getCompartment implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) getCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/compartments/{compartmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetCompartmentResponse
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

// GetDomain Get the specified domain's information.
// - If the domain doesn't exists, returns 404 NOT FOUND.
// - If any internal error occurs, returns 500 INTERNAL SERVER ERROR.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/GetDomain.go.html to see an example of how to use GetDomain API.
func (client IdentityClient) GetDomain(ctx context.Context, request GetDomainRequest) (response GetDomainResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDomain, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDomainResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDomainResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDomainResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDomainResponse")
	}
	return
}

// getDomain implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) getDomain(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/domains/{domainId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDomainResponse
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

// GetDynamicGroup Gets the specified dynamic group's information.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/GetDynamicGroup.go.html to see an example of how to use GetDynamicGroup API.
func (client IdentityClient) GetDynamicGroup(ctx context.Context, request GetDynamicGroupRequest) (response GetDynamicGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDynamicGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDynamicGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDynamicGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDynamicGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDynamicGroupResponse")
	}
	return
}

// getDynamicGroup implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) getDynamicGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dynamicGroups/{dynamicGroupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDynamicGroupResponse
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

// GetGroup Gets the specified group's information.
// This operation does not return a list of all the users in the group. To do that, use
// ListUserGroupMemberships and
// provide the group's OCID as a query parameter in the request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/GetGroup.go.html to see an example of how to use GetGroup API.
func (client IdentityClient) GetGroup(ctx context.Context, request GetGroupRequest) (response GetGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetGroupResponse")
	}
	return
}

// getGroup implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) getGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/groups/{groupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetGroupResponse
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

// GetIamWorkRequest Gets details on a specified IAM work request. For asynchronous operations in Identity and Access Management service, opc-work-request-id header values contains
// iam work request id that can be provided in this API to track the current status of the operation.
// - If workrequest exists, returns 202 ACCEPTED
// - If workrequest does not exist, returns 404 NOT FOUND
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/GetIamWorkRequest.go.html to see an example of how to use GetIamWorkRequest API.
func (client IdentityClient) GetIamWorkRequest(ctx context.Context, request GetIamWorkRequestRequest) (response GetIamWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getIamWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetIamWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetIamWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetIamWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetIamWorkRequestResponse")
	}
	return
}

// getIamWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) getIamWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/iamWorkRequests/{iamWorkRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetIamWorkRequestResponse
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

// GetIdentityProvider **Deprecated.** For more information, see Deprecated IAM Service APIs (https://docs.cloud.oracle.com/Content/Identity/Reference/deprecatediamapis.htm).
// Gets the specified identity provider's information.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/GetIdentityProvider.go.html to see an example of how to use GetIdentityProvider API.
func (client IdentityClient) GetIdentityProvider(ctx context.Context, request GetIdentityProviderRequest) (response GetIdentityProviderResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getIdentityProvider, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetIdentityProviderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetIdentityProviderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetIdentityProviderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetIdentityProviderResponse")
	}
	return
}

// getIdentityProvider implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) getIdentityProvider(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/identityProviders/{identityProviderId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetIdentityProviderResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &identityprovider{})
	return response, err
}

// GetIdpGroupMapping **Deprecated.** For more information, see Deprecated IAM Service APIs (https://docs.cloud.oracle.com/Content/Identity/Reference/deprecatediamapis.htm).
// Gets the specified group mapping.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/GetIdpGroupMapping.go.html to see an example of how to use GetIdpGroupMapping API.
func (client IdentityClient) GetIdpGroupMapping(ctx context.Context, request GetIdpGroupMappingRequest) (response GetIdpGroupMappingResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getIdpGroupMapping, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetIdpGroupMappingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetIdpGroupMappingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetIdpGroupMappingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetIdpGroupMappingResponse")
	}
	return
}

// getIdpGroupMapping implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) getIdpGroupMapping(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/identityProviders/{identityProviderId}/groupMappings/{mappingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetIdpGroupMappingResponse
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

// GetMfaTotpDevice Get the specified MFA TOTP device for the specified user.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/GetMfaTotpDevice.go.html to see an example of how to use GetMfaTotpDevice API.
func (client IdentityClient) GetMfaTotpDevice(ctx context.Context, request GetMfaTotpDeviceRequest) (response GetMfaTotpDeviceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getMfaTotpDevice, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetMfaTotpDeviceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetMfaTotpDeviceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetMfaTotpDeviceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetMfaTotpDeviceResponse")
	}
	return
}

// getMfaTotpDevice implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) getMfaTotpDevice(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/users/{userId}/mfaTotpDevices/{mfaTotpDeviceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetMfaTotpDeviceResponse
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

// GetNetworkSource Gets the specified network source's information.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/GetNetworkSource.go.html to see an example of how to use GetNetworkSource API.
func (client IdentityClient) GetNetworkSource(ctx context.Context, request GetNetworkSourceRequest) (response GetNetworkSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getNetworkSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetNetworkSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetNetworkSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetNetworkSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetNetworkSourceResponse")
	}
	return
}

// getNetworkSource implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) getNetworkSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkSources/{networkSourceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetNetworkSourceResponse
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

// GetPolicy Gets the specified policy's information.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/GetPolicy.go.html to see an example of how to use GetPolicy API.
func (client IdentityClient) GetPolicy(ctx context.Context, request GetPolicyRequest) (response GetPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPolicyResponse")
	}
	return
}

// getPolicy implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) getPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/policies/{policyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPolicyResponse
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

// GetTag Gets the specified tag's information.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/GetTag.go.html to see an example of how to use GetTag API.
func (client IdentityClient) GetTag(ctx context.Context, request GetTagRequest) (response GetTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTag, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTagResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTagResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTagResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTagResponse")
	}
	return
}

// getTag implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) getTag(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/tagNamespaces/{tagNamespaceId}/tags/{tagName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTagResponse
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

// GetTagDefault Retrieves the specified tag default.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/GetTagDefault.go.html to see an example of how to use GetTagDefault API.
func (client IdentityClient) GetTagDefault(ctx context.Context, request GetTagDefaultRequest) (response GetTagDefaultResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTagDefault, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTagDefaultResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTagDefaultResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTagDefaultResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTagDefaultResponse")
	}
	return
}

// getTagDefault implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) getTagDefault(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/tagDefaults/{tagDefaultId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTagDefaultResponse
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

// GetTagNamespace Gets the specified tag namespace's information.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/GetTagNamespace.go.html to see an example of how to use GetTagNamespace API.
func (client IdentityClient) GetTagNamespace(ctx context.Context, request GetTagNamespaceRequest) (response GetTagNamespaceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTagNamespace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTagNamespaceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTagNamespaceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTagNamespaceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTagNamespaceResponse")
	}
	return
}

// getTagNamespace implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) getTagNamespace(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/tagNamespaces/{tagNamespaceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTagNamespaceResponse
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

// GetTaggingWorkRequest Gets details on a specified work request. The workRequestID is returned in the opc-workrequest-id header
// for any asynchronous operation in tagging service.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/GetTaggingWorkRequest.go.html to see an example of how to use GetTaggingWorkRequest API.
func (client IdentityClient) GetTaggingWorkRequest(ctx context.Context, request GetTaggingWorkRequestRequest) (response GetTaggingWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTaggingWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTaggingWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTaggingWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTaggingWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTaggingWorkRequestResponse")
	}
	return
}

// getTaggingWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) getTaggingWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/taggingWorkRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTaggingWorkRequestResponse
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

// GetTenancy Get the specified tenancy's information.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/GetTenancy.go.html to see an example of how to use GetTenancy API.
func (client IdentityClient) GetTenancy(ctx context.Context, request GetTenancyRequest) (response GetTenancyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTenancy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTenancyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTenancyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTenancyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTenancyResponse")
	}
	return
}

// getTenancy implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) getTenancy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/tenancies/{tenancyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTenancyResponse
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

// GetUser Gets the specified user's information.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/GetUser.go.html to see an example of how to use GetUser API.
func (client IdentityClient) GetUser(ctx context.Context, request GetUserRequest) (response GetUserResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getUser, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetUserResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetUserResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetUserResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetUserResponse")
	}
	return
}

// getUser implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) getUser(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/users/{userId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetUserResponse
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

// GetUserGroupMembership Gets the specified UserGroupMembership's information.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/GetUserGroupMembership.go.html to see an example of how to use GetUserGroupMembership API.
func (client IdentityClient) GetUserGroupMembership(ctx context.Context, request GetUserGroupMembershipRequest) (response GetUserGroupMembershipResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getUserGroupMembership, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetUserGroupMembershipResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetUserGroupMembershipResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetUserGroupMembershipResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetUserGroupMembershipResponse")
	}
	return
}

// getUserGroupMembership implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) getUserGroupMembership(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/userGroupMemberships/{userGroupMembershipId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetUserGroupMembershipResponse
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

// GetUserUIPasswordInformation Gets the specified user's console password information. The returned object contains the user's OCID,
// but not the password itself. The actual password is returned only when created or reset.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/GetUserUIPasswordInformation.go.html to see an example of how to use GetUserUIPasswordInformation API.
func (client IdentityClient) GetUserUIPasswordInformation(ctx context.Context, request GetUserUIPasswordInformationRequest) (response GetUserUIPasswordInformationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getUserUIPasswordInformation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetUserUIPasswordInformationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetUserUIPasswordInformationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetUserUIPasswordInformationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetUserUIPasswordInformationResponse")
	}
	return
}

// getUserUIPasswordInformation implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) getUserUIPasswordInformation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/users/{userId}/uiPassword", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetUserUIPasswordInformationResponse
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

// GetWorkRequest Gets details on a specified work request. The workRequestID is returned in the opc-workrequest-id header
// for any asynchronous operation in the compartment service.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/GetWorkRequest.go.html to see an example of how to use GetWorkRequest API.
func (client IdentityClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
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
func (client IdentityClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAllowedDomainLicenseTypes List the allowed domain license types supported by OCI
// If {@code currentLicenseTypeName} provided, returns allowed license types a domain with the specified license type name can migrate to.
// If {@code name} is provided, it filters and returns resources that match the given license type name.
// Otherwise, returns all valid license types that are currently supported.
// - If license type details are retrieved sucessfully, return 202 ACCEPTED.
// - If any internal error occurs, return 500 INTERNAL SERVER ERROR.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListAllowedDomainLicenseTypes.go.html to see an example of how to use ListAllowedDomainLicenseTypes API.
func (client IdentityClient) ListAllowedDomainLicenseTypes(ctx context.Context, request ListAllowedDomainLicenseTypesRequest) (response ListAllowedDomainLicenseTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAllowedDomainLicenseTypes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAllowedDomainLicenseTypesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAllowedDomainLicenseTypesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAllowedDomainLicenseTypesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAllowedDomainLicenseTypesResponse")
	}
	return
}

// listAllowedDomainLicenseTypes implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) listAllowedDomainLicenseTypes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/allowedDomainLicenseTypes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAllowedDomainLicenseTypesResponse
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

// ListApiKeys Lists the API signing keys for the specified user. A user can have a maximum of three keys.
// Every user has permission to use this API call for *their own user ID*.  An administrator in your
// organization does not need to write a policy to give users this ability.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListApiKeys.go.html to see an example of how to use ListApiKeys API.
func (client IdentityClient) ListApiKeys(ctx context.Context, request ListApiKeysRequest) (response ListApiKeysResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listApiKeys, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListApiKeysResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListApiKeysResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListApiKeysResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListApiKeysResponse")
	}
	return
}

// listApiKeys implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) listApiKeys(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/users/{userId}/apiKeys", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListApiKeysResponse
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

// ListAuthTokens Lists the auth tokens for the specified user. The returned object contains the token's OCID, but not
// the token itself. The actual token is returned only upon creation.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListAuthTokens.go.html to see an example of how to use ListAuthTokens API.
func (client IdentityClient) ListAuthTokens(ctx context.Context, request ListAuthTokensRequest) (response ListAuthTokensResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAuthTokens, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAuthTokensResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAuthTokensResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAuthTokensResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAuthTokensResponse")
	}
	return
}

// listAuthTokens implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) listAuthTokens(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/users/{userId}/authTokens", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAuthTokensResponse
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

// ListAvailabilityDomains Lists the availability domains in your tenancy. Specify the OCID of either the tenancy or another
// of your compartments as the value for the compartment ID (remember that the tenancy is simply the root compartment).
// See Where to Get the Tenancy's OCID and User's OCID (https://docs.cloud.oracle.com/Content/API/Concepts/apisigningkey.htm#five).
// Note that the order of the results returned can change if availability domains are added or removed; therefore, do not
// create a dependency on the list order.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListAvailabilityDomains.go.html to see an example of how to use ListAvailabilityDomains API.
func (client IdentityClient) ListAvailabilityDomains(ctx context.Context, request ListAvailabilityDomainsRequest) (response ListAvailabilityDomainsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAvailabilityDomains, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAvailabilityDomainsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAvailabilityDomainsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAvailabilityDomainsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAvailabilityDomainsResponse")
	}
	return
}

// listAvailabilityDomains implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) listAvailabilityDomains(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/availabilityDomains", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAvailabilityDomainsResponse
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

// ListBulkActionResourceTypes Lists the resource-types supported by compartment bulk actions. Use this API to help you provide the correct
// resource-type information to the BulkDeleteResources
// and BulkMoveResources operations. The returned list of
// resource-types provides the appropriate resource-type names to use with the bulk action operations along with
// the type of identifying information you'll need to provide for each resource-type. Most resource-types just
// require an OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) to identify a specific resource, but some resource-types,
// such as buckets, require you to provide other identifying information.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListBulkActionResourceTypes.go.html to see an example of how to use ListBulkActionResourceTypes API.
func (client IdentityClient) ListBulkActionResourceTypes(ctx context.Context, request ListBulkActionResourceTypesRequest) (response ListBulkActionResourceTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listBulkActionResourceTypes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListBulkActionResourceTypesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListBulkActionResourceTypesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListBulkActionResourceTypesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListBulkActionResourceTypesResponse")
	}
	return
}

// listBulkActionResourceTypes implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) listBulkActionResourceTypes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/compartments/bulkActionResourceTypes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListBulkActionResourceTypesResponse
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

// ListBulkEditTagsResourceTypes Lists the resource types that support bulk tag editing.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListBulkEditTagsResourceTypes.go.html to see an example of how to use ListBulkEditTagsResourceTypes API.
func (client IdentityClient) ListBulkEditTagsResourceTypes(ctx context.Context, request ListBulkEditTagsResourceTypesRequest) (response ListBulkEditTagsResourceTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listBulkEditTagsResourceTypes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListBulkEditTagsResourceTypesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListBulkEditTagsResourceTypesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListBulkEditTagsResourceTypesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListBulkEditTagsResourceTypesResponse")
	}
	return
}

// listBulkEditTagsResourceTypes implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) listBulkEditTagsResourceTypes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/tags/bulkEditResourceTypes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListBulkEditTagsResourceTypesResponse
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

// ListCompartments Lists the compartments in a specified compartment. The members of the list
// returned depends on the values set for several parameters.
// With the exception of the tenancy (root compartment), the ListCompartments operation
// returns only the first-level child compartments in the parent compartment specified in
// `compartmentId`. The list does not include any subcompartments of the child
// compartments (grandchildren).
// The parameter `accessLevel` specifies whether to return only those compartments for which the
// requestor has INSPECT permissions on at least one resource directly
// or indirectly (the resource can be in a subcompartment).
// The parameter `compartmentIdInSubtree` applies only when you perform ListCompartments on the
// tenancy (root compartment). When set to true, the entire hierarchy of compartments can be returned.
// To get a full list of all compartments and subcompartments in the tenancy (root compartment),
// set the parameter `compartmentIdInSubtree` to true and `accessLevel` to ANY.
// See Where to Get the Tenancy's OCID and User's OCID (https://docs.cloud.oracle.com/Content/API/Concepts/apisigningkey.htm#five).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListCompartments.go.html to see an example of how to use ListCompartments API.
func (client IdentityClient) ListCompartments(ctx context.Context, request ListCompartmentsRequest) (response ListCompartmentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCompartments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCompartmentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCompartmentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCompartmentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCompartmentsResponse")
	}
	return
}

// listCompartments implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) listCompartments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/compartments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCompartmentsResponse
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

// ListCostTrackingTags Lists all the tags enabled for cost-tracking in the specified tenancy. For information about
// cost-tracking tags, see Using Cost-tracking Tags (https://docs.cloud.oracle.com/Content/Identity/Concepts/taggingoverview.htm#costs).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListCostTrackingTags.go.html to see an example of how to use ListCostTrackingTags API.
func (client IdentityClient) ListCostTrackingTags(ctx context.Context, request ListCostTrackingTagsRequest) (response ListCostTrackingTagsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCostTrackingTags, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCostTrackingTagsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCostTrackingTagsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCostTrackingTagsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCostTrackingTagsResponse")
	}
	return
}

// listCostTrackingTags implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) listCostTrackingTags(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/tagNamespaces/actions/listCostTrackingTags", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCostTrackingTagsResponse
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

// ListCustomerSecretKeys Lists the secret keys for the specified user. The returned object contains the secret key's OCID, but not
// the secret key itself. The actual secret key is returned only upon creation.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListCustomerSecretKeys.go.html to see an example of how to use ListCustomerSecretKeys API.
func (client IdentityClient) ListCustomerSecretKeys(ctx context.Context, request ListCustomerSecretKeysRequest) (response ListCustomerSecretKeysResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listCustomerSecretKeys, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListCustomerSecretKeysResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListCustomerSecretKeysResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListCustomerSecretKeysResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListCustomerSecretKeysResponse")
	}
	return
}

// listCustomerSecretKeys implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) listCustomerSecretKeys(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/users/{userId}/customerSecretKeys", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListCustomerSecretKeysResponse
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

// ListDomains List all domains that are homed or have a replica region in current region.
// - If any internal error occurs, return 500 INTERNAL SERVER ERROR.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListDomains.go.html to see an example of how to use ListDomains API.
func (client IdentityClient) ListDomains(ctx context.Context, request ListDomainsRequest) (response ListDomainsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDomains, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDomainsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDomainsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDomainsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDomainsResponse")
	}
	return
}

// listDomains implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) listDomains(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/domains", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDomainsResponse
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

// ListDynamicGroups Lists the dynamic groups in your tenancy. You must specify your tenancy's OCID as the value for
// the compartment ID (remember that the tenancy is simply the root compartment).
// See Where to Get the Tenancy's OCID and User's OCID (https://docs.cloud.oracle.com/Content/API/Concepts/apisigningkey.htm#five).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListDynamicGroups.go.html to see an example of how to use ListDynamicGroups API.
func (client IdentityClient) ListDynamicGroups(ctx context.Context, request ListDynamicGroupsRequest) (response ListDynamicGroupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDynamicGroups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDynamicGroupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDynamicGroupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDynamicGroupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDynamicGroupsResponse")
	}
	return
}

// listDynamicGroups implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) listDynamicGroups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/dynamicGroups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDynamicGroupsResponse
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

// ListFaultDomains Lists the Fault Domains in your tenancy. Specify the OCID of either the tenancy or another
// of your compartments as the value for the compartment ID (remember that the tenancy is simply the root compartment).
// See Where to Get the Tenancy's OCID and User's OCID (https://docs.cloud.oracle.com/Content/API/Concepts/apisigningkey.htm#five).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListFaultDomains.go.html to see an example of how to use ListFaultDomains API.
func (client IdentityClient) ListFaultDomains(ctx context.Context, request ListFaultDomainsRequest) (response ListFaultDomainsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listFaultDomains, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListFaultDomainsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListFaultDomainsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListFaultDomainsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListFaultDomainsResponse")
	}
	return
}

// listFaultDomains implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) listFaultDomains(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/faultDomains", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListFaultDomainsResponse
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

// ListGroups Lists the groups in your tenancy. You must specify your tenancy's OCID as the value for
// the compartment ID (remember that the tenancy is simply the root compartment).
// See Where to Get the Tenancy's OCID and User's OCID (https://docs.cloud.oracle.com/Content/API/Concepts/apisigningkey.htm#five).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListGroups.go.html to see an example of how to use ListGroups API.
func (client IdentityClient) ListGroups(ctx context.Context, request ListGroupsRequest) (response ListGroupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listGroups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListGroupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListGroupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListGroupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListGroupsResponse")
	}
	return
}

// listGroups implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) listGroups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/groups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListGroupsResponse
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

// ListIamWorkRequestErrors Gets error details for a specified IAM work request. For asynchronous operations in Identity and Access Management service, opc-work-request-id header values contains
// iam work request id that can be provided in this API to track the current status of the operation.
// - If workrequest exists, returns 202 ACCEPTED
// - If workrequest does not exist, returns 404 NOT FOUND
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListIamWorkRequestErrors.go.html to see an example of how to use ListIamWorkRequestErrors API.
func (client IdentityClient) ListIamWorkRequestErrors(ctx context.Context, request ListIamWorkRequestErrorsRequest) (response ListIamWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listIamWorkRequestErrors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListIamWorkRequestErrorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListIamWorkRequestErrorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListIamWorkRequestErrorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListIamWorkRequestErrorsResponse")
	}
	return
}

// listIamWorkRequestErrors implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) listIamWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/iamWorkRequests/{iamWorkRequestId}/errors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListIamWorkRequestErrorsResponse
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

// ListIamWorkRequestLogs Gets logs for a specified IAM work request. For asynchronous operations in Identity and Access Management service, opc-work-request-id header values contains
// iam work request id that can be provided in this API to track the current status of the operation.
// - If workrequest exists, returns 202 ACCEPTED
// - If workrequest does not exist, returns 404 NOT FOUND
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListIamWorkRequestLogs.go.html to see an example of how to use ListIamWorkRequestLogs API.
func (client IdentityClient) ListIamWorkRequestLogs(ctx context.Context, request ListIamWorkRequestLogsRequest) (response ListIamWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listIamWorkRequestLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListIamWorkRequestLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListIamWorkRequestLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListIamWorkRequestLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListIamWorkRequestLogsResponse")
	}
	return
}

// listIamWorkRequestLogs implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) listIamWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/iamWorkRequests/{iamWorkRequestId}/logs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListIamWorkRequestLogsResponse
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

// ListIamWorkRequests List the IAM work requests in compartment
// - If IAM workrequest  details are retrieved sucessfully, return 202 ACCEPTED.
// - If any internal error occurs, return 500 INTERNAL SERVER ERROR.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListIamWorkRequests.go.html to see an example of how to use ListIamWorkRequests API.
func (client IdentityClient) ListIamWorkRequests(ctx context.Context, request ListIamWorkRequestsRequest) (response ListIamWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listIamWorkRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListIamWorkRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListIamWorkRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListIamWorkRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListIamWorkRequestsResponse")
	}
	return
}

// listIamWorkRequests implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) listIamWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/iamWorkRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListIamWorkRequestsResponse
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

// ListIdentityProviderGroups **Deprecated.** For more information, see Deprecated IAM Service APIs (https://docs.cloud.oracle.com/Content/Identity/Reference/deprecatediamapis.htm).
// Lists the identity provider groups.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListIdentityProviderGroups.go.html to see an example of how to use ListIdentityProviderGroups API.
func (client IdentityClient) ListIdentityProviderGroups(ctx context.Context, request ListIdentityProviderGroupsRequest) (response ListIdentityProviderGroupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listIdentityProviderGroups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListIdentityProviderGroupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListIdentityProviderGroupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListIdentityProviderGroupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListIdentityProviderGroupsResponse")
	}
	return
}

// listIdentityProviderGroups implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) listIdentityProviderGroups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/identityProviders/{identityProviderId}/groups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListIdentityProviderGroupsResponse
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

//listidentityprovider allows to unmarshal list of polymorphic IdentityProvider
type listidentityprovider []identityprovider

//UnmarshalPolymorphicJSON unmarshals polymorphic json list of items
func (m *listidentityprovider) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {
	res := make([]IdentityProvider, len(*m))
	for i, v := range *m {
		nn, err := v.UnmarshalPolymorphicJSON(v.JsonData)
		if err != nil {
			return nil, err
		}
		res[i] = nn.(IdentityProvider)
	}
	return res, nil
}

// ListIdentityProviders **Deprecated.** For more information, see Deprecated IAM Service APIs (https://docs.cloud.oracle.com/Content/Identity/Reference/deprecatediamapis.htm).
// Lists all the identity providers in your tenancy. You must specify the identity provider type (e.g., `SAML2` for
// identity providers using the SAML2.0 protocol). You must specify your tenancy's OCID as the value for the
// compartment ID (remember that the tenancy is simply the root compartment).
// See Where to Get the Tenancy's OCID and User's OCID (https://docs.cloud.oracle.com/Content/API/Concepts/apisigningkey.htm#five).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListIdentityProviders.go.html to see an example of how to use ListIdentityProviders API.
func (client IdentityClient) ListIdentityProviders(ctx context.Context, request ListIdentityProvidersRequest) (response ListIdentityProvidersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listIdentityProviders, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListIdentityProvidersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListIdentityProvidersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListIdentityProvidersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListIdentityProvidersResponse")
	}
	return
}

// listIdentityProviders implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) listIdentityProviders(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/identityProviders", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListIdentityProvidersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &listidentityprovider{})
	return response, err
}

// ListIdpGroupMappings **Deprecated.** For more information, see Deprecated IAM Service APIs (https://docs.cloud.oracle.com/Content/Identity/Reference/deprecatediamapis.htm).
// Lists the group mappings for the specified identity provider.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListIdpGroupMappings.go.html to see an example of how to use ListIdpGroupMappings API.
func (client IdentityClient) ListIdpGroupMappings(ctx context.Context, request ListIdpGroupMappingsRequest) (response ListIdpGroupMappingsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listIdpGroupMappings, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListIdpGroupMappingsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListIdpGroupMappingsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListIdpGroupMappingsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListIdpGroupMappingsResponse")
	}
	return
}

// listIdpGroupMappings implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) listIdpGroupMappings(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/identityProviders/{identityProviderId}/groupMappings", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListIdpGroupMappingsResponse
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

// ListMfaTotpDevices Lists the MFA TOTP devices for the specified user. The returned object contains the device's OCID, but not
// the seed. The seed is returned only upon creation or when the IAM service regenerates the MFA seed for the device.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListMfaTotpDevices.go.html to see an example of how to use ListMfaTotpDevices API.
func (client IdentityClient) ListMfaTotpDevices(ctx context.Context, request ListMfaTotpDevicesRequest) (response ListMfaTotpDevicesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMfaTotpDevices, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMfaTotpDevicesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMfaTotpDevicesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMfaTotpDevicesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMfaTotpDevicesResponse")
	}
	return
}

// listMfaTotpDevices implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) listMfaTotpDevices(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/users/{userId}/mfaTotpDevices", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListMfaTotpDevicesResponse
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

// ListNetworkSources Lists the network sources in your tenancy. You must specify your tenancy's OCID as the value for
// the compartment ID (remember that the tenancy is simply the root compartment).
// See Where to Get the Tenancy's OCID and User's OCID (https://docs.cloud.oracle.com/Content/API/Concepts/apisigningkey.htm#five).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListNetworkSources.go.html to see an example of how to use ListNetworkSources API.
func (client IdentityClient) ListNetworkSources(ctx context.Context, request ListNetworkSourcesRequest) (response ListNetworkSourcesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listNetworkSources, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListNetworkSourcesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListNetworkSourcesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListNetworkSourcesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListNetworkSourcesResponse")
	}
	return
}

// listNetworkSources implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) listNetworkSources(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkSources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListNetworkSourcesResponse
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

// ListOAuthClientCredentials List of Oauth tokens for the user
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListOAuthClientCredentials.go.html to see an example of how to use ListOAuthClientCredentials API.
func (client IdentityClient) ListOAuthClientCredentials(ctx context.Context, request ListOAuthClientCredentialsRequest) (response ListOAuthClientCredentialsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listOAuthClientCredentials, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListOAuthClientCredentialsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListOAuthClientCredentialsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListOAuthClientCredentialsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListOAuthClientCredentialsResponse")
	}
	return
}

// listOAuthClientCredentials implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) listOAuthClientCredentials(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/users/{userId}/oauth2ClientCredentials", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListOAuthClientCredentialsResponse
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

// ListPolicies Lists the policies in the specified compartment (either the tenancy or another of your compartments).
// See Where to Get the Tenancy's OCID and User's OCID (https://docs.cloud.oracle.com/Content/API/Concepts/apisigningkey.htm#five).
// To determine which policies apply to a particular group or compartment, you must view the individual
// statements inside all your policies. There isn't a way to automatically obtain that information via the API.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListPolicies.go.html to see an example of how to use ListPolicies API.
func (client IdentityClient) ListPolicies(ctx context.Context, request ListPoliciesRequest) (response ListPoliciesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listPolicies, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListPoliciesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListPoliciesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListPoliciesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListPoliciesResponse")
	}
	return
}

// listPolicies implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) listPolicies(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/policies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListPoliciesResponse
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

// ListRegionSubscriptions Lists the region subscriptions for the specified tenancy.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListRegionSubscriptions.go.html to see an example of how to use ListRegionSubscriptions API.
func (client IdentityClient) ListRegionSubscriptions(ctx context.Context, request ListRegionSubscriptionsRequest) (response ListRegionSubscriptionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRegionSubscriptions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRegionSubscriptionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRegionSubscriptionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRegionSubscriptionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRegionSubscriptionsResponse")
	}
	return
}

// listRegionSubscriptions implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) listRegionSubscriptions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/tenancies/{tenancyId}/regionSubscriptions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListRegionSubscriptionsResponse
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

// ListRegions Lists all the regions offered by Oracle Cloud Infrastructure.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListRegions.go.html to see an example of how to use ListRegions API.
func (client IdentityClient) ListRegions(ctx context.Context) (response ListRegionsResponse, err error) {
	var ociResponse common.OCIResponse
	ociResponse, err = client.listRegions(ctx)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRegionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRegionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRegionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRegionsResponse")
	}
	return
}

// listRegions performs the request (retry policy is not enabled without a request object)
func (client IdentityClient) listRegions(ctx context.Context) (common.OCIResponse, error) {
	httpRequest := common.MakeDefaultHTTPRequest(http.MethodGet, "/regions")
	var err error

	var response ListRegionsResponse
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

// ListSmtpCredentials Lists the SMTP credentials for the specified user. The returned object contains the credential's OCID,
// the SMTP user name but not the SMTP password. The SMTP password is returned only upon creation.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListSmtpCredentials.go.html to see an example of how to use ListSmtpCredentials API.
func (client IdentityClient) ListSmtpCredentials(ctx context.Context, request ListSmtpCredentialsRequest) (response ListSmtpCredentialsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSmtpCredentials, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSmtpCredentialsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSmtpCredentialsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSmtpCredentialsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSmtpCredentialsResponse")
	}
	return
}

// listSmtpCredentials implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) listSmtpCredentials(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/users/{userId}/smtpCredentials", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSmtpCredentialsResponse
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

// ListSwiftPasswords **Deprecated. Use ListAuthTokens instead.**
// Lists the Swift passwords for the specified user. The returned object contains the password's OCID, but not
// the password itself. The actual password is returned only upon creation.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListSwiftPasswords.go.html to see an example of how to use ListSwiftPasswords API.
func (client IdentityClient) ListSwiftPasswords(ctx context.Context, request ListSwiftPasswordsRequest) (response ListSwiftPasswordsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSwiftPasswords, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSwiftPasswordsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSwiftPasswordsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSwiftPasswordsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSwiftPasswordsResponse")
	}
	return
}

// listSwiftPasswords implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) listSwiftPasswords(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/users/{userId}/swiftPasswords", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSwiftPasswordsResponse
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

// ListTagDefaults Lists the tag defaults for tag definitions in the specified compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListTagDefaults.go.html to see an example of how to use ListTagDefaults API.
func (client IdentityClient) ListTagDefaults(ctx context.Context, request ListTagDefaultsRequest) (response ListTagDefaultsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTagDefaults, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTagDefaultsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTagDefaultsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTagDefaultsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTagDefaultsResponse")
	}
	return
}

// listTagDefaults implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) listTagDefaults(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/tagDefaults", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTagDefaultsResponse
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

// ListTagNamespaces Lists the tag namespaces in the specified compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListTagNamespaces.go.html to see an example of how to use ListTagNamespaces API.
func (client IdentityClient) ListTagNamespaces(ctx context.Context, request ListTagNamespacesRequest) (response ListTagNamespacesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTagNamespaces, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTagNamespacesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTagNamespacesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTagNamespacesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTagNamespacesResponse")
	}
	return
}

// listTagNamespaces implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) listTagNamespaces(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/tagNamespaces", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTagNamespacesResponse
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

// ListTaggingWorkRequestErrors Gets the errors for a work request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListTaggingWorkRequestErrors.go.html to see an example of how to use ListTaggingWorkRequestErrors API.
func (client IdentityClient) ListTaggingWorkRequestErrors(ctx context.Context, request ListTaggingWorkRequestErrorsRequest) (response ListTaggingWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTaggingWorkRequestErrors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTaggingWorkRequestErrorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTaggingWorkRequestErrorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTaggingWorkRequestErrorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTaggingWorkRequestErrorsResponse")
	}
	return
}

// listTaggingWorkRequestErrors implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) listTaggingWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/taggingWorkRequests/{workRequestId}/errors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTaggingWorkRequestErrorsResponse
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

// ListTaggingWorkRequestLogs Gets the logs for a work request.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListTaggingWorkRequestLogs.go.html to see an example of how to use ListTaggingWorkRequestLogs API.
func (client IdentityClient) ListTaggingWorkRequestLogs(ctx context.Context, request ListTaggingWorkRequestLogsRequest) (response ListTaggingWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTaggingWorkRequestLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTaggingWorkRequestLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTaggingWorkRequestLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTaggingWorkRequestLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTaggingWorkRequestLogsResponse")
	}
	return
}

// listTaggingWorkRequestLogs implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) listTaggingWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/taggingWorkRequests/{workRequestId}/logs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTaggingWorkRequestLogsResponse
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

// ListTaggingWorkRequests Lists the tagging work requests in compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListTaggingWorkRequests.go.html to see an example of how to use ListTaggingWorkRequests API.
func (client IdentityClient) ListTaggingWorkRequests(ctx context.Context, request ListTaggingWorkRequestsRequest) (response ListTaggingWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTaggingWorkRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTaggingWorkRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTaggingWorkRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTaggingWorkRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTaggingWorkRequestsResponse")
	}
	return
}

// listTaggingWorkRequests implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) listTaggingWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/taggingWorkRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTaggingWorkRequestsResponse
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

// ListTags Lists the tag definitions in the specified tag namespace.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListTags.go.html to see an example of how to use ListTags API.
func (client IdentityClient) ListTags(ctx context.Context, request ListTagsRequest) (response ListTagsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTags, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTagsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTagsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTagsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTagsResponse")
	}
	return
}

// listTags implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) listTags(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/tagNamespaces/{tagNamespaceId}/tags", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTagsResponse
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

// ListUserGroupMemberships Lists the `UserGroupMembership` objects in your tenancy. You must specify your tenancy's OCID
// as the value for the compartment ID
// (see Where to Get the Tenancy's OCID and User's OCID (https://docs.cloud.oracle.com/Content/API/Concepts/apisigningkey.htm#five)).
// You must also then filter the list in one of these ways:
// - You can limit the results to just the memberships for a given user by specifying a `userId`.
// - Similarly, you can limit the results to just the memberships for a given group by specifying a `groupId`.
// - You can set both the `userId` and `groupId` to determine if the specified user is in the specified group.
// If the answer is no, the response is an empty list.
// - Although`userId` and `groupId` are not individually required, you must set one of them.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListUserGroupMemberships.go.html to see an example of how to use ListUserGroupMemberships API.
func (client IdentityClient) ListUserGroupMemberships(ctx context.Context, request ListUserGroupMembershipsRequest) (response ListUserGroupMembershipsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listUserGroupMemberships, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListUserGroupMembershipsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListUserGroupMembershipsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListUserGroupMembershipsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListUserGroupMembershipsResponse")
	}
	return
}

// listUserGroupMemberships implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) listUserGroupMemberships(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/userGroupMemberships", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListUserGroupMembershipsResponse
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

// ListUsers Lists the users in your tenancy. You must specify your tenancy's OCID as the value for the
// compartment ID (remember that the tenancy is simply the root compartment).
// See Where to Get the Tenancy's OCID and User's OCID (https://docs.cloud.oracle.com/Content/API/Concepts/apisigningkey.htm#five).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListUsers.go.html to see an example of how to use ListUsers API.
func (client IdentityClient) ListUsers(ctx context.Context, request ListUsersRequest) (response ListUsersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listUsers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListUsersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListUsersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListUsersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListUsersResponse")
	}
	return
}

// listUsers implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) listUsers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/users", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListUsersResponse
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

// ListWorkRequests Lists the work requests in compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ListWorkRequests.go.html to see an example of how to use ListWorkRequests API.
func (client IdentityClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
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
func (client IdentityClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

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
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// MoveCompartment Move the compartment to a different parent compartment in the same tenancy. When you move a
// compartment, all its contents (subcompartments and resources) are moved with it. Note that
// the `CompartmentId` that you specify in the path is the compartment that you want to move.
// **IMPORTANT**: After you move a compartment to a new parent compartment, the access policies of
// the new parent take effect and the policies of the previous parent no longer apply. Ensure that you
// are aware of the implications for the compartment contents before you move it. For more
// information, see Moving a Compartment (https://docs.cloud.oracle.com/Content/Identity/Tasks/managingcompartments.htm#MoveCompartment).
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/MoveCompartment.go.html to see an example of how to use MoveCompartment API.
func (client IdentityClient) MoveCompartment(ctx context.Context, request MoveCompartmentRequest) (response MoveCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.moveCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = MoveCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = MoveCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(MoveCompartmentResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into MoveCompartmentResponse")
	}
	return
}

// moveCompartment implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) moveCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/compartments/{compartmentId}/actions/moveCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response MoveCompartmentResponse
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

// RecoverCompartment Recover the compartment from DELETED state to ACTIVE state.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/RecoverCompartment.go.html to see an example of how to use RecoverCompartment API.
func (client IdentityClient) RecoverCompartment(ctx context.Context, request RecoverCompartmentRequest) (response RecoverCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.recoverCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RecoverCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RecoverCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RecoverCompartmentResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RecoverCompartmentResponse")
	}
	return
}

// recoverCompartment implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) recoverCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/compartments/{compartmentId}/actions/recoverCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RecoverCompartmentResponse
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

// RemoveUserFromGroup Removes a user from a group by deleting the corresponding `UserGroupMembership`.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/RemoveUserFromGroup.go.html to see an example of how to use RemoveUserFromGroup API.
func (client IdentityClient) RemoveUserFromGroup(ctx context.Context, request RemoveUserFromGroupRequest) (response RemoveUserFromGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.removeUserFromGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveUserFromGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveUserFromGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveUserFromGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveUserFromGroupResponse")
	}
	return
}

// removeUserFromGroup implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) removeUserFromGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/userGroupMemberships/{userGroupMembershipId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveUserFromGroupResponse
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

// ResetIdpScimClient Resets the OAuth2 client credentials for the SCIM client associated with this identity provider.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/ResetIdpScimClient.go.html to see an example of how to use ResetIdpScimClient API.
func (client IdentityClient) ResetIdpScimClient(ctx context.Context, request ResetIdpScimClientRequest) (response ResetIdpScimClientResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.resetIdpScimClient, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ResetIdpScimClientResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ResetIdpScimClientResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ResetIdpScimClientResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ResetIdpScimClientResponse")
	}
	return
}

// resetIdpScimClient implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) resetIdpScimClient(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/identityProviders/{identityProviderId}/actions/resetScimClient", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ResetIdpScimClientResponse
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

// UpdateAuthToken Updates the specified auth token's description.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/UpdateAuthToken.go.html to see an example of how to use UpdateAuthToken API.
func (client IdentityClient) UpdateAuthToken(ctx context.Context, request UpdateAuthTokenRequest) (response UpdateAuthTokenResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAuthToken, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAuthTokenResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAuthTokenResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAuthTokenResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAuthTokenResponse")
	}
	return
}

// updateAuthToken implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) updateAuthToken(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/users/{userId}/authTokens/{authTokenId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAuthTokenResponse
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

// UpdateAuthenticationPolicy Updates authentication policy for the specified tenancy
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/UpdateAuthenticationPolicy.go.html to see an example of how to use UpdateAuthenticationPolicy API.
func (client IdentityClient) UpdateAuthenticationPolicy(ctx context.Context, request UpdateAuthenticationPolicyRequest) (response UpdateAuthenticationPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAuthenticationPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAuthenticationPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAuthenticationPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAuthenticationPolicyResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAuthenticationPolicyResponse")
	}
	return
}

// updateAuthenticationPolicy implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) updateAuthenticationPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/authenticationPolicies/{compartmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAuthenticationPolicyResponse
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

// UpdateCompartment Updates the specified compartment's description or name. You can't update the root compartment.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/UpdateCompartment.go.html to see an example of how to use UpdateCompartment API.
func (client IdentityClient) UpdateCompartment(ctx context.Context, request UpdateCompartmentRequest) (response UpdateCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateCompartmentResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateCompartmentResponse")
	}
	return
}

// updateCompartment implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) updateCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/compartments/{compartmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateCompartmentResponse
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

// UpdateCustomerSecretKey Updates the specified secret key's description.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/UpdateCustomerSecretKey.go.html to see an example of how to use UpdateCustomerSecretKey API.
func (client IdentityClient) UpdateCustomerSecretKey(ctx context.Context, request UpdateCustomerSecretKeyRequest) (response UpdateCustomerSecretKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateCustomerSecretKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateCustomerSecretKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateCustomerSecretKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateCustomerSecretKeyResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateCustomerSecretKeyResponse")
	}
	return
}

// updateCustomerSecretKey implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) updateCustomerSecretKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/users/{userId}/customerSecretKeys/{customerSecretKeyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateCustomerSecretKeyResponse
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

// UpdateDomain Updates domain information and associated stripe. This is an asynchronous call where
// the associated stripe and domain are updated.
// To track progress, HTTP GET on /iamWorkRequests/{iamWorkRequestsId} endpoint will provide
// the async operation's status.
// - If the {@code displayName} is not unique within the tenancy, returns 400 BAD REQUEST.
// - If any field other than {@code description} is requested to be updated for DEFAULT domain,
// returns 400 BAD REQUEST.
// - If Domain is not active or being updated, returns 400 BAD REQUEST.
// - If Domain {@code type} is DEFAULT or DEFAULT_LIGHTWEIGHT, return 400 BAD Request
// - If the domain doesn't exists, returns 404 NOT FOUND.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/UpdateDomain.go.html to see an example of how to use UpdateDomain API.
func (client IdentityClient) UpdateDomain(ctx context.Context, request UpdateDomainRequest) (response UpdateDomainResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDomain, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDomainResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDomainResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDomainResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDomainResponse")
	}
	return
}

// updateDomain implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) updateDomain(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/domains/{domainId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDomainResponse
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

// UpdateDynamicGroup Updates the specified dynamic group.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/UpdateDynamicGroup.go.html to see an example of how to use UpdateDynamicGroup API.
func (client IdentityClient) UpdateDynamicGroup(ctx context.Context, request UpdateDynamicGroupRequest) (response UpdateDynamicGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDynamicGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDynamicGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDynamicGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDynamicGroupResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDynamicGroupResponse")
	}
	return
}

// updateDynamicGroup implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) updateDynamicGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/dynamicGroups/{dynamicGroupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDynamicGroupResponse
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

// UpdateGroup Updates the specified group.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/UpdateGroup.go.html to see an example of how to use UpdateGroup API.
func (client IdentityClient) UpdateGroup(ctx context.Context, request UpdateGroupRequest) (response UpdateGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateGroupResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateGroupResponse")
	}
	return
}

// updateGroup implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) updateGroup(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/groups/{groupId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateGroupResponse
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

// UpdateIdentityProvider **Deprecated.** For more information, see Deprecated IAM Service APIs (https://docs.cloud.oracle.com/Content/Identity/Reference/deprecatediamapis.htm).
// Updates the specified identity provider.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/UpdateIdentityProvider.go.html to see an example of how to use UpdateIdentityProvider API.
func (client IdentityClient) UpdateIdentityProvider(ctx context.Context, request UpdateIdentityProviderRequest) (response UpdateIdentityProviderResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateIdentityProvider, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateIdentityProviderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateIdentityProviderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateIdentityProviderResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateIdentityProviderResponse")
	}
	return
}

// updateIdentityProvider implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) updateIdentityProvider(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/identityProviders/{identityProviderId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateIdentityProviderResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &identityprovider{})
	return response, err
}

// UpdateIdpGroupMapping **Deprecated.** For more information, see Deprecated IAM Service APIs (https://docs.cloud.oracle.com/Content/Identity/Reference/deprecatediamapis.htm).
// Updates the specified group mapping.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/UpdateIdpGroupMapping.go.html to see an example of how to use UpdateIdpGroupMapping API.
func (client IdentityClient) UpdateIdpGroupMapping(ctx context.Context, request UpdateIdpGroupMappingRequest) (response UpdateIdpGroupMappingResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateIdpGroupMapping, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateIdpGroupMappingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateIdpGroupMappingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateIdpGroupMappingResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateIdpGroupMappingResponse")
	}
	return
}

// updateIdpGroupMapping implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) updateIdpGroupMapping(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/identityProviders/{identityProviderId}/groupMappings/{mappingId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateIdpGroupMappingResponse
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

// UpdateNetworkSource Updates the specified network source.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/UpdateNetworkSource.go.html to see an example of how to use UpdateNetworkSource API.
func (client IdentityClient) UpdateNetworkSource(ctx context.Context, request UpdateNetworkSourceRequest) (response UpdateNetworkSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateNetworkSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateNetworkSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateNetworkSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateNetworkSourceResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateNetworkSourceResponse")
	}
	return
}

// updateNetworkSource implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) updateNetworkSource(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/networkSources/{networkSourceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateNetworkSourceResponse
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

// UpdateOAuthClientCredential Updates Oauth token for the user
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/UpdateOAuthClientCredential.go.html to see an example of how to use UpdateOAuthClientCredential API.
func (client IdentityClient) UpdateOAuthClientCredential(ctx context.Context, request UpdateOAuthClientCredentialRequest) (response UpdateOAuthClientCredentialResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateOAuthClientCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateOAuthClientCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateOAuthClientCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateOAuthClientCredentialResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateOAuthClientCredentialResponse")
	}
	return
}

// updateOAuthClientCredential implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) updateOAuthClientCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/users/{userId}/oauth2ClientCredentials/{oauth2ClientCredentialId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateOAuthClientCredentialResponse
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

// UpdatePolicy Updates the specified policy. You can update the description or the policy statements themselves.
// Policy changes take effect typically within 10 seconds.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/UpdatePolicy.go.html to see an example of how to use UpdatePolicy API.
func (client IdentityClient) UpdatePolicy(ctx context.Context, request UpdatePolicyRequest) (response UpdatePolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updatePolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdatePolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdatePolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdatePolicyResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdatePolicyResponse")
	}
	return
}

// updatePolicy implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) updatePolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/policies/{policyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdatePolicyResponse
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

// UpdateSmtpCredential Updates the specified SMTP credential's description.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/UpdateSmtpCredential.go.html to see an example of how to use UpdateSmtpCredential API.
func (client IdentityClient) UpdateSmtpCredential(ctx context.Context, request UpdateSmtpCredentialRequest) (response UpdateSmtpCredentialResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSmtpCredential, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSmtpCredentialResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSmtpCredentialResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSmtpCredentialResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSmtpCredentialResponse")
	}
	return
}

// updateSmtpCredential implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) updateSmtpCredential(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/users/{userId}/smtpCredentials/{smtpCredentialId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSmtpCredentialResponse
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

// UpdateSwiftPassword **Deprecated. Use UpdateAuthToken instead.**
// Updates the specified Swift password's description.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/UpdateSwiftPassword.go.html to see an example of how to use UpdateSwiftPassword API.
func (client IdentityClient) UpdateSwiftPassword(ctx context.Context, request UpdateSwiftPasswordRequest) (response UpdateSwiftPasswordResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSwiftPassword, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSwiftPasswordResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSwiftPasswordResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSwiftPasswordResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSwiftPasswordResponse")
	}
	return
}

// updateSwiftPassword implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) updateSwiftPassword(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/users/{userId}/swiftPasswords/{swiftPasswordId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSwiftPasswordResponse
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

// UpdateTag Updates the specified tag definition.
// Setting `validator` determines the value type. Tags can use either a static value or a
// list of possible values. Static values are entered by a user applying the tag to a resource.
// Lists are created by you and the user must apply a value from the list. On update, any values
// in a list that were previously set do not change, but new values must pass validation. Values
// already applied to a resource do not change.
// You cannot remove list values that appear in a TagDefault. To remove a list value that
// appears in a TagDefault, first update the TagDefault to use a different value.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/UpdateTag.go.html to see an example of how to use UpdateTag API.
func (client IdentityClient) UpdateTag(ctx context.Context, request UpdateTagRequest) (response UpdateTagResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTag, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTagResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTagResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTagResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTagResponse")
	}
	return
}

// updateTag implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) updateTag(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/tagNamespaces/{tagNamespaceId}/tags/{tagName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateTagResponse
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

// UpdateTagDefault Updates the specified tag default. If you specify that a value is required, a value is set
// during resource creation (either by the user creating the resource or another tag defualt).
// If no value is set, resource creation is blocked.
// * If the `isRequired` flag is set to "true", the value is set during resource creation.
// * If the `isRequired` flag is set to "false", the value you enter is set during resource creation.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/UpdateTagDefault.go.html to see an example of how to use UpdateTagDefault API.
func (client IdentityClient) UpdateTagDefault(ctx context.Context, request UpdateTagDefaultRequest) (response UpdateTagDefaultResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTagDefault, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTagDefaultResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTagDefaultResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTagDefaultResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTagDefaultResponse")
	}
	return
}

// updateTagDefault implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) updateTagDefault(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/tagDefaults/{tagDefaultId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateTagDefaultResponse
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

// UpdateTagNamespace Updates the the specified tag namespace. You can't update the namespace name.
// Updating `isRetired` to 'true' retires the namespace and all the tag definitions in the namespace. Reactivating a
// namespace (changing `isRetired` from 'true' to 'false') does not reactivate tag definitions.
// To reactivate the tag definitions, you must reactivate each one individually *after* you reactivate the namespace,
// using UpdateTag. For more information about retiring tag namespaces, see
// Retiring Key Definitions and Namespace Definitions (https://docs.cloud.oracle.com/Content/Identity/Concepts/taggingoverview.htm#Retiring).
// You can't add a namespace with the same name as a retired namespace in the same tenancy.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/UpdateTagNamespace.go.html to see an example of how to use UpdateTagNamespace API.
func (client IdentityClient) UpdateTagNamespace(ctx context.Context, request UpdateTagNamespaceRequest) (response UpdateTagNamespaceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTagNamespace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTagNamespaceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTagNamespaceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTagNamespaceResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTagNamespaceResponse")
	}
	return
}

// updateTagNamespace implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) updateTagNamespace(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/tagNamespaces/{tagNamespaceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateTagNamespaceResponse
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

// UpdateUser Updates the description of the specified user.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/UpdateUser.go.html to see an example of how to use UpdateUser API.
func (client IdentityClient) UpdateUser(ctx context.Context, request UpdateUserRequest) (response UpdateUserResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateUser, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateUserResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateUserResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateUserResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateUserResponse")
	}
	return
}

// updateUser implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) updateUser(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/users/{userId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateUserResponse
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

// UpdateUserCapabilities Updates the capabilities of the specified user.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/UpdateUserCapabilities.go.html to see an example of how to use UpdateUserCapabilities API.
func (client IdentityClient) UpdateUserCapabilities(ctx context.Context, request UpdateUserCapabilitiesRequest) (response UpdateUserCapabilitiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateUserCapabilities, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateUserCapabilitiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateUserCapabilitiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateUserCapabilitiesResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateUserCapabilitiesResponse")
	}
	return
}

// updateUserCapabilities implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) updateUserCapabilities(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/users/{userId}/capabilities", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateUserCapabilitiesResponse
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

// UpdateUserState Updates the state of the specified user.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/UpdateUserState.go.html to see an example of how to use UpdateUserState API.
func (client IdentityClient) UpdateUserState(ctx context.Context, request UpdateUserStateRequest) (response UpdateUserStateResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateUserState, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateUserStateResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateUserStateResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateUserStateResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateUserStateResponse")
	}
	return
}

// updateUserState implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) updateUserState(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/users/{userId}/state", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateUserStateResponse
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

// UploadApiKey Uploads an API signing key for the specified user.
// Every user has permission to use this operation to upload a key for *their own user ID*. An
// administrator in your organization does not need to write a policy to give users this ability.
// To compare, administrators who have permission to the tenancy can use this operation to upload a
// key for any user, including themselves.
// **Important:** Even though you have permission to upload an API key, you might not yet
// have permission to do much else. If you try calling an operation unrelated to your own credential
// management (e.g., `ListUsers`, `LaunchInstance`) and receive an "unauthorized" error,
// check with an administrator to confirm which IAM Service group(s) you're in and what access
// you have. Also confirm you're working in the correct compartment.
// After you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before using
// the object, first make sure its `lifecycleState` has changed to ACTIVE.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identity/UploadApiKey.go.html to see an example of how to use UploadApiKey API.
func (client IdentityClient) UploadApiKey(ctx context.Context, request UploadApiKeyRequest) (response UploadApiKeyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.uploadApiKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UploadApiKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UploadApiKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UploadApiKeyResponse); ok {
		common.EcContext.UpdateEndOfWindow(time.Duration(240 * time.Second))
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UploadApiKeyResponse")
	}
	return
}

// uploadApiKey implements the OCIOperation interface (enables retrying operations)
func (client IdentityClient) uploadApiKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/users/{userId}/apiKeys", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UploadApiKeyResponse
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
