// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Security Attribute API
//
// Use the Security Attributes API to manage security attributes and security attribute namespaces. For more information, see the documentation for Security Attributes (https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/managing-security-attributes.htm) and Security Attribute Nampespaces (https://docs.cloud.oracle.com/iaas/Content/zero-trust-packet-routing/managing-security-attribute-namespaces.htm).
//

package securityattribute

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// SecurityAttributeClient a client for SecurityAttribute
type SecurityAttributeClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewSecurityAttributeClientWithConfigurationProvider Creates a new default SecurityAttribute client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewSecurityAttributeClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client SecurityAttributeClient, err error) {
	if enabled := common.CheckForEnabledServices("securityattribute"); !enabled {
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
	return newSecurityAttributeClientFromBaseClient(baseClient, provider)
}

// NewSecurityAttributeClientWithOboToken Creates a new default SecurityAttribute client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewSecurityAttributeClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client SecurityAttributeClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newSecurityAttributeClientFromBaseClient(baseClient, configProvider)
}

func newSecurityAttributeClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client SecurityAttributeClient, err error) {
	// SecurityAttribute service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("SecurityAttribute"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = SecurityAttributeClient{BaseClient: baseClient}
	client.BasePath = "20240815"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *SecurityAttributeClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("securityattribute", "https://security-attribute.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *SecurityAttributeClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *SecurityAttributeClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// BulkDeleteSecurityAttributes Deletes the specified security attribute definitions. This operation triggers a process that removes the
// security attributes from all resources in your tenancy. The security attributes must be within the same security attribute namespace.
//
// The following actions happen immediately:
//
// After you start this operation, the state of the tag changes to DELETING, and security attribute removal
// from resources begins. This process can take up to 48 hours depending on the number of resources that
// are tagged and the regions in which those resources reside.
//
// When all security attributes have been removed, the state changes to DELETED. You cannot restore a deleted security attribute. After the security attribute state
// changes to DELETED, you can use the same security attribute name again.
//
// After you start this operation, you cannot start either the DeleteSecurityAttribute or the CascadeDeleteSecurityAttributeNamespace operation until this process completes.
//
// In order to delete security attribute, you must first retire the security attribute. Use UpdateSecurityAttribute
// to retire a security attribute.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/securityattribute/BulkDeleteSecurityAttributes.go.html to see an example of how to use BulkDeleteSecurityAttributes API.
// A default retry strategy applies to this operation BulkDeleteSecurityAttributes()
func (client SecurityAttributeClient) BulkDeleteSecurityAttributes(ctx context.Context, request BulkDeleteSecurityAttributesRequest) (response BulkDeleteSecurityAttributesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.bulkDeleteSecurityAttributes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BulkDeleteSecurityAttributesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BulkDeleteSecurityAttributesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BulkDeleteSecurityAttributesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BulkDeleteSecurityAttributesResponse")
	}
	return
}

// bulkDeleteSecurityAttributes implements the OCIOperation interface (enables retrying operations)
func (client SecurityAttributeClient) bulkDeleteSecurityAttributes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/securityAttributes/actions/bulkDelete", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BulkDeleteSecurityAttributesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/security-attribute/20240815/SecurityAttribute/BulkDeleteSecurityAttributes"
		err = common.PostProcessServiceError(err, "SecurityAttribute", "BulkDeleteSecurityAttributes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// BulkEditSecurityAttributes Edits the specified list of security attributes for the selected resources.
// This operation triggers a process that edits the attributes on all selected resources. The possible actions are:
//   - Add a security attribute when it does not already exist on the resource.
//   - Update the value for a security attribute when it is present on the resource.
//   - Add a security attribute when it does not already exist on the resource or update the value when it is present on the resource.
//   - Remove a security attribute from a resource. The security attribute is removed from the resource regardless of the value.
//
// The edits can include a combination of operations and attributes.
// However, multiple operations cannot apply to the same attribute in the same request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/securityattribute/BulkEditSecurityAttributes.go.html to see an example of how to use BulkEditSecurityAttributes API.
// A default retry strategy applies to this operation BulkEditSecurityAttributes()
func (client SecurityAttributeClient) BulkEditSecurityAttributes(ctx context.Context, request BulkEditSecurityAttributesRequest) (response BulkEditSecurityAttributesResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.bulkEditSecurityAttributes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BulkEditSecurityAttributesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BulkEditSecurityAttributesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BulkEditSecurityAttributesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BulkEditSecurityAttributesResponse")
	}
	return
}

// bulkEditSecurityAttributes implements the OCIOperation interface (enables retrying operations)
func (client SecurityAttributeClient) bulkEditSecurityAttributes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/securityAttributes/actions/bulkEdit", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response BulkEditSecurityAttributesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/security-attribute/20240815/SecurityAttribute/BulkEditSecurityAttributes"
		err = common.PostProcessServiceError(err, "SecurityAttribute", "BulkEditSecurityAttributes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CascadingDeleteSecurityAttributeNamespace Deletes the specified security attribute namespace. This operation triggers a process that removes all of the security attributes
// defined in the specified security attribute namespace from all resources in your tenancy and then deletes the security attribute namespace.
// After you start the delete operation:
//   - New security attribute key definitions cannot be created under the namespace.
//   - The state of the security attribute namespace changes to DELETING.
//   - Security attribute removal from the resources begins.
//
// This process can take up to 48 hours depending on the number of security attributes in the namespace, the number of resources
// that are tagged, and the locations of the regions in which those resources reside.
// After all security attributes are removed, the state changes to DELETED. You cannot restore a deleted security attribute namespace. After the deleted security attribute namespace
// changes its state to DELETED, you can use the name of the deleted security attribute namespace again.
// After you start this operation, you cannot start either the DeleteSecurityAttribute or the BulkDeleteSecurityAttributes operation until this process completes.
// To delete a security attribute namespace, you must first retire it. Use UpdateSecurityAttributeNamespace
// to retire a security attribute namespace.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/securityattribute/CascadingDeleteSecurityAttributeNamespace.go.html to see an example of how to use CascadingDeleteSecurityAttributeNamespace API.
// A default retry strategy applies to this operation CascadingDeleteSecurityAttributeNamespace()
func (client SecurityAttributeClient) CascadingDeleteSecurityAttributeNamespace(ctx context.Context, request CascadingDeleteSecurityAttributeNamespaceRequest) (response CascadingDeleteSecurityAttributeNamespaceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.cascadingDeleteSecurityAttributeNamespace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CascadingDeleteSecurityAttributeNamespaceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CascadingDeleteSecurityAttributeNamespaceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CascadingDeleteSecurityAttributeNamespaceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CascadingDeleteSecurityAttributeNamespaceResponse")
	}
	return
}

// cascadingDeleteSecurityAttributeNamespace implements the OCIOperation interface (enables retrying operations)
func (client SecurityAttributeClient) cascadingDeleteSecurityAttributeNamespace(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/securityAttributeNamespaces/{securityAttributeNamespaceId}/actions/cascadeDelete", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CascadingDeleteSecurityAttributeNamespaceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/security-attribute/20240815/SecurityAttributeNamespace/CascadingDeleteSecurityAttributeNamespace"
		err = common.PostProcessServiceError(err, "SecurityAttribute", "CascadingDeleteSecurityAttributeNamespace", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeSecurityAttributeNamespaceCompartment Moves the specified security attribute namespace to the specified compartment within the same tenancy.
// To move the security attribute namespace, you must have the manage security-attributes permission on both compartments.
// For more information about IAM policies, see Details for IAM (https://docs.cloud.oracle.com/Content/Identity/policyreference/iampolicyreference.htm).
// Moving a security attribute namespace moves all the security attributes contained in the security attribute namespace.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/securityattribute/ChangeSecurityAttributeNamespaceCompartment.go.html to see an example of how to use ChangeSecurityAttributeNamespaceCompartment API.
// A default retry strategy applies to this operation ChangeSecurityAttributeNamespaceCompartment()
func (client SecurityAttributeClient) ChangeSecurityAttributeNamespaceCompartment(ctx context.Context, request ChangeSecurityAttributeNamespaceCompartmentRequest) (response ChangeSecurityAttributeNamespaceCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeSecurityAttributeNamespaceCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeSecurityAttributeNamespaceCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeSecurityAttributeNamespaceCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeSecurityAttributeNamespaceCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeSecurityAttributeNamespaceCompartmentResponse")
	}
	return
}

// changeSecurityAttributeNamespaceCompartment implements the OCIOperation interface (enables retrying operations)
func (client SecurityAttributeClient) changeSecurityAttributeNamespaceCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/securityAttributeNamespaces/{securityAttributeNamespaceId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeSecurityAttributeNamespaceCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/security-attribute/20240815/SecurityAttributeNamespace/ChangeSecurityAttributeNamespaceCompartment"
		err = common.PostProcessServiceError(err, "SecurityAttribute", "ChangeSecurityAttributeNamespaceCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSecurityAttribute Creates a new security attribute in the specified security attribute namespace.
// The security attribute requires either the OCID or the name of the security attribute namespace that will contain this
// security attribute.
// You must specify a *name* for the attribute, which must be unique across all attributes in the security attribute namespace
// and cannot be changed. The only valid characters for security attribute names are: 0-9, A-Z, a-z, -, _ characters.
// Names are case insensitive. That means, for example, "mySecurityAttribute" and "mysecurityattribute" are not allowed in the same namespace.
// If you specify a name that's already in use in the security attribute namespace, a 409 error is returned.
// The security attribute must have a *description*. It does not have to be unique, and you can change it with
// UpdateSecurityAttribute.
// The security attribute must have a value type, which is specified with a validator. Security attribute can use either a
// static value or a list of possible values. Static values are entered by a user applying the security attribute
// to a resource. Lists are created by the user and the user must apply a value from the list. Lists
// are validated.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/securityattribute/CreateSecurityAttribute.go.html to see an example of how to use CreateSecurityAttribute API.
// A default retry strategy applies to this operation CreateSecurityAttribute()
func (client SecurityAttributeClient) CreateSecurityAttribute(ctx context.Context, request CreateSecurityAttributeRequest) (response CreateSecurityAttributeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createSecurityAttribute, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSecurityAttributeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSecurityAttributeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSecurityAttributeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSecurityAttributeResponse")
	}
	return
}

// createSecurityAttribute implements the OCIOperation interface (enables retrying operations)
func (client SecurityAttributeClient) createSecurityAttribute(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/securityAttributeNamespaces/{securityAttributeNamespaceId}/securityAttributes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSecurityAttributeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/security-attribute/20240815/SecurityAttribute/CreateSecurityAttribute"
		err = common.PostProcessServiceError(err, "SecurityAttribute", "CreateSecurityAttribute", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSecurityAttributeNamespace Creates a new security attribute namespace in the specified compartment.
// You must specify the compartment ID in the request object (remember that the tenancy is simply the root
// compartment).
// You must also specify a *name* for the namespace, which must be unique across all namespaces in your tenancy
// and cannot be changed. The only valid characters for security attribute names are:  0-9, A-Z, a-z, -, _ characters.
// Names are case insensitive. That means, for example, "myNamespace" and "mynamespace" are not allowed
// in the same tenancy. Once you created a namespace, you cannot change the name.
// If you specify a name that's already in use in the tenancy, a 409 error is returned.
// You must also specify a *description* for the namespace.
// It does not have to be unique, and you can change it with
// SecurityAttributeNamespace.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/securityattribute/CreateSecurityAttributeNamespace.go.html to see an example of how to use CreateSecurityAttributeNamespace API.
// A default retry strategy applies to this operation CreateSecurityAttributeNamespace()
func (client SecurityAttributeClient) CreateSecurityAttributeNamespace(ctx context.Context, request CreateSecurityAttributeNamespaceRequest) (response CreateSecurityAttributeNamespaceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createSecurityAttributeNamespace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSecurityAttributeNamespaceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSecurityAttributeNamespaceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSecurityAttributeNamespaceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSecurityAttributeNamespaceResponse")
	}
	return
}

// createSecurityAttributeNamespace implements the OCIOperation interface (enables retrying operations)
func (client SecurityAttributeClient) createSecurityAttributeNamespace(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/securityAttributeNamespaces", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSecurityAttributeNamespaceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/security-attribute/20240815/SecurityAttributeNamespace/CreateSecurityAttributeNamespace"
		err = common.PostProcessServiceError(err, "SecurityAttribute", "CreateSecurityAttributeNamespace", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSecurityAttribute Deletes the specified security attribute. This operation triggers a process that removes the
// security attribute from all resources in your tenancy.
// When you start the delete operation, the state of the security attribute changes to DELETING and security attribute removal
// from resources begins. This can take up to 48 hours depending on the number of resources that
// were tagged as well as the regions in which those resources reside.
// When all attributes have been removed, the state changes to DELETED. You cannot restore a deleted attribute. Once the deleted attribute
// changes its state to DELETED, you can use the same security attribute name again.
// After you start this operation, you cannot start either the BulkDeleteSecurityAttributes or the CascadeDeleteTagNamespace operation until this process completes.
// To delete a security attribute, you must first retire it. Use UpdateSecurityAttribute
// to retire a security attribute.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/securityattribute/DeleteSecurityAttribute.go.html to see an example of how to use DeleteSecurityAttribute API.
// A default retry strategy applies to this operation DeleteSecurityAttribute()
func (client SecurityAttributeClient) DeleteSecurityAttribute(ctx context.Context, request DeleteSecurityAttributeRequest) (response DeleteSecurityAttributeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteSecurityAttribute, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSecurityAttributeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSecurityAttributeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSecurityAttributeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSecurityAttributeResponse")
	}
	return
}

// deleteSecurityAttribute implements the OCIOperation interface (enables retrying operations)
func (client SecurityAttributeClient) deleteSecurityAttribute(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/securityAttributeNamespaces/{securityAttributeNamespaceId}/securityAttributes/{securityAttributeName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSecurityAttributeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/security-attribute/20240815/SecurityAttribute/DeleteSecurityAttribute"
		err = common.PostProcessServiceError(err, "SecurityAttribute", "DeleteSecurityAttribute", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSecurityAttributeNamespace Deletes the specified security attribute namespace. Only an empty security attribute namespace can be deleted with this operation. To use this operation
// to delete a security attribute namespace that contains security attributes, first delete all of its security attributes.
// Use DeleteSecurityAttribute to delete a security attribute.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/securityattribute/DeleteSecurityAttributeNamespace.go.html to see an example of how to use DeleteSecurityAttributeNamespace API.
// A default retry strategy applies to this operation DeleteSecurityAttributeNamespace()
func (client SecurityAttributeClient) DeleteSecurityAttributeNamespace(ctx context.Context, request DeleteSecurityAttributeNamespaceRequest) (response DeleteSecurityAttributeNamespaceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteSecurityAttributeNamespace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSecurityAttributeNamespaceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSecurityAttributeNamespaceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSecurityAttributeNamespaceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSecurityAttributeNamespaceResponse")
	}
	return
}

// deleteSecurityAttributeNamespace implements the OCIOperation interface (enables retrying operations)
func (client SecurityAttributeClient) deleteSecurityAttributeNamespace(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/securityAttributeNamespaces/{securityAttributeNamespaceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSecurityAttributeNamespaceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/security-attribute/20240815/SecurityAttributeNamespace/DeleteSecurityAttributeNamespace"
		err = common.PostProcessServiceError(err, "SecurityAttribute", "DeleteSecurityAttributeNamespace", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSecurityAttribute Gets the specified security attribute's information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/securityattribute/GetSecurityAttribute.go.html to see an example of how to use GetSecurityAttribute API.
// A default retry strategy applies to this operation GetSecurityAttribute()
func (client SecurityAttributeClient) GetSecurityAttribute(ctx context.Context, request GetSecurityAttributeRequest) (response GetSecurityAttributeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSecurityAttribute, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSecurityAttributeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSecurityAttributeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSecurityAttributeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSecurityAttributeResponse")
	}
	return
}

// getSecurityAttribute implements the OCIOperation interface (enables retrying operations)
func (client SecurityAttributeClient) getSecurityAttribute(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityAttributeNamespaces/{securityAttributeNamespaceId}/securityAttributes/{securityAttributeName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSecurityAttributeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/security-attribute/20240815/SecurityAttribute/GetSecurityAttribute"
		err = common.PostProcessServiceError(err, "SecurityAttribute", "GetSecurityAttribute", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSecurityAttributeNamespace Gets the specified security attribute namespace's information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/securityattribute/GetSecurityAttributeNamespace.go.html to see an example of how to use GetSecurityAttributeNamespace API.
// A default retry strategy applies to this operation GetSecurityAttributeNamespace()
func (client SecurityAttributeClient) GetSecurityAttributeNamespace(ctx context.Context, request GetSecurityAttributeNamespaceRequest) (response GetSecurityAttributeNamespaceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSecurityAttributeNamespace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSecurityAttributeNamespaceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSecurityAttributeNamespaceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSecurityAttributeNamespaceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSecurityAttributeNamespaceResponse")
	}
	return
}

// getSecurityAttributeNamespace implements the OCIOperation interface (enables retrying operations)
func (client SecurityAttributeClient) getSecurityAttributeNamespace(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityAttributeNamespaces/{securityAttributeNamespaceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSecurityAttributeNamespaceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/security-attribute/20240815/SecurityAttributeNamespace/GetSecurityAttributeNamespace"
		err = common.PostProcessServiceError(err, "SecurityAttribute", "GetSecurityAttributeNamespace", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSecurityAttributeWorkRequest Gets details on a specified work request. The workRequestID is returned in the opc-work-request-id header
// for any asynchronous operation in security attributes service.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/securityattribute/GetSecurityAttributeWorkRequest.go.html to see an example of how to use GetSecurityAttributeWorkRequest API.
// A default retry strategy applies to this operation GetSecurityAttributeWorkRequest()
func (client SecurityAttributeClient) GetSecurityAttributeWorkRequest(ctx context.Context, request GetSecurityAttributeWorkRequestRequest) (response GetSecurityAttributeWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSecurityAttributeWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSecurityAttributeWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSecurityAttributeWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSecurityAttributeWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSecurityAttributeWorkRequestResponse")
	}
	return
}

// getSecurityAttributeWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client SecurityAttributeClient) getSecurityAttributeWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityAttributeWorkRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSecurityAttributeWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/security-attribute/20240815/SecurityAttributeWorkRequest/GetSecurityAttributeWorkRequest"
		err = common.PostProcessServiceError(err, "SecurityAttribute", "GetSecurityAttributeWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSecurityAttributeNamespaces Lists the security attribute namespaces in the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/securityattribute/ListSecurityAttributeNamespaces.go.html to see an example of how to use ListSecurityAttributeNamespaces API.
// A default retry strategy applies to this operation ListSecurityAttributeNamespaces()
func (client SecurityAttributeClient) ListSecurityAttributeNamespaces(ctx context.Context, request ListSecurityAttributeNamespacesRequest) (response ListSecurityAttributeNamespacesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSecurityAttributeNamespaces, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSecurityAttributeNamespacesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSecurityAttributeNamespacesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSecurityAttributeNamespacesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSecurityAttributeNamespacesResponse")
	}
	return
}

// listSecurityAttributeNamespaces implements the OCIOperation interface (enables retrying operations)
func (client SecurityAttributeClient) listSecurityAttributeNamespaces(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityAttributeNamespaces", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSecurityAttributeNamespacesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/security-attribute/20240815/SecurityAttributeNamespaceSummary/ListSecurityAttributeNamespaces"
		err = common.PostProcessServiceError(err, "SecurityAttribute", "ListSecurityAttributeNamespaces", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSecurityAttributeWorkRequestErrors Gets the errors for a work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/securityattribute/ListSecurityAttributeWorkRequestErrors.go.html to see an example of how to use ListSecurityAttributeWorkRequestErrors API.
// A default retry strategy applies to this operation ListSecurityAttributeWorkRequestErrors()
func (client SecurityAttributeClient) ListSecurityAttributeWorkRequestErrors(ctx context.Context, request ListSecurityAttributeWorkRequestErrorsRequest) (response ListSecurityAttributeWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSecurityAttributeWorkRequestErrors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSecurityAttributeWorkRequestErrorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSecurityAttributeWorkRequestErrorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSecurityAttributeWorkRequestErrorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSecurityAttributeWorkRequestErrorsResponse")
	}
	return
}

// listSecurityAttributeWorkRequestErrors implements the OCIOperation interface (enables retrying operations)
func (client SecurityAttributeClient) listSecurityAttributeWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityAttributeWorkRequests/{workRequestId}/errors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSecurityAttributeWorkRequestErrorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/security-attribute/20240815/SecurityAttributeWorkRequestErrorSummary/ListSecurityAttributeWorkRequestErrors"
		err = common.PostProcessServiceError(err, "SecurityAttribute", "ListSecurityAttributeWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSecurityAttributeWorkRequestLogs Gets the logs for a work request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/securityattribute/ListSecurityAttributeWorkRequestLogs.go.html to see an example of how to use ListSecurityAttributeWorkRequestLogs API.
// A default retry strategy applies to this operation ListSecurityAttributeWorkRequestLogs()
func (client SecurityAttributeClient) ListSecurityAttributeWorkRequestLogs(ctx context.Context, request ListSecurityAttributeWorkRequestLogsRequest) (response ListSecurityAttributeWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSecurityAttributeWorkRequestLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSecurityAttributeWorkRequestLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSecurityAttributeWorkRequestLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSecurityAttributeWorkRequestLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSecurityAttributeWorkRequestLogsResponse")
	}
	return
}

// listSecurityAttributeWorkRequestLogs implements the OCIOperation interface (enables retrying operations)
func (client SecurityAttributeClient) listSecurityAttributeWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityAttributeWorkRequests/{workRequestId}/logs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSecurityAttributeWorkRequestLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/security-attribute/20240815/SecurityAttributeWorkRequestLogSummary/ListSecurityAttributeWorkRequestLogs"
		err = common.PostProcessServiceError(err, "SecurityAttribute", "ListSecurityAttributeWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSecurityAttributeWorkRequests Lists the security attribute work requests in compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/securityattribute/ListSecurityAttributeWorkRequests.go.html to see an example of how to use ListSecurityAttributeWorkRequests API.
// A default retry strategy applies to this operation ListSecurityAttributeWorkRequests()
func (client SecurityAttributeClient) ListSecurityAttributeWorkRequests(ctx context.Context, request ListSecurityAttributeWorkRequestsRequest) (response ListSecurityAttributeWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSecurityAttributeWorkRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSecurityAttributeWorkRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSecurityAttributeWorkRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSecurityAttributeWorkRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSecurityAttributeWorkRequestsResponse")
	}
	return
}

// listSecurityAttributeWorkRequests implements the OCIOperation interface (enables retrying operations)
func (client SecurityAttributeClient) listSecurityAttributeWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityAttributeWorkRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSecurityAttributeWorkRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/security-attribute/20240815/SecurityAttributeWorkRequestSummary/ListSecurityAttributeWorkRequests"
		err = common.PostProcessServiceError(err, "SecurityAttribute", "ListSecurityAttributeWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSecurityAttributes Lists the security attributes in the specified namespace.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/securityattribute/ListSecurityAttributes.go.html to see an example of how to use ListSecurityAttributes API.
// A default retry strategy applies to this operation ListSecurityAttributes()
func (client SecurityAttributeClient) ListSecurityAttributes(ctx context.Context, request ListSecurityAttributesRequest) (response ListSecurityAttributesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSecurityAttributes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSecurityAttributesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSecurityAttributesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSecurityAttributesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSecurityAttributesResponse")
	}
	return
}

// listSecurityAttributes implements the OCIOperation interface (enables retrying operations)
func (client SecurityAttributeClient) listSecurityAttributes(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/securityAttributeNamespaces/{securityAttributeNamespaceId}/securityAttributes", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSecurityAttributesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/security-attribute/20240815/SecurityAttributeSummary/ListSecurityAttributes"
		err = common.PostProcessServiceError(err, "SecurityAttribute", "ListSecurityAttributes", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSecurityAttribute Updates the specified security attribute. You can only update `description`, and `isRetired`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/securityattribute/UpdateSecurityAttribute.go.html to see an example of how to use UpdateSecurityAttribute API.
// A default retry strategy applies to this operation UpdateSecurityAttribute()
func (client SecurityAttributeClient) UpdateSecurityAttribute(ctx context.Context, request UpdateSecurityAttributeRequest) (response UpdateSecurityAttributeResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateSecurityAttribute, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSecurityAttributeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSecurityAttributeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSecurityAttributeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSecurityAttributeResponse")
	}
	return
}

// updateSecurityAttribute implements the OCIOperation interface (enables retrying operations)
func (client SecurityAttributeClient) updateSecurityAttribute(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/securityAttributeNamespaces/{securityAttributeNamespaceId}/securityAttributes/{securityAttributeName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSecurityAttributeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/security-attribute/20240815/SecurityAttribute/UpdateSecurityAttribute"
		err = common.PostProcessServiceError(err, "SecurityAttribute", "UpdateSecurityAttribute", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSecurityAttributeNamespace Updates the specified security attribute namespace. You can't update the namespace name.
// Updating `isRetired` to 'true' retires the namespace and all the security attributes in the namespace. Reactivating a
// namespace (changing `isRetired` from 'true' to 'false') does not reactivate security attributes.
// To reactivate the security attributes, you must reactivate each one individually *after* you reactivate the namespace,
// using UpdateTag. For more information about retiring security attribute namespaces, see
// Managing Security Attribute Namespaces (https://docs.cloud.oracle.com/Content/zero-trust-packet-routing/managing-security-attribute-namespaces.htm).
// You can't add a namespace with the same name as a retired namespace in the same tenancy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/securityattribute/UpdateSecurityAttributeNamespace.go.html to see an example of how to use UpdateSecurityAttributeNamespace API.
// A default retry strategy applies to this operation UpdateSecurityAttributeNamespace()
func (client SecurityAttributeClient) UpdateSecurityAttributeNamespace(ctx context.Context, request UpdateSecurityAttributeNamespaceRequest) (response UpdateSecurityAttributeNamespaceResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateSecurityAttributeNamespace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSecurityAttributeNamespaceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSecurityAttributeNamespaceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSecurityAttributeNamespaceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSecurityAttributeNamespaceResponse")
	}
	return
}

// updateSecurityAttributeNamespace implements the OCIOperation interface (enables retrying operations)
func (client SecurityAttributeClient) updateSecurityAttributeNamespace(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/securityAttributeNamespaces/{securityAttributeNamespaceId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSecurityAttributeNamespaceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/security-attribute/20240815/SecurityAttributeNamespace/UpdateSecurityAttributeNamespace"
		err = common.PostProcessServiceError(err, "SecurityAttribute", "UpdateSecurityAttributeNamespace", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
