// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Domains API
//
// Use the Identity Domains API to manage resources within an identity domain, for example, users, dynamic resource groups, groups, and identity providers. For information about managing resources within identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm). This REST API is SCIM compliant.
// Use the table of contents and search tool to explore the Identity Domains API.
//

package identitydomains

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IdentityProvider Federation trusted partner Identity Provider
type IdentityProvider struct {

	// REQUIRED. The schemas attribute is an array of Strings which allows introspection of the supported schema version for a SCIM representation as well any schema extensions supported by that representation. Each String value must be a unique URI. This specification defines URIs for User, Group, and a standard \"enterprise\" extension. All representations of SCIM schema MUST include a non-zero value array with value(s) of the URIs supported by that representation. Duplicate values MUST NOT be included. Value order is not specified and MUST not impact behavior.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Schemas []string `mandatory:"true" json:"schemas"`

	// Unique name of the trusted Identity Provider.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: always
	//  - type: string
	//  - uniqueness: server
	PartnerName *string `mandatory:"true" json:"partnerName"`

	// Set to true to indicate Partner enabled.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	Enabled *bool `mandatory:"true" json:"enabled"`

	// Unique identifier for the SCIM Resource as defined by the Service Provider. Each representation of the Resource MUST include a non-empty id value. This identifier MUST be unique across the Service Provider's entire set of Resources. It MUST be a stable, non-reassignable identifier that does not change when the same Resource is returned in subsequent requests. The value of the id attribute is always issued by the Service Provider and MUST never be specified by the Service Consumer. bulkId: is a reserved keyword and MUST NOT be used in the unique identifier.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: always
	//  - type: string
	//  - uniqueness: global
	Id *string `mandatory:"false" json:"id"`

	// Unique OCI identifier for the SCIM Resource.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: global
	Ocid *string `mandatory:"false" json:"ocid"`

	Meta *Meta `mandatory:"false" json:"meta"`

	IdcsCreatedBy *IdcsCreatedBy `mandatory:"false" json:"idcsCreatedBy"`

	IdcsLastModifiedBy *IdcsLastModifiedBy `mandatory:"false" json:"idcsLastModifiedBy"`

	// Each value of this attribute specifies an operation that only an internal client may perform on this particular resource.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	IdcsPreventedOperations []IdcsPreventedOperationsEnum `mandatory:"false" json:"idcsPreventedOperations,omitempty"`

	// A list of tags on this resource.
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [key, value]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: request
	//  - type: complex
	//  - uniqueness: none
	Tags []Tags `mandatory:"false" json:"tags"`

	// A boolean flag indicating this resource in the process of being deleted. Usually set to true when synchronous deletion of the resource would take too long.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	DeleteInProgress *bool `mandatory:"false" json:"deleteInProgress"`

	// The release number when the resource was upgraded.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	IdcsLastUpgradedInRelease *string `mandatory:"false" json:"idcsLastUpgradedInRelease"`

	// OCI Domain Id (ocid) in which the resource lives.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	DomainOcid *string `mandatory:"false" json:"domainOcid"`

	// OCI Compartment Id (ocid) in which the resource lives.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	CompartmentOcid *string `mandatory:"false" json:"compartmentOcid"`

	// OCI Tenant Id (ocid) in which the resource lives.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	TenancyOcid *string `mandatory:"false" json:"tenancyOcid"`

	// An identifier for the Resource as defined by the Service Consumer. The externalId may simplify identification of the Resource between Service Consumer and Service Provider by allowing the Consumer to refer to the Resource with its own identifier, obviating the need to store a local mapping between the local identifier of the Resource and the identifier used by the Service Provider. Each Resource MAY include a non-empty externalId value. The value of the externalId attribute is always issued by the Service Consumer and can never be specified by the Service Provider. The Service Provider MUST always interpret the externalId as scoped to the Service Consumer's tenant.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	ExternalId *string `mandatory:"false" json:"externalId"`

	// Description
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Description *string `mandatory:"false" json:"description"`

	// Metadata
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Metadata *string `mandatory:"false" json:"metadata"`

	// Provider ID
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: server
	PartnerProviderId *string `mandatory:"false" json:"partnerProviderId"`

	// The alternate Provider ID to be used as the Oracle Identity Cloud Service providerID (instead of the one in SamlSettings) when interacting with this IdP.
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	TenantProviderId *string `mandatory:"false" json:"tenantProviderId"`

	// Succinct ID
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: server
	SuccinctId *string `mandatory:"false" json:"succinctId"`

	// Identity Provider SSO URL
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdpSsoUrl *string `mandatory:"false" json:"idpSsoUrl"`

	// Logout request URL
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	LogoutRequestUrl *string `mandatory:"false" json:"logoutRequestUrl"`

	// Logout response URL
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	LogoutResponseUrl *string `mandatory:"false" json:"logoutResponseUrl"`

	// Signing certificate
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	SigningCertificate *string `mandatory:"false" json:"signingCertificate"`

	// Encryption certificate
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	EncryptionCertificate *string `mandatory:"false" json:"encryptionCertificate"`

	// Default authentication request name ID format.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	NameIdFormat *string `mandatory:"false" json:"nameIdFormat"`

	// Set to true to include the signing certificate in the signature.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IncludeSigningCertInSignature *bool `mandatory:"false" json:"includeSigningCertInSignature"`

	// HTTP binding to use for authentication requests.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	AuthnRequestBinding IdentityProviderAuthnRequestBindingEnum `mandatory:"false" json:"authnRequestBinding,omitempty"`

	// HTTP binding to use for logout.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	LogoutBinding IdentityProviderLogoutBindingEnum `mandatory:"false" json:"logoutBinding,omitempty"`

	// Set to true to enable logout.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	LogoutEnabled *bool `mandatory:"false" json:"logoutEnabled"`

	// Signature hash algorithm.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	SignatureHashAlgorithm IdentityProviderSignatureHashAlgorithmEnum `mandatory:"false" json:"signatureHashAlgorithm,omitempty"`

	// Identity Provider Icon URL.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IconUrl *string `mandatory:"false" json:"iconUrl"`

	// Set to true to indicate whether to show IdP in login page or not.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	ShownOnLoginPage *bool `mandatory:"false" json:"shownOnLoginPage"`

	// Set to true to indicate JIT User Provisioning is enabled
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	JitUserProvEnabled *bool `mandatory:"false" json:"jitUserProvEnabled"`

	// Set to true to indicate JIT User Provisioning Groups should be assigned based on assertion attribute
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	JitUserProvGroupAssertionAttributeEnabled *bool `mandatory:"false" json:"jitUserProvGroupAssertionAttributeEnabled"`

	// Set to true to indicate JIT User Provisioning Groups should be assigned from a static list
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	JitUserProvGroupStaticListEnabled *bool `mandatory:"false" json:"jitUserProvGroupStaticListEnabled"`

	// Set to true to indicate JIT User Creation is enabled
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	JitUserProvCreateUserEnabled *bool `mandatory:"false" json:"jitUserProvCreateUserEnabled"`

	// Set to true to indicate JIT User Creation is enabled
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	JitUserProvAttributeUpdateEnabled *bool `mandatory:"false" json:"jitUserProvAttributeUpdateEnabled"`

	// The default value is 'Overwrite', which tells Just-In-Time user-provisioning to replace any current group-assignments for a User with those assigned by assertions and/or those assigned statically. Specify 'Merge' if you want Just-In-Time user-provisioning to combine its group-assignments with those the user already has.
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	JitUserProvGroupAssignmentMethod IdentityProviderJitUserProvGroupAssignmentMethodEnum `mandatory:"false" json:"jitUserProvGroupAssignmentMethod,omitempty"`

	// Property to indicate the mode of group mapping
	// **Added In:** 2205120021
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	JitUserProvGroupMappingMode IdentityProviderJitUserProvGroupMappingModeEnum `mandatory:"false" json:"jitUserProvGroupMappingMode,omitempty"`

	// Name of the assertion attribute containing the users groups
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	JitUserProvGroupSAMLAttributeName *string `mandatory:"false" json:"jitUserProvGroupSAMLAttributeName"`

	// The serviceInstanceIdentifier of the App that hosts this IdP. This value will match the opcServiceInstanceGUID of any service-instance that the IdP represents.
	// **Added In:** 18.2.6
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: never
	//  - type: string
	//  - uniqueness: server
	ServiceInstanceIdentifier *string `mandatory:"false" json:"serviceInstanceIdentifier"`

	// User mapping method.
	// **Deprecated Since: 20.1.3**
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	//  - idcsValuePersistedInOtherAttribute: true
	UserMappingMethod IdentityProviderUserMappingMethodEnum `mandatory:"false" json:"userMappingMethod,omitempty"`

	// This property specifies the userstore attribute value that must match the incoming assertion attribute value or the incoming nameid attribute value in order to identify the user during SSO.<br>You can construct the userMappingStoreAttribute value by specifying attributes from the Oracle Identity Cloud Service Core Users schema. For examples of how to construct the userMappingStoreAttribute value, see the <b>Example of a Request Body</b> section of the Examples tab for the <a href='./op-admin-v1-identityproviders-post.html'>POST</a> and <a href='./op-admin-v1-identityproviders-id-put.html'>PUT</a> methods of the /IdentityProviders endpoint.
	// **Deprecated Since: 20.1.3**
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	//  - idcsValuePersistedInOtherAttribute: true
	UserMappingStoreAttribute *string `mandatory:"false" json:"userMappingStoreAttribute"`

	// Assertion attribute name.
	// **Deprecated Since: 20.1.3**
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	//  - idcsValuePersistedInOtherAttribute: true
	AssertionAttribute *string `mandatory:"false" json:"assertionAttribute"`

	// Identity Provider Type
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: false
	//  - returned: always
	//  - type: string
	//  - uniqueness: none
	Type IdentityProviderTypeEnum `mandatory:"false" json:"type,omitempty"`

	// This SP requires requests SAML IdP to enforce re-authentication.
	// **Added In:** 2102181953
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	RequireForceAuthn *bool `mandatory:"false" json:"requireForceAuthn"`

	// SAML SP must accept encrypted assertion only.
	// **Added In:** 2102181953
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	RequiresEncryptedAssertion *bool `mandatory:"false" json:"requiresEncryptedAssertion"`

	// SAML SP HoK Enabled.
	// **Added In:** 2102181953
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	SamlHoKRequired *bool `mandatory:"false" json:"samlHoKRequired"`

	// SAML SP authentication type.
	// **Added In:** 2102181953
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	RequestedAuthenticationContext []string `mandatory:"false" json:"requestedAuthenticationContext"`

	// Set to true to indicate ignoring absence of group while provisioning
	// **Added In:** 2111112015
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsAddedSinceVersion: 30
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	JitUserProvIgnoreErrorOnAbsentGroups *bool `mandatory:"false" json:"jitUserProvIgnoreErrorOnAbsentGroups"`

	// Records the notification timestamp for the IdP whose signing certificate is about to expire
	// **Added In:** 2302092332
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: dateTime
	//  - uniqueness: none
	LastNotificationSentTime *string `mandatory:"false" json:"lastNotificationSentTime"`

	// The list of mappings between the Identity Domain Group and the IDP group.
	// **Added In:** 2205120021
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [idpGroup]
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	//  - uniqueness: none
	JitUserProvGroupMappings []IdentityProviderJitUserProvGroupMappings `mandatory:"false" json:"jitUserProvGroupMappings"`

	JitUserProvAttributes *IdentityProviderJitUserProvAttributes `mandatory:"false" json:"jitUserProvAttributes"`

	// Refers to every group of which a JIT-provisioned User should be a member.  Just-in-Time user-provisioning applies this static list when jitUserProvGroupStaticListEnabled:true.
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [value]
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	//  - uniqueness: none
	JitUserProvAssignedGroups []IdentityProviderJitUserProvAssignedGroups `mandatory:"false" json:"jitUserProvAssignedGroups"`

	CorrelationPolicy *IdentityProviderCorrelationPolicy `mandatory:"false" json:"correlationPolicy"`

	UrnIetfParamsScimSchemasOracleIdcsExtensionSocialIdentityProvider *ExtensionSocialIdentityProvider `mandatory:"false" json:"urn:ietf:params:scim:schemas:oracle:idcs:extension:social:IdentityProvider"`

	UrnIetfParamsScimSchemasOracleIdcsExtensionX509IdentityProvider *ExtensionX509IdentityProvider `mandatory:"false" json:"urn:ietf:params:scim:schemas:oracle:idcs:extension:x509:IdentityProvider"`
}

