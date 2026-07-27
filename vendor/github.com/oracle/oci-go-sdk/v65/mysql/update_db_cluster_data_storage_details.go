// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateDbClusterDataStorageDetails Data storage settings for the shared-storage DB cluster.
// Updates allow increasing the storage size only; storage size cannot be decreased.
type UpdateDbClusterDataStorageDetails struct {

	// Initial storage, in GBs, allocated for the shared-storage DB cluster when it is launched. The default value is 200 GB.
	ReservedStorageSizeInGBs *int `mandatory:"false" json:"reservedStorageSizeInGBs"`

	// The maximum storage limit, in GBs, for automatic expansion of shared-storage DB cluster storage.
	// This value is always greater than or equal to the value of reservedStorageSizeInGBs. If the maximum size limit is set to the same value as that of reservedStorageSizeInGBs, the shared-storage DB cluster storage does not expand automatically.
	MaxStorageSizeInGBs *int `mandatory:"false" json:"maxStorageSizeInGBs"`
}

func (m UpdateDbClusterDataStorageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDbClusterDataStorageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
