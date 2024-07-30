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

// GroupExtGrants Grants assigned to group
// **SCIM++ Properties:**
//   - idcsSearchable: true
//   - multiValued: true
//   - mutability: readOnly
//   - required: false
//   - returned: request
//   - type: complex
//   - uniqueness: none
type GroupExtGrants struct {

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

	// App identifier
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	AppId *string `mandatory:"false" json:"appId"`

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
	GrantMechanism GroupExtGrantsGrantMechanismEnum `mandatory:"false" json:"grantMechanism,omitempty"`
}

func (m GroupExtGrants) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GroupExtGrants) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingGroupExtGrantsGrantMechanismEnum(string(m.GrantMechanism)); !ok && m.GrantMechanism != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GrantMechanism: %s. Supported values are: %s.", m.GrantMechanism, strings.Join(GetGroupExtGrantsGrantMechanismEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GroupExtGrantsGrantMechanismEnum Enum with underlying type: string
type GroupExtGrantsGrantMechanismEnum string

// Set of constants representing the allowable values for GroupExtGrantsGrantMechanismEnum
const (
	GroupExtGrantsGrantMechanismImportApproleMembers GroupExtGrantsGrantMechanismEnum = "IMPORT_APPROLE_MEMBERS"
	GroupExtGrantsGrantMechanismAdministratorToUser  GroupExtGrantsGrantMechanismEnum = "ADMINISTRATOR_TO_USER"
	GroupExtGrantsGrantMechanismAdministratorToGroup GroupExtGrantsGrantMechanismEnum = "ADMINISTRATOR_TO_GROUP"
	GroupExtGrantsGrantMechanismServiceManagerToUser GroupExtGrantsGrantMechanismEnum = "SERVICE_MANAGER_TO_USER"
	GroupExtGrantsGrantMechanismAdministratorToApp   GroupExtGrantsGrantMechanismEnum = "ADMINISTRATOR_TO_APP"
	GroupExtGrantsGrantMechanismServiceManagerToApp  GroupExtGrantsGrantMechanismEnum = "SERVICE_MANAGER_TO_APP"
	GroupExtGrantsGrantMechanismOpcInfraToApp        GroupExtGrantsGrantMechanismEnum = "OPC_INFRA_TO_APP"
	GroupExtGrantsGrantMechanismGroupMembership      GroupExtGrantsGrantMechanismEnum = "GROUP_MEMBERSHIP"
)

var mappingGroupExtGrantsGrantMechanismEnum = map[string]GroupExtGrantsGrantMechanismEnum{
	"IMPORT_APPROLE_MEMBERS":  GroupExtGrantsGrantMechanismImportApproleMembers,
	"ADMINISTRATOR_TO_USER":   GroupExtGrantsGrantMechanismAdministratorToUser,
	"ADMINISTRATOR_TO_GROUP":  GroupExtGrantsGrantMechanismAdministratorToGroup,
	"SERVICE_MANAGER_TO_USER": GroupExtGrantsGrantMechanismServiceManagerToUser,
	"ADMINISTRATOR_TO_APP":    GroupExtGrantsGrantMechanismAdministratorToApp,
	"SERVICE_MANAGER_TO_APP":  GroupExtGrantsGrantMechanismServiceManagerToApp,
	"OPC_INFRA_TO_APP":        GroupExtGrantsGrantMechanismOpcInfraToApp,
	"GROUP_MEMBERSHIP":        GroupExtGrantsGrantMechanismGroupMembership,
}

var mappingGroupExtGrantsGrantMechanismEnumLowerCase = map[string]GroupExtGrantsGrantMechanismEnum{
	"import_approle_members":  GroupExtGrantsGrantMechanismImportApproleMembers,
	"administrator_to_user":   GroupExtGrantsGrantMechanismAdministratorToUser,
	"administrator_to_group":  GroupExtGrantsGrantMechanismAdministratorToGroup,
	"service_manager_to_user": GroupExtGrantsGrantMechanismServiceManagerToUser,
	"administrator_to_app":    GroupExtGrantsGrantMechanismAdministratorToApp,
	"service_manager_to_app":  GroupExtGrantsGrantMechanismServiceManagerToApp,
	"opc_infra_to_app":        GroupExtGrantsGrantMechanismOpcInfraToApp,
	"group_membership":        GroupExtGrantsGrantMechanismGroupMembership,
}

// GetGroupExtGrantsGrantMechanismEnumValues Enumerates the set of values for GroupExtGrantsGrantMechanismEnum
func GetGroupExtGrantsGrantMechanismEnumValues() []GroupExtGrantsGrantMechanismEnum {
	values := make([]GroupExtGrantsGrantMechanismEnum, 0)
	for _, v := range mappingGroupExtGrantsGrantMechanismEnum {
		values = append(values, v)
	}
	return values
}

// GetGroupExtGrantsGrantMechanismEnumStringValues Enumerates the set of values in String for GroupExtGrantsGrantMechanismEnum
func GetGroupExtGrantsGrantMechanismEnumStringValues() []string {
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

// GetMappingGroupExtGrantsGrantMechanismEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGroupExtGrantsGrantMechanismEnum(val string) (GroupExtGrantsGrantMechanismEnum, bool) {
	enum, ok := mappingGroupExtGrantsGrantMechanismEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
