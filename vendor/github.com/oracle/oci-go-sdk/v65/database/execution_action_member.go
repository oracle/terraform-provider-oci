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

// ExecutionActionMember The action member details.
type ExecutionActionMember struct {

	// The priority order of the execution action member.
	MemberOrder *int `mandatory:"true" json:"memberOrder"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the parent resource the execution action belongs to.
	MemberId *string `mandatory:"true" json:"memberId"`

	// The current status of the execution action member. Valid states are SCHEDULED, IN_PROGRESS, FAILED, CANCELED, DURATION_EXCEEDED, RESCHEDULED and COMPLETED.
	// enum:
	// - SCHEDULED
	// - IN_PROGRESS
	// - FAILED
	// - CANCELED
	// - DURATION_EXCEEDED
	// - RESCHEDULED
	// - SUCCEEDED
	Status *string `mandatory:"false" json:"status"`

	// The estimated time of the execution action member in minutes.
	EstimatedTimeInMins *int `mandatory:"false" json:"estimatedTimeInMins"`

	// The total time taken by corresponding resource activity in minutes.
	TotalTimeTakenInMins *int `mandatory:"false" json:"totalTimeTakenInMins"`
}

func (m ExecutionActionMember) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExecutionActionMember) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
