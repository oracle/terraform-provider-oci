// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ParameterFileVersion A parameter file detatails
type ParameterFileVersion struct {

	// A unique name associated with the current migration/job and extract/replicat name
	Name *string `mandatory:"true" json:"name"`

	// Return boolean true/false for the currently in-use parameter file (factory or a versioned file)
	IsCurrent *bool `mandatory:"true" json:"isCurrent"`

	// Return true/false for whether the parameter file is oracle provided (Factory)
	IsFactory *bool `mandatory:"true" json:"isFactory"`

	// Indicator of Parameter File 'kind' (for an EXTRACT or a REPLICAT)
	Kind JobParameterFileVersionKindEnum `mandatory:"true" json:"kind"`

	// The time when this parameter file was applied on the process
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The content in base64 encoded character string containing the value of the parameter file
	Content *string `mandatory:"true" json:"content"`

	// Describes the current parameter file version
	Description *string `mandatory:"false" json:"description"`
}

func (m ParameterFileVersion) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ParameterFileVersion) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingJobParameterFileVersionKindEnum(string(m.Kind)); !ok && m.Kind != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Kind: %s. Supported values are: %s.", m.Kind, strings.Join(GetJobParameterFileVersionKindEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
