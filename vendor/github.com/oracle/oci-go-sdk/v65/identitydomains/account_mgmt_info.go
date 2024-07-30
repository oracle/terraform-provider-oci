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

// AccountMgmtInfo Schema for AccountMgmtInfo resource.
type AccountMgmtInfo struct {

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

	App *AccountMgmtInfoApp `mandatory:"true" json:"app"`

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

	// Unique identifier of the Account
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Uid *string `mandatory:"false" json:"uid"`

	// Name of the Account
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Name *string `mandatory:"false" json:"name"`

	// Unique key for this AccountMgmtInfo, which is used to prevent duplicate AccountMgmtInfo resources. Key is composed of a subset of app, owner and accountType.
	// **Added In:** 18.1.2
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: server
	CompositeKey *string `mandatory:"false" json:"compositeKey"`

	// If true, the account is activated
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	Active *bool `mandatory:"false" json:"active"`

	// Type of Account
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	AccountType *string `mandatory:"false" json:"accountType"`

	// If true, indicates that this managed object is an account, which is an identity that represents a user in the context of a specific application
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IsAccount *bool `mandatory:"false" json:"isAccount"`

	// If true, this account has been marked as a favorite of the User who owns it
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	Favorite *bool `mandatory:"false" json:"favorite"`

	// If true, the operation will not be performed on the target
	// **Added In:** 17.4.6
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	DoNotPerformActionOnTarget *bool `mandatory:"false" json:"doNotPerformActionOnTarget"`

	// If true, a back-fill grant will not be created for a connected managed app as part of account creation.
	// **Added In:** 18.2.6
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	DoNotBackFillGrants *bool `mandatory:"false" json:"doNotBackFillGrants"`

	// Last accessed timestamp of an application
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: dateTime
	//  - uniqueness: none
	LastAccessed *string `mandatory:"false" json:"lastAccessed"`

	// Last sync timestamp of the account
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: dateTime
	//  - uniqueness: none
	SyncTimestamp *string `mandatory:"false" json:"syncTimestamp"`

	// Last recorded sync situation for the account
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	SyncSituation AccountMgmtInfoSyncSituationEnum `mandatory:"false" json:"syncSituation,omitempty"`

	// Last recorded sync response for the account
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	SyncResponse *string `mandatory:"false" json:"syncResponse"`

	// If true, then the response to the account creation operation on a connected managed app returns a preview of the account data that is evaluated by the attribute value generation policy. Note that an account will not be created on the target application when this attribute is set to true.
	// **Added In:** 18.2.6
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	PreviewOnly *bool `mandatory:"false" json:"previewOnly"`

	// The context in which the operation is performed on the account.
	// **Added In:** 19.1.4
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	OperationContext AccountMgmtInfoOperationContextEnum `mandatory:"false" json:"operationContext,omitempty"`

	ObjectClass *AccountMgmtInfoObjectClass `mandatory:"false" json:"objectClass"`

	ResourceType *AccountMgmtInfoResourceType `mandatory:"false" json:"resourceType"`

	// Matching owning users of the account
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [value]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: complex
	//  - uniqueness: none
	MatchingOwners []AccountMgmtInfoMatchingOwners `mandatory:"false" json:"matchingOwners"`

	UserWalletArtifact *AccountMgmtInfoUserWalletArtifact `mandatory:"false" json:"userWalletArtifact"`

	Owner *AccountMgmtInfoOwner `mandatory:"false" json:"owner"`
}