func (m IdentityProvider) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IdentityProvider) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.IdcsPreventedOperations {
		if _, ok := GetMappingIdcsPreventedOperationsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdcsPreventedOperations: %s. Supported values are: %s.", val, strings.Join(GetIdcsPreventedOperationsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingIdentityProviderAuthnRequestBindingEnum(string(m.AuthnRequestBinding)); !ok && m.AuthnRequestBinding != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuthnRequestBinding: %s. Supported values are: %s.", m.AuthnRequestBinding, strings.Join(GetIdentityProviderAuthnRequestBindingEnumStringValues(), ",")))
	}
	if _, ok := GetMappingIdentityProviderLogoutBindingEnum(string(m.LogoutBinding)); !ok && m.LogoutBinding != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LogoutBinding: %s. Supported values are: %s.", m.LogoutBinding, strings.Join(GetIdentityProviderLogoutBindingEnumStringValues(), ",")))
	}
	if _, ok := GetMappingIdentityProviderSignatureHashAlgorithmEnum(string(m.SignatureHashAlgorithm)); !ok && m.SignatureHashAlgorithm != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SignatureHashAlgorithm: %s. Supported values are: %s.", m.SignatureHashAlgorithm, strings.Join(GetIdentityProviderSignatureHashAlgorithmEnumStringValues(), ",")))
	}
	if _, ok := GetMappingIdentityProviderJitUserProvGroupAssignmentMethodEnum(string(m.JitUserProvGroupAssignmentMethod)); !ok && m.JitUserProvGroupAssignmentMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for JitUserProvGroupAssignmentMethod: %s. Supported values are: %s.", m.JitUserProvGroupAssignmentMethod, strings.Join(GetIdentityProviderJitUserProvGroupAssignmentMethodEnumStringValues(), ",")))
	}
	if _, ok := GetMappingIdentityProviderJitUserProvGroupMappingModeEnum(string(m.JitUserProvGroupMappingMode)); !ok && m.JitUserProvGroupMappingMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for JitUserProvGroupMappingMode: %s. Supported values are: %s.", m.JitUserProvGroupMappingMode, strings.Join(GetIdentityProviderJitUserProvGroupMappingModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingIdentityProviderUserMappingMethodEnum(string(m.UserMappingMethod)); !ok && m.UserMappingMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UserMappingMethod: %s. Supported values are: %s.", m.UserMappingMethod, strings.Join(GetIdentityProviderUserMappingMethodEnumStringValues(), ",")))
	}
	if _, ok := GetMappingIdentityProviderTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetIdentityProviderTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IdentityProviderAuthnRequestBindingEnum Enum with underlying type: string
