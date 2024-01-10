// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateMaintenanceWindowDetails Defines the maintenance window for create operation, when automatic actions can be performed.
type CreateMaintenanceWindowDetails struct {

	// Days of the week.
	Day DayEnum `mandatory:"true" json:"day"`

	// Start hour for maintenance period. Hour is in UTC.
	StartHour *int `mandatory:"true" json:"startHour"`
}

func (m CreateMaintenanceWindowDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateMaintenanceWindowDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDayEnum(string(m.Day)); !ok && m.Day != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Day: %s. Supported values are: %s.", m.Day, strings.Join(GetDayEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
