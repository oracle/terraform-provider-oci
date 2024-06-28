// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MigrationParameterDetails Migration parameter details object.
type MigrationParameterDetails struct {

	// Parameter name.
	Name *string `mandatory:"true" json:"name"`

	// Parameter data type.
	DataType AdvancedParameterDataTypesEnum `mandatory:"true" json:"dataType"`

	// If a STRING data type then the value should be an array of characters,
	// if a INTEGER data type then the value should be an integer value,
	// if a FLOAT data type then the value should be an float value,
	// if a BOOLEAN data type then the value should be TRUE or FALSE.
	Value *string `mandatory:"true" json:"value"`
}

func (m MigrationParameterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MigrationParameterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAdvancedParameterDataTypesEnum(string(m.DataType)); !ok && m.DataType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataType: %s. Supported values are: %s.", m.DataType, strings.Join(GetAdvancedParameterDataTypesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
