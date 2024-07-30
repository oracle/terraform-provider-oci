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

// Setting Settings schema
type Setting struct {

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

	// This value indicates whether Customer Service Representatives can login and have readOnly or readWrite access.  A value of 'none' means CSR cannot login to the services.
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	CsrAccess SettingCsrAccessEnum `mandatory:"true" json:"csrAccess"`

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

	// Contact emails used to notify tenants. Can be one or more user or group alias emails.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	ContactEmails []string `mandatory:"false" json:"contactEmails"`

	// Indicates if the branding is default or custom
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	CustomBranding *bool `mandatory:"false" json:"customBranding"`

	// Preferred written or spoken language used for localized user interfaces
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCanonicalValueSourceFilter: attrName eq "languages" and attrValues.value eq "$(preferredLanguage)"
	//  - idcsCanonicalValueSourceResourceType: AllowedValue
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	PreferredLanguage *string `mandatory:"false" json:"preferredLanguage"`

	// User's timezone
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCanonicalValueSourceFilter: attrName eq "timezones" and attrValues.value eq "$(timezone)"
	//  - idcsCanonicalValueSourceResourceType: AllowedValue
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Timezone *string `mandatory:"false" json:"timezone"`

	// Controls whether DiagnosticRecords for external search-operations (against SCIM resource-types in the Admin service) identify returned resources.  If true, indicates that for each successful external search-operation at least one DiagnosticRecord will include at least one identifier for each matching resource that is returned in that search-response.  If false, no DiagnosticRecord should be expected to identify returned resources for a search-operation.  The default value is false.
	// **Added In:** 2011192329
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	DiagnosticRecordForSearchIdentifiesReturnedResources *bool `mandatory:"false" json:"diagnosticRecordForSearchIdentifiesReturnedResources"`

	// Specifies whether re-authentication is required or not when a user changes one of their security factors such as password or email. Default is true to ensure more secure behavior.
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	ReAuthWhenChangingMyAuthenticationFactors *bool `mandatory:"false" json:"reAuthWhenChangingMyAuthenticationFactors"`

	// If reAuthWhenChangingMyAuthenticationFactors is true (default), this attribute specifies which re-authentication factor to use. Allowed value is \"password\".
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	ReAuthFactor []SettingReAuthFactorEnum `mandatory:"false" json:"reAuthFactor,omitempty"`

	// Default location for purposes of localizing items such as currency, date and time format, numerical representations, and so on.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCanonicalValueSourceFilter: attrName eq "locales" and attrValues.value eq "$(locale)"
	//  - idcsCanonicalValueSourceResourceType: AllowedValue
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Locale *string `mandatory:"false" json:"locale"`

	// Indicates if access on SigningCert is allowed to public or not
	// **Added In:** 17.3.4
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	SigningCertPublicAccess *bool `mandatory:"false" json:"signingCertPublicAccess"`

	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	// - caseExact: false
	// - multiValued: false
	// - mutability: readWrite
	// - required: false
	// - returned: default
	// - type: string
	// - uniqueness: none
	// Subject mapping user profile attribute. The input format should be SCIM compliant. This attribute should be of type String and multivalued to false.
	SubMappingAttr *string `mandatory:"false" json:"subMappingAttr"`

	// Indicates whether all the Apps in this customer tenancy should trust each other. A value of true overrides the 'defaultTrustScope' attribute here in Settings, as well as any App-specific 'trustScope' attribute, to force in effect 'trustScope=Account' for every App in this customer tenancy.
	// **Added In:** 18.1.6
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	AccountAlwaysTrustScope *bool `mandatory:"false" json:"accountAlwaysTrustScope"`

	// **Deprecated Since: 18.3.6**
	// **SCIM++ Properties:**
	// - multiValued: false
	// - mutability: readWrite
	// - required: false
	// - returned: default
	// - type: string
	// Indicates the default trust scope for all apps
	DefaultTrustScope SettingDefaultTrustScopeEnum `mandatory:"false" json:"defaultTrustScope,omitempty"`

	// Tenant issuer.
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Issuer *string `mandatory:"false" json:"issuer"`

	// Previous Tenant issuer. This is an Oracle Identity Cloud Service internal attribute which is not meant to be directly modified by ID Admin. Even if the request body (Settings) contains this attribute, the actual value will be set according to the Oracle Identity Cloud Service internal logic rather than solely based on the value provided in the request payload.
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	PrevIssuer *string `mandatory:"false" json:"prevIssuer"`

	// The level of diagnostic logging that is currently in effect. A level of 0 (zero) indicates that diagnostic logging is disabled. A level of 1 (one) indicates that diagnostic logging is enabled.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	DiagnosticLevel *int `mandatory:"false" json:"diagnosticLevel"`

	// The end time up to which diagnostic recording is switched on
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: dateTime
	//  - uniqueness: none
	DiagnosticTracingUpto *string `mandatory:"false" json:"diagnosticTracingUpto"`

	// One or more email domains allowed in a user's email field. If unassigned, any domain is allowed.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	AllowedDomains []string `mandatory:"false" json:"allowedDomains"`

	// Indicates if Terms of Use is enabled in UI
	// **Added In:** 18.2.4
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	EnableTermsOfUse *bool `mandatory:"false" json:"enableTermsOfUse"`

	// Terms of Use URL
	// **Added In:** 18.2.4
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	TermsOfUseUrl *string `mandatory:"false" json:"termsOfUseUrl"`

	// Privacy Policy URL
	// **Added In:** 18.2.4
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	PrivacyPolicyUrl *string `mandatory:"false" json:"privacyPolicyUrl"`

	// Database Migration Status
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	MigrationStatus *string `mandatory:"false" json:"migrationStatus"`

	// On-Premises provisioning feature toggle.
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	OnPremisesProvisioning *bool `mandatory:"false" json:"onPremisesProvisioning"`

	// If specified, indicates the set of Urls which can be returned to after successful forgot password flow
	// **Added In:** 19.3.3
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: true
	//  - required: false
	//  - mutability: readWrite
	//  - returned: default
	//  - uniqueness: none
	//  - caseExact: false
	AllowedForgotPasswordFlowReturnUrls []string `mandatory:"false" json:"allowedForgotPasswordFlowReturnUrls"`

	// If specified, indicates the set of allowed notification redirect Urls which can be specified as the value of \"notificationRedirectUrl\" in the POST .../admin/v1/MePasswordResetRequestor request payload, which will then be included in the reset password email notification sent to a user as part of the forgot password / password reset flow.
	// **Added In:** 2009041201
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: true
	//  - required: false
	//  - mutability: readWrite
	//  - returned: default
	//  - uniqueness: none
	//  - caseExact: false
	AllowedNotificationRedirectUrls []string `mandatory:"false" json:"allowedNotificationRedirectUrls"`

	// Audit Event retention period. If set, overrides default of 30 days after which Audit Events will be purged
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	AuditEventRetentionPeriod *int `mandatory:"false" json:"auditEventRetentionPeriod"`

	// Indicates if 'hosted' option was selected
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IsHostedPage *bool `mandatory:"false" json:"isHostedPage"`

	// Storage URL location where the sanitized custom html is located
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	CustomHtmlLocation *string `mandatory:"false" json:"customHtmlLocation"`

	// Storage URL location where the sanitized custom css is located
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	CustomCssLocation *string `mandatory:"false" json:"customCssLocation"`

	// Custom translations (JSON String)
	// **Added In:** 20.1.3
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	CustomTranslation *string `mandatory:"false" json:"customTranslation"`

	// The attribute to store the cloud account name
	// **Deprecated Since: 2011192329**
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	CloudAccountName *string `mandatory:"false" json:"cloudAccountName"`

	// CloudAccountMigration: Enable Custom SIM Migrator Url.
	// **Added In:** 2012271618
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	CloudMigrationUrlEnabled *bool `mandatory:"false" json:"cloudMigrationUrlEnabled"`

	// If specified, indicates the custom SIM Migrator Url which can be used while SIM to Oracle Identity Cloud Service CloudAccount Migration.
	// **Added In:** 2012271618
	// **SCIM++ Properties:**
	//  - type: string
	//  - multiValued: false
	//  - required: false
	//  - mutability: readWrite
	//  - returned: default
	//  - uniqueness: none
	//  - caseExact: false
	CloudMigrationCustomUrl *string `mandatory:"false" json:"cloudMigrationCustomUrl"`

	// By default, a service admin can list all users in stripe. If true, a service admin cannot list other users.
	// **Added In:** 2108190438
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	ServiceAdminCannotListOtherUsers *bool `mandatory:"false" json:"serviceAdminCannotListOtherUsers"`

	// Limit the maximum return of members for an AppRole
	// **Added In:** 2111112015
	// **SCIM++ Properties:**
	//  - idcsMinValue: 0
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	MaxNoOfAppRoleMembersToReturn *int `mandatory:"false" json:"maxNoOfAppRoleMembersToReturn"`

	// Limit the maximum return of CMVA for an App
	// **Added In:** 2111112015
	// **SCIM++ Properties:**
	//  - idcsMinValue: 0
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	MaxNoOfAppCMVAToReturn *int `mandatory:"false" json:"maxNoOfAppCMVAToReturn"`

	// Maximum duration for IAM User Principal Session Token expiry
	// **Added In:** 2307071836
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	IamUpstSessionExpiry *int `mandatory:"false" json:"iamUpstSessionExpiry"`

	CloudGateCorsSettings *SettingsCloudGateCorsSettings `mandatory:"false" json:"cloudGateCorsSettings"`

	CertificateValidation *SettingsCertificateValidation `mandatory:"false" json:"certificateValidation"`

	// Custom claims associated with the specific tenant
	// **Added In:** 18.4.2
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [name]
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	//  - uniqueness: none
	TenantCustomClaims []SettingsTenantCustomClaims `mandatory:"false" json:"tenantCustomClaims"`

	// Purge Configs for different Resource Types
	// **Deprecated Since: 19.1.6**
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [resourceName]
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	PurgeConfigs []SettingsPurgeConfigs `mandatory:"false" json:"purgeConfigs"`

	// Default name of the Company in different locales
	// **Added In:** 18.2.2
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [locale]
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: complex
	DefaultCompanyNames []SettingsDefaultCompanyNames `mandatory:"false" json:"defaultCompanyNames"`

	// Default Login text in different locales
	// **Added In:** 18.2.2
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [locale]
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: complex
	DefaultLoginTexts []SettingsDefaultLoginTexts `mandatory:"false" json:"defaultLoginTexts"`

	// References to various images
	// **Added In:** 18.2.2
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [type]
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: complex
	DefaultImages []SettingsDefaultImages `mandatory:"false" json:"defaultImages"`

	// Name of the company in different locales
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [locale]
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	CompanyNames []SettingsCompanyNames `mandatory:"false" json:"companyNames"`

	// Login text in different locales
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [locale]
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	LoginTexts []SettingsLoginTexts `mandatory:"false" json:"loginTexts"`

	// References to various images
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [type]
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: complex
	Images []SettingsImages `mandatory:"false" json:"images"`
}