type IdentityProviderAuthnRequestBindingEnum string

// Set of constants representing the allowable values for IdentityProviderAuthnRequestBindingEnum
const (
	IdentityProviderAuthnRequestBindingRedirect IdentityProviderAuthnRequestBindingEnum = "Redirect"
	IdentityProviderAuthnRequestBindingPost     IdentityProviderAuthnRequestBindingEnum = "Post"
)

var mappingIdentityProviderAuthnRequestBindingEnum = map[string]IdentityProviderAuthnRequestBindingEnum{
	"Redirect": IdentityProviderAuthnRequestBindingRedirect,
	"Post":     IdentityProviderAuthnRequestBindingPost,
}

var mappingIdentityProviderAuthnRequestBindingEnumLowerCase = map[string]IdentityProviderAuthnRequestBindingEnum{
	"redirect": IdentityProviderAuthnRequestBindingRedirect,
	"post":     IdentityProviderAuthnRequestBindingPost,
}

// GetIdentityProviderAuthnRequestBindingEnumValues Enumerates the set of values for IdentityProviderAuthnRequestBindingEnum
func GetIdentityProviderAuthnRequestBindingEnumValues() []IdentityProviderAuthnRequestBindingEnum {
	values := make([]IdentityProviderAuthnRequestBindingEnum, 0)
	for _, v := range mappingIdentityProviderAuthnRequestBindingEnum {
		values = append(values, v)
	}
	return values
}

