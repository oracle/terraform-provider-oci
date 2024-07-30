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

// AppGrantedAppRoles A list of AppRoles that are granted to this App (and that are defined by other Apps). Within the Oracle Public Cloud infrastructure, this allows AppID-based association. Such an association allows this App to act as a consumer and thus to access resources of another App that acts as a producer.
type AppGrantedAppRoles struct {

	// The id of an AppRole that is granted to this App.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Value *string `mandatory:"true" json:"value"`

	// The URI of an AppRole that is granted to this App.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: reference
	//  - uniqueness: none
	Ref *string `mandatory:"false" json:"$ref"`

	// A label that indicates whether this AppRole was granted directly to the App (or indirectly through a Group). For an App, the value of this attribute will always be 'direct' (because an App cannot be a member of a Group).
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Type AppGrantedAppRolesTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The display-name of an AppRole that is granted to this App.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Display *string `mandatory:"false" json:"display"`

	// The id of the App that defines this AppRole, which is granted to this App. The App that defines the AppRole acts as the producer; the App to which the AppRole is granted acts as a consumer.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	AppId *string `mandatory:"false" json:"appId"`

	// The name of the App that defines this AppRole, which is granted to this App. The App that defines the AppRole acts as the producer; the App to which the AppRole is granted acts as a consumer.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	AppName *string `mandatory:"false" json:"appName"`

	// If true, then this granted AppRole confers administrative privileges within the App that defines it. Otherwise, the granted AppRole confers only functional privileges.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	AdminRole *bool `mandatory:"false" json:"adminRole"`

	// The name of the legacy group associated with this AppRole.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	LegacyGroupName *string `mandatory:"false" json:"legacyGroupName"`

	// If true, indicates that this value must be protected.
	// **Added In:** 18.2.2
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: boolean
	//  - uniqueness: none
	ReadOnly *bool `mandatory:"false" json:"readOnly"`
}

func (m AppGrantedAppRoles) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AppGrantedAppRoles) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAppGrantedAppRolesTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetAppGrantedAppRolesTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AppGrantedAppRolesTypeEnum Enum with underlying type: string
type AppGrantedAppRolesTypeEnum string

// Set of constants representing the allowable values for AppGrantedAppRolesTypeEnum
const (
	AppGrantedAppRolesTypeDirect   AppGrantedAppRolesTypeEnum = "direct"
	AppGrantedAppRolesTypeIndirect AppGrantedAppRolesTypeEnum = "indirect"
)

var mappingAppGrantedAppRolesTypeEnum = map[string]AppGrantedAppRolesTypeEnum{
	"direct":   AppGrantedAppRolesTypeDirect,
	"indirect": AppGrantedAppRolesTypeIndirect,
}

var mappingAppGrantedAppRolesTypeEnumLowerCase = map[string]AppGrantedAppRolesTypeEnum{
	"direct":   AppGrantedAppRolesTypeDirect,
	"indirect": AppGrantedAppRolesTypeIndirect,
}

// GetAppGrantedAppRolesTypeEnumValues Enumerates the set of values for AppGrantedAppRolesTypeEnum
func GetAppGrantedAppRolesTypeEnumValues() []AppGrantedAppRolesTypeEnum {
	values := make([]AppGrantedAppRolesTypeEnum, 0)
	for _, v := range mappingAppGrantedAppRolesTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAppGrantedAppRolesTypeEnumStringValues Enumerates the set of values in String for AppGrantedAppRolesTypeEnum
func GetAppGrantedAppRolesTypeEnumStringValues() []string {
	return []string{
		"direct",
		"indirect",
	}
}

// GetMappingAppGrantedAppRolesTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAppGrantedAppRolesTypeEnum(val string) (AppGrantedAppRolesTypeEnum, bool) {
	enum, ok := mappingAppGrantedAppRolesTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
