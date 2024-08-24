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

// UpdateScheduledActionDetails Describes the modification parameters for the Scheduled Action.
type UpdateScheduledActionDetails struct {

	// Map<ParamName, ParamValue> where a key value pair describes the specific action parameter.
	// Example: `{"count": "3"}`
	ActionParams map[string]string `mandatory:"false" json:"actionParams"`

	// The list of action members in a scheduled action.
	ActionMembers []ActionMember `mandatory:"false" json:"actionMembers"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Scheduling Window.
	SchedulingWindowId *string `mandatory:"false" json:"schedulingWindowId"`
}

func (m UpdateScheduledActionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateScheduledActionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
