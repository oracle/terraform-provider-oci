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

// AppGrants Grants assigned to the app
type AppGrants struct {

	// Grant identifier
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Value *string `mandatory:"false" json:"value"`

	// Grant URI
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: reference
	//  - uniqueness: none
	Ref *string `mandatory:"false" json:"$ref"`

	// Grantee identifier
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	GranteeId *string `mandatory:"false" json:"granteeId"`

	// Grantee resource type. Allowed values are User and Group.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	GranteeType AppGrantsGranteeTypeEnum `mandatory:"false" json:"granteeType,omitempty"`

	// Each value of grantMechanism indicates how (or by what component) some App (or App-Entitlement) was granted.
	// A customer or the UI should use only grantMechanism values that start with 'ADMINISTRATOR':
	//   - 'ADMINISTRATOR_TO_USER' is for a direct grant to a specific User.
	//   - 'ADMINISTRATOR_TO_GROUP' is for a grant to a specific Group, which results in indirect grants to Users who are members of that Group.
	//   - 'ADMINISTRATOR_TO_APP' is for a grant to a specific App.  The grantee (client) App gains access to the granted (server) App.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	GrantMechanism AppGrantsGrantMechanismEnum `mandatory:"false" json:"grantMechanism,omitempty"`
}

func (m AppGrants) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AppGrants) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAppGrantsGranteeTypeEnum(string(m.GranteeType)); !ok && m.GranteeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GranteeType: %s. Supported values are: %s.", m.GranteeType, strings.Join(GetAppGrantsGranteeTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAppGrantsGrantMechanismEnum(string(m.GrantMechanism)); !ok && m.GrantMechanism != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GrantMechanism: %s. Supported values are: %s.", m.GrantMechanism, strings.Join(GetAppGrantsGrantMechanismEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AppGrantsGranteeTypeEnum Enum with underlying type: string
type AppGrantsGranteeTypeEnum string

// Set of constants representing the allowable values for AppGrantsGranteeTypeEnum
const (
	AppGrantsGranteeTypeUser  AppGrantsGranteeTypeEnum = "User"
	AppGrantsGranteeTypeGroup AppGrantsGranteeTypeEnum = "Group"
)

var mappingAppGrantsGranteeTypeEnum = map[string]AppGrantsGranteeTypeEnum{
	"User":  AppGrantsGranteeTypeUser,
	"Group": AppGrantsGranteeTypeGroup,
}

var mappingAppGrantsGranteeTypeEnumLowerCase = map[string]AppGrantsGranteeTypeEnum{
	"user":  AppGrantsGranteeTypeUser,
	"group": AppGrantsGranteeTypeGroup,
}

// GetAppGrantsGranteeTypeEnumValues Enumerates the set of values for AppGrantsGranteeTypeEnum
func GetAppGrantsGranteeTypeEnumValues() []AppGrantsGranteeTypeEnum {
	values := make([]AppGrantsGranteeTypeEnum, 0)
	for _, v := range mappingAppGrantsGranteeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAppGrantsGranteeTypeEnumStringValues Enumerates the set of values in String for AppGrantsGranteeTypeEnum
func GetAppGrantsGranteeTypeEnumStringValues() []string {
	return []string{
		"User",
		"Group",
	}
}

// GetMappingAppGrantsGranteeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAppGrantsGranteeTypeEnum(val string) (AppGrantsGranteeTypeEnum, bool) {
	enum, ok := mappingAppGrantsGranteeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AppGrantsGrantMechanismEnum Enum with underlying type: string
type AppGrantsGrantMechanismEnum string

// Set of constants representing the allowable values for AppGrantsGrantMechanismEnum
const (
	AppGrantsGrantMechanismImportApproleMembers AppGrantsGrantMechanismEnum = "IMPORT_APPROLE_MEMBERS"
	AppGrantsGrantMechanismAdministratorToUser  AppGrantsGrantMechanismEnum = "ADMINISTRATOR_TO_USER"
	AppGrantsGrantMechanismAdministratorToGroup AppGrantsGrantMechanismEnum = "ADMINISTRATOR_TO_GROUP"
	AppGrantsGrantMechanismServiceManagerToUser AppGrantsGrantMechanismEnum = "SERVICE_MANAGER_TO_USER"
	AppGrantsGrantMechanismAdministratorToApp   AppGrantsGrantMechanismEnum = "ADMINISTRATOR_TO_APP"
	AppGrantsGrantMechanismServiceManagerToApp  AppGrantsGrantMechanismEnum = "SERVICE_MANAGER_TO_APP"
	AppGrantsGrantMechanismOpcInfraToApp        AppGrantsGrantMechanismEnum = "OPC_INFRA_TO_APP"
	AppGrantsGrantMechanismGroupMembership      AppGrantsGrantMechanismEnum = "GROUP_MEMBERSHIP"
)

var mappingAppGrantsGrantMechanismEnum = map[string]AppGrantsGrantMechanismEnum{
	"IMPORT_APPROLE_MEMBERS":  AppGrantsGrantMechanismImportApproleMembers,
	"ADMINISTRATOR_TO_USER":   AppGrantsGrantMechanismAdministratorToUser,
	"ADMINISTRATOR_TO_GROUP":  AppGrantsGrantMechanismAdministratorToGroup,
	"SERVICE_MANAGER_TO_USER": AppGrantsGrantMechanismServiceManagerToUser,
	"ADMINISTRATOR_TO_APP":    AppGrantsGrantMechanismAdministratorToApp,
	"SERVICE_MANAGER_TO_APP":  AppGrantsGrantMechanismServiceManagerToApp,
	"OPC_INFRA_TO_APP":        AppGrantsGrantMechanismOpcInfraToApp,
	"GROUP_MEMBERSHIP":        AppGrantsGrantMechanismGroupMembership,
}

var mappingAppGrantsGrantMechanismEnumLowerCase = map[string]AppGrantsGrantMechanismEnum{
	"import_approle_members":  AppGrantsGrantMechanismImportApproleMembers,
	"administrator_to_user":   AppGrantsGrantMechanismAdministratorToUser,
	"administrator_to_group":  AppGrantsGrantMechanismAdministratorToGroup,
	"service_manager_to_user": AppGrantsGrantMechanismServiceManagerToUser,
	"administrator_to_app":    AppGrantsGrantMechanismAdministratorToApp,
	"service_manager_to_app":  AppGrantsGrantMechanismServiceManagerToApp,
	"opc_infra_to_app":        AppGrantsGrantMechanismOpcInfraToApp,
	"group_membership":        AppGrantsGrantMechanismGroupMembership,
}

// GetAppGrantsGrantMechanismEnumValues Enumerates the set of values for AppGrantsGrantMechanismEnum
func GetAppGrantsGrantMechanismEnumValues() []AppGrantsGrantMechanismEnum {
	values := make([]AppGrantsGrantMechanismEnum, 0)
	for _, v := range mappingAppGrantsGrantMechanismEnum {
		values = append(values, v)
	}
	return values
}

// GetAppGrantsGrantMechanismEnumStringValues Enumerates the set of values in String for AppGrantsGrantMechanismEnum
func GetAppGrantsGrantMechanismEnumStringValues() []string {
	return []string{
		"IMPORT_APPROLE_MEMBERS",
		"ADMINISTRATOR_TO_USER",
		"ADMINISTRATOR_TO_GROUP",
		"SERVICE_MANAGER_TO_USER",
		"ADMINISTRATOR_TO_APP",
		"SERVICE_MANAGER_TO_APP",
		"OPC_INFRA_TO_APP",
		"GROUP_MEMBERSHIP",
	}
}

// GetMappingAppGrantsGrantMechanismEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAppGrantsGrantMechanismEnum(val string) (AppGrantsGrantMechanismEnum, bool) {
	enum, ok := mappingAppGrantsGrantMechanismEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
