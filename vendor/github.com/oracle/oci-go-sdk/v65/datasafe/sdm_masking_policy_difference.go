// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// SdmMaskingPolicyDifference A resource that tracks the differences between sensitive columns in the sensitive data model and masking columns in the masking policy
type SdmMaskingPolicyDifference struct {

	// The OCID of the Sensitive data model and masking policy difference resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains the Sensitive data model and masking policy difference resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The type of the SDM masking policy difference. It defines the difference scope.
	// NEW identifies new sensitive columns in the sensitive data model that are not in the masking policy.
	// DELETED identifies columns that are present in the masking policy but have been deleted from the sensitive data model.
	// MODIFIED identifies columns that are present in the sensitive data model as well as the masking policy but some of their attributes have been modified.
	// ALL covers all the above three scenarios and reports new, deleted and modified columns.
	DifferenceType SdmMaskingPolicyDifferenceDifferenceTypeEnum `mandatory:"true" json:"differenceType"`

	// The display name of the SDM masking policy difference.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The date and time the SDM masking policy difference was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the SDM masking policy difference creation started, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreationStarted *common.SDKTime `mandatory:"true" json:"timeCreationStarted"`

	// The current state of the SDM masking policy difference.
	LifecycleState SdmMaskingPolicyDifferenceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID of the sensitive data model associated with the SDM masking policy difference.
	SensitiveDataModelId *string `mandatory:"true" json:"sensitiveDataModelId"`

	// The OCID of the masking policy associated with the SDM masking policy difference.
	MaskingPolicyId *string `mandatory:"true" json:"maskingPolicyId"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m SdmMaskingPolicyDifference) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SdmMaskingPolicyDifference) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSdmMaskingPolicyDifferenceDifferenceTypeEnum(string(m.DifferenceType)); !ok && m.DifferenceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DifferenceType: %s. Supported values are: %s.", m.DifferenceType, strings.Join(GetSdmMaskingPolicyDifferenceDifferenceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSdmMaskingPolicyDifferenceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSdmMaskingPolicyDifferenceLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SdmMaskingPolicyDifferenceDifferenceTypeEnum Enum with underlying type: string
type SdmMaskingPolicyDifferenceDifferenceTypeEnum string

// Set of constants representing the allowable values for SdmMaskingPolicyDifferenceDifferenceTypeEnum
const (
	SdmMaskingPolicyDifferenceDifferenceTypeAll      SdmMaskingPolicyDifferenceDifferenceTypeEnum = "ALL"
	SdmMaskingPolicyDifferenceDifferenceTypeNew      SdmMaskingPolicyDifferenceDifferenceTypeEnum = "NEW"
	SdmMaskingPolicyDifferenceDifferenceTypeModified SdmMaskingPolicyDifferenceDifferenceTypeEnum = "MODIFIED"
	SdmMaskingPolicyDifferenceDifferenceTypeDeleted  SdmMaskingPolicyDifferenceDifferenceTypeEnum = "DELETED"
)

var mappingSdmMaskingPolicyDifferenceDifferenceTypeEnum = map[string]SdmMaskingPolicyDifferenceDifferenceTypeEnum{
	"ALL":      SdmMaskingPolicyDifferenceDifferenceTypeAll,
	"NEW":      SdmMaskingPolicyDifferenceDifferenceTypeNew,
	"MODIFIED": SdmMaskingPolicyDifferenceDifferenceTypeModified,
	"DELETED":  SdmMaskingPolicyDifferenceDifferenceTypeDeleted,
}

var mappingSdmMaskingPolicyDifferenceDifferenceTypeEnumLowerCase = map[string]SdmMaskingPolicyDifferenceDifferenceTypeEnum{
	"all":      SdmMaskingPolicyDifferenceDifferenceTypeAll,
	"new":      SdmMaskingPolicyDifferenceDifferenceTypeNew,
	"modified": SdmMaskingPolicyDifferenceDifferenceTypeModified,
	"deleted":  SdmMaskingPolicyDifferenceDifferenceTypeDeleted,
}

// GetSdmMaskingPolicyDifferenceDifferenceTypeEnumValues Enumerates the set of values for SdmMaskingPolicyDifferenceDifferenceTypeEnum
func GetSdmMaskingPolicyDifferenceDifferenceTypeEnumValues() []SdmMaskingPolicyDifferenceDifferenceTypeEnum {
	values := make([]SdmMaskingPolicyDifferenceDifferenceTypeEnum, 0)
	for _, v := range mappingSdmMaskingPolicyDifferenceDifferenceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSdmMaskingPolicyDifferenceDifferenceTypeEnumStringValues Enumerates the set of values in String for SdmMaskingPolicyDifferenceDifferenceTypeEnum
func GetSdmMaskingPolicyDifferenceDifferenceTypeEnumStringValues() []string {
	return []string{
		"ALL",
		"NEW",
		"MODIFIED",
		"DELETED",
	}
}

// GetMappingSdmMaskingPolicyDifferenceDifferenceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSdmMaskingPolicyDifferenceDifferenceTypeEnum(val string) (SdmMaskingPolicyDifferenceDifferenceTypeEnum, bool) {
	enum, ok := mappingSdmMaskingPolicyDifferenceDifferenceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SdmMaskingPolicyDifferenceLifecycleStateEnum Enum with underlying type: string
type SdmMaskingPolicyDifferenceLifecycleStateEnum string

// Set of constants representing the allowable values for SdmMaskingPolicyDifferenceLifecycleStateEnum
const (
	SdmMaskingPolicyDifferenceLifecycleStateCreating SdmMaskingPolicyDifferenceLifecycleStateEnum = "CREATING"
	SdmMaskingPolicyDifferenceLifecycleStateActive   SdmMaskingPolicyDifferenceLifecycleStateEnum = "ACTIVE"
	SdmMaskingPolicyDifferenceLifecycleStateUpdating SdmMaskingPolicyDifferenceLifecycleStateEnum = "UPDATING"
	SdmMaskingPolicyDifferenceLifecycleStateDeleting SdmMaskingPolicyDifferenceLifecycleStateEnum = "DELETING"
	SdmMaskingPolicyDifferenceLifecycleStateDeleted  SdmMaskingPolicyDifferenceLifecycleStateEnum = "DELETED"
	SdmMaskingPolicyDifferenceLifecycleStateFailed   SdmMaskingPolicyDifferenceLifecycleStateEnum = "FAILED"
)

var mappingSdmMaskingPolicyDifferenceLifecycleStateEnum = map[string]SdmMaskingPolicyDifferenceLifecycleStateEnum{
	"CREATING": SdmMaskingPolicyDifferenceLifecycleStateCreating,
	"ACTIVE":   SdmMaskingPolicyDifferenceLifecycleStateActive,
	"UPDATING": SdmMaskingPolicyDifferenceLifecycleStateUpdating,
	"DELETING": SdmMaskingPolicyDifferenceLifecycleStateDeleting,
	"DELETED":  SdmMaskingPolicyDifferenceLifecycleStateDeleted,
	"FAILED":   SdmMaskingPolicyDifferenceLifecycleStateFailed,
}

var mappingSdmMaskingPolicyDifferenceLifecycleStateEnumLowerCase = map[string]SdmMaskingPolicyDifferenceLifecycleStateEnum{
	"creating": SdmMaskingPolicyDifferenceLifecycleStateCreating,
	"active":   SdmMaskingPolicyDifferenceLifecycleStateActive,
	"updating": SdmMaskingPolicyDifferenceLifecycleStateUpdating,
	"deleting": SdmMaskingPolicyDifferenceLifecycleStateDeleting,
	"deleted":  SdmMaskingPolicyDifferenceLifecycleStateDeleted,
	"failed":   SdmMaskingPolicyDifferenceLifecycleStateFailed,
}

// GetSdmMaskingPolicyDifferenceLifecycleStateEnumValues Enumerates the set of values for SdmMaskingPolicyDifferenceLifecycleStateEnum
func GetSdmMaskingPolicyDifferenceLifecycleStateEnumValues() []SdmMaskingPolicyDifferenceLifecycleStateEnum {
	values := make([]SdmMaskingPolicyDifferenceLifecycleStateEnum, 0)
	for _, v := range mappingSdmMaskingPolicyDifferenceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSdmMaskingPolicyDifferenceLifecycleStateEnumStringValues Enumerates the set of values in String for SdmMaskingPolicyDifferenceLifecycleStateEnum
func GetSdmMaskingPolicyDifferenceLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingSdmMaskingPolicyDifferenceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSdmMaskingPolicyDifferenceLifecycleStateEnum(val string) (SdmMaskingPolicyDifferenceLifecycleStateEnum, bool) {
	enum, ok := mappingSdmMaskingPolicyDifferenceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
