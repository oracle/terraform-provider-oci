// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// TargetAlertPolicyAssociationDimensions Details of aggregation dimensions used for summarizing target alert policy associations.
type TargetAlertPolicyAssociationDimensions struct {

	// Indicates whether alert policy was disabled for target due to one of rules caused the
	// generation of more than 100 alerts  per minute.
	SystemStatus TargetAlertPolicyAssociationDimensionsSystemStatusEnum `mandatory:"false" json:"systemStatus,omitempty"`

	// Indicates if the target-alert policy association is enabled or disabled by user.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`
}

func (m TargetAlertPolicyAssociationDimensions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TargetAlertPolicyAssociationDimensions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTargetAlertPolicyAssociationDimensionsSystemStatusEnum(string(m.SystemStatus)); !ok && m.SystemStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SystemStatus: %s. Supported values are: %s.", m.SystemStatus, strings.Join(GetTargetAlertPolicyAssociationDimensionsSystemStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TargetAlertPolicyAssociationDimensionsSystemStatusEnum Enum with underlying type: string
type TargetAlertPolicyAssociationDimensionsSystemStatusEnum string

// Set of constants representing the allowable values for TargetAlertPolicyAssociationDimensionsSystemStatusEnum
const (
	TargetAlertPolicyAssociationDimensionsSystemStatusEnabled  TargetAlertPolicyAssociationDimensionsSystemStatusEnum = "ENABLED"
	TargetAlertPolicyAssociationDimensionsSystemStatusDisabled TargetAlertPolicyAssociationDimensionsSystemStatusEnum = "DISABLED"
)

var mappingTargetAlertPolicyAssociationDimensionsSystemStatusEnum = map[string]TargetAlertPolicyAssociationDimensionsSystemStatusEnum{
	"ENABLED":  TargetAlertPolicyAssociationDimensionsSystemStatusEnabled,
	"DISABLED": TargetAlertPolicyAssociationDimensionsSystemStatusDisabled,
}

var mappingTargetAlertPolicyAssociationDimensionsSystemStatusEnumLowerCase = map[string]TargetAlertPolicyAssociationDimensionsSystemStatusEnum{
	"enabled":  TargetAlertPolicyAssociationDimensionsSystemStatusEnabled,
	"disabled": TargetAlertPolicyAssociationDimensionsSystemStatusDisabled,
}

// GetTargetAlertPolicyAssociationDimensionsSystemStatusEnumValues Enumerates the set of values for TargetAlertPolicyAssociationDimensionsSystemStatusEnum
func GetTargetAlertPolicyAssociationDimensionsSystemStatusEnumValues() []TargetAlertPolicyAssociationDimensionsSystemStatusEnum {
	values := make([]TargetAlertPolicyAssociationDimensionsSystemStatusEnum, 0)
	for _, v := range mappingTargetAlertPolicyAssociationDimensionsSystemStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetTargetAlertPolicyAssociationDimensionsSystemStatusEnumStringValues Enumerates the set of values in String for TargetAlertPolicyAssociationDimensionsSystemStatusEnum
func GetTargetAlertPolicyAssociationDimensionsSystemStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingTargetAlertPolicyAssociationDimensionsSystemStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTargetAlertPolicyAssociationDimensionsSystemStatusEnum(val string) (TargetAlertPolicyAssociationDimensionsSystemStatusEnum, bool) {
	enum, ok := mappingTargetAlertPolicyAssociationDimensionsSystemStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
