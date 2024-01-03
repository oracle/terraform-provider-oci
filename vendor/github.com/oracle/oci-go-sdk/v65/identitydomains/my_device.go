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

// MyDevice Device Resource.
type MyDevice struct {

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

	User *MyDeviceUser `mandatory:"true" json:"user"`

	// Authentication Factors
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsCompositeKey: [type]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: complex
	AuthenticationFactors []MyDeviceAuthenticationFactors `mandatory:"true" json:"authenticationFactors"`

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

	// An identifier for the Resource as defined by the Service Consumer. The externalId may simplify identification of the Resource between Service Consumer and Service Provider by allowing the Consumer to refer to the Resource with its own identifier, obviating the need to store a local mapping between the local identifier of the Resource and the identifier used by the Service Provider. Each Resource MAY include a non-empty externalId value. The value of the externalId attribute is always issued be the Service Consumer and can never be specified by the Service Provider. The Service Provider MUST always interpret the externalId as scoped to the Service Consumer's tenant.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	ExternalId *string `mandatory:"false" json:"externalId"`

	// Device friendly display name
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	//  - idcsRequiresWriteForAccessFlows: true
	//  - idcsRequiresImmediateReadAfterWriteForAccessFlows: true
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Device Platform
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: immutable
	//  - idcsRequiresWriteForAccessFlows: true
	//  - idcsRequiresImmediateReadAfterWriteForAccessFlows: true
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Platform MyDevicePlatformEnum `mandatory:"false" json:"platform,omitempty"`

	// Device Status
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	//  - idcsRequiresWriteForAccessFlows: true
	//  - idcsRequiresImmediateReadAfterWriteForAccessFlows: true
	Status MyDeviceStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Additional comments/reasons for the change in device status
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Reason *string `mandatory:"false" json:"reason"`

	// Device hardware name/model
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	//  - idcsRequiresWriteForAccessFlows: true
	//  - idcsRequiresImmediateReadAfterWriteForAccessFlows: true
	DeviceType *string `mandatory:"false" json:"deviceType"`

	// Mobile Authenticator App Version
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	//  - idcsRequiresWriteForAccessFlows: true
	//  - idcsRequiresImmediateReadAfterWriteForAccessFlows: true
	AppVersion *string `mandatory:"false" json:"appVersion"`

	// Mobile Authenticator App Package Id
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	//  - idcsRequiresWriteForAccessFlows: true
	//  - idcsRequiresImmediateReadAfterWriteForAccessFlows: true
	PackageId *string `mandatory:"false" json:"packageId"`

	// Last Sync time for device
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - idcsRequiresWriteForAccessFlows: true
	//  - idcsRequiresImmediateReadAfterWriteForAccessFlows: true
	//  - required: false
	//  - returned: default
	//  - type: dateTime
	//  - uniqueness: none
	LastSyncTime *string `mandatory:"false" json:"lastSyncTime"`

	// The most recent timestamp when the device was successfully validated using one time passcode
	// **Added In:** 17.3.6
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - idcsAllowUpdatesInReadOnlyMode: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - idcsRequiresWriteForAccessFlows: true
	//  - idcsRequiresImmediateReadAfterWriteForAccessFlows: true
	//  - required: false
	//  - returned: default
	//  - type: dateTime
	//  - uniqueness: none
	LastValidatedTime *string `mandatory:"false" json:"lastValidatedTime"`

	// Device Compliance Status
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	IsCompliant *bool `mandatory:"false" json:"isCompliant"`

	// Country code of user's Phone Number
	// **Added In:** 19.1.4
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	//  - idcsRequiresWriteForAccessFlows: true
	//  - idcsRequiresImmediateReadAfterWriteForAccessFlows: true
	CountryCode *string `mandatory:"false" json:"countryCode"`

	// User's Phone Number
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	//  - idcsRequiresWriteForAccessFlows: true
	//  - idcsRequiresImmediateReadAfterWriteForAccessFlows: true
	PhoneNumber *string `mandatory:"false" json:"phoneNumber"`

	// Flag that indicates whether the device is enrolled for account recovery
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	//  - idcsRequiresWriteForAccessFlows: true
	//  - idcsRequiresImmediateReadAfterWriteForAccessFlows: true
	IsAccRecEnabled *bool `mandatory:"false" json:"isAccRecEnabled"`

	// Unique id sent from device
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	//  - idcsRequiresWriteForAccessFlows: true
	//  - idcsRequiresImmediateReadAfterWriteForAccessFlows: true
	DeviceUUID *string `mandatory:"false" json:"deviceUUID"`

	// Device base public Key
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	//  - idcsRequiresWriteForAccessFlows: true
	//  - idcsRequiresImmediateReadAfterWriteForAccessFlows: true
	BasePublicKey *string `mandatory:"false" json:"basePublicKey"`

	// Authentication method used in device. For FIDO, it will contain SECURITY_KEY/WINDOWS_HELLO etc
	// **Added In:** 2009232244
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	//  - idcsRequiresWriteForAccessFlows: true
	//  - idcsRequiresImmediateReadAfterWriteForAccessFlows: true
	AuthenticationMethod *string `mandatory:"false" json:"authenticationMethod"`

	// Attribute added for replication log, it is not used by IDCS, just added as place holder
	// **Added In:** 2111040242
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	ExpiresOn *int `mandatory:"false" json:"expiresOn"`

	// Attribute added for replication log, it is not used by IDCS, the DEK that encrypts the specific seed for that user
	// **Added In:** 2111040242
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	SeedDekId *string `mandatory:"false" json:"seedDekId"`

	// Attribute added for replication log, it is not used by IDCS, it is actual encrypted TOTP seed for the user
	// **Added In:** 2111040242
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Seed *string `mandatory:"false" json:"seed"`

	ThirdPartyFactor *MyDeviceThirdPartyFactor `mandatory:"false" json:"thirdPartyFactor"`

	PushNotificationTarget *MyDevicePushNotificationTarget `mandatory:"false" json:"pushNotificationTarget"`

	// Device additional attributes
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [key, value]
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: complex
	AdditionalAttributes []MyDeviceAdditionalAttributes `mandatory:"false" json:"additionalAttributes"`

	// Device Non Compliances
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [name, value]
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: complex
	NonCompliances []MyDeviceNonCompliances `mandatory:"false" json:"nonCompliances"`
}