// GetIdentityProviderAuthnRequestBindingEnumStringValues Enumerates the set of values in String for IdentityProviderAuthnRequestBindingEnum
func GetIdentityProviderAuthnRequestBindingEnumStringValues() []string {
	return []string{
		"Redirect",
		"Post",
	}
}

// GetMappingIdentityProviderAuthnRequestBindingEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIdentityProviderAuthnRequestBindingEnum(val string) (IdentityProviderAuthnRequestBindingEnum, bool) {
	enum, ok := mappingIdentityProviderAuthnRequestBindingEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// IdentityProviderLogoutBindingEnum Enum with underlying type: string
type IdentityProviderLogoutBindingEnum string

// Set of constants representing the allowable values for IdentityProviderLogoutBindingEnum
const (
	IdentityProviderLogoutBindingRedirect IdentityProviderLogoutBindingEnum = "Redirect"
	IdentityProviderLogoutBindingPost     IdentityProviderLogoutBindingEnum = "Post"
)

var mappingIdentityProviderLogoutBindingEnum = map[string]IdentityProviderLogoutBindingEnum{
	"Redirect": IdentityProviderLogoutBindingRedirect,
	"Post":     IdentityProviderLogoutBindingPost,
}

var mappingIdentityProviderLogoutBindingEnumLowerCase = map[string]IdentityProviderLogoutBindingEnum{
	"redirect": IdentityProviderLogoutBindingRedirect,
	"post":     IdentityProviderLogoutBindingPost,
}

