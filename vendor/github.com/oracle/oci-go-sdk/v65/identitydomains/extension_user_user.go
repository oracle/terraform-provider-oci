// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// ExtensionUserUser Oracle Identity Cloud Service User
type ExtensionUserUser struct {

	// A Boolean value indicating whether or not the user is federated.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCsvAttributeName: Federated
	//  - idcsCsvAttributeNameMappings: [[columnHeaderName:Federated]]
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - idcsRequiresWriteForAccessFlows: true
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IsFederatedUser *bool `mandatory:"false" json:"isFederatedUser"`

	// A Boolean value indicating whether or not authentication request by this user should be delegated to a remote app. This value should be true only when the User was originally synced from an app which is enabled for delegated authentication
	// **Added In:** 17.4.6
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: false
	//  - returned: never
	//  - type: boolean
	//  - uniqueness: none
	IsAuthenticationDelegated *bool `mandatory:"false" json:"isAuthenticationDelegated"`

	// A supplemental status indicating the reason why a user is disabled
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	Status ExtensionUserUserStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Registration provider
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Provider ExtensionUserUserProviderEnum `mandatory:"false" json:"provider,omitempty"`

	// User's preferred landing page following login, logout and reset password.
	// **Added In:** 2302092332
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	PreferredUiLandingPage ExtensionUserUserPreferredUiLandingPageEnum `mandatory:"false" json:"preferredUiLandingPage,omitempty"`

	// User creation mechanism
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCsvAttributeNameMappings: [[defaultValue:import]]
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - idcsRequiresWriteForAccessFlows: true
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	CreationMechanism ExtensionUserUserCreationMechanismEnum `mandatory:"false" json:"creationMechanism,omitempty"`

	// Specifies date time when a User's group membership was last modified.
	// **Added In:** 2304270343
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: dateTime
	//  - uniqueness: none
	GroupMembershipLastModified *string `mandatory:"false" json:"groupMembershipLastModified"`

	// A Boolean value indicating whether or not to hide the getting started page
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	DoNotShowGettingStarted *bool `mandatory:"false" json:"doNotShowGettingStarted"`

	// A Boolean value indicating whether or not to send email notification after creating the user. This attribute is not used in update/replace operations.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCsvAttributeNameMappings: [[columnHeaderName:ByPass Notification]]
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: immutable
	//  - idcsRequiresWriteForAccessFlows: true
	//  - required: false
	//  - returned: never
	//  - type: boolean
	//  - uniqueness: none
	BypassNotification *bool `mandatory:"false" json:"bypassNotification"`

	// A Boolean value indicating whether or not a user is enrolled for account recovery
	// **Added In:** 19.1.4
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: boolean
	//  - uniqueness: none
	IsAccountRecoveryEnrolled *bool `mandatory:"false" json:"isAccountRecoveryEnrolled"`

	// Boolean value to prompt user to setup account recovery during login.
	// **Added In:** 19.1.4
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: request
	//  - type: boolean
	//  - uniqueness: none
	AccountRecoveryRequired *bool `mandatory:"false" json:"accountRecoveryRequired"`

	// A Boolean value indicating whether to bypass notification and return user token to be used by an external client to control the user flow.
	// **Added In:** 18.4.2
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: false
	//  - returned: never
	//  - type: boolean
	//  - uniqueness: none
	UserFlowControlledByExternalClient *bool `mandatory:"false" json:"userFlowControlledByExternalClient"`

	// A Boolean value indicating whether or not group membership is normalized for this user.
	// **Deprecated Since: 19.3.3**
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: false
	//  - returned: never
	//  - type: boolean
	//  - uniqueness: none
	IsGroupMembershipNormalized *bool `mandatory:"false" json:"isGroupMembershipNormalized"`

	// A Boolean value Indicates whether this User's group membership has been sync'ed from Group.members to UsersGroups.
	// **Added In:** 19.3.3
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: false
	//  - returned: never
	//  - type: boolean
	//  - uniqueness: none
	IsGroupMembershipSyncedToUsersGroups *bool `mandatory:"false" json:"isGroupMembershipSyncedToUsersGroups"`

	// Specifies the EmailTemplate to be used when sending notification to the user this request is for. If specified, it overrides the default EmailTemplate for this event.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: writeOnly
	//  - required: false
	//  - returned: never
	//  - type: string
	//  - uniqueness: none
	NotificationEmailTemplateId *string `mandatory:"false" json:"notificationEmailTemplateId"`

	// Indicates if User is a Service User
	// **Added In:** 2306131901
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCsvAttributeName: Service User
	//  - idcsCsvAttributeNameMappings: [[columnHeaderName:Service User]]
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	ServiceUser *bool `mandatory:"false" json:"serviceUser"`

	// A list of Support Accounts corresponding to user.
	// **Added In:** 2103141444
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [value]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: complex
	//  - uniqueness: none
	SupportAccounts []UserExtSupportAccounts `mandatory:"false" json:"supportAccounts"`

	// Description:
	// **Added In:** 19.2.1
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [value, idcsAppRoleId]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: complex
	IdcsAppRolesLimitedToGroups []UserExtIdcsAppRolesLimitedToGroups `mandatory:"false" json:"idcsAppRolesLimitedToGroups"`

	UserToken *UserExtUserToken `mandatory:"false" json:"userToken"`

	SyncedFromApp *UserExtSyncedFromApp `mandatory:"false" json:"syncedFromApp"`

	ApplicableAuthenticationTargetApp *UserExtApplicableAuthenticationTargetApp `mandatory:"false" json:"applicableAuthenticationTargetApp"`

	DelegatedAuthenticationTargetApp *UserExtDelegatedAuthenticationTargetApp `mandatory:"false" json:"delegatedAuthenticationTargetApp"`

	// Accounts assigned to this User. Each value of this attribute refers to an app-specific identity that is owned by this User. Therefore, this attribute is a convenience that allows one to see on each User the Apps to which that User has access.
	// **SCIM++ Properties:**
	//  - idcsPii: true
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: complex
	//  - uniqueness: none
	Accounts []UserExtAccounts `mandatory:"false" json:"accounts"`

	// Grants to this User. Each value of this attribute refers to a Grant to this User of some App (and optionally of some entitlement). Therefore, this attribute is a convenience that allows one to see on each User all of the Grants to that User.
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: complex
	//  - uniqueness: none
	Grants []UserExtGrants `mandatory:"false" json:"grants"`

	// A list of all AppRoles to which this User belongs directly, indirectly or implicitly. The User could belong directly because the User is a member of the AppRole, could belong indirectly because the User is a member of a Group that is a member of the AppRole, or could belong implicitly because the AppRole is public.
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [value]
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: complex
	//  - uniqueness: none
	AppRoles []UserExtAppRoles `mandatory:"false" json:"appRoles"`
}

