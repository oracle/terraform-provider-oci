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

// CloneAssessmentDetails Customizable details when performing cloning of an Assessment.
type CloneAssessmentDetails interface {
	GetSourceDatabaseConnection() *SourceAssessmentConnection

	GetTargetDatabaseConnection() *TargetAssessmentConnection

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	GetDisplayName() *string

	// The OCID of the resource being referenced.
	GetCompartmentId() *string

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags. Example: {"Department": "Finance"}
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type cloneassessmentdetails struct {
	JsonData                 []byte
	DisplayName              *string                           `mandatory:"false" json:"displayName"`
	CompartmentId            *string                           `mandatory:"false" json:"compartmentId"`
	FreeformTags             map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags              map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SourceDatabaseConnection *SourceAssessmentConnection       `mandatory:"true" json:"sourceDatabaseConnection"`
	TargetDatabaseConnection *TargetAssessmentConnection       `mandatory:"true" json:"targetDatabaseConnection"`
	DatabaseCombination      string                            `json:"databaseCombination"`
}

// UnmarshalJSON unmarshals json
func (m *cloneassessmentdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercloneassessmentdetails cloneassessmentdetails
	s := struct {
		Model Unmarshalercloneassessmentdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SourceDatabaseConnection = s.Model.SourceDatabaseConnection
	m.TargetDatabaseConnection = s.Model.TargetDatabaseConnection
	m.DisplayName = s.Model.DisplayName
	m.CompartmentId = s.Model.CompartmentId
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.DatabaseCombination = s.Model.DatabaseCombination

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *cloneassessmentdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DatabaseCombination {
	case "MYSQL":
		mm := MySqlCloneAssessmentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE":
		mm := OracleCloneAssessmentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CloneAssessmentDetails: %s.", m.DatabaseCombination)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m cloneassessmentdetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m cloneassessmentdetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m cloneassessmentdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m cloneassessmentdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSourceDatabaseConnection returns SourceDatabaseConnection
func (m cloneassessmentdetails) GetSourceDatabaseConnection() *SourceAssessmentConnection {
	return m.SourceDatabaseConnection
}

// GetTargetDatabaseConnection returns TargetDatabaseConnection
func (m cloneassessmentdetails) GetTargetDatabaseConnection() *TargetAssessmentConnection {
	return m.TargetDatabaseConnection
}

func (m cloneassessmentdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m cloneassessmentdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
