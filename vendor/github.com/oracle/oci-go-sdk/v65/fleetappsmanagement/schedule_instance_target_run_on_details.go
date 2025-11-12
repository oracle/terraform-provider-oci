// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ScheduleInstanceTargetRunOnDetails Schedule Instance Target Run On Details.
type ScheduleInstanceTargetRunOnDetails struct {

	// Build control flow conditions that determine the relevance of the task execution against targets.
	Condition *string `mandatory:"true" json:"condition"`
}

func (m ScheduleInstanceTargetRunOnDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScheduleInstanceTargetRunOnDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ScheduleInstanceTargetRunOnDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeScheduleInstanceTargetRunOnDetails ScheduleInstanceTargetRunOnDetails
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeScheduleInstanceTargetRunOnDetails
	}{
		"SCHEDULE_INSTANCE_TARGETS",
		(MarshalTypeScheduleInstanceTargetRunOnDetails)(m),
	}

	return json.Marshal(&s)
}
