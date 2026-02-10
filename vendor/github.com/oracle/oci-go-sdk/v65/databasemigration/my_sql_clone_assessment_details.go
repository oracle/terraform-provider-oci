// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// MySqlCloneAssessmentDetails MySQL Clone Assessment Summary
type MySqlCloneAssessmentDetails struct {
	SourceDatabaseConnection *SourceAssessmentConnection `mandatory:"true" json:"sourceDatabaseConnection"`

	TargetDatabaseConnection *TargetAssessmentConnection `mandatory:"true" json:"targetDatabaseConnection"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID of the resource being referenced.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags. Example: {"Department": "Finance"}
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

// GetDisplayName returns DisplayName
func (m MySqlCloneAssessmentDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m MySqlCloneAssessmentDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetSourceDatabaseConnection returns SourceDatabaseConnection
func (m MySqlCloneAssessmentDetails) GetSourceDatabaseConnection() *SourceAssessmentConnection {
	return m.SourceDatabaseConnection
}

// GetTargetDatabaseConnection returns TargetDatabaseConnection
func (m MySqlCloneAssessmentDetails) GetTargetDatabaseConnection() *TargetAssessmentConnection {
	return m.TargetDatabaseConnection
}

// GetFreeformTags returns FreeformTags
func (m MySqlCloneAssessmentDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m MySqlCloneAssessmentDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m MySqlCloneAssessmentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MySqlCloneAssessmentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m MySqlCloneAssessmentDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMySqlCloneAssessmentDetails MySqlCloneAssessmentDetails
	s := struct {
		DiscriminatorParam string `json:"databaseCombination"`
		MarshalTypeMySqlCloneAssessmentDetails
	}{
		"MYSQL",
		(MarshalTypeMySqlCloneAssessmentDetails)(m),
	}

	return json.Marshal(&s)
}
