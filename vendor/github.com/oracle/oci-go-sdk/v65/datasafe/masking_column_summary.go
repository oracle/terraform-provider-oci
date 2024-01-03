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

// MaskingColumnSummary Summary of a masking column.
type MaskingColumnSummary struct {

	// The unique key that identifies a masking column. The key is numeric and unique within a masking policy.
	Key *string `mandatory:"true" json:"key"`

	// The OCID of the masking policy that contains the masking column.
	MaskingPolicyId *string `mandatory:"true" json:"maskingPolicyId"`

	// The current state of the masking column.
	LifecycleState MaskingColumnLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the masking column was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the masking column was last updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The name of the schema that contains the database column.
	SchemaName *string `mandatory:"true" json:"schemaName"`

	// The name of the object (table or editioning view) that contains the database column.
	ObjectName *string `mandatory:"true" json:"objectName"`

	// The name of the database column. Note that the same name is used for the masking column.
	// There is no separate displayName attribute for the masking column.
	ColumnName *string `mandatory:"true" json:"columnName"`

	// Indicates whether data masking is enabled for the masking column.
	IsMaskingEnabled *bool `mandatory:"true" json:"isMaskingEnabled"`

	// Details about the current state of the masking column.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The type of the object that contains the database column.
	ObjectType ObjectTypeEnum `mandatory:"false" json:"objectType,omitempty"`

	// An array of child columns that are in referential relationship with the masking column.
	ChildColumns []string `mandatory:"false" json:"childColumns"`

	// The group of the masking column. All the columns in a group are masked together to ensure
	// that the masked data across these columns continue
	// to retain the same logical relationship. For more details, check <a href=https://docs.oracle.com/en/cloud/paas/data-safe/udscs/group-masking1.html#GUID-755056B9-9540-48C0-9491-262A44A85037>Group Masking in the Data Safe documentation.</a>
	MaskingColumnGroup *string `mandatory:"false" json:"maskingColumnGroup"`

	// The OCID of the sensitive type associated with the masking column.
	SensitiveTypeId *string `mandatory:"false" json:"sensitiveTypeId"`

	// The data type of the masking column.
	DataType *string `mandatory:"false" json:"dataType"`

	// An array of masking formats assigned to the masking column.
	MaskingFormats []MaskingFormat `mandatory:"false" json:"maskingFormats"`
}

func (m MaskingColumnSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MaskingColumnSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMaskingColumnLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMaskingColumnLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingObjectTypeEnum(string(m.ObjectType)); !ok && m.ObjectType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ObjectType: %s. Supported values are: %s.", m.ObjectType, strings.Join(GetObjectTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
