// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// ValidateBackupDetails Details required to validate backup.
type ValidateBackupDetails struct {

	// Specifies whether the backup needs to be prepared for fast restore or not.
	// Set to true to prepare the backup, set to false (default) if not required.
	// Note: The prepared backup will replace the original backup and will not generate a new backup copy.
	// The cost associated with the backup may vary, as the prepared backup will consistently be a full backup,
	// it may also change the storage size of the original backup.
	IsPreparedBackupRequired *bool `mandatory:"true" json:"isPreparedBackupRequired"`
}

func (m ValidateBackupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ValidateBackupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
