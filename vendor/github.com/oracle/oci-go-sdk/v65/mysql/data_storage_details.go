// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DataStorageDetails Data Storage configuration properties.
type DataStorageDetails struct {

	// Enable/disable automatic storage expansion. When set to true, the DB System will automatically
	// add storage incrementally up to the value specified in maxStorageSizeInGBs.
	IsAutoExpandStorageEnabled *bool `mandatory:"false" json:"isAutoExpandStorageEnabled"`

	// Maximum storage size this DB System can expand to. When isAutoExpandStorageEnabled
	// is set to true, the DB System will add storage incrementally up to this value.
	// DB Systems with an initial storage size of 400 GB or less can be expanded up to 32 TB.
	// DB Systems with an initial storage size between 401-800 GB can be expanded up to 64 TB.
	// DB Systems with an initial storage size between 801-1200 GB can be expanded up to 96 TB.
	// DB Systems with an initial storage size of 1201 GB or more can be expanded up to 128 TB.
	// It is not possible to decrease data storage size. You cannot set the maximum data storage size to less
	// than either current DB System dataStorageSizeInGBs or allocatedStorageSizeInGBs.
	MaxStorageSizeInGBs *int `mandatory:"false" json:"maxStorageSizeInGBs"`
}

func (m DataStorageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataStorageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
