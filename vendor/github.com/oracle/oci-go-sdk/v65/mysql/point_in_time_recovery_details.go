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

// PointInTimeRecoveryDetails Point-in-time Recovery details like earliest and latest recovery time point for the DB System.
type PointInTimeRecoveryDetails struct {

	// Earliest recovery time point for the DB System, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeEarliestRecoveryPoint *common.SDKTime `mandatory:"true" json:"timeEarliestRecoveryPoint"`

	// Latest recovery time point for the DB System, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeLatestRecoveryPoint *common.SDKTime `mandatory:"true" json:"timeLatestRecoveryPoint"`
}

func (m PointInTimeRecoveryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PointInTimeRecoveryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
