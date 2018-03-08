// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"context"
	"fmt"
	"net/http"

	"github.com/oracle/oci-go-sdk/common"
)

//IdentityClient a client for Identity
type IdentityClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewIdentityClientWithConfigurationProvider Creates a new default Identity client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewIdentityClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client IdentityClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	client = IdentityClient{BaseClient: baseClient}
	client.BasePath = "20160918"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *IdentityClient) SetRegion(region string) {
	client.Host = fmt.Sprintf(common.DefaultHostURLTemplate, "identity", region)
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *IdentityClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.config = &configProvider
	client.SetRegion(region)
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *IdentityClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AddUserToGroup Adds the specified user to the specified group and returns a `UserGroupMembership` object with its own OCID.
// After you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before using the
// object, first make sure its `lifecycleState` has changed to ACTIVE.
func (client IdentityClient) AddUserToGroup(ctx context.Context, request AddUserToGroupRequest, options ...common.RetryPolicyOption) (response AddUserToGroupResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/userGroupMemberships/", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// CreateCompartment Creates a new compartment in your tenancy.
// **Important:** Compartments cannot be deleted.
// You must specify your tenancy's OCID as the compartment ID in the request object. Remember that the tenancy
// is simply the root compartment. For information about OCIDs, see
// [Resource Identifiers]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
// You must also specify a *name* for the compartment, which must be unique across all compartments in
// your tenancy. You can use this name or the OCID when writing policies that apply
// to the compartment. For more information about policies, see
// [How Policies Work]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policies.htm).
// You must also specify a *description* for the compartment (although it can be an empty string). It does
// not have to be unique, and you can change it anytime with
// UpdateCompartment.
// After you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before using the
// object, first make sure its `lifecycleState` has changed to ACTIVE.
func (client IdentityClient) CreateCompartment(ctx context.Context, request CreateCompartmentRequest, options ...common.RetryPolicyOption) (response CreateCompartmentResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/compartments/", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// CreateCustomerSecretKey Creates a new secret key for the specified user. Secret keys are used for authentication with the Object Storage Service's Amazon S3
// compatible API. For information, see
// [Managing User Credentials]({{DOC_SERVER_URL}}/Content/Identity/Tasks/managingcredentials.htm).
// You must specify a *description* for the secret key (although it can be an empty string). It does not
// have to be unique, and you can change it anytime with
// UpdateCustomerSecretKey.
// Every user has permission to create a secret key for *their own user ID*. An administrator in your organization
// does not need to write a policy to give users this ability. To compare, administrators who have permission to the
// tenancy can use this operation to create a secret key for any user, including themselves.
func (client IdentityClient) CreateCustomerSecretKey(ctx context.Context, request CreateCustomerSecretKeyRequest, options ...common.RetryPolicyOption) (response CreateCustomerSecretKeyResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/users/{userId}/customerSecretKeys/", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// CreateGroup Creates a new group in your tenancy.
// You must specify your tenancy's OCID as the compartment ID in the request object (remember that the tenancy
// is simply the root compartment). Notice that IAM resources (users, groups, compartments, and some policies)
// reside within the tenancy itself, unlike cloud resources such as compute instances, which typically
// reside within compartments inside the tenancy. For information about OCIDs, see
// [Resource Identifiers]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
// You must also specify a *name* for the group, which must be unique across all groups in your tenancy and
// cannot be changed. You can use this name or the OCID when writing policies that apply to the group. For more
// information about policies, see [How Policies Work]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policies.htm).
// You must also specify a *description* for the group (although it can be an empty string). It does not
// have to be unique, and you can change it anytime with UpdateGroup.
// After you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before using the
// object, first make sure its `lifecycleState` has changed to ACTIVE.
// After creating the group, you need to put users in it and write policies for it.
// See AddUserToGroup and
// CreatePolicy.
func (client IdentityClient) CreateGroup(ctx context.Context, request CreateGroupRequest, options ...common.RetryPolicyOption) (response CreateGroupResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/groups/", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// CreateIdentityProvider Creates a new identity provider in your tenancy. For more information, see
// [Identity Providers and Federation]({{DOC_SERVER_URL}}/Content/Identity/Concepts/federation.htm).
// You must specify your tenancy's OCID as the compartment ID in the request object.
// Remember that the tenancy is simply the root compartment. For information about
// OCIDs, see [Resource Identifiers]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
// You must also specify a *name* for the `IdentityProvider`, which must be unique
// across all `IdentityProvider` objects in your tenancy and cannot be changed.
// You must also specify a *description* for the `IdentityProvider` (although
// it can be an empty string). It does not have to be unique, and you can change
// it anytime with
// UpdateIdentityProvider.
// After you send your request, the new object's `lifecycleState` will temporarily
// be CREATING. Before using the object, first make sure its `lifecycleState` has
// changed to ACTIVE.
func (client IdentityClient) CreateIdentityProvider(ctx context.Context, request CreateIdentityProviderRequest, options ...common.RetryPolicyOption) (response CreateIdentityProviderResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/identityProviders/", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &identityprovider{})
		},
		RetryPolicyOptions: options,
	})
	return
}

// CreateIdpGroupMapping Creates a single mapping between an IdP group and an IAM Service
// Group.
func (client IdentityClient) CreateIdpGroupMapping(ctx context.Context, request CreateIdpGroupMappingRequest, options ...common.RetryPolicyOption) (response CreateIdpGroupMappingResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/identityProviders/{identityProviderId}/groupMappings/", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// CreateOrResetUIPassword Creates a new Console one-time password for the specified user. For more information about user
// credentials, see [User Credentials]({{DOC_SERVER_URL}}/Content/Identity/Concepts/usercredentials.htm).
// Use this operation after creating a new user, or if a user forgets their password. The new one-time
// password is returned to you in the response, and you must securely deliver it to the user. They'll
// be prompted to change this password the next time they sign in to the Console. If they don't change
// it within 7 days, the password will expire and you'll need to create a new one-time password for the
// user.
// **Note:** The user's Console login is the unique name you specified when you created the user
// (see CreateUser).
func (client IdentityClient) CreateOrResetUIPassword(ctx context.Context, request CreateOrResetUIPasswordRequest, options ...common.RetryPolicyOption) (response CreateOrResetUIPasswordResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/users/{userId}/uiPassword", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// CreatePolicy Creates a new policy in the specified compartment (either the tenancy or another of your compartments).
// If you're new to policies, see [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
// You must specify a *name* for the policy, which must be unique across all policies in your tenancy
// and cannot be changed.
// You must also specify a *description* for the policy (although it can be an empty string). It does not
// have to be unique, and you can change it anytime with UpdatePolicy.
// You must specify one or more policy statements in the statements array. For information about writing
// policies, see [How Policies Work]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policies.htm) and
// [Common Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/commonpolicies.htm).
// After you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before using the
// object, first make sure its `lifecycleState` has changed to ACTIVE.
// New policies take effect typically within 10 seconds.
func (client IdentityClient) CreatePolicy(ctx context.Context, request CreatePolicyRequest, options ...common.RetryPolicyOption) (response CreatePolicyResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/policies/", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// CreateRegionSubscription Creates a subscription to a region for a tenancy.
func (client IdentityClient) CreateRegionSubscription(ctx context.Context, request CreateRegionSubscriptionRequest, options ...common.RetryPolicyOption) (response CreateRegionSubscriptionResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/tenancies/{tenancyId}/regionSubscriptions", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// CreateSwiftPassword Creates a new Swift password for the specified user. For information about what Swift passwords are for, see
// [Managing User Credentials]({{DOC_SERVER_URL}}/Content/Identity/Tasks/managingcredentials.htm).
// You must specify a *description* for the Swift password (although it can be an empty string). It does not
// have to be unique, and you can change it anytime with
// UpdateSwiftPassword.
// Every user has permission to create a Swift password for *their own user ID*. An administrator in your organization
// does not need to write a policy to give users this ability. To compare, administrators who have permission to the
// tenancy can use this operation to create a Swift password for any user, including themselves.
func (client IdentityClient) CreateSwiftPassword(ctx context.Context, request CreateSwiftPasswordRequest, options ...common.RetryPolicyOption) (response CreateSwiftPasswordResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/users/{userId}/swiftPasswords/", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// CreateUser Creates a new user in your tenancy. For conceptual information about users, your tenancy, and other
// IAM Service components, see [Overview of the IAM Service]({{DOC_SERVER_URL}}/Content/Identity/Concepts/overview.htm).
// You must specify your tenancy's OCID as the compartment ID in the request object (remember that the
// tenancy is simply the root compartment). Notice that IAM resources (users, groups, compartments, and
// some policies) reside within the tenancy itself, unlike cloud resources such as compute instances,
// which typically reside within compartments inside the tenancy. For information about OCIDs, see
// [Resource Identifiers]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
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
// [Required Keys and OCIDs]({{DOC_SERVER_URL}}/Content/API/Concepts/apisigningkey.htm) and also
// UploadApiKey).
// **Important:** Make sure to inform the new user which compartment(s) they have access to.
func (client IdentityClient) CreateUser(ctx context.Context, request CreateUserRequest, options ...common.RetryPolicyOption) (response CreateUserResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/users/", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// DeleteApiKey Deletes the specified API signing key for the specified user.
// Every user has permission to use this operation to delete a key for *their own user ID*. An
// administrator in your organization does not need to write a policy to give users this ability.
// To compare, administrators who have permission to the tenancy can use this operation to delete
// a key for any user, including themselves.
func (client IdentityClient) DeleteApiKey(ctx context.Context, request DeleteApiKeyRequest, options ...common.RetryPolicyOption) (response DeleteApiKeyResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/users/{userId}/apiKeys/{fingerprint}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// DeleteCustomerSecretKey Deletes the specified secret key for the specified user.
func (client IdentityClient) DeleteCustomerSecretKey(ctx context.Context, request DeleteCustomerSecretKeyRequest, options ...common.RetryPolicyOption) (response DeleteCustomerSecretKeyResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/users/{userId}/customerSecretKeys/{customerSecretKeyId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// DeleteGroup Deletes the specified group. The group must be empty.
func (client IdentityClient) DeleteGroup(ctx context.Context, request DeleteGroupRequest, options ...common.RetryPolicyOption) (response DeleteGroupResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/groups/{groupId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// DeleteIdentityProvider Deletes the specified identity provider. The identity provider must not have
// any group mappings (see IdpGroupMapping).
func (client IdentityClient) DeleteIdentityProvider(ctx context.Context, request DeleteIdentityProviderRequest, options ...common.RetryPolicyOption) (response DeleteIdentityProviderResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/identityProviders/{identityProviderId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// DeleteIdpGroupMapping Deletes the specified group mapping.
func (client IdentityClient) DeleteIdpGroupMapping(ctx context.Context, request DeleteIdpGroupMappingRequest, options ...common.RetryPolicyOption) (response DeleteIdpGroupMappingResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/identityProviders/{identityProviderId}/groupMappings/{mappingId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// DeletePolicy Deletes the specified policy. The deletion takes effect typically within 10 seconds.
func (client IdentityClient) DeletePolicy(ctx context.Context, request DeletePolicyRequest, options ...common.RetryPolicyOption) (response DeletePolicyResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/policies/{policyId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// DeleteSwiftPassword Deletes the specified Swift password for the specified user.
func (client IdentityClient) DeleteSwiftPassword(ctx context.Context, request DeleteSwiftPasswordRequest, options ...common.RetryPolicyOption) (response DeleteSwiftPasswordResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/users/{userId}/swiftPasswords/{swiftPasswordId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// DeleteUser Deletes the specified user. The user must not be in any groups.
func (client IdentityClient) DeleteUser(ctx context.Context, request DeleteUserRequest, options ...common.RetryPolicyOption) (response DeleteUserResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/users/{userId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetCompartment Gets the specified compartment's information.
// This operation does not return a list of all the resources inside the compartment. There is no single
// API operation that does that. Compartments can contain multiple types of resources (instances, block
// storage volumes, etc.). To find out what's in a compartment, you must call the "List" operation for
// each resource type and specify the compartment's OCID as a query parameter in the request. For example,
// call the ListInstances operation in the Cloud Compute
// Service or the ListVolumes operation in Cloud Block Storage.
func (client IdentityClient) GetCompartment(ctx context.Context, request GetCompartmentRequest, options ...common.RetryPolicyOption) (response GetCompartmentResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/compartments/{compartmentId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetGroup Gets the specified group's information.
// This operation does not return a list of all the users in the group. To do that, use
// ListUserGroupMemberships and
// provide the group's OCID as a query parameter in the request.
func (client IdentityClient) GetGroup(ctx context.Context, request GetGroupRequest, options ...common.RetryPolicyOption) (response GetGroupResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/groups/{groupId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetIdentityProvider Gets the specified identity provider's information.
func (client IdentityClient) GetIdentityProvider(ctx context.Context, request GetIdentityProviderRequest, options ...common.RetryPolicyOption) (response GetIdentityProviderResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/identityProviders/{identityProviderId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &identityprovider{})
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetIdpGroupMapping Gets the specified group mapping.
func (client IdentityClient) GetIdpGroupMapping(ctx context.Context, request GetIdpGroupMappingRequest, options ...common.RetryPolicyOption) (response GetIdpGroupMappingResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/identityProviders/{identityProviderId}/groupMappings/{mappingId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetPolicy Gets the specified policy's information.
func (client IdentityClient) GetPolicy(ctx context.Context, request GetPolicyRequest, options ...common.RetryPolicyOption) (response GetPolicyResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/policies/{policyId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetTenancy Get the specified tenancy's information.
func (client IdentityClient) GetTenancy(ctx context.Context, request GetTenancyRequest, options ...common.RetryPolicyOption) (response GetTenancyResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/tenancies/{tenancyId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetUser Gets the specified user's information.
func (client IdentityClient) GetUser(ctx context.Context, request GetUserRequest, options ...common.RetryPolicyOption) (response GetUserResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/users/{userId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// GetUserGroupMembership Gets the specified UserGroupMembership's information.
func (client IdentityClient) GetUserGroupMembership(ctx context.Context, request GetUserGroupMembershipRequest, options ...common.RetryPolicyOption) (response GetUserGroupMembershipResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/userGroupMemberships/{userGroupMembershipId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListApiKeys Lists the API signing keys for the specified user. A user can have a maximum of three keys.
// Every user has permission to use this API call for *their own user ID*.  An administrator in your
// organization does not need to write a policy to give users this ability.
func (client IdentityClient) ListApiKeys(ctx context.Context, request ListApiKeysRequest, options ...common.RetryPolicyOption) (response ListApiKeysResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/users/{userId}/apiKeys/", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListAvailabilityDomains Lists the Availability Domains in your tenancy. Specify the OCID of either the tenancy or another
// of your compartments as the value for the compartment ID (remember that the tenancy is simply the root compartment).
// See [Where to Get the Tenancy's OCID and User's OCID]({{DOC_SERVER_URL}}/Content/API/Concepts/apisigningkey.htm#five).
func (client IdentityClient) ListAvailabilityDomains(ctx context.Context, request ListAvailabilityDomainsRequest, options ...common.RetryPolicyOption) (response ListAvailabilityDomainsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/availabilityDomains/", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListCompartments Lists the compartments in your tenancy. You must specify your tenancy's OCID as the value
// for the compartment ID (remember that the tenancy is simply the root compartment).
// See [Where to Get the Tenancy's OCID and User's OCID]({{DOC_SERVER_URL}}/Content/API/Concepts/apisigningkey.htm#five).
func (client IdentityClient) ListCompartments(ctx context.Context, request ListCompartmentsRequest, options ...common.RetryPolicyOption) (response ListCompartmentsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/compartments/", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListCustomerSecretKeys Lists the secret keys for the specified user. The returned object contains the secret key's OCID, but not
// the secret key itself. The actual secret key is returned only upon creation.
func (client IdentityClient) ListCustomerSecretKeys(ctx context.Context, request ListCustomerSecretKeysRequest, options ...common.RetryPolicyOption) (response ListCustomerSecretKeysResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/users/{userId}/customerSecretKeys/", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListFaultDomains Lists the Fault Domains in your tenancy. Specify the OCID of either the tenancy or another
// of your compartments as the value for the compartment ID (remember that the tenancy is simply the root compartment).
// See [Where to Get the Tenancy's OCID and User's OCID]({{DOC_SERVER_URL}}/Content/API/Concepts/apisigningkey.htm#five).
func (client IdentityClient) ListFaultDomains(ctx context.Context, request ListFaultDomainsRequest, options ...common.RetryPolicyOption) (response ListFaultDomainsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/faultDomains/", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListGroups Lists the groups in your tenancy. You must specify your tenancy's OCID as the value for
// the compartment ID (remember that the tenancy is simply the root compartment).
// See [Where to Get the Tenancy's OCID and User's OCID]({{DOC_SERVER_URL}}/Content/API/Concepts/apisigningkey.htm#five).
func (client IdentityClient) ListGroups(ctx context.Context, request ListGroupsRequest, options ...common.RetryPolicyOption) (response ListGroupsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/groups/", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
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

// ListIdentityProviders Lists all the identity providers in your tenancy. You must specify the identity provider type (e.g., `SAML2` for
// identity providers using the SAML2.0 protocol). You must specify your tenancy's OCID as the value for the
// compartment ID (remember that the tenancy is simply the root compartment).
// See [Where to Get the Tenancy's OCID and User's OCID]({{DOC_SERVER_URL}}/Content/API/Concepts/apisigningkey.htm#five).
func (client IdentityClient) ListIdentityProviders(ctx context.Context, request ListIdentityProvidersRequest, options ...common.RetryPolicyOption) (response ListIdentityProvidersResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/identityProviders/", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &listidentityprovider{})
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListIdpGroupMappings Lists the group mappings for the specified identity provider.
func (client IdentityClient) ListIdpGroupMappings(ctx context.Context, request ListIdpGroupMappingsRequest, options ...common.RetryPolicyOption) (response ListIdpGroupMappingsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/identityProviders/{identityProviderId}/groupMappings/", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListPolicies Lists the policies in the specified compartment (either the tenancy or another of your compartments).
// See [Where to Get the Tenancy's OCID and User's OCID]({{DOC_SERVER_URL}}/Content/API/Concepts/apisigningkey.htm#five).
// To determine which policies apply to a particular group or compartment, you must view the individual
// statements inside all your policies. There isn't a way to automatically obtain that information via the API.
func (client IdentityClient) ListPolicies(ctx context.Context, request ListPoliciesRequest, options ...common.RetryPolicyOption) (response ListPoliciesResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/policies/", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListRegionSubscriptions Lists the region subscriptions for the specified tenancy.
func (client IdentityClient) ListRegionSubscriptions(ctx context.Context, request ListRegionSubscriptionsRequest, options ...common.RetryPolicyOption) (response ListRegionSubscriptionsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/tenancies/{tenancyId}/regionSubscriptions", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListRegions Lists all the regions offered by Oracle Cloud Infrastructure.
func (client IdentityClient) ListRegions(ctx context.Context, options ...common.RetryPolicyOption) (response ListRegionsResponse, err error) {
	httpRequest := common.MakeDefaultHTTPRequest(http.MethodGet, "/regions")

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListSwiftPasswords Lists the Swift passwords for the specified user. The returned object contains the password's OCID, but not
// the password itself. The actual password is returned only upon creation.
func (client IdentityClient) ListSwiftPasswords(ctx context.Context, request ListSwiftPasswordsRequest, options ...common.RetryPolicyOption) (response ListSwiftPasswordsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/users/{userId}/swiftPasswords/", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListUserGroupMemberships Lists the `UserGroupMembership` objects in your tenancy. You must specify your tenancy's OCID
// as the value for the compartment ID
// (see [Where to Get the Tenancy's OCID and User's OCID]({{DOC_SERVER_URL}}/Content/API/Concepts/apisigningkey.htm#five)).
// You must also then filter the list in one of these ways:
// - You can limit the results to just the memberships for a given user by specifying a `userId`.
// - Similarly, you can limit the results to just the memberships for a given group by specifying a `groupId`.
// - You can set both the `userId` and `groupId` to determine if the specified user is in the specified group.
// If the answer is no, the response is an empty list.
func (client IdentityClient) ListUserGroupMemberships(ctx context.Context, request ListUserGroupMembershipsRequest, options ...common.RetryPolicyOption) (response ListUserGroupMembershipsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/userGroupMemberships/", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// ListUsers Lists the users in your tenancy. You must specify your tenancy's OCID as the value for the
// compartment ID (remember that the tenancy is simply the root compartment).
// See [Where to Get the Tenancy's OCID and User's OCID]({{DOC_SERVER_URL}}/Content/API/Concepts/apisigningkey.htm#five).
func (client IdentityClient) ListUsers(ctx context.Context, request ListUsersRequest, options ...common.RetryPolicyOption) (response ListUsersResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/users/", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// RemoveUserFromGroup Removes a user from a group by deleting the corresponding `UserGroupMembership`.
func (client IdentityClient) RemoveUserFromGroup(ctx context.Context, request RemoveUserFromGroupRequest, options ...common.RetryPolicyOption) (response RemoveUserFromGroupResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/userGroupMemberships/{userGroupMembershipId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// UpdateCompartment Updates the specified compartment's description or name. You can't update the root compartment.
func (client IdentityClient) UpdateCompartment(ctx context.Context, request UpdateCompartmentRequest, options ...common.RetryPolicyOption) (response UpdateCompartmentResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/compartments/{compartmentId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// UpdateCustomerSecretKey Updates the specified secret key's description.
func (client IdentityClient) UpdateCustomerSecretKey(ctx context.Context, request UpdateCustomerSecretKeyRequest, options ...common.RetryPolicyOption) (response UpdateCustomerSecretKeyResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/users/{userId}/customerSecretKeys/{customerSecretKeyId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// UpdateGroup Updates the specified group.
func (client IdentityClient) UpdateGroup(ctx context.Context, request UpdateGroupRequest, options ...common.RetryPolicyOption) (response UpdateGroupResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/groups/{groupId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// UpdateIdentityProvider Updates the specified identity provider.
func (client IdentityClient) UpdateIdentityProvider(ctx context.Context, request UpdateIdentityProviderRequest, options ...common.RetryPolicyOption) (response UpdateIdentityProviderResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/identityProviders/{identityProviderId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &identityprovider{})
		},
		RetryPolicyOptions: options,
	})
	return
}

// UpdateIdpGroupMapping Updates the specified group mapping.
func (client IdentityClient) UpdateIdpGroupMapping(ctx context.Context, request UpdateIdpGroupMappingRequest, options ...common.RetryPolicyOption) (response UpdateIdpGroupMappingResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/identityProviders/{identityProviderId}/groupMappings/{mappingId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// UpdatePolicy Updates the specified policy. You can update the description or the policy statements themselves.
// Policy changes take effect typically within 10 seconds.
func (client IdentityClient) UpdatePolicy(ctx context.Context, request UpdatePolicyRequest, options ...common.RetryPolicyOption) (response UpdatePolicyResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/policies/{policyId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// UpdateSwiftPassword Updates the specified Swift password's description.
func (client IdentityClient) UpdateSwiftPassword(ctx context.Context, request UpdateSwiftPasswordRequest, options ...common.RetryPolicyOption) (response UpdateSwiftPasswordResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/users/{userId}/swiftPasswords/{swiftPasswordId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// UpdateUser Updates the description of the specified user.
func (client IdentityClient) UpdateUser(ctx context.Context, request UpdateUserRequest, options ...common.RetryPolicyOption) (response UpdateUserResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/users/{userId}", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}

// UpdateUserState Updates the state of the specified user.
func (client IdentityClient) UpdateUserState(ctx context.Context, request UpdateUserStateRequest, options ...common.RetryPolicyOption) (response UpdateUserStateResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/users/{userId}/state/", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
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
func (client IdentityClient) UploadApiKey(ctx context.Context, request UploadApiKeyRequest, options ...common.RetryPolicyOption) (response UploadApiKeyResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/users/{userId}/apiKeys/", request)
	if err != nil {
		return
	}

	err = client.Call(ctx, &httpRequest, common.CallConfig{
		ResponseCallback: func(httpResponse *http.Response, e error) error {
			response.RawResponse = httpResponse
			if e != nil {
				return e
			}

			return common.UnmarshalResponse(httpResponse, &response)
		},
		RetryPolicyOptions: options,
	})
	return
}
