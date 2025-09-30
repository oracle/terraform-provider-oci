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

// KmsKeyVariable KMS key variable.
type KmsKeyVariable struct {

	// The display name for the variable as shown in the UI.
	Title *string `mandatory:"false" json:"title"`

	// Detailed information about this variable's purpose and usage.
	Description *string `mandatory:"false" json:"description"`

	// Indicates if this input variable is required for stack execution.
	IsRequired *bool `mandatory:"false" json:"isRequired"`

	// Hint to control whether this variable is visible.
	Visible *string `mandatory:"false" json:"visible"`

	DependsOn *KmsKeyVariableDependsOn `mandatory:"false" json:"dependsOn"`
}

// GetTitle returns Title
func (m KmsKeyVariable) GetTitle() *string {
	return m.Title
}

// GetDescription returns Description
func (m KmsKeyVariable) GetDescription() *string {
	return m.Description
}

// GetIsRequired returns IsRequired
func (m KmsKeyVariable) GetIsRequired() *bool {
	return m.IsRequired
}

// GetVisible returns Visible
func (m KmsKeyVariable) GetVisible() *string {
	return m.Visible
}

func (m KmsKeyVariable) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m KmsKeyVariable) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m KmsKeyVariable) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeKmsKeyVariable KmsKeyVariable
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeKmsKeyVariable
	}{
		"OCI_KMS_KEY_ID",
		(MarshalTypeKmsKeyVariable)(m),
	}

	return json.Marshal(&s)
}
