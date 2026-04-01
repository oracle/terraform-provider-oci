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

// ControlledUpdate The details of the last successfully processed controlled update action.
type ControlledUpdate struct {

	// The MySQL version to be applied to the selected instances.
	TargetMysqlVersion *string `mandatory:"false" json:"targetMysqlVersion"`

	// Defines the MySQL instances to be operated during a controlled update.
	//  - ALL_BUT_PRIMARY: Update all MySQL instances in a highly available DB System except the primary group member,
	//    without triggering a controlled failover.
	//  - PRIMARY_ONLY: Update the primary group member in a highly available DB System
	//    after a controlled failover (downtime is expected). This operation requires that the other
	//    MySQL instances have been previously updated using the ALL_BUT_PRIMARY option.
	TargetDbInstances ControlledUpdateTargetDbInstancesEnum `mandatory:"false" json:"targetDbInstances,omitempty"`
}

func (m ControlledUpdate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ControlledUpdate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingControlledUpdateTargetDbInstancesEnum(string(m.TargetDbInstances)); !ok && m.TargetDbInstances != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetDbInstances: %s. Supported values are: %s.", m.TargetDbInstances, strings.Join(GetControlledUpdateTargetDbInstancesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
