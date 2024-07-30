// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Domains API
//
// Use the Identity Domains API to manage resources within an identity domain, for example, users, dynamic resource groups, groups, and identity providers. For information about managing resources within identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm).
// Use this pattern to construct endpoints for identity domains: `https://<domainURL>/admin/v1/`. See Finding an Identity Domain URL (https://docs.oracle.com/en-us/iaas/Content/Identity/api-getstarted/locate-identity-domain-url.htm) to locate the domain URL you need.
// Use the table of contents and search tool to explore the Identity Domains API.
//

package identitydomains

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// App Schema for App resource.
type App struct {

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

	// Display name of the application. Display name is intended to be user-friendly, and an administrator can change the value at any time.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: always
	//  - type: string
	//  - uniqueness: server
	DisplayName *string `mandatory:"true" json:"displayName"`

	BasedOnTemplate *AppBasedOnTemplate `mandatory:"true" json:"basedOnTemplate"`

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

	// Indicates whether the application is billed as an OPCService. If true, customer is not billed for runtime operations of the app.
	// **Added In:** 18.4.2
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: always
	//  - type: boolean
	//  - uniqueness: none
	MeterAsOPCService *bool `mandatory:"false" json:"meterAsOPCService"`

	// Name of the application. Also serves as username if the application authenticates to Oracle Public Cloud infrastructure. This name may not be user-friendly and cannot be changed once an App is created.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: server
	Name *string `mandatory:"false" json:"name"`

	// This value is the credential of this App, which this App supplies as a password when this App authenticates to the Oracle Public Cloud infrastructure. This value is also the client secret of this App when it acts as an OAuthClient.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - idcsSensitive: none
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	ClientSecret *string `mandatory:"false" json:"clientSecret"`

	// Hashed Client Secret. This hash-value is used to verify the 'clientSecret' credential of this App
	// **Added In:** 2106240046
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - idcsSensitive: hash_sc
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	HashedClientSecret *string `mandatory:"false" json:"hashedClientSecret"`

	// Description of the application.
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

	// Encryption Alogrithm to use for encrypting ID token.
	// **Added In:** 2010242156
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdTokenEncAlgo *string `mandatory:"false" json:"idTokenEncAlgo"`

	// Service Names allow to use OCI signature for client authentication instead of client credentials
	// **Added In:** 2207040824
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	DelegatedServiceNames []string `mandatory:"false" json:"delegatedServiceNames"`

	// If true, this App is able to participate in runtime services, such as automatic-login, OAuth, and SAML. If false, all runtime services are disabled for this App, and only administrative operations can be performed.
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	Active *bool `mandatory:"false" json:"active"`

	// Application icon.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	AppIcon *string `mandatory:"false" json:"appIcon"`

	// Application thumbnail.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	AppThumbnail *string `mandatory:"false" json:"appThumbnail"`

	// If true, this App was migrated from an earlier version of Oracle Public Cloud infrastructure (and may therefore require special handling from runtime services such as OAuth or SAML). If false, this App requires no special handling from runtime services.
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	Migrated *bool `mandatory:"false" json:"migrated"`

	// If true, this App is an internal infrastructure App.
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	Infrastructure *bool `mandatory:"false" json:"infrastructure"`

	// If true, this App allows runtime services to log end users into this App automatically.
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IsLoginTarget *bool `mandatory:"false" json:"isLoginTarget"`

	// If true, this app will be displayed in the MyApps page of each end-user who has access to the App.
	// **Added In:** 18.1.2
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	ShowInMyApps *bool `mandatory:"false" json:"showInMyApps"`

	// The protocol that runtime services will use to log end users in to this App automatically. If 'OIDC', then runtime services use the OpenID Connect protocol. If 'SAML', then runtime services use Security Assertion Markup Language protocol.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	LoginMechanism AppLoginMechanismEnum `mandatory:"false" json:"loginMechanism,omitempty"`

	// The URL of the landing page for this App, which is the first page that an end user should see if runtime services log that end user in to this App automatically.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	LandingPageUrl *string `mandatory:"false" json:"landingPageUrl"`

	// Application Logo URL
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	ProductLogoUrl *string `mandatory:"false" json:"productLogoUrl"`

	// Privacy Policy URL
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	PrivacyPolicyUrl *string `mandatory:"false" json:"privacyPolicyUrl"`

	// Terms of Service URL
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	TermsOfServiceUrl *string `mandatory:"false" json:"termsOfServiceUrl"`

	// Contact Email Address
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	ContactEmailAddress *string `mandatory:"false" json:"contactEmailAddress"`

	// Product Name
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	ProductName *string `mandatory:"false" json:"productName"`

	// Home Page URL
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	HomePageUrl *string `mandatory:"false" json:"homePageUrl"`

	// If true, this application acts as FormFill Application
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IsFormFill *bool `mandatory:"false" json:"isFormFill"`

	// If true, this application acts as an OAuth Client
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IsOAuthClient *bool `mandatory:"false" json:"isOAuthClient"`

	// If true, this application acts as an Radius App
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IsRadiusApp *bool `mandatory:"false" json:"isRadiusApp"`

	// Specifies the type of access that this App has when it acts as an OAuthClient.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	ClientType AppClientTypeEnum `mandatory:"false" json:"clientType,omitempty"`

	// OPTIONAL. Each value is a URI within this App. This attribute is required when this App acts as an OAuthClient and is involved in three-legged flows (authorization-code flows).
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	RedirectUris []string `mandatory:"false" json:"redirectUris"`

	// If true, indicates that the system should allow all URL-schemes within each value of the 'redirectUris' attribute.  Also indicates that the system should not attempt to confirm that each value of the 'redirectUris' attribute is a valid URI.  In particular, the system should not confirm that the domain component of the URI is a top-level domain and the system should not confirm that the hostname portion is a valid system that is reachable over the network.
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	AllUrlSchemesAllowed *bool `mandatory:"false" json:"allUrlSchemesAllowed"`

	// OAuth will use this URI to logout if this App wants to participate in SSO, and if this App's session gets cleared as part of global logout. Note: This attribute is used only if this App acts as an OAuthClient.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	LogoutUri *string `mandatory:"false" json:"logoutUri"`

	// Each value of this attribute is the URI of a landing page within this App. It is used only when this App, acting as an OAuthClient, initiates the logout flow and wants to be redirected back to one of its landing pages.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	PostLogoutRedirectUris []string `mandatory:"false" json:"postLogoutRedirectUris"`

	// List of grant-types that this App is allowed to use when it acts as an OAuthClient.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	AllowedGrants []string `mandatory:"false" json:"allowedGrants"`

	// OPTIONAL. Required only when this App acts as an OAuthClient. Supported values are 'introspect' and 'onBehalfOfUser'. The value 'introspect' allows the client to look inside the access-token. The value 'onBehalfOfUser' overrides how the client's privileges are combined with the privileges of the Subject User. Ordinarily, authorization calculates the set of effective privileges as the intersection of the client's privileges and the user's privileges. The value 'onBehalfOf' indicates that authorization should ignore the privileges of the client and use only the user's privileges to calculate the effective privileges.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	AllowedOperations []AppAllowedOperationsEnum `mandatory:"false" json:"allowedOperations,omitempty"`

	// Network Perimeters checking mode
	// **Added In:** 2010242156
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	ClientIPChecking AppClientIPCheckingEnum `mandatory:"false" json:"clientIPChecking,omitempty"`

	// If true, this application is an Oracle Public Cloud service-instance.
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IsOPCService *bool `mandatory:"false" json:"isOPCService"`

	// If true, indicates that this application accepts an Oracle Cloud Identity Service User as a login-identity (does not require an account) and relies for authorization on the User's memberships in AppRoles.
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IsUnmanagedApp *bool `mandatory:"false" json:"isUnmanagedApp"`

	// If true, any managed App that is based on this template is checked for access control that is, access to this app is subject to successful authorization at SSO service, viz. app grants to start with.
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	AllowAccessControl *bool `mandatory:"false" json:"allowAccessControl"`

	// If true, indicates that this application acts as an OAuth Resource.
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IsOAuthResource *bool `mandatory:"false" json:"isOAuthResource"`

	// Expiry-time in seconds for an Access Token. Any token that allows access to this App will expire after the specified duration.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	AccessTokenExpiry *int `mandatory:"false" json:"accessTokenExpiry"`

	// Expiry-time in seconds for a Refresh Token.  Any token that allows access to this App, once refreshed, will expire after the specified duration.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	RefreshTokenExpiry *int `mandatory:"false" json:"refreshTokenExpiry"`

	// If true, indicates that the Refresh Token is allowed when this App acts as an OAuth Resource.
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	AllowOffline *bool `mandatory:"false" json:"allowOffline"`

	// Callback Service URL
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	CallbackServiceUrl *string `mandatory:"false" json:"callbackServiceUrl"`

	// The base URI for all of the scopes defined in this App. The value of 'audience' is combined with the 'value' of each scope to form an 'fqs' or fully qualified scope.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Audience *string `mandatory:"false" json:"audience"`

	// If true, indicates that the App should be visible in each end-user's mobile application.
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IsMobileTarget *bool `mandatory:"false" json:"isMobileTarget"`

	// This attribute specifies the URL of the page that the App uses when an end-user signs in to that App.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	LoginPageUrl *string `mandatory:"false" json:"loginPageUrl"`

	// This attribute specifies the callback URL for the social linking operation.
	// **Added In:** 18.2.4
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	LinkingCallbackUrl *string `mandatory:"false" json:"linkingCallbackUrl"`

	// This attribute specifies the URL of the page that the App uses when an end-user signs out.
	// **Added In:** 17.4.2
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	LogoutPageUrl *string `mandatory:"false" json:"logoutPageUrl"`

	// This attribute specifies the URL of the page to which an application will redirect an end-user in case of error.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	ErrorPageUrl *string `mandatory:"false" json:"errorPageUrl"`

	// If true, then this App acts as a SAML Service Provider.
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IsSamlServiceProvider *bool `mandatory:"false" json:"isSamlServiceProvider"`

	// If true, the webtier policy is active
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IsWebTierPolicy *bool `mandatory:"false" json:"isWebTierPolicy"`

	// If true, indicates that this App supports Kerberos Authentication
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IsKerberosRealm *bool `mandatory:"false" json:"isKerberosRealm"`

	// URL of application icon.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: reference
	//  - uniqueness: none
	Icon *string `mandatory:"false" json:"icon"`

	// If true, this App is an AliasApp and it cannot be granted to an end-user directly.
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: false
	//  - returned: always
	//  - type: boolean
	//  - uniqueness: none
	IsAliasApp *bool `mandatory:"false" json:"isAliasApp"`

	// If true, indicates that access to this App requires an account. That is, in order to log in to the App, a User must use an application-specific identity that is maintained in the remote identity-repository of that App.
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IsManagedApp *bool `mandatory:"false" json:"isManagedApp"`

	// This Uniform Resource Name (URN) value identifies the type of Oracle Public Cloud service of which this app is an instance.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	ServiceTypeURN *string `mandatory:"false" json:"serviceTypeURN"`

	// This value specifies the version of the Oracle Public Cloud service of which this App is an instance
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	ServiceTypeVersion *string `mandatory:"false" json:"serviceTypeVersion"`

	// This flag indicates if the App is capable of validating obligations with the token for allowing access to the App.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IsObligationCapable *bool `mandatory:"false" json:"isObligationCapable"`

	// If true, this App requires an upgrade and mandates attention from application administrator. The flag is used by UI to indicate this app is ready to upgrade.
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	ReadyToUpgrade *bool `mandatory:"false" json:"readyToUpgrade"`

	// Indicates the scope of trust for this App when acting as an OAuthClient. A value of 'Explicit' indicates that the App is allowed to access only the scopes of OAuthResources that are explicitly specified as 'allowedScopes'. A value of 'Account' indicates that the App is allowed implicitly to access any scope of any OAuthResource within the same Oracle Cloud Account. A value of 'Tags' indicates that the App is allowed to access any scope of any OAuthResource with a matching tag within the same Oracle Cloud Account. A value of 'Default' indicates that the Tenant default trust scope configured in the Tenant Settings is used.
	// **Added In:** 17.4.2
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	TrustScope AppTrustScopeEnum `mandatory:"false" json:"trustScope,omitempty"`

	// If true, this application acts as database service Application
	// **Added In:** 18.2.2
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - type: boolean
	IsDatabaseService *bool `mandatory:"false" json:"isDatabaseService"`

	// A list of secondary audiences--additional URIs to be added automatically to any OAuth token that allows access to this App. Note: This attribute is used mainly for backward compatibility in certain Oracle Public Cloud Apps.
	// **Deprecated Since: 18.2.6**
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	SecondaryAudiences []string `mandatory:"false" json:"secondaryAudiences"`

	// If true, this app acts as Enterprise app with Authentication and URL Authz policy.
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IsEnterpriseApp *bool `mandatory:"false" json:"isEnterpriseApp"`

	// If true, indicates that consent should be skipped for all scopes
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	BypassConsent *bool `mandatory:"false" json:"bypassConsent"`

	// Indicates whether the application is allowed to be access using kmsi token.
	// **Added In:** 2111190457
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: always
	//  - type: boolean
	//  - uniqueness: none
	DisableKmsiTokenAuthentication *bool `mandatory:"false" json:"disableKmsiTokenAuthentication"`

	// If true, indicates the app is used for multicloud service integration.
	// **Added In:** 2301202328
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IsMulticloudServiceApp *bool `mandatory:"false" json:"isMulticloudServiceApp"`

	RadiusPolicy *AppRadiusPolicy `mandatory:"false" json:"radiusPolicy"`

	// Network Perimeter
	// **Added In:** 2010242156
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [value]
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	AppsNetworkPerimeters []AppAppsNetworkPerimeters `mandatory:"false" json:"appsNetworkPerimeters"`

	// A collection of arbitrary properties that scope the privileges of a cloud-control App.
	// **Added In:** 18.4.2
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [name]
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: complex
	//  - uniqueness: none
	CloudControlProperties []AppCloudControlProperties `mandatory:"false" json:"cloudControlProperties"`

	// App attributes editable by subject
	// **Added In:** 18.2.6
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCompositeKey: [name]
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: complex
	//  - uniqueness: none
	EditableAttributes []AppEditableAttributes `mandatory:"false" json:"editableAttributes"`

	TermsOfUse *AppTermsOfUse `mandatory:"false" json:"termsOfUse"`

	// A list of secondary audiences--additional URIs to be added automatically to any OAuth token that allows access to this App. Note: This attribute is used mainly for backward compatibility in certain Oracle Public Cloud Apps.
	// **Added In:** 18.2.2
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCompositeKey: [value]
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	//  - uniqueness: none
	ProtectableSecondaryAudiences []AppProtectableSecondaryAudiences `mandatory:"false" json:"protectableSecondaryAudiences"`

	IdpPolicy *AppIdpPolicy `mandatory:"false" json:"idpPolicy"`

	// A list of tags, acting as an OAuthClient, this App is allowed to access.
	// **Added In:** 17.4.6
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [key, value]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	//  - uniqueness: none
	AllowedTags []AppAllowedTags `mandatory:"false" json:"allowedTags"`

	AppSignonPolicy *AppAppSignonPolicy `mandatory:"false" json:"appSignonPolicy"`

	// Trust Policies.
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [value]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	TrustPolicies []AppTrustPolicies `mandatory:"false" json:"trustPolicies"`

	SignonPolicy *AppSignonPolicy `mandatory:"false" json:"signonPolicy"`

	// A list of IdentityProvider assigned to app. A user trying to access this app will be automatically redirected to configured IdP during the authentication phase, before being able to access App.
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [value]
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: request
	//  - type: complex
	IdentityProviders []AppIdentityProviders `mandatory:"false" json:"identityProviders"`

	// Accounts of App
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [value]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: complex
	//  - uniqueness: none
	Accounts []AppAccounts `mandatory:"false" json:"accounts"`

	// Grants assigned to the app
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [value]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: complex
	//  - uniqueness: none
	Grants []AppGrants `mandatory:"false" json:"grants"`

	// Custom attribute that is required to compute other attribute values during app creation.
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [name]
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: always
	//  - type: complex
	//  - uniqueness: none
	ServiceParams []AppServiceParams `mandatory:"false" json:"serviceParams"`

	// Label for the attribute to be shown in the UI.
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [name]
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: immutable
	//  - required: false
	//  - returned: default
	//  - type: complex
	//  - uniqueness: none
	AttrRenderingMetadata []AppAttrRenderingMetadata `mandatory:"false" json:"attrRenderingMetadata"`

	// A list of AppRoles that are granted to this App (and that are defined by other Apps). Within the Oracle Public Cloud infrastructure, this allows AppID-based association. Such an association allows this App to act as a consumer and thus to access resources of another App that acts as a producer.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsCompositeKey: [value]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: complex
	//  - uniqueness: none
	GrantedAppRoles []AppGrantedAppRoles `mandatory:"false" json:"grantedAppRoles"`

	SamlServiceProvider *AppSamlServiceProvider `mandatory:"false" json:"samlServiceProvider"`

	// A list of scopes (exposed by this App or by other Apps) that this App is allowed to access when it acts as an OAuthClient.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsCompositeKey: [fqs]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	//  - uniqueness: none
	AllowedScopes []AppAllowedScopes `mandatory:"false" json:"allowedScopes"`

	// Each value of this attribute represent a certificate that this App uses when it acts as an OAuthClient.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCompositeKey: [certAlias]
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	//  - uniqueness: none
	Certificates []AppCertificates `mandatory:"false" json:"certificates"`

	// Each value of this internal attribute refers to an Oracle Public Cloud infrastructure App on which this App depends.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsCompositeKey: [value]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	//  - uniqueness: none
	AliasApps []AppAliasApps `mandatory:"false" json:"aliasApps"`

	AsOPCService *AppAsOpcService `mandatory:"false" json:"asOPCService"`

	// A list of AppRoles defined by this UnmanagedApp. Membership in each of these AppRoles confers administrative privilege within this App.
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [value]
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: complex
	AdminRoles []AppAdminRoles `mandatory:"false" json:"adminRoles"`

	// A list of AppRoles defined by this UnmanagedApp. Membership in each of these AppRoles confers end-user privilege within this App.
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [value]
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: complex
	UserRoles []AppUserRoles `mandatory:"false" json:"userRoles"`

	// Scopes defined by this App. Used when this App acts as an OAuth Resource.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsCompositeKey: [value]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	//  - uniqueness: none
	Scopes []AppScopes `mandatory:"false" json:"scopes"`

	UrnIetfParamsScimSchemasOracleIdcsExtensionRadiusAppApp *AppExtensionRadiusAppApp `mandatory:"false" json:"urn:ietf:params:scim:schemas:oracle:idcs:extension:radiusApp:App"`

	UrnIetfParamsScimSchemasOracleIdcsExtensionSamlServiceProviderApp *AppExtensionSamlServiceProviderApp `mandatory:"false" json:"urn:ietf:params:scim:schemas:oracle:idcs:extension:samlServiceProvider:App"`

	UrnIetfParamsScimSchemasOracleIdcsExtensionWebTierPolicyApp *AppExtensionWebTierPolicyApp `mandatory:"false" json:"urn:ietf:params:scim:schemas:oracle:idcs:extension:webTierPolicy:App"`

	UrnIetfParamsScimSchemasOracleIdcsExtensionManagedappApp *AppExtensionManagedappApp `mandatory:"false" json:"urn:ietf:params:scim:schemas:oracle:idcs:extension:managedapp:App"`

	UrnIetfParamsScimSchemasOracleIdcsExtensionFormFillAppTemplateAppTemplate *AppExtensionFormFillAppTemplateAppTemplate `mandatory:"false" json:"urn:ietf:params:scim:schemas:oracle:idcs:extension:formFillAppTemplate:AppTemplate"`

	UrnIetfParamsScimSchemasOracleIdcsExtensionOpcServiceApp *AppExtensionOpcServiceApp `mandatory:"false" json:"urn:ietf:params:scim:schemas:oracle:idcs:extension:opcService:App"`

	UrnIetfParamsScimSchemasOracleIdcsExtensionKerberosRealmApp *AppExtensionKerberosRealmApp `mandatory:"false" json:"urn:ietf:params:scim:schemas:oracle:idcs:extension:kerberosRealm:App"`

	UrnIetfParamsScimSchemasOracleIdcsExtensionRequestableApp *AppExtensionRequestableApp `mandatory:"false" json:"urn:ietf:params:scim:schemas:oracle:idcs:extension:requestable:App"`

	UrnIetfParamsScimSchemasOracleIdcsExtensionFormFillAppApp *AppExtensionFormFillAppApp `mandatory:"false" json:"urn:ietf:params:scim:schemas:oracle:idcs:extension:formFillApp:App"`

	UrnIetfParamsScimSchemasOracleIdcsExtensionDbcsApp *AppExtensionDbcsApp `mandatory:"false" json:"urn:ietf:params:scim:schemas:oracle:idcs:extension:dbcs:App"`

	UrnIetfParamsScimSchemasOracleIdcsExtensionEnterpriseAppApp *AppExtensionEnterpriseAppApp `mandatory:"false" json:"urn:ietf:params:scim:schemas:oracle:idcs:extension:enterpriseApp:App"`

	UrnIetfParamsScimSchemasOracleIdcsExtensionOciTags *ExtensionOciTags `mandatory:"false" json:"urn:ietf:params:scim:schemas:oracle:idcs:extension:OCITags"`

	UrnIetfParamsScimSchemasOracleIdcsExtensionMulticloudServiceAppApp *AppExtensionMulticloudServiceAppApp `mandatory:"false" json:"urn:ietf:params:scim:schemas:oracle:idcs:extension:multicloudServiceApp:App"`
}