func (m AccountMgmtInfo) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AccountMgmtInfo) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.IdcsPreventedOperations {
		if _, ok := GetMappingIdcsPreventedOperationsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdcsPreventedOperations: %s. Supported values are: %s.", val, strings.Join(GetIdcsPreventedOperationsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingAccountMgmtInfoSyncSituationEnum(string(m.SyncSituation)); !ok && m.SyncSituation != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SyncSituation: %s. Supported values are: %s.", m.SyncSituation, strings.Join(GetAccountMgmtInfoSyncSituationEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAccountMgmtInfoOperationContextEnum(string(m.OperationContext)); !ok && m.OperationContext != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationContext: %s. Supported values are: %s.", m.OperationContext, strings.Join(GetAccountMgmtInfoOperationContextEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AccountMgmtInfoSyncSituationEnum Enum with underlying type: string
type AccountMgmtInfoSyncSituationEnum string

// Set of constants representing the allowable values for AccountMgmtInfoSyncSituationEnum
const (
	AccountMgmtInfoSyncSituationFound     AccountMgmtInfoSyncSituationEnum = "Found"
	AccountMgmtInfoSyncSituationMatched   AccountMgmtInfoSyncSituationEnum = "Matched"
	AccountMgmtInfoSyncSituationUnmatched AccountMgmtInfoSyncSituationEnum = "Unmatched"
	AccountMgmtInfoSyncSituationLost      AccountMgmtInfoSyncSituationEnum = "Lost"
	AccountMgmtInfoSyncSituationDisputed  AccountMgmtInfoSyncSituationEnum = "Disputed"
	AccountMgmtInfoSyncSituationConfirmed AccountMgmtInfoSyncSituationEnum = "Confirmed"
)

var mappingAccountMgmtInfoSyncSituationEnum = map[string]AccountMgmtInfoSyncSituationEnum{
	"Found":     AccountMgmtInfoSyncSituationFound,
	"Matched":   AccountMgmtInfoSyncSituationMatched,
	"Unmatched": AccountMgmtInfoSyncSituationUnmatched,
	"Lost":      AccountMgmtInfoSyncSituationLost,
	"Disputed":  AccountMgmtInfoSyncSituationDisputed,
	"Confirmed": AccountMgmtInfoSyncSituationConfirmed,
}

var mappingAccountMgmtInfoSyncSituationEnumLowerCase = map[string]AccountMgmtInfoSyncSituationEnum{
	"found":     AccountMgmtInfoSyncSituationFound,
	"matched":   AccountMgmtInfoSyncSituationMatched,
	"unmatched": AccountMgmtInfoSyncSituationUnmatched,
	"lost":      AccountMgmtInfoSyncSituationLost,
	"disputed":  AccountMgmtInfoSyncSituationDisputed,
	"confirmed": AccountMgmtInfoSyncSituationConfirmed,
}

// GetAccountMgmtInfoSyncSituationEnumValues Enumerates the set of values for AccountMgmtInfoSyncSituationEnum
func GetAccountMgmtInfoSyncSituationEnumValues() []AccountMgmtInfoSyncSituationEnum {
	values := make([]AccountMgmtInfoSyncSituationEnum, 0)
	for _, v := range mappingAccountMgmtInfoSyncSituationEnum {
		values = append(values, v)
	}
	return values
}

// GetAccountMgmtInfoSyncSituationEnumStringValues Enumerates the set of values in String for AccountMgmtInfoSyncSituationEnum
func GetAccountMgmtInfoSyncSituationEnumStringValues() []string {
	return []string{
		"Found",
		"Matched",
		"Unmatched",
		"Lost",
		"Disputed",
		"Confirmed",
	}
}

// GetMappingAccountMgmtInfoSyncSituationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAccountMgmtInfoSyncSituationEnum(val string) (AccountMgmtInfoSyncSituationEnum, bool) {
	enum, ok := mappingAccountMgmtInfoSyncSituationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AccountMgmtInfoOperationContextEnum Enum with underlying type: string
type AccountMgmtInfoOperationContextEnum string

// Set of constants representing the allowable values for AccountMgmtInfoOperationContextEnum
const (
	AccountMgmtInfoOperationContextLifecycleevent AccountMgmtInfoOperationContextEnum = "LifeCycleEvent"
)

var mappingAccountMgmtInfoOperationContextEnum = map[string]AccountMgmtInfoOperationContextEnum{
	"LifeCycleEvent": AccountMgmtInfoOperationContextLifecycleevent,
}

var mappingAccountMgmtInfoOperationContextEnumLowerCase = map[string]AccountMgmtInfoOperationContextEnum{
	"lifecycleevent": AccountMgmtInfoOperationContextLifecycleevent,
}

// GetAccountMgmtInfoOperationContextEnumValues Enumerates the set of values for AccountMgmtInfoOperationContextEnum
func GetAccountMgmtInfoOperationContextEnumValues() []AccountMgmtInfoOperationContextEnum {
	values := make([]AccountMgmtInfoOperationContextEnum, 0)
	for _, v := range mappingAccountMgmtInfoOperationContextEnum {
		values = append(values, v)
	}
	return values
}

// GetAccountMgmtInfoOperationContextEnumStringValues Enumerates the set of values in String for AccountMgmtInfoOperationContextEnum
func GetAccountMgmtInfoOperationContextEnumStringValues() []string {
	return []string{
		"LifeCycleEvent",
	}
}

// GetMappingAccountMgmtInfoOperationContextEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAccountMgmtInfoOperationContextEnum(val string) (AccountMgmtInfoOperationContextEnum, bool) {
	enum, ok := mappingAccountMgmtInfoOperationContextEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
