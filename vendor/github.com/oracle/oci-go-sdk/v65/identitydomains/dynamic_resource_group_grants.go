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

// DynamicResourceGroupGrants Grants assigned to group
type DynamicResourceGroupGrants struct {

	// Grant identifier
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsAddedSinceVersion: 3
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
	//  - idcsAddedSinceVersion: 3
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
	//  - idcsAddedSinceVersion: 3
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
	//  - idcsAddedSinceVersion: 3
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	GrantMechanism DynamicResourceGroupGrantsGrantMechanismEnum `mandatory:"false" json:"grantMechanism,omitempty"`
}

func (m DynamicResourceGroupGrants) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DynamicResourceGroupGrants) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDynamicResourceGroupGrantsGrantMechanismEnum(string(m.GrantMechanism)); !ok && m.GrantMechanism != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GrantMechanism: %s. Supported values are: %s.", m.GrantMechanism, strings.Join(GetDynamicResourceGroupGrantsGrantMechanismEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DynamicResourceGroupGrantsGrantMechanismEnum Enum with underlying type: string
type DynamicResourceGroupGrantsGrantMechanismEnum string

// Set of constants representing the allowable values for DynamicResourceGroupGrantsGrantMechanismEnum
const (
	DynamicResourceGroupGrantsGrantMechanismImportApproleMembers                DynamicResourceGroupGrantsGrantMechanismEnum = "IMPORT_APPROLE_MEMBERS"
	DynamicResourceGroupGrantsGrantMechanismAdministratorToDynamicResourceGroup DynamicResourceGroupGrantsGrantMechanismEnum = "ADMINISTRATOR_TO_DYNAMIC_RESOURCE_GROUP"
	DynamicResourceGroupGrantsGrantMechanismAdministratorToUser                 DynamicResourceGroupGrantsGrantMechanismEnum = "ADMINISTRATOR_TO_USER"
	DynamicResourceGroupGrantsGrantMechanismAdministratorToGroup                DynamicResourceGroupGrantsGrantMechanismEnum = "ADMINISTRATOR_TO_GROUP"
	DynamicResourceGroupGrantsGrantMechanismServiceManagerToUser                DynamicResourceGroupGrantsGrantMechanismEnum = "SERVICE_MANAGER_TO_USER"
	DynamicResourceGroupGrantsGrantMechanismAdministratorToApp                  DynamicResourceGroupGrantsGrantMechanismEnum = "ADMINISTRATOR_TO_APP"
	DynamicResourceGroupGrantsGrantMechanismServiceManagerToApp                 DynamicResourceGroupGrantsGrantMechanismEnum = "SERVICE_MANAGER_TO_APP"
	DynamicResourceGroupGrantsGrantMechanismOpcInfraToApp                       DynamicResourceGroupGrantsGrantMechanismEnum = "OPC_INFRA_TO_APP"
	DynamicResourceGroupGrantsGrantMechanismGroupMembership                     DynamicResourceGroupGrantsGrantMechanismEnum = "GROUP_MEMBERSHIP"
)

var mappingDynamicResourceGroupGrantsGrantMechanismEnum = map[string]DynamicResourceGroupGrantsGrantMechanismEnum{
	"IMPORT_APPROLE_MEMBERS":                  DynamicResourceGroupGrantsGrantMechanismImportApproleMembers,
	"ADMINISTRATOR_TO_DYNAMIC_RESOURCE_GROUP": DynamicResourceGroupGrantsGrantMechanismAdministratorToDynamicResourceGroup,
	"ADMINISTRATOR_TO_USER":                   DynamicResourceGroupGrantsGrantMechanismAdministratorToUser,
	"ADMINISTRATOR_TO_GROUP":                  DynamicResourceGroupGrantsGrantMechanismAdministratorToGroup,
	"SERVICE_MANAGER_TO_USER":                 DynamicResourceGroupGrantsGrantMechanismServiceManagerToUser,
	"ADMINISTRATOR_TO_APP":                    DynamicResourceGroupGrantsGrantMechanismAdministratorToApp,
	"SERVICE_MANAGER_TO_APP":                  DynamicResourceGroupGrantsGrantMechanismServiceManagerToApp,
	"OPC_INFRA_TO_APP":                        DynamicResourceGroupGrantsGrantMechanismOpcInfraToApp,
	"GROUP_MEMBERSHIP":                        DynamicResourceGroupGrantsGrantMechanismGroupMembership,
}

var mappingDynamicResourceGroupGrantsGrantMechanismEnumLowerCase = map[string]DynamicResourceGroupGrantsGrantMechanismEnum{
	"import_approle_members":                  DynamicResourceGroupGrantsGrantMechanismImportApproleMembers,
	"administrator_to_dynamic_resource_group": DynamicResourceGroupGrantsGrantMechanismAdministratorToDynamicResourceGroup,
	"administrator_to_user":                   DynamicResourceGroupGrantsGrantMechanismAdministratorToUser,
	"administrator_to_group":                  DynamicResourceGroupGrantsGrantMechanismAdministratorToGroup,
	"service_manager_to_user":                 DynamicResourceGroupGrantsGrantMechanismServiceManagerToUser,
	"administrator_to_app":                    DynamicResourceGroupGrantsGrantMechanismAdministratorToApp,
	"service_manager_to_app":                  DynamicResourceGroupGrantsGrantMechanismServiceManagerToApp,
	"opc_infra_to_app":                        DynamicResourceGroupGrantsGrantMechanismOpcInfraToApp,
	"group_membership":                        DynamicResourceGroupGrantsGrantMechanismGroupMembership,
}

// GetDynamicResourceGroupGrantsGrantMechanismEnumValues Enumerates the set of values for DynamicResourceGroupGrantsGrantMechanismEnum
func GetDynamicResourceGroupGrantsGrantMechanismEnumValues() []DynamicResourceGroupGrantsGrantMechanismEnum {
	values := make([]DynamicResourceGroupGrantsGrantMechanismEnum, 0)
	for _, v := range mappingDynamicResourceGroupGrantsGrantMechanismEnum {
		values = append(values, v)
	}
	return values
}

// GetDynamicResourceGroupGrantsGrantMechanismEnumStringValues Enumerates the set of values in String for DynamicResourceGroupGrantsGrantMechanismEnum
func GetDynamicResourceGroupGrantsGrantMechanismEnumStringValues() []string {
	return []string{
		"IMPORT_APPROLE_MEMBERS",
		"ADMINISTRATOR_TO_DYNAMIC_RESOURCE_GROUP",
		"ADMINISTRATOR_TO_USER",
		"ADMINISTRATOR_TO_GROUP",
		"SERVICE_MANAGER_TO_USER",
		"ADMINISTRATOR_TO_APP",
		"SERVICE_MANAGER_TO_APP",
		"OPC_INFRA_TO_APP",
		"GROUP_MEMBERSHIP",
	}
}

// GetMappingDynamicResourceGroupGrantsGrantMechanismEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDynamicResourceGroupGrantsGrantMechanismEnum(val string) (DynamicResourceGroupGrantsGrantMechanismEnum, bool) {
	enum, ok := mappingDynamicResourceGroupGrantsGrantMechanismEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
