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

// ExtensionGroupGroup Oracle Identity Cloud Service Group
type ExtensionGroupGroup struct {

	// Group description
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCsvAttributeName: Description
	//  - idcsCsvAttributeNameMappings: [[columnHeaderName:Description]]
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Description *string `mandatory:"false" json:"description"`

	// Source from which this group got created.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCsvAttributeNameMappings: [[defaultValue:import]]
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	CreationMechanism ExtensionGroupGroupCreationMechanismEnum `mandatory:"false" json:"creationMechanism,omitempty"`

	PasswordPolicy *GroupExtPasswordPolicy `mandatory:"false" json:"passwordPolicy"`

	SyncedFromApp *GroupExtSyncedFromApp `mandatory:"false" json:"syncedFromApp"`

	// Grants assigned to group
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: complex
	//  - uniqueness: none
	Grants []GroupExtGrants `mandatory:"false" json:"grants"`

	// Group owners
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCompositeKey: [value, type]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: request
	//  - type: complex
	//  - uniqueness: none
	Owners []GroupExtOwners `mandatory:"false" json:"owners"`

	// A list of appRoles that the user belongs to, either thorough direct membership, nested groups, or dynamically calculated
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCompositeKey: [value]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: complex
	//  - uniqueness: none
	AppRoles []GroupExtAppRoles `mandatory:"false" json:"appRoles"`
}

func (m ExtensionGroupGroup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExtensionGroupGroup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingExtensionGroupGroupCreationMechanismEnum(string(m.CreationMechanism)); !ok && m.CreationMechanism != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CreationMechanism: %s. Supported values are: %s.", m.CreationMechanism, strings.Join(GetExtensionGroupGroupCreationMechanismEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExtensionGroupGroupCreationMechanismEnum Enum with underlying type: string
type ExtensionGroupGroupCreationMechanismEnum string

// Set of constants representing the allowable values for ExtensionGroupGroupCreationMechanismEnum
const (
	ExtensionGroupGroupCreationMechanismBulk     ExtensionGroupGroupCreationMechanismEnum = "bulk"
	ExtensionGroupGroupCreationMechanismApi      ExtensionGroupGroupCreationMechanismEnum = "api"
	ExtensionGroupGroupCreationMechanismAdsync   ExtensionGroupGroupCreationMechanismEnum = "adsync"
	ExtensionGroupGroupCreationMechanismAuthsync ExtensionGroupGroupCreationMechanismEnum = "authsync"
	ExtensionGroupGroupCreationMechanismIdcsui   ExtensionGroupGroupCreationMechanismEnum = "idcsui"
	ExtensionGroupGroupCreationMechanismImport   ExtensionGroupGroupCreationMechanismEnum = "import"
)

var mappingExtensionGroupGroupCreationMechanismEnum = map[string]ExtensionGroupGroupCreationMechanismEnum{
	"bulk":     ExtensionGroupGroupCreationMechanismBulk,
	"api":      ExtensionGroupGroupCreationMechanismApi,
	"adsync":   ExtensionGroupGroupCreationMechanismAdsync,
	"authsync": ExtensionGroupGroupCreationMechanismAuthsync,
	"idcsui":   ExtensionGroupGroupCreationMechanismIdcsui,
	"import":   ExtensionGroupGroupCreationMechanismImport,
}

var mappingExtensionGroupGroupCreationMechanismEnumLowerCase = map[string]ExtensionGroupGroupCreationMechanismEnum{
	"bulk":     ExtensionGroupGroupCreationMechanismBulk,
	"api":      ExtensionGroupGroupCreationMechanismApi,
	"adsync":   ExtensionGroupGroupCreationMechanismAdsync,
	"authsync": ExtensionGroupGroupCreationMechanismAuthsync,
	"idcsui":   ExtensionGroupGroupCreationMechanismIdcsui,
	"import":   ExtensionGroupGroupCreationMechanismImport,
}

// GetExtensionGroupGroupCreationMechanismEnumValues Enumerates the set of values for ExtensionGroupGroupCreationMechanismEnum
func GetExtensionGroupGroupCreationMechanismEnumValues() []ExtensionGroupGroupCreationMechanismEnum {
	values := make([]ExtensionGroupGroupCreationMechanismEnum, 0)
	for _, v := range mappingExtensionGroupGroupCreationMechanismEnum {
		values = append(values, v)
	}
	return values
}

// GetExtensionGroupGroupCreationMechanismEnumStringValues Enumerates the set of values in String for ExtensionGroupGroupCreationMechanismEnum
func GetExtensionGroupGroupCreationMechanismEnumStringValues() []string {
	return []string{
		"bulk",
		"api",
		"adsync",
		"authsync",
		"idcsui",
		"import",
	}
}

// GetMappingExtensionGroupGroupCreationMechanismEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExtensionGroupGroupCreationMechanismEnum(val string) (ExtensionGroupGroupCreationMechanismEnum, bool) {
	enum, ok := mappingExtensionGroupGroupCreationMechanismEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
