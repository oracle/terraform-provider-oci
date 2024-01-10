// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateDumpTransferDetails Optional additional properties for dump transfer.
type UpdateDumpTransferDetails struct {
	Source UpdateHostDumpTransferDetails `mandatory:"false" json:"source"`

	Target UpdateHostDumpTransferDetails `mandatory:"false" json:"target"`

	// OCID of the shared storage mount target
	SharedStorageMountTargetId *string `mandatory:"false" json:"sharedStorageMountTargetId"`
}

func (m UpdateDumpTransferDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDumpTransferDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateDumpTransferDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Source                     updatehostdumptransferdetails `json:"source"`
		Target                     updatehostdumptransferdetails `json:"target"`
		SharedStorageMountTargetId *string                       `json:"sharedStorageMountTargetId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.Source.UnmarshalPolymorphicJSON(model.Source.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Source = nn.(UpdateHostDumpTransferDetails)
	} else {
		m.Source = nil
	}

	nn, e = model.Target.UnmarshalPolymorphicJSON(model.Target.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Target = nn.(UpdateHostDumpTransferDetails)
	} else {
		m.Target = nil
	}

	m.SharedStorageMountTargetId = model.SharedStorageMountTargetId

	return
}
