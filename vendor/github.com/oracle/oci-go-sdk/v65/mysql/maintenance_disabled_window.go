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

// MaintenanceDisabledWindow Time window during which downtime-inducing maintenance shall not be performed.
// Downtime-free maintenance may be performed to apply required security patches.
type MaintenanceDisabledWindow struct {

	// The time from when maintenance is disabled.
	// Must be set together with timeEnd and must be before timeEnd.
	// as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeStart *common.SDKTime `mandatory:"true" json:"timeStart"`

	// The time until when maintenance is disabled.
	// Must be set together with timeStart and must be after timeStart.
	// as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeEnd *common.SDKTime `mandatory:"true" json:"timeEnd"`
}

func (m MaintenanceDisabledWindow) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MaintenanceDisabledWindow) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
