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

// DataBaseVariable Database variable.
type DataBaseVariable struct {

	// The display name for the variable as shown in the UI.
	Title *string `mandatory:"false" json:"title"`

	// Detailed information about this variable's purpose and usage.
	Description *string `mandatory:"false" json:"description"`

	// Indicates if this input variable is required for stack execution.
	IsRequired *bool `mandatory:"false" json:"isRequired"`

	// Hint to control whether this variable is visible.
	Visible *string `mandatory:"false" json:"visible"`

	DependsOn *DataBaseVariableDependsOn `mandatory:"false" json:"dependsOn"`
}

// GetTitle returns Title
func (m DataBaseVariable) GetTitle() *string {
	return m.Title
}

// GetDescription returns Description
func (m DataBaseVariable) GetDescription() *string {
	return m.Description
}

// GetIsRequired returns IsRequired
func (m DataBaseVariable) GetIsRequired() *bool {
	return m.IsRequired
}

// GetVisible returns Visible
func (m DataBaseVariable) GetVisible() *string {
	return m.Visible
}

func (m DataBaseVariable) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataBaseVariable) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DataBaseVariable) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataBaseVariable DataBaseVariable
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDataBaseVariable
	}{
		"OCI_DATABASE_DATABASE_ID",
		(MarshalTypeDataBaseVariable)(m),
	}

	return json.Marshal(&s)
}