func (m ExtensionUserUser) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExtensionUserUser) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingExtensionUserUserStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetExtensionUserUserStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExtensionUserUserProviderEnum(string(m.Provider)); !ok && m.Provider != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Provider: %s. Supported values are: %s.", m.Provider, strings.Join(GetExtensionUserUserProviderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExtensionUserUserPreferredUiLandingPageEnum(string(m.PreferredUiLandingPage)); !ok && m.PreferredUiLandingPage != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PreferredUiLandingPage: %s. Supported values are: %s.", m.PreferredUiLandingPage, strings.Join(GetExtensionUserUserPreferredUiLandingPageEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExtensionUserUserCreationMechanismEnum(string(m.CreationMechanism)); !ok && m.CreationMechanism != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CreationMechanism: %s. Supported values are: %s.", m.CreationMechanism, strings.Join(GetExtensionUserUserCreationMechanismEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExtensionUserUserStatusEnum Enum with underlying type: string
type ExtensionUserUserStatusEnum string

// Set of constants representing the allowable values for ExtensionUserUserStatusEnum
const (
	ExtensionUserUserStatusPendingverification ExtensionUserUserStatusEnum = "pendingVerification"
	ExtensionUserUserStatusVerified            ExtensionUserUserStatusEnum = "verified"
)

var mappingExtensionUserUserStatusEnum = map[string]ExtensionUserUserStatusEnum{
	"pendingVerification": ExtensionUserUserStatusPendingverification,
	"verified":            ExtensionUserUserStatusVerified,
}

var mappingExtensionUserUserStatusEnumLowerCase = map[string]ExtensionUserUserStatusEnum{
	"pendingverification": ExtensionUserUserStatusPendingverification,
	"verified":            ExtensionUserUserStatusVerified,
}

// GetExtensionUserUserStatusEnumValues Enumerates the set of values for ExtensionUserUserStatusEnum
func GetExtensionUserUserStatusEnumValues() []ExtensionUserUserStatusEnum {
	values := make([]ExtensionUserUserStatusEnum, 0)
	for _, v := range mappingExtensionUserUserStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetExtensionUserUserStatusEnumStringValues Enumerates the set of values in String for ExtensionUserUserStatusEnum
func GetExtensionUserUserStatusEnumStringValues() []string {
	return []string{
		"pendingVerification",
		"verified",
	}
}

// GetMappingExtensionUserUserStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExtensionUserUserStatusEnum(val string) (ExtensionUserUserStatusEnum, bool) {
	enum, ok := mappingExtensionUserUserStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExtensionUserUserProviderEnum Enum with underlying type: string
type ExtensionUserUserProviderEnum string

// Set of constants representing the allowable values for ExtensionUserUserProviderEnum
const (
	ExtensionUserUserProviderFacebook ExtensionUserUserProviderEnum = "facebook"
	ExtensionUserUserProviderGoogle   ExtensionUserUserProviderEnum = "google"
	ExtensionUserUserProviderIdcs     ExtensionUserUserProviderEnum = "IDCS"
	ExtensionUserUserProviderTwitter  ExtensionUserUserProviderEnum = "twitter"
)

var mappingExtensionUserUserProviderEnum = map[string]ExtensionUserUserProviderEnum{
	"facebook": ExtensionUserUserProviderFacebook,
	"google":   ExtensionUserUserProviderGoogle,
	"IDCS":     ExtensionUserUserProviderIdcs,
	"twitter":  ExtensionUserUserProviderTwitter,
}

var mappingExtensionUserUserProviderEnumLowerCase = map[string]ExtensionUserUserProviderEnum{
	"facebook": ExtensionUserUserProviderFacebook,
	"google":   ExtensionUserUserProviderGoogle,
	"idcs":     ExtensionUserUserProviderIdcs,
	"twitter":  ExtensionUserUserProviderTwitter,
}

// GetExtensionUserUserProviderEnumValues Enumerates the set of values for ExtensionUserUserProviderEnum
func GetExtensionUserUserProviderEnumValues() []ExtensionUserUserProviderEnum {
	values := make([]ExtensionUserUserProviderEnum, 0)
	for _, v := range mappingExtensionUserUserProviderEnum {
		values = append(values, v)
	}
	return values
}

// GetExtensionUserUserProviderEnumStringValues Enumerates the set of values in String for ExtensionUserUserProviderEnum
func GetExtensionUserUserProviderEnumStringValues() []string {
	return []string{
		"facebook",
		"google",
		"IDCS",
		"twitter",
	}
}

// GetMappingExtensionUserUserProviderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExtensionUserUserProviderEnum(val string) (ExtensionUserUserProviderEnum, bool) {
	enum, ok := mappingExtensionUserUserProviderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExtensionUserUserPreferredUiLandingPageEnum Enum with underlying type: string
type ExtensionUserUserPreferredUiLandingPageEnum string

// Set of constants representing the allowable values for ExtensionUserUserPreferredUiLandingPageEnum
const (
	ExtensionUserUserPreferredUiLandingPageMyapps     ExtensionUserUserPreferredUiLandingPageEnum = "MyApps"
	ExtensionUserUserPreferredUiLandingPageMyprofile  ExtensionUserUserPreferredUiLandingPageEnum = "MyProfile"
	ExtensionUserUserPreferredUiLandingPageOciconsole ExtensionUserUserPreferredUiLandingPageEnum = "OciConsole"
)

var mappingExtensionUserUserPreferredUiLandingPageEnum = map[string]ExtensionUserUserPreferredUiLandingPageEnum{
	"MyApps":     ExtensionUserUserPreferredUiLandingPageMyapps,
	"MyProfile":  ExtensionUserUserPreferredUiLandingPageMyprofile,
	"OciConsole": ExtensionUserUserPreferredUiLandingPageOciconsole,
}

var mappingExtensionUserUserPreferredUiLandingPageEnumLowerCase = map[string]ExtensionUserUserPreferredUiLandingPageEnum{
	"myapps":     ExtensionUserUserPreferredUiLandingPageMyapps,
	"myprofile":  ExtensionUserUserPreferredUiLandingPageMyprofile,
	"ociconsole": ExtensionUserUserPreferredUiLandingPageOciconsole,
}

// GetExtensionUserUserPreferredUiLandingPageEnumValues Enumerates the set of values for ExtensionUserUserPreferredUiLandingPageEnum
func GetExtensionUserUserPreferredUiLandingPageEnumValues() []ExtensionUserUserPreferredUiLandingPageEnum {
	values := make([]ExtensionUserUserPreferredUiLandingPageEnum, 0)
	for _, v := range mappingExtensionUserUserPreferredUiLandingPageEnum {
		values = append(values, v)
	}
	return values
}

// GetExtensionUserUserPreferredUiLandingPageEnumStringValues Enumerates the set of values in String for ExtensionUserUserPreferredUiLandingPageEnum
func GetExtensionUserUserPreferredUiLandingPageEnumStringValues() []string {
	return []string{
		"MyApps",
		"MyProfile",
		"OciConsole",
	}
}

// GetMappingExtensionUserUserPreferredUiLandingPageEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExtensionUserUserPreferredUiLandingPageEnum(val string) (ExtensionUserUserPreferredUiLandingPageEnum, bool) {
	enum, ok := mappingExtensionUserUserPreferredUiLandingPageEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExtensionUserUserCreationMechanismEnum Enum with underlying type: string
type ExtensionUserUserCreationMechanismEnum string

// Set of constants representing the allowable values for ExtensionUserUserCreationMechanismEnum
const (
	ExtensionUserUserCreationMechanismBulk     ExtensionUserUserCreationMechanismEnum = "bulk"
	ExtensionUserUserCreationMechanismApi      ExtensionUserUserCreationMechanismEnum = "api"
	ExtensionUserUserCreationMechanismAdsync   ExtensionUserUserCreationMechanismEnum = "adsync"
	ExtensionUserUserCreationMechanismIdcsui   ExtensionUserUserCreationMechanismEnum = "idcsui"
	ExtensionUserUserCreationMechanismImport   ExtensionUserUserCreationMechanismEnum = "import"
	ExtensionUserUserCreationMechanismAuthsync ExtensionUserUserCreationMechanismEnum = "authsync"
	ExtensionUserUserCreationMechanismSelfreg  ExtensionUserUserCreationMechanismEnum = "selfreg"
	ExtensionUserUserCreationMechanismSamljit  ExtensionUserUserCreationMechanismEnum = "samljit"
)

var mappingExtensionUserUserCreationMechanismEnum = map[string]ExtensionUserUserCreationMechanismEnum{
	"bulk":     ExtensionUserUserCreationMechanismBulk,
	"api":      ExtensionUserUserCreationMechanismApi,
	"adsync":   ExtensionUserUserCreationMechanismAdsync,
	"idcsui":   ExtensionUserUserCreationMechanismIdcsui,
	"import":   ExtensionUserUserCreationMechanismImport,
	"authsync": ExtensionUserUserCreationMechanismAuthsync,
	"selfreg":  ExtensionUserUserCreationMechanismSelfreg,
	"samljit":  ExtensionUserUserCreationMechanismSamljit,
}

var mappingExtensionUserUserCreationMechanismEnumLowerCase = map[string]ExtensionUserUserCreationMechanismEnum{
	"bulk":     ExtensionUserUserCreationMechanismBulk,
	"api":      ExtensionUserUserCreationMechanismApi,
	"adsync":   ExtensionUserUserCreationMechanismAdsync,
	"idcsui":   ExtensionUserUserCreationMechanismIdcsui,
	"import":   ExtensionUserUserCreationMechanismImport,
	"authsync": ExtensionUserUserCreationMechanismAuthsync,
	"selfreg":  ExtensionUserUserCreationMechanismSelfreg,
	"samljit":  ExtensionUserUserCreationMechanismSamljit,
}

// GetExtensionUserUserCreationMechanismEnumValues Enumerates the set of values for ExtensionUserUserCreationMechanismEnum
func GetExtensionUserUserCreationMechanismEnumValues() []ExtensionUserUserCreationMechanismEnum {
	values := make([]ExtensionUserUserCreationMechanismEnum, 0)
	for _, v := range mappingExtensionUserUserCreationMechanismEnum {
		values = append(values, v)
	}
	return values
}

// GetExtensionUserUserCreationMechanismEnumStringValues Enumerates the set of values in String for ExtensionUserUserCreationMechanismEnum
func GetExtensionUserUserCreationMechanismEnumStringValues() []string {
	return []string{
		"bulk",
		"api",
		"adsync",
		"idcsui",
		"import",
		"authsync",
		"selfreg",
		"samljit",
	}
}

// GetMappingExtensionUserUserCreationMechanismEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExtensionUserUserCreationMechanismEnum(val string) (ExtensionUserUserCreationMechanismEnum, bool) {
	enum, ok := mappingExtensionUserUserCreationMechanismEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