// GetIdentityProviderLogoutBindingEnumValues Enumerates the set of values for IdentityProviderLogoutBindingEnum
func GetIdentityProviderLogoutBindingEnumValues() []IdentityProviderLogoutBindingEnum {
	values := make([]IdentityProviderLogoutBindingEnum, 0)
	for _, v := range mappingIdentityProviderLogoutBindingEnum {
		values = append(values, v)
	}
	return values
}

// GetIdentityProviderLogoutBindingEnumStringValues Enumerates the set of values in String for IdentityProviderLogoutBindingEnum
func GetIdentityProviderLogoutBindingEnumStringValues() []string {
	return []string{
		"Redirect",
		"Post",
	}
}

// GetMappingIdentityProviderLogoutBindingEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIdentityProviderLogoutBindingEnum(val string) (IdentityProviderLogoutBindingEnum, bool) {
	enum, ok := mappingIdentityProviderLogoutBindingEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// IdentityProviderSignatureHashAlgorithmEnum Enum with underlying type: string
type IdentityProviderSignatureHashAlgorithmEnum string

// Set of constants representing the allowable values for IdentityProviderSignatureHashAlgorithmEnum
const (
	IdentityProviderSignatureHashAlgorithm1   IdentityProviderSignatureHashAlgorithmEnum = "SHA-1"
	IdentityProviderSignatureHashAlgorithm256 IdentityProviderSignatureHashAlgorithmEnum = "SHA-256"
)

