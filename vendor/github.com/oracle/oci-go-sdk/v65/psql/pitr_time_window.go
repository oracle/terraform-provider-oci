// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// Use the OCI Database with PostgreSQL API to manage resources such as database systems, database nodes, backups, and configurations.
// For information, see the user guide documentation for the service (https://docs.oracle.com/iaas/Content/postgresql/home.htm).
//

package psql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PitrTimeWindow Time interval for which point-in-time recovery (PITR) is supported.
// The database can be restored to any timestamp between
// `timeRecoveryWindowStart` and `timeRecoveryWindowEnd` (inclusive).
type PitrTimeWindow struct {

	// Earliest timestamp in the PITR window to which the database can be
	// restored. Timestamps earlier than this are not recoverable.
	// The value must be an RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp.
	// Example: `2016-08-25T21:10:29Z`
	TimeRecoveryWindowStart *common.SDKTime `mandatory:"true" json:"timeRecoveryWindowStart"`

	// Latest timestamp in the PITR window to which the database can be
	// restored. Timestamps later than this are not recoverable.
	// The value must be an RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp.
	// Example: `2016-08-25T21:10:29Z`
	TimeRecoveryWindowEnd *common.SDKTime `mandatory:"true" json:"timeRecoveryWindowEnd"`
}

func (m PitrTimeWindow) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PitrTimeWindow) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
