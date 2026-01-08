// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// SessionRevoker Manage Revoked Sessions.
type SessionRevoker struct {

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

	RevokeBy *SessionRevokerRevokeBy `mandatory:"true" json:"revokeBy"`

	// Reason for the Session Termination/Revocation
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: writeOnly
	//  - required: true
	//  - returned: never
	//  - type: string
	//  - uniqueness: none
	Reason SessionRevokerReasonEnum `mandatory:"true" json:"reason"`

	// Unique identifier for the SCIM Resource as defined by the Service Provider. Each representation of the Resource MUST include a non-empty id value. This identifier MUST be unique across the Service Provider's entire set of Resources. It MUST be a stable, non-reassignable identifier that does not change when the same Resource is returned in subsequent requests. The value of the id attribute is always issued by the Service Provider and MUST never be specified by the Service Consumer. bulkId: is a reserved keyword and MUST NOT be used in the unique identifier.
	// **SCIM++ Properties:**
	//  - caseExact: true
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
	//  - idcsCsvAttributeNameMappings: [[columnHeaderName:Tag Key, mapsTo:tags.key], [columnHeaderName:Tag Value, mapsTo:tags.value]]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: request
	//  - type: complex
	//  - uniqueness: none
	Tags []Tags `mandatory:"false" json:"tags"`

	// Operations that are locked on the resource.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - idcsCsvAttributeNameMappings: [[columnHeaderName:Locked Operations, multiValueDelimiter:;]]
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	IdcsLockedOperations []SessionRevokerIdcsLockedOperationsEnum `mandatory:"false" json:"idcsLockedOperations,omitempty"`

	IdcsLockedBy *SessionRevokerIdcsLockedBy `mandatory:"false" json:"idcsLockedBy"`

	// The most recent DateTime the resource was locked.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: dateTime
	//  - uniqueness: none
	IdcsLockedOn *string `mandatory:"false" json:"idcsLockedOn"`

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

	// If true, in addition to revoking the Session, any OAuth tokens (Programmatic AT or Refresh), referencing the revoke criteria (i.e app/user/sessionId), will also be revoked.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: writeOnly
	//  - required: false
	//  - returned: never
	//  - type: boolean
	//  - uniqueness: none
	RevokeProgrammaticOAuthTokens *bool `mandatory:"false" json:"revokeProgrammaticOAuthTokens"`
}