var mappingIdentityProviderSignatureHashAlgorithmEnum = map[string]IdentityProviderSignatureHashAlgorithmEnum{
	"SHA-1":   IdentityProviderSignatureHashAlgorithm1,
	"SHA-256": IdentityProviderSignatureHashAlgorithm256,
}

var mappingIdentityProviderSignatureHashAlgorithmEnumLowerCase = map[string]IdentityProviderSignatureHashAlgorithmEnum{
	"sha-1":   IdentityProviderSignatureHashAlgorithm1,
	"sha-256": IdentityProviderSignatureHashAlgorithm256,
}

// GetIdentityProviderSignatureHashAlgorithmEnumValues Enumerates the set of values for IdentityProviderSignatureHashAlgorithmEnum
func GetIdentityProviderSignatureHashAlgorithmEnumValues() []IdentityProviderSignatureHashAlgorithmEnum {
	values := make([]IdentityProviderSignatureHashAlgorithmEnum, 0)
	for _, v := range mappingIdentityProviderSignatureHashAlgorithmEnum {
		values = append(values, v)
	}
	return values
}

// GetIdentityProviderSignatureHashAlgorithmEnumStringValues Enumerates the set of values in String for IdentityProviderSignatureHashAlgorithmEnum
func GetIdentityProviderSignatureHashAlgorithmEnumStringValues() []string {
	return []string{
		"SHA-1",
		"SHA-256",
	}
}

// GetMappingIdentityProviderSignatureHashAlgorithmEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIdentityProviderSignatureHashAlgorithmEnum(val string) (IdentityProviderSignatureHashAlgorithmEnum, bool) {
	enum, ok := mappingIdentityProviderSignatureHashAlgorithmEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// IdentityProviderJitUserProvGroupAssignmentMethodEnum Enum with underlying type: string
type IdentityProviderJitUserProvGroupAssignmentMethodEnum string

// Set of constants representing the allowable values for IdentityProviderJitUserProvGroupAssignmentMethodEnum
const (
	IdentityProviderJitUserProvGroupAssignmentMethodOverwrite IdentityProviderJitUserProvGroupAssignmentMethodEnum = "Overwrite"
	IdentityProviderJitUserProvGroupAssignmentMethodMerge     IdentityProviderJitUserProvGroupAssignmentMethodEnum = "Merge"
)

var mappingIdentityProviderJitUserProvGroupAssignmentMethodEnum = map[string]IdentityProviderJitUserProvGroupAssignmentMethodEnum{
	"Overwrite": IdentityProviderJitUserProvGroupAssignmentMethodOverwrite,
	"Merge":     IdentityProviderJitUserProvGroupAssignmentMethodMerge,
}

var mappingIdentityProviderJitUserProvGroupAssignmentMethodEnumLowerCase = map[string]IdentityProviderJitUserProvGroupAssignmentMethodEnum{
	"overwrite": IdentityProviderJitUserProvGroupAssignmentMethodOverwrite,
	"merge":     IdentityProviderJitUserProvGroupAssignmentMethodMerge,
}

// GetIdentityProviderJitUserProvGroupAssignmentMethodEnumValues Enumerates the set of values for IdentityProviderJitUserProvGroupAssignmentMethodEnum
func GetIdentityProviderJitUserProvGroupAssignmentMethodEnumValues() []IdentityProviderJitUserProvGroupAssignmentMethodEnum {
	values := make([]IdentityProviderJitUserProvGroupAssignmentMethodEnum, 0)
	for _, v := range mappingIdentityProviderJitUserProvGroupAssignmentMethodEnum {
		values = append(values, v)
	}
	return values
}

