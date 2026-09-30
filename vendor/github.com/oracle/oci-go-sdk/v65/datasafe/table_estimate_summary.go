// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// TableEstimateSummary Estimated row count and size details for a table in a subsetting policy.
type TableEstimateSummary struct {

	// The OCID of the target database associated with the table estimate object.
	TargetId *string `mandatory:"true" json:"targetId"`

	// The name of the schema that contains the table.
	SchemaName *string `mandatory:"true" json:"schemaName"`

	// The name of the table.
	TableName *string `mandatory:"true" json:"tableName"`

	// The type of the database object.
	ObjectType ObjectTypeEnum `mandatory:"true" json:"objectType"`

	// The initial number of rows in the table.
	InitialRowCount *int64 `mandatory:"true" json:"initialRowCount"`

	// The estimated number of rows in the table after subsetting.
	EstimatedRowCount *int64 `mandatory:"true" json:"estimatedRowCount"`

	// The initial size of the table in KBs.
	InitialSizeInKBs *string `mandatory:"true" json:"initialSizeInKBs"`

	// The estimated size of the table in KBs after subsetting.
	EstimatedSizeInKBs *string `mandatory:"true" json:"estimatedSizeInKBs"`
}

func (m TableEstimateSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TableEstimateSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingObjectTypeEnum(string(m.ObjectType)); !ok && m.ObjectType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ObjectType: %s. Supported values are: %s.", m.ObjectType, strings.Join(GetObjectTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