func (m App) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m App) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.IdcsPreventedOperations {
		if _, ok := GetMappingIdcsPreventedOperationsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdcsPreventedOperations: %s. Supported values are: %s.", val, strings.Join(GetIdcsPreventedOperationsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingAppLoginMechanismEnum(string(m.LoginMechanism)); !ok && m.LoginMechanism != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LoginMechanism: %s. Supported values are: %s.", m.LoginMechanism, strings.Join(GetAppLoginMechanismEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAppClientTypeEnum(string(m.ClientType)); !ok && m.ClientType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ClientType: %s. Supported values are: %s.", m.ClientType, strings.Join(GetAppClientTypeEnumStringValues(), ",")))
	}
	for _, val := range m.AllowedOperations {
		if _, ok := GetMappingAppAllowedOperationsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AllowedOperations: %s. Supported values are: %s.", val, strings.Join(GetAppAllowedOperationsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingAppClientIPCheckingEnum(string(m.ClientIPChecking)); !ok && m.ClientIPChecking != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ClientIPChecking: %s. Supported values are: %s.", m.ClientIPChecking, strings.Join(GetAppClientIPCheckingEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAppTrustScopeEnum(string(m.TrustScope)); !ok && m.TrustScope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TrustScope: %s. Supported values are: %s.", m.TrustScope, strings.Join(GetAppTrustScopeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AppLoginMechanismEnum Enum with underlying type: string
type AppLoginMechanismEnum string

// Set of constants representing the allowable values for AppLoginMechanismEnum
const (
	AppLoginMechanismOidc     AppLoginMechanismEnum = "OIDC"
	AppLoginMechanismSaml     AppLoginMechanismEnum = "SAML"
	AppLoginMechanismFormfill AppLoginMechanismEnum = "FORMFILL"
	AppLoginMechanismRadius   AppLoginMechanismEnum = "RADIUS"
)

var mappingAppLoginMechanismEnum = map[string]AppLoginMechanismEnum{
	"OIDC":     AppLoginMechanismOidc,
	"SAML":     AppLoginMechanismSaml,
	"FORMFILL": AppLoginMechanismFormfill,
	"RADIUS":   AppLoginMechanismRadius,
}

var mappingAppLoginMechanismEnumLowerCase = map[string]AppLoginMechanismEnum{
	"oidc":     AppLoginMechanismOidc,
	"saml":     AppLoginMechanismSaml,
	"formfill": AppLoginMechanismFormfill,
	"radius":   AppLoginMechanismRadius,
}

// GetAppLoginMechanismEnumValues Enumerates the set of values for AppLoginMechanismEnum
func GetAppLoginMechanismEnumValues() []AppLoginMechanismEnum {
	values := make([]AppLoginMechanismEnum, 0)
	for _, v := range mappingAppLoginMechanismEnum {
		values = append(values, v)
	}
	return values
}

// GetAppLoginMechanismEnumStringValues Enumerates the set of values in String for AppLoginMechanismEnum
func GetAppLoginMechanismEnumStringValues() []string {
	return []string{
		"OIDC",
		"SAML",
		"FORMFILL",
		"RADIUS",
	}
}

// GetMappingAppLoginMechanismEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAppLoginMechanismEnum(val string) (AppLoginMechanismEnum, bool) {
	enum, ok := mappingAppLoginMechanismEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AppClientTypeEnum Enum with underlying type: string
type AppClientTypeEnum string

// Set of constants representing the allowable values for AppClientTypeEnum
const (
	AppClientTypeConfidential AppClientTypeEnum = "confidential"
	AppClientTypePublic       AppClientTypeEnum = "public"
	AppClientTypeTrusted      AppClientTypeEnum = "trusted"
)

var mappingAppClientTypeEnum = map[string]AppClientTypeEnum{
	"confidential": AppClientTypeConfidential,
	"public":       AppClientTypePublic,
	"trusted":      AppClientTypeTrusted,
}

var mappingAppClientTypeEnumLowerCase = map[string]AppClientTypeEnum{
	"confidential": AppClientTypeConfidential,
	"public":       AppClientTypePublic,
	"trusted":      AppClientTypeTrusted,
}

// GetAppClientTypeEnumValues Enumerates the set of values for AppClientTypeEnum
func GetAppClientTypeEnumValues() []AppClientTypeEnum {
	values := make([]AppClientTypeEnum, 0)
	for _, v := range mappingAppClientTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAppClientTypeEnumStringValues Enumerates the set of values in String for AppClientTypeEnum
func GetAppClientTypeEnumStringValues() []string {
	return []string{
		"confidential",
		"public",
		"trusted",
	}
}

// GetMappingAppClientTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAppClientTypeEnum(val string) (AppClientTypeEnum, bool) {
	enum, ok := mappingAppClientTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AppAllowedOperationsEnum Enum with underlying type: string
type AppAllowedOperationsEnum string

// Set of constants representing the allowable values for AppAllowedOperationsEnum
const (
	AppAllowedOperationsIntrospect     AppAllowedOperationsEnum = "introspect"
	AppAllowedOperationsOnbehalfofuser AppAllowedOperationsEnum = "onBehalfOfUser"
)

var mappingAppAllowedOperationsEnum = map[string]AppAllowedOperationsEnum{
	"introspect":     AppAllowedOperationsIntrospect,
	"onBehalfOfUser": AppAllowedOperationsOnbehalfofuser,
}

var mappingAppAllowedOperationsEnumLowerCase = map[string]AppAllowedOperationsEnum{
	"introspect":     AppAllowedOperationsIntrospect,
	"onbehalfofuser": AppAllowedOperationsOnbehalfofuser,
}

// GetAppAllowedOperationsEnumValues Enumerates the set of values for AppAllowedOperationsEnum
func GetAppAllowedOperationsEnumValues() []AppAllowedOperationsEnum {
	values := make([]AppAllowedOperationsEnum, 0)
	for _, v := range mappingAppAllowedOperationsEnum {
		values = append(values, v)
	}
	return values
}

// GetAppAllowedOperationsEnumStringValues Enumerates the set of values in String for AppAllowedOperationsEnum
func GetAppAllowedOperationsEnumStringValues() []string {
	return []string{
		"introspect",
		"onBehalfOfUser",
	}
}

// GetMappingAppAllowedOperationsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAppAllowedOperationsEnum(val string) (AppAllowedOperationsEnum, bool) {
	enum, ok := mappingAppAllowedOperationsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AppClientIPCheckingEnum Enum with underlying type: string
type AppClientIPCheckingEnum string

// Set of constants representing the allowable values for AppClientIPCheckingEnum
const (
	AppClientIPCheckingAnywhere    AppClientIPCheckingEnum = "anywhere"
	AppClientIPCheckingWhitelisted AppClientIPCheckingEnum = "whitelisted"
)

var mappingAppClientIPCheckingEnum = map[string]AppClientIPCheckingEnum{
	"anywhere":    AppClientIPCheckingAnywhere,
	"whitelisted": AppClientIPCheckingWhitelisted,
}

var mappingAppClientIPCheckingEnumLowerCase = map[string]AppClientIPCheckingEnum{
	"anywhere":    AppClientIPCheckingAnywhere,
	"whitelisted": AppClientIPCheckingWhitelisted,
}

// GetAppClientIPCheckingEnumValues Enumerates the set of values for AppClientIPCheckingEnum
func GetAppClientIPCheckingEnumValues() []AppClientIPCheckingEnum {
	values := make([]AppClientIPCheckingEnum, 0)
	for _, v := range mappingAppClientIPCheckingEnum {
		values = append(values, v)
	}
	return values
}

// GetAppClientIPCheckingEnumStringValues Enumerates the set of values in String for AppClientIPCheckingEnum
func GetAppClientIPCheckingEnumStringValues() []string {
	return []string{
		"anywhere",
		"whitelisted",
	}
}

// GetMappingAppClientIPCheckingEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAppClientIPCheckingEnum(val string) (AppClientIPCheckingEnum, bool) {
	enum, ok := mappingAppClientIPCheckingEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AppTrustScopeEnum Enum with underlying type: string
type AppTrustScopeEnum string

// Set of constants representing the allowable values for AppTrustScopeEnum
const (
	AppTrustScopeExplicit AppTrustScopeEnum = "Explicit"
	AppTrustScopeAccount  AppTrustScopeEnum = "Account"
	AppTrustScopeTags     AppTrustScopeEnum = "Tags"
	AppTrustScopeDefault  AppTrustScopeEnum = "Default"
)

var mappingAppTrustScopeEnum = map[string]AppTrustScopeEnum{
	"Explicit": AppTrustScopeExplicit,
	"Account":  AppTrustScopeAccount,
	"Tags":     AppTrustScopeTags,
	"Default":  AppTrustScopeDefault,
}

var mappingAppTrustScopeEnumLowerCase = map[string]AppTrustScopeEnum{
	"explicit": AppTrustScopeExplicit,
	"account":  AppTrustScopeAccount,
	"tags":     AppTrustScopeTags,
	"default":  AppTrustScopeDefault,
}

// GetAppTrustScopeEnumValues Enumerates the set of values for AppTrustScopeEnum
func GetAppTrustScopeEnumValues() []AppTrustScopeEnum {
	values := make([]AppTrustScopeEnum, 0)
	for _, v := range mappingAppTrustScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetAppTrustScopeEnumStringValues Enumerates the set of values in String for AppTrustScopeEnum
func GetAppTrustScopeEnumStringValues() []string {
	return []string{
		"Explicit",
		"Account",
		"Tags",
		"Default",
	}
}

// GetMappingAppTrustScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAppTrustScopeEnum(val string) (AppTrustScopeEnum, bool) {
	enum, ok := mappingAppTrustScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