// GetIdentityProviderJitUserProvGroupAssignmentMethodEnumStringValues Enumerates the set of values in String for IdentityProviderJitUserProvGroupAssignmentMethodEnum
func GetIdentityProviderJitUserProvGroupAssignmentMethodEnumStringValues() []string {
	return []string{
		"Overwrite",
		"Merge",
	}
}

// GetMappingIdentityProviderJitUserProvGroupAssignmentMethodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIdentityProviderJitUserProvGroupAssignmentMethodEnum(val string) (IdentityProviderJitUserProvGroupAssignmentMethodEnum, bool) {
	enum, ok := mappingIdentityProviderJitUserProvGroupAssignmentMethodEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// IdentityProviderJitUserProvGroupMappingModeEnum Enum with underlying type: string
type IdentityProviderJitUserProvGroupMappingModeEnum string

// Set of constants representing the allowable values for IdentityProviderJitUserProvGroupMappingModeEnum
const (
	IdentityProviderJitUserProvGroupMappingModeImplicit IdentityProviderJitUserProvGroupMappingModeEnum = "implicit"
	IdentityProviderJitUserProvGroupMappingModeExplicit IdentityProviderJitUserProvGroupMappingModeEnum = "explicit"
)

var mappingIdentityProviderJitUserProvGroupMappingModeEnum = map[string]IdentityProviderJitUserProvGroupMappingModeEnum{
	"implicit": IdentityProviderJitUserProvGroupMappingModeImplicit,
	"explicit": IdentityProviderJitUserProvGroupMappingModeExplicit,
}

var mappingIdentityProviderJitUserProvGroupMappingModeEnumLowerCase = map[string]IdentityProviderJitUserProvGroupMappingModeEnum{
	"implicit": IdentityProviderJitUserProvGroupMappingModeImplicit,
	"explicit": IdentityProviderJitUserProvGroupMappingModeExplicit,
}

// GetIdentityProviderJitUserProvGroupMappingModeEnumValues Enumerates the set of values for IdentityProviderJitUserProvGroupMappingModeEnum
func GetIdentityProviderJitUserProvGroupMappingModeEnumValues() []IdentityProviderJitUserProvGroupMappingModeEnum {
	values := make([]IdentityProviderJitUserProvGroupMappingModeEnum, 0)
	for _, v := range mappingIdentityProviderJitUserProvGroupMappingModeEnum {
		values = append(values, v)
	}
	return values
}

// GetIdentityProviderJitUserProvGroupMappingModeEnumStringValues Enumerates the set of values in String for IdentityProviderJitUserProvGroupMappingModeEnum
func GetIdentityProviderJitUserProvGroupMappingModeEnumStringValues() []string {
	return []string{
		"implicit",
		"explicit",
	}
}

// GetMappingIdentityProviderJitUserProvGroupMappingModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIdentityProviderJitUserProvGroupMappingModeEnum(val string) (IdentityProviderJitUserProvGroupMappingModeEnum, bool) {
	enum, ok := mappingIdentityProviderJitUserProvGroupMappingModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// IdentityProviderUserMappingMethodEnum Enum with underlying type: string
type IdentityProviderUserMappingMethodEnum string

// Set of constants representing the allowable values for IdentityProviderUserMappingMethodEnum
const (
	IdentityProviderUserMappingMethodNameidtouserattribute             IdentityProviderUserMappingMethodEnum = "NameIDToUserAttribute"
	IdentityProviderUserMappingMethodAssertionattributetouserattribute IdentityProviderUserMappingMethodEnum = "AssertionAttributeToUserAttribute"
	IdentityProviderUserMappingMethodCorrelationpolicyrule             IdentityProviderUserMappingMethodEnum = "CorrelationPolicyRule"
)