func (m MyDevice) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MyDevice) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.IdcsPreventedOperations {
		if _, ok := GetMappingIdcsPreventedOperationsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdcsPreventedOperations: %s. Supported values are: %s.", val, strings.Join(GetIdcsPreventedOperationsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingMyDevicePlatformEnum(string(m.Platform)); !ok && m.Platform != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Platform: %s. Supported values are: %s.", m.Platform, strings.Join(GetMyDevicePlatformEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMyDeviceStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetMyDeviceStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MyDevicePlatformEnum Enum with underlying type: string
type MyDevicePlatformEnum string

// Set of constants representing the allowable values for MyDevicePlatformEnum
const (
	MyDevicePlatformIos      MyDevicePlatformEnum = "IOS"
	MyDevicePlatformAndroid  MyDevicePlatformEnum = "ANDROID"
	MyDevicePlatformWindows  MyDevicePlatformEnum = "WINDOWS"
	MyDevicePlatformCellular MyDevicePlatformEnum = "CELLULAR"
)

var mappingMyDevicePlatformEnum = map[string]MyDevicePlatformEnum{
	"IOS":      MyDevicePlatformIos,
	"ANDROID":  MyDevicePlatformAndroid,
	"WINDOWS":  MyDevicePlatformWindows,
	"CELLULAR": MyDevicePlatformCellular,
}

var mappingMyDevicePlatformEnumLowerCase = map[string]MyDevicePlatformEnum{
	"ios":      MyDevicePlatformIos,
	"android":  MyDevicePlatformAndroid,
	"windows":  MyDevicePlatformWindows,
	"cellular": MyDevicePlatformCellular,
}

// GetMyDevicePlatformEnumValues Enumerates the set of values for MyDevicePlatformEnum
func GetMyDevicePlatformEnumValues() []MyDevicePlatformEnum {
	values := make([]MyDevicePlatformEnum, 0)
	for _, v := range mappingMyDevicePlatformEnum {
		values = append(values, v)
	}
	return values
}

// GetMyDevicePlatformEnumStringValues Enumerates the set of values in String for MyDevicePlatformEnum
func GetMyDevicePlatformEnumStringValues() []string {
	return []string{
		"IOS",
		"ANDROID",
		"WINDOWS",
		"CELLULAR",
	}
}

// GetMappingMyDevicePlatformEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMyDevicePlatformEnum(val string) (MyDevicePlatformEnum, bool) {
	enum, ok := mappingMyDevicePlatformEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MyDeviceStatusEnum Enum with underlying type: string
type MyDeviceStatusEnum string

// Set of constants representing the allowable values for MyDeviceStatusEnum
const (
	MyDeviceStatusInitiated  MyDeviceStatusEnum = "INITIATED"
	MyDeviceStatusInprogress MyDeviceStatusEnum = "INPROGRESS"
	MyDeviceStatusInactive   MyDeviceStatusEnum = "INACTIVE"
	MyDeviceStatusEnrolled   MyDeviceStatusEnum = "ENROLLED"
	MyDeviceStatusLocked     MyDeviceStatusEnum = "LOCKED"
	MyDeviceStatusBlocked    MyDeviceStatusEnum = "BLOCKED"
)

var mappingMyDeviceStatusEnum = map[string]MyDeviceStatusEnum{
	"INITIATED":  MyDeviceStatusInitiated,
	"INPROGRESS": MyDeviceStatusInprogress,
	"INACTIVE":   MyDeviceStatusInactive,
	"ENROLLED":   MyDeviceStatusEnrolled,
	"LOCKED":     MyDeviceStatusLocked,
	"BLOCKED":    MyDeviceStatusBlocked,
}

var mappingMyDeviceStatusEnumLowerCase = map[string]MyDeviceStatusEnum{
	"initiated":  MyDeviceStatusInitiated,
	"inprogress": MyDeviceStatusInprogress,
	"inactive":   MyDeviceStatusInactive,
	"enrolled":   MyDeviceStatusEnrolled,
	"locked":     MyDeviceStatusLocked,
	"blocked":    MyDeviceStatusBlocked,
}

// GetMyDeviceStatusEnumValues Enumerates the set of values for MyDeviceStatusEnum
func GetMyDeviceStatusEnumValues() []MyDeviceStatusEnum {
	values := make([]MyDeviceStatusEnum, 0)
	for _, v := range mappingMyDeviceStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetMyDeviceStatusEnumStringValues Enumerates the set of values in String for MyDeviceStatusEnum
func GetMyDeviceStatusEnumStringValues() []string {
	return []string{
		"INITIATED",
		"INPROGRESS",
		"INACTIVE",
		"ENROLLED",
		"LOCKED",
		"BLOCKED",
	}
}

// GetMappingMyDeviceStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMyDeviceStatusEnum(val string) (MyDeviceStatusEnum, bool) {
	enum, ok := mappingMyDeviceStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
