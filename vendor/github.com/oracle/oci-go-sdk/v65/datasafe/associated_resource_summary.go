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

// AssociatedResourceSummary Summary details of the associated resource of an attribute set.
type AssociatedResourceSummary struct {

	// The resource type that is associated with the attribute set.
	AssociatedResourceType AssociatedResourceSummaryAssociatedResourceTypeEnum `mandatory:"false" json:"associatedResourceType,omitempty"`

	// The OCID of the resource that is associated with the attribute set.
	AssociatedResourceId *string `mandatory:"false" json:"associatedResourceId"`

	// The display name of the resource that is associated with the attribute set. The name does not have to be unique, and is changeable.
	AssociatedResourceName *string `mandatory:"false" json:"associatedResourceName"`

	// The date and time when associated started between resource and the attribute set, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time when associated is removed between resources and the attribute set, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m AssociatedResourceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AssociatedResourceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAssociatedResourceSummaryAssociatedResourceTypeEnum(string(m.AssociatedResourceType)); !ok && m.AssociatedResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AssociatedResourceType: %s. Supported values are: %s.", m.AssociatedResourceType, strings.Join(GetAssociatedResourceSummaryAssociatedResourceTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AssociatedResourceSummaryAssociatedResourceTypeEnum Enum with underlying type: string
type AssociatedResourceSummaryAssociatedResourceTypeEnum string

// Set of constants representing the allowable values for AssociatedResourceSummaryAssociatedResourceTypeEnum
const (
	AssociatedResourceSummaryAssociatedResourceTypeAuditPolicy AssociatedResourceSummaryAssociatedResourceTypeEnum = "AUDIT_POLICY"
)

var mappingAssociatedResourceSummaryAssociatedResourceTypeEnum = map[string]AssociatedResourceSummaryAssociatedResourceTypeEnum{
	"AUDIT_POLICY": AssociatedResourceSummaryAssociatedResourceTypeAuditPolicy,
}

var mappingAssociatedResourceSummaryAssociatedResourceTypeEnumLowerCase = map[string]AssociatedResourceSummaryAssociatedResourceTypeEnum{
	"audit_policy": AssociatedResourceSummaryAssociatedResourceTypeAuditPolicy,
}

// GetAssociatedResourceSummaryAssociatedResourceTypeEnumValues Enumerates the set of values for AssociatedResourceSummaryAssociatedResourceTypeEnum
func GetAssociatedResourceSummaryAssociatedResourceTypeEnumValues() []AssociatedResourceSummaryAssociatedResourceTypeEnum {
	values := make([]AssociatedResourceSummaryAssociatedResourceTypeEnum, 0)
	for _, v := range mappingAssociatedResourceSummaryAssociatedResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAssociatedResourceSummaryAssociatedResourceTypeEnumStringValues Enumerates the set of values in String for AssociatedResourceSummaryAssociatedResourceTypeEnum
func GetAssociatedResourceSummaryAssociatedResourceTypeEnumStringValues() []string {
	return []string{
		"AUDIT_POLICY",
	}
}

// GetMappingAssociatedResourceSummaryAssociatedResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAssociatedResourceSummaryAssociatedResourceTypeEnum(val string) (AssociatedResourceSummaryAssociatedResourceTypeEnum, bool) {
	enum, ok := mappingAssociatedResourceSummaryAssociatedResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