var mappingIdentityProviderUserMappingMethodEnum = map[string]IdentityProviderUserMappingMethodEnum{
	"NameIDToUserAttribute":             IdentityProviderUserMappingMethodNameidtouserattribute,
	"AssertionAttributeToUserAttribute": IdentityProviderUserMappingMethodAssertionattributetouserattribute,
	"CorrelationPolicyRule":             IdentityProviderUserMappingMethodCorrelationpolicyrule,
}

var mappingIdentityProviderUserMappingMethodEnumLowerCase = map[string]IdentityProviderUserMappingMethodEnum{
	"nameidtouserattribute":             IdentityProviderUserMappingMethodNameidtouserattribute,
	"assertionattributetouserattribute": IdentityProviderUserMappingMethodAssertionattributetouserattribute,
	"correlationpolicyrule":             IdentityProviderUserMappingMethodCorrelationpolicyrule,
}

// GetIdentityProviderUserMappingMethodEnumValues Enumerates the set of values for IdentityProviderUserMappingMethodEnum
func GetIdentityProviderUserMappingMethodEnumValues() []IdentityProviderUserMappingMethodEnum {
	values := make([]IdentityProviderUserMappingMethodEnum, 0)
	for _, v := range mappingIdentityProviderUserMappingMethodEnum {
		values = append(values, v)
	}
	return values
}

// GetIdentityProviderUserMappingMethodEnumStringValues Enumerates the set of values in String for IdentityProviderUserMappingMethodEnum
func GetIdentityProviderUserMappingMethodEnumStringValues() []string {
	return []string{
		"NameIDToUserAttribute",
		"AssertionAttributeToUserAttribute",
		"CorrelationPolicyRule",
	}
}

// GetMappingIdentityProviderUserMappingMethodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIdentityProviderUserMappingMethodEnum(val string) (IdentityProviderUserMappingMethodEnum, bool) {
	enum, ok := mappingIdentityProviderUserMappingMethodEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// IdentityProviderTypeEnum Enum with underlying type: string
type IdentityProviderTypeEnum string

// Set of constants representing the allowable values for IdentityProviderTypeEnum
const (
	IdentityProviderTypeSaml   IdentityProviderTypeEnum = "SAML"
	IdentityProviderTypeSocial IdentityProviderTypeEnum = "SOCIAL"
	IdentityProviderTypeIwa    IdentityProviderTypeEnum = "IWA"
	IdentityProviderTypeX509   IdentityProviderTypeEnum = "X509"
	IdentityProviderTypeLocal  IdentityProviderTypeEnum = "LOCAL"
)

var mappingIdentityProviderTypeEnum = map[string]IdentityProviderTypeEnum{
	"SAML":   IdentityProviderTypeSaml,
	"SOCIAL": IdentityProviderTypeSocial,
	"IWA":    IdentityProviderTypeIwa,
	"X509":   IdentityProviderTypeX509,
	"LOCAL":  IdentityProviderTypeLocal,
}

var mappingIdentityProviderTypeEnumLowerCase = map[string]IdentityProviderTypeEnum{
	"saml":   IdentityProviderTypeSaml,
	"social": IdentityProviderTypeSocial,
	"iwa":    IdentityProviderTypeIwa,
	"x509":   IdentityProviderTypeX509,
	"local":  IdentityProviderTypeLocal,
}

// GetIdentityProviderTypeEnumValues Enumerates the set of values for IdentityProviderTypeEnum
func GetIdentityProviderTypeEnumValues() []IdentityProviderTypeEnum {
	values := make([]IdentityProviderTypeEnum, 0)
	for _, v := range mappingIdentityProviderTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetIdentityProviderTypeEnumStringValues Enumerates the set of values in String for IdentityProviderTypeEnum
func GetIdentityProviderTypeEnumStringValues() []string {
	return []string{
		"SAML",
		"SOCIAL",
		"IWA",
		"X509",
		"LOCAL",
	}
}

// GetMappingIdentityProviderTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIdentityProviderTypeEnum(val string) (IdentityProviderTypeEnum, bool) {
	enum, ok := mappingIdentityProviderTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
