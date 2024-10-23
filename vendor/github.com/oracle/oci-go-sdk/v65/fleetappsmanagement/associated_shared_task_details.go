// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// AssociatedSharedTaskDetails The details of the shared task.
// Tasks that are part of the task library and can be reused across runbooks.
type AssociatedSharedTaskDetails struct {

	// The ID of taskRecord.
	TaskRecordId *string `mandatory:"true" json:"taskRecordId"`
}

func (m AssociatedSharedTaskDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AssociatedSharedTaskDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AssociatedSharedTaskDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAssociatedSharedTaskDetails AssociatedSharedTaskDetails
	s := struct {
		DiscriminatorParam string `json:"scope"`
		MarshalTypeAssociatedSharedTaskDetails
	}{
		"SHARED",
		(MarshalTypeAssociatedSharedTaskDetails)(m),
	}

	return json.Marshal(&s)
}
