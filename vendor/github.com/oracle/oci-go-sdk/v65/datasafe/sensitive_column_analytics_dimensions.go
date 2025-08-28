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

// SensitiveColumnAnalyticsDimensions The dimensions available for sensitive column analytics.
type SensitiveColumnAnalyticsDimensions struct {

	// The OCID of the target database associated with the sensitive column.
	TargetId *string `mandatory:"false" json:"targetId"`

	// The database schema that contains the sensitive column.
	SchemaName *string `mandatory:"false" json:"schemaName"`

	// The database object that contains the sensitive column.
	ObjectName *string `mandatory:"false" json:"objectName"`

	// The name of the sensitive column.
	ColumnName *string `mandatory:"false" json:"columnName"`

	// The OCID of the sensitive type associated with the sensitive column.
	SensitiveTypeId *string `mandatory:"false" json:"sensitiveTypeId"`

	// The OCID of the sensitive data model which contains the sensitive column.
	SensitiveDataModelId *string `mandatory:"false" json:"sensitiveDataModelId"`
}

func (m SensitiveColumnAnalyticsDimensions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SensitiveColumnAnalyticsDimensions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
