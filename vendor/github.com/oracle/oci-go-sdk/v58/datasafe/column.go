// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// Column Description of a column
type Column struct {

	// Name of the column displayed on UI.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Specifies the corresponding field name in the data source.
	FieldName *string `mandatory:"true" json:"fieldName"`

	// Indicates if the column is hidden. Values can either be 'true' or 'false'.
	IsHidden *bool `mandatory:"true" json:"isHidden"`

	// Specifies the display order of the column.
	DisplayOrder *int `mandatory:"true" json:"displayOrder"`

	// Specifies the data type of the column.
	DataType *string `mandatory:"false" json:"dataType"`
}

func (m Column) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Column) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
