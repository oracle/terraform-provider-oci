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

// MaskedColumnSummary Summary of a masked column. A masked column is a database column masked by a data masking request.
type MaskedColumnSummary struct {

	// The unique key that identifies the masked column. It's numeric and unique within a masking policy.
	Key *string `mandatory:"true" json:"key"`

	// The name of the schema that contains the masked column.
	SchemaName *string `mandatory:"true" json:"schemaName"`

	// The name of the object (table or editioning view) that contains the masked column.
	ObjectName *string `mandatory:"true" json:"objectName"`

	// The type of the object (table or editioning view) that contains the masked column.
	ObjectType ObjectTypeEnum `mandatory:"true" json:"objectType"`

	// The name of the masked column.
	ColumnName *string `mandatory:"true" json:"columnName"`

	// The masking format used for masking the column.
	MaskingFormatUsed *string `mandatory:"true" json:"maskingFormatUsed"`

	// The total number of values masked in the column.
	TotalMaskedValues *int64 `mandatory:"true" json:"totalMaskedValues"`

	// The unique key that identifies the parent column of the masked column.
	ParentColumnKey *string `mandatory:"false" json:"parentColumnKey"`

	// The OCID of the sensitive type associated with the masked column.
	SensitiveTypeId *string `mandatory:"false" json:"sensitiveTypeId"`

	// The masking group of the masked column.
	MaskingColumnGroup *string `mandatory:"false" json:"maskingColumnGroup"`
}

func (m MaskedColumnSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MaskedColumnSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingObjectTypeEnum(string(m.ObjectType)); !ok && m.ObjectType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ObjectType: %s. Supported values are: %s.", m.ObjectType, strings.Join(GetObjectTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
