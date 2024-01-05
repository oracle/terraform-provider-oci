// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AssignTargetUuidHandling Enables assignment of IDs on the target to anonymous transactions coming from the source. The target server UUID
// is added as a prefix to the ID.
type AssignTargetUuidHandling struct {

	// Specifies one of the coordinates (file) at which the replica should begin
	// reading the source's log. As this value specifies the point where replication
	// starts from, it is only used once, when it starts. It is never used again,
	// unless a new UpdateChannel operation modifies it.
	LastConfiguredLogFilename *string `mandatory:"false" json:"lastConfiguredLogFilename"`

	// Specifies one of the coordinates (offset) at which the replica should begin
	// reading the source's log. As this value specifies the point where replication
	// starts from, it is only used once, when it starts. It is never used again,
	// unless a new UpdateChannel operation modifies it.
	LastConfiguredLogOffset *int64 `mandatory:"false" json:"lastConfiguredLogOffset"`
}

func (m AssignTargetUuidHandling) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AssignTargetUuidHandling) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AssignTargetUuidHandling) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAssignTargetUuidHandling AssignTargetUuidHandling
	s := struct {
		DiscriminatorParam string `json:"policy"`
		MarshalTypeAssignTargetUuidHandling
	}{
		"ASSIGN_TARGET_UUID",
		(MarshalTypeAssignTargetUuidHandling)(m),
	}

	return json.Marshal(&s)
}
