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

// CreateMaskingColumnDetails Details to create a new masking column. You can use this to add a parent column only.
// It automatically adds the child columns from the associated sensitive data model or
// target database. If you provide the sensitiveTypeId attribute but not the maskingFormats
// attribute, it automatically assigns the default masking format associated with the
// specified sensitive type. Alternatively, if you provide the maskingFormats attribute,
// the specified masking formats are assigned to the column.
// Using the maskingFormats attribute, you can assign one or more masking formats to a
// column. You can specify a condition as part of each masking format. It enables you
// to do <a href="https://docs.oracle.com/en/cloud/paas/data-safe/udscs/conditional-masking.html">conditional masking</a>
// so that you can mask the column data values differently using different masking formats
// and the associated conditions. A masking format can have one or more format entries.
// The combined output of all the format entries is used for masking. It provides the
// flexibility to define a masking format that can generate different parts of a data
// value separately and then combine them to get the final data value for masking.
// You can use the maskingColumnGroup attribute to group the columns that you would like
// to mask together. It enables you to do
// <a href="https://docs.oracle.com/en/cloud/paas/data-safe/udscs/group-masking1.html#GUID-755056B9-9540-48C0-9491-262A44A85037">group or compound masking</a>
// that ensures that the masked data across the columns in a group continue to retain
// the same logical relationship.
type CreateMaskingColumnDetails struct {

	// The name of the schema that contains the database column. This attribute cannot be updated for an existing masking column.
	SchemaName *string `mandatory:"true" json:"schemaName"`

	// The name of the object (table or editioning view) that contains the database column. This attribute cannot be updated for an existing masking column.
	ObjectName *string `mandatory:"true" json:"objectName"`

	// The name of the database column. This attribute cannot be updated for an existing
	//  masking column. Note that the same name is used for the masking column. There
	//  is no separate displayName attribute for the masking column.
	ColumnName *string `mandatory:"true" json:"columnName"`

	// The type of the object that contains the database column.
	ObjectType ObjectTypeEnum `mandatory:"false" json:"objectType,omitempty"`

	// The group of the masking column. It's a masking group identifier and can be any string
	// of acceptable length. All the columns in a group are masked together to ensure that
	// the masked data across these columns continue to retain the same logical relationship.
	// For more details, check
	// <a href=https://docs.oracle.com/en/cloud/paas/data-safe/udscs/group-masking1.html#GUID-755056B9-9540-48C0-9491-262A44A85037>Group Masking in the Data Safe documentation.</a>
	MaskingColumnGroup *string `mandatory:"false" json:"maskingColumnGroup"`

	// The OCID of the sensitive type to be associated with the masking column. Note that
	// if the maskingFormats attribute isn't provided while creating a masking column,
	// the default masking format associated with the specified sensitive type is assigned
	// to the masking column.
	SensitiveTypeId *string `mandatory:"false" json:"sensitiveTypeId"`

	// Indicates whether data masking is enabled for the masking column. Set it to false if
	// you don't want to mask the column.
	IsMaskingEnabled *bool `mandatory:"false" json:"isMaskingEnabled"`

	// The masking formats to be assigned to the masking column. You can specify a condition
	// as part of each masking format. It enables you to do
	// <a href="https://docs.oracle.com/en/cloud/paas/data-safe/udscs/conditional-masking.html">conditional masking</a>
	// so that you can mask the column data values differently using different masking
	// formats and the associated conditions. A masking format can have one or more format
	// entries. The combined output of all the format entries is used for masking. It
	// provides the flexibility to define a masking format that can generate different parts
	// of a data value separately and then combine them to get the final data value for masking.
	MaskingFormats []MaskingFormat `mandatory:"false" json:"maskingFormats"`
}

func (m CreateMaskingColumnDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateMaskingColumnDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingObjectTypeEnum(string(m.ObjectType)); !ok && m.ObjectType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ObjectType: %s. Supported values are: %s.", m.ObjectType, strings.Join(GetObjectTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