func (m SessionRevoker) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SessionRevoker) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSessionRevokerReasonEnum(string(m.Reason)); !ok && m.Reason != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Reason: %s. Supported values are: %s.", m.Reason, strings.Join(GetSessionRevokerReasonEnumStringValues(), ",")))
	}

	for _, val := range m.IdcsPreventedOperations {
		if _, ok := GetMappingIdcsPreventedOperationsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdcsPreventedOperations: %s. Supported values are: %s.", val, strings.Join(GetIdcsPreventedOperationsEnumStringValues(), ",")))
		}
	}

	for _, val := range m.IdcsLockedOperations {
		if _, ok := GetMappingSessionRevokerIdcsLockedOperationsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdcsLockedOperations: %s. Supported values are: %s.", val, strings.Join(GetSessionRevokerIdcsLockedOperationsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SessionRevokerIdcsLockedOperationsEnum Enum with underlying type: string
type SessionRevokerIdcsLockedOperationsEnum string

// Set of constants representing the allowable values for SessionRevokerIdcsLockedOperationsEnum
const (
	SessionRevokerIdcsLockedOperationsReplace SessionRevokerIdcsLockedOperationsEnum = "replace"
	SessionRevokerIdcsLockedOperationsUpdate  SessionRevokerIdcsLockedOperationsEnum = "update"
	SessionRevokerIdcsLockedOperationsDelete  SessionRevokerIdcsLockedOperationsEnum = "delete"
)

var mappingSessionRevokerIdcsLockedOperationsEnum = map[string]SessionRevokerIdcsLockedOperationsEnum{
	"replace": SessionRevokerIdcsLockedOperationsReplace,
	"update":  SessionRevokerIdcsLockedOperationsUpdate,
	"delete":  SessionRevokerIdcsLockedOperationsDelete,
}

var mappingSessionRevokerIdcsLockedOperationsEnumLowerCase = map[string]SessionRevokerIdcsLockedOperationsEnum{
	"replace": SessionRevokerIdcsLockedOperationsReplace,
	"update":  SessionRevokerIdcsLockedOperationsUpdate,
	"delete":  SessionRevokerIdcsLockedOperationsDelete,
}

// GetSessionRevokerIdcsLockedOperationsEnumValues Enumerates the set of values for SessionRevokerIdcsLockedOperationsEnum
func GetSessionRevokerIdcsLockedOperationsEnumValues() []SessionRevokerIdcsLockedOperationsEnum {
	values := make([]SessionRevokerIdcsLockedOperationsEnum, 0)
	for _, v := range mappingSessionRevokerIdcsLockedOperationsEnum {
		values = append(values, v)
	}
	return values
}

// GetSessionRevokerIdcsLockedOperationsEnumStringValues Enumerates the set of values in String for SessionRevokerIdcsLockedOperationsEnum
func GetSessionRevokerIdcsLockedOperationsEnumStringValues() []string {
	return []string{
		"replace",
		"update",
		"delete",
	}
}

// GetMappingSessionRevokerIdcsLockedOperationsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSessionRevokerIdcsLockedOperationsEnum(val string) (SessionRevokerIdcsLockedOperationsEnum, bool) {
	enum, ok := mappingSessionRevokerIdcsLockedOperationsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SessionRevokerReasonEnum Enum with underlying type: string
type SessionRevokerReasonEnum string

// Set of constants representing the allowable values for SessionRevokerReasonEnum
const (
	SessionRevokerReasonExpired            SessionRevokerReasonEnum = "Expired"
	SessionRevokerReasonLogout             SessionRevokerReasonEnum = "Logout"
	SessionRevokerReasonRevoked            SessionRevokerReasonEnum = "Revoked"
	SessionRevokerReasonIdletimeout        SessionRevokerReasonEnum = "IdleTimeout"
	SessionRevokerReasonAppdeactivation    SessionRevokerReasonEnum = "AppDeactivation"
	SessionRevokerReasonAppdeletion        SessionRevokerReasonEnum = "AppDeletion"
	SessionRevokerReasonUserdeletion       SessionRevokerReasonEnum = "UserDeletion"
	SessionRevokerReasonUserdeactivation   SessionRevokerReasonEnum = "UserDeactivation"
	SessionRevokerReasonUserpasswordchange SessionRevokerReasonEnum = "UserPasswordChange"
	SessionRevokerReasonUserpasswordreset  SessionRevokerReasonEnum = "UserPasswordReset"
)

var mappingSessionRevokerReasonEnum = map[string]SessionRevokerReasonEnum{
	"Expired":            SessionRevokerReasonExpired,
	"Logout":             SessionRevokerReasonLogout,
	"Revoked":            SessionRevokerReasonRevoked,
	"IdleTimeout":        SessionRevokerReasonIdletimeout,
	"AppDeactivation":    SessionRevokerReasonAppdeactivation,
	"AppDeletion":        SessionRevokerReasonAppdeletion,
	"UserDeletion":       SessionRevokerReasonUserdeletion,
	"UserDeactivation":   SessionRevokerReasonUserdeactivation,
	"UserPasswordChange": SessionRevokerReasonUserpasswordchange,
	"UserPasswordReset":  SessionRevokerReasonUserpasswordreset,
}

var mappingSessionRevokerReasonEnumLowerCase = map[string]SessionRevokerReasonEnum{
	"expired":            SessionRevokerReasonExpired,
	"logout":             SessionRevokerReasonLogout,
	"revoked":            SessionRevokerReasonRevoked,
	"idletimeout":        SessionRevokerReasonIdletimeout,
	"appdeactivation":    SessionRevokerReasonAppdeactivation,
	"appdeletion":        SessionRevokerReasonAppdeletion,
	"userdeletion":       SessionRevokerReasonUserdeletion,
	"userdeactivation":   SessionRevokerReasonUserdeactivation,
	"userpasswordchange": SessionRevokerReasonUserpasswordchange,
	"userpasswordreset":  SessionRevokerReasonUserpasswordreset,
}

// GetSessionRevokerReasonEnumValues Enumerates the set of values for SessionRevokerReasonEnum
func GetSessionRevokerReasonEnumValues() []SessionRevokerReasonEnum {
	values := make([]SessionRevokerReasonEnum, 0)
	for _, v := range mappingSessionRevokerReasonEnum {
		values = append(values, v)
	}
	return values
}

// GetSessionRevokerReasonEnumStringValues Enumerates the set of values in String for SessionRevokerReasonEnum
func GetSessionRevokerReasonEnumStringValues() []string {
	return []string{
		"Expired",
		"Logout",
		"Revoked",
		"IdleTimeout",
		"AppDeactivation",
		"AppDeletion",
		"UserDeletion",
		"UserDeactivation",
		"UserPasswordChange",
		"UserPasswordReset",
	}
}

// GetMappingSessionRevokerReasonEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSessionRevokerReasonEnum(val string) (SessionRevokerReasonEnum, bool) {
	enum, ok := mappingSessionRevokerReasonEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
