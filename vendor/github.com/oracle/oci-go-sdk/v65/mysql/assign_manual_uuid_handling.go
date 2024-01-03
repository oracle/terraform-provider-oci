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

// AssignManualUuidHandling Enables assignment of IDs on the target to anonymous transactions coming from the source. A manually defined UUID
// is added as a prefix to the ID.
type AssignManualUuidHandling struct {

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

	// The UUID that is used as a prefix when generating transaction identifiers for anonymous transactions
	// coming from the source. You can change the UUID later.
	Uuid *string `mandatory:"false" json:"uuid"`
}

func (m AssignManualUuidHandling) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AssignManualUuidHandling) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AssignManualUuidHandling) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAssignManualUuidHandling AssignManualUuidHandling
	s := struct {
		DiscriminatorParam string `json:"policy"`
		MarshalTypeAssignManualUuidHandling
	}{
		"ASSIGN_MANUAL_UUID",
		(MarshalTypeAssignManualUuidHandling)(m),
	}

	return json.Marshal(&s)
}