func (m Setting) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Setting) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSettingCsrAccessEnum(string(m.CsrAccess)); !ok && m.CsrAccess != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CsrAccess: %s. Supported values are: %s.", m.CsrAccess, strings.Join(GetSettingCsrAccessEnumStringValues(), ",")))
	}

	for _, val := range m.IdcsPreventedOperations {
		if _, ok := GetMappingIdcsPreventedOperationsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdcsPreventedOperations: %s. Supported values are: %s.", val, strings.Join(GetIdcsPreventedOperationsEnumStringValues(), ",")))
		}
	}

	for _, val := range m.ReAuthFactor {
		if _, ok := GetMappingSettingReAuthFactorEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReAuthFactor: %s. Supported values are: %s.", val, strings.Join(GetSettingReAuthFactorEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingSettingDefaultTrustScopeEnum(string(m.DefaultTrustScope)); !ok && m.DefaultTrustScope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DefaultTrustScope: %s. Supported values are: %s.", m.DefaultTrustScope, strings.Join(GetSettingDefaultTrustScopeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SettingReAuthFactorEnum Enum with underlying type: string
type SettingReAuthFactorEnum string

// Set of constants representing the allowable values for SettingReAuthFactorEnum
const (
	SettingReAuthFactorPassword SettingReAuthFactorEnum = "password"
)

var mappingSettingReAuthFactorEnum = map[string]SettingReAuthFactorEnum{
	"password": SettingReAuthFactorPassword,
}

var mappingSettingReAuthFactorEnumLowerCase = map[string]SettingReAuthFactorEnum{
	"password": SettingReAuthFactorPassword,
}

// GetSettingReAuthFactorEnumValues Enumerates the set of values for SettingReAuthFactorEnum
func GetSettingReAuthFactorEnumValues() []SettingReAuthFactorEnum {
	values := make([]SettingReAuthFactorEnum, 0)
	for _, v := range mappingSettingReAuthFactorEnum {
		values = append(values, v)
	}
	return values
}

// GetSettingReAuthFactorEnumStringValues Enumerates the set of values in String for SettingReAuthFactorEnum
func GetSettingReAuthFactorEnumStringValues() []string {
	return []string{
		"password",
	}
}

// GetMappingSettingReAuthFactorEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSettingReAuthFactorEnum(val string) (SettingReAuthFactorEnum, bool) {
	enum, ok := mappingSettingReAuthFactorEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SettingCsrAccessEnum Enum with underlying type: string
type SettingCsrAccessEnum string

// Set of constants representing the allowable values for SettingCsrAccessEnum
const (
	SettingCsrAccessReadonly  SettingCsrAccessEnum = "readOnly"
	SettingCsrAccessReadwrite SettingCsrAccessEnum = "readWrite"
	SettingCsrAccessNone      SettingCsrAccessEnum = "none"
)

var mappingSettingCsrAccessEnum = map[string]SettingCsrAccessEnum{
	"readOnly":  SettingCsrAccessReadonly,
	"readWrite": SettingCsrAccessReadwrite,
	"none":      SettingCsrAccessNone,
}

var mappingSettingCsrAccessEnumLowerCase = map[string]SettingCsrAccessEnum{
	"readonly":  SettingCsrAccessReadonly,
	"readwrite": SettingCsrAccessReadwrite,
	"none":      SettingCsrAccessNone,
}

// GetSettingCsrAccessEnumValues Enumerates the set of values for SettingCsrAccessEnum
func GetSettingCsrAccessEnumValues() []SettingCsrAccessEnum {
	values := make([]SettingCsrAccessEnum, 0)
	for _, v := range mappingSettingCsrAccessEnum {
		values = append(values, v)
	}
	return values
}

// GetSettingCsrAccessEnumStringValues Enumerates the set of values in String for SettingCsrAccessEnum
func GetSettingCsrAccessEnumStringValues() []string {
	return []string{
		"readOnly",
		"readWrite",
		"none",
	}
}

// GetMappingSettingCsrAccessEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSettingCsrAccessEnum(val string) (SettingCsrAccessEnum, bool) {
	enum, ok := mappingSettingCsrAccessEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SettingDefaultTrustScopeEnum Enum with underlying type: string
type SettingDefaultTrustScopeEnum string

// Set of constants representing the allowable values for SettingDefaultTrustScopeEnum
const (
	SettingDefaultTrustScopeExplicit SettingDefaultTrustScopeEnum = "Explicit"
	SettingDefaultTrustScopeAccount  SettingDefaultTrustScopeEnum = "Account"
	SettingDefaultTrustScopeTags     SettingDefaultTrustScopeEnum = "Tags"
)

var mappingSettingDefaultTrustScopeEnum = map[string]SettingDefaultTrustScopeEnum{
	"Explicit": SettingDefaultTrustScopeExplicit,
	"Account":  SettingDefaultTrustScopeAccount,
	"Tags":     SettingDefaultTrustScopeTags,
}

var mappingSettingDefaultTrustScopeEnumLowerCase = map[string]SettingDefaultTrustScopeEnum{
	"explicit": SettingDefaultTrustScopeExplicit,
	"account":  SettingDefaultTrustScopeAccount,
	"tags":     SettingDefaultTrustScopeTags,
}

// GetSettingDefaultTrustScopeEnumValues Enumerates the set of values for SettingDefaultTrustScopeEnum
func GetSettingDefaultTrustScopeEnumValues() []SettingDefaultTrustScopeEnum {
	values := make([]SettingDefaultTrustScopeEnum, 0)
	for _, v := range mappingSettingDefaultTrustScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetSettingDefaultTrustScopeEnumStringValues Enumerates the set of values in String for SettingDefaultTrustScopeEnum
func GetSettingDefaultTrustScopeEnumStringValues() []string {
	return []string{
		"Explicit",
		"Account",
		"Tags",
	}
}

// GetMappingSettingDefaultTrustScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSettingDefaultTrustScopeEnum(val string) (SettingDefaultTrustScopeEnum, bool) {
	enum, ok := mappingSettingDefaultTrustScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
