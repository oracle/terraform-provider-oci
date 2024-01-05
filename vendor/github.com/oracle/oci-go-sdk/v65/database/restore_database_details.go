// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RestoreDatabaseDetails The representation of RestoreDatabaseDetails
type RestoreDatabaseDetails struct {

	// Restores using the backup with the System Change Number (SCN) specified.
	// This field is applicable for both use cases - Restoring Container Database or Restoring specific Pluggable Database.
	DatabaseSCN *string `mandatory:"false" json:"databaseSCN"`

	// Restores to the timestamp specified.
	Timestamp *common.SDKTime `mandatory:"false" json:"timestamp"`

	// Restores to the last known good state with the least possible data loss.
	Latest *bool `mandatory:"false" json:"latest"`

	// Restores only the Pluggable Database (if specified) using the inputs provided in request.
	PluggableDatabaseName *string `mandatory:"false" json:"pluggableDatabaseName"`
}

func (m RestoreDatabaseDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RestoreDatabaseDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
