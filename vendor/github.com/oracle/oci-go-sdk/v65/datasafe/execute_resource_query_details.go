// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExecuteResourceQueryDetails The details of the query that will be executed.
type ExecuteResourceQueryDetails struct {

	// The OCID of the compartment on which the query will be executed.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The view on which the query will be executed.
	ResourceQueryView ExecuteResourceQueryDetailsResourceQueryViewEnum `mandatory:"true" json:"resourceQueryView"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" json:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ExecuteResourceQueryDetailsAccessLevelEnum `mandatory:"false" json:"accessLevel,omitempty"`

	// The scimQuery query parameter accepts filter expressions that use the syntax described in Section 3.2.2.2
	// of the System for Cross-Domain Identity Management (SCIM) specification, which is available
	// at RFC3339 (https://tools.ietf.org/html/draft-ietf-scim-api-12). In SCIM filtering expressions,
	// text, date, and time values must be enclosed in quotation marks, with date and time values using ISO-8601 format.
	// (Numeric and boolean values should not be quoted.)
	ScimQuery *string `mandatory:"false" json:"scimQuery"`

	// The list of fields on which the group by has to be done.
	GroupBy []string `mandatory:"false" json:"groupBy"`
}

func (m ExecuteResourceQueryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExecuteResourceQueryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExecuteResourceQueryDetailsResourceQueryViewEnum(string(m.ResourceQueryView)); !ok && m.ResourceQueryView != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceQueryView: %s. Supported values are: %s.", m.ResourceQueryView, strings.Join(GetExecuteResourceQueryDetailsResourceQueryViewEnumStringValues(), ",")))
	}

	if _, ok := GetMappingExecuteResourceQueryDetailsAccessLevelEnum(string(m.AccessLevel)); !ok && m.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", m.AccessLevel, strings.Join(GetExecuteResourceQueryDetailsAccessLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExecuteResourceQueryDetailsAccessLevelEnum Enum with underlying type: string
type ExecuteResourceQueryDetailsAccessLevelEnum string

// Set of constants representing the allowable values for ExecuteResourceQueryDetailsAccessLevelEnum
const (
	ExecuteResourceQueryDetailsAccessLevelRestricted ExecuteResourceQueryDetailsAccessLevelEnum = "RESTRICTED"
	ExecuteResourceQueryDetailsAccessLevelAccessible ExecuteResourceQueryDetailsAccessLevelEnum = "ACCESSIBLE"
)

var mappingExecuteResourceQueryDetailsAccessLevelEnum = map[string]ExecuteResourceQueryDetailsAccessLevelEnum{
	"RESTRICTED": ExecuteResourceQueryDetailsAccessLevelRestricted,
	"ACCESSIBLE": ExecuteResourceQueryDetailsAccessLevelAccessible,
}

var mappingExecuteResourceQueryDetailsAccessLevelEnumLowerCase = map[string]ExecuteResourceQueryDetailsAccessLevelEnum{
	"restricted": ExecuteResourceQueryDetailsAccessLevelRestricted,
	"accessible": ExecuteResourceQueryDetailsAccessLevelAccessible,
}

// GetExecuteResourceQueryDetailsAccessLevelEnumValues Enumerates the set of values for ExecuteResourceQueryDetailsAccessLevelEnum
func GetExecuteResourceQueryDetailsAccessLevelEnumValues() []ExecuteResourceQueryDetailsAccessLevelEnum {
	values := make([]ExecuteResourceQueryDetailsAccessLevelEnum, 0)
	for _, v := range mappingExecuteResourceQueryDetailsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetExecuteResourceQueryDetailsAccessLevelEnumStringValues Enumerates the set of values in String for ExecuteResourceQueryDetailsAccessLevelEnum
func GetExecuteResourceQueryDetailsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingExecuteResourceQueryDetailsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExecuteResourceQueryDetailsAccessLevelEnum(val string) (ExecuteResourceQueryDetailsAccessLevelEnum, bool) {
	enum, ok := mappingExecuteResourceQueryDetailsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExecuteResourceQueryDetailsResourceQueryViewEnum Enum with underlying type: string
type ExecuteResourceQueryDetailsResourceQueryViewEnum string

// Set of constants representing the allowable values for ExecuteResourceQueryDetailsResourceQueryViewEnum
const (
	ExecuteResourceQueryDetailsResourceQueryViewAuditPolicy              ExecuteResourceQueryDetailsResourceQueryViewEnum = "AUDIT_POLICY"
	ExecuteResourceQueryDetailsResourceQueryViewSecurityPolicyDeployment ExecuteResourceQueryDetailsResourceQueryViewEnum = "SECURITY_POLICY_DEPLOYMENT"
	ExecuteResourceQueryDetailsResourceQueryViewAttributeSet             ExecuteResourceQueryDetailsResourceQueryViewEnum = "ATTRIBUTE_SET"
	ExecuteResourceQueryDetailsResourceQueryViewSecurityPolicyConfig     ExecuteResourceQueryDetailsResourceQueryViewEnum = "SECURITY_POLICY_CONFIG"
)

var mappingExecuteResourceQueryDetailsResourceQueryViewEnum = map[string]ExecuteResourceQueryDetailsResourceQueryViewEnum{
	"AUDIT_POLICY":               ExecuteResourceQueryDetailsResourceQueryViewAuditPolicy,
	"SECURITY_POLICY_DEPLOYMENT": ExecuteResourceQueryDetailsResourceQueryViewSecurityPolicyDeployment,
	"ATTRIBUTE_SET":              ExecuteResourceQueryDetailsResourceQueryViewAttributeSet,
	"SECURITY_POLICY_CONFIG":     ExecuteResourceQueryDetailsResourceQueryViewSecurityPolicyConfig,
}

var mappingExecuteResourceQueryDetailsResourceQueryViewEnumLowerCase = map[string]ExecuteResourceQueryDetailsResourceQueryViewEnum{
	"audit_policy":               ExecuteResourceQueryDetailsResourceQueryViewAuditPolicy,
	"security_policy_deployment": ExecuteResourceQueryDetailsResourceQueryViewSecurityPolicyDeployment,
	"attribute_set":              ExecuteResourceQueryDetailsResourceQueryViewAttributeSet,
	"security_policy_config":     ExecuteResourceQueryDetailsResourceQueryViewSecurityPolicyConfig,
}

// GetExecuteResourceQueryDetailsResourceQueryViewEnumValues Enumerates the set of values for ExecuteResourceQueryDetailsResourceQueryViewEnum
func GetExecuteResourceQueryDetailsResourceQueryViewEnumValues() []ExecuteResourceQueryDetailsResourceQueryViewEnum {
	values := make([]ExecuteResourceQueryDetailsResourceQueryViewEnum, 0)
	for _, v := range mappingExecuteResourceQueryDetailsResourceQueryViewEnum {
		values = append(values, v)
	}
	return values
}

// GetExecuteResourceQueryDetailsResourceQueryViewEnumStringValues Enumerates the set of values in String for ExecuteResourceQueryDetailsResourceQueryViewEnum
func GetExecuteResourceQueryDetailsResourceQueryViewEnumStringValues() []string {
	return []string{
		"AUDIT_POLICY",
		"SECURITY_POLICY_DEPLOYMENT",
		"ATTRIBUTE_SET",
		"SECURITY_POLICY_CONFIG",
	}
}

// GetMappingExecuteResourceQueryDetailsResourceQueryViewEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExecuteResourceQueryDetailsResourceQueryViewEnum(val string) (ExecuteResourceQueryDetailsResourceQueryViewEnum, bool) {
	enum, ok := mappingExecuteResourceQueryDetailsResourceQueryViewEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
