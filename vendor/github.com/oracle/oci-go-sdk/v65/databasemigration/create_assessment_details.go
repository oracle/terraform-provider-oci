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

// CreateAssessmentDetails Common Create Assessment details.
type CreateAssessmentDetails interface {

	// The OCID of the resource being referenced.
	GetCompartmentId() *string

	// A network speed in Megabits per second.
	GetNetworkSpeedMegabitPerSecond() NetworkSpeedMegabitPerSecondEnum

	// Time allowed for the application downtime.
	GetAcceptableDowntime() AcceptableDowntimeEnum

	// The size of a source database.
	GetDatabaseDataSize() DatabaseDataSizeEnum

	// DDL expectation values.
	GetDdlExpectation() DdlExpectationEnum

	GetSourceDatabaseConnection() *SourceAssessmentConnection

	GetTargetDatabaseConnection() *TargetAssessmentConnection

	// A user-friendly description. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	GetDescription() *string

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	GetDisplayName() *string

	// The type of assessment creation.
	GetCreationType() CreationTypeEnum

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags. Example: {"Department": "Finance"}
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type createassessmentdetails struct {
	JsonData                     []byte
	Description                  *string                           `mandatory:"false" json:"description"`
	DisplayName                  *string                           `mandatory:"false" json:"displayName"`
	CreationType                 CreationTypeEnum                  `mandatory:"false" json:"creationType,omitempty"`
	FreeformTags                 map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags                  map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	CompartmentId                *string                           `mandatory:"true" json:"compartmentId"`
	NetworkSpeedMegabitPerSecond NetworkSpeedMegabitPerSecondEnum  `mandatory:"true" json:"networkSpeedMegabitPerSecond"`
	AcceptableDowntime           AcceptableDowntimeEnum            `mandatory:"true" json:"acceptableDowntime"`
	DatabaseDataSize             DatabaseDataSizeEnum              `mandatory:"true" json:"databaseDataSize"`
	DdlExpectation               DdlExpectationEnum                `mandatory:"true" json:"ddlExpectation"`
	SourceDatabaseConnection     *SourceAssessmentConnection       `mandatory:"true" json:"sourceDatabaseConnection"`
	TargetDatabaseConnection     *TargetAssessmentConnection       `mandatory:"true" json:"targetDatabaseConnection"`
	DatabaseCombination          string                            `json:"databaseCombination"`
}

// UnmarshalJSON unmarshals json
func (m *createassessmentdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateassessmentdetails createassessmentdetails
	s := struct {
		Model Unmarshalercreateassessmentdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CompartmentId = s.Model.CompartmentId
	m.NetworkSpeedMegabitPerSecond = s.Model.NetworkSpeedMegabitPerSecond
	m.AcceptableDowntime = s.Model.AcceptableDowntime
	m.DatabaseDataSize = s.Model.DatabaseDataSize
	m.DdlExpectation = s.Model.DdlExpectation
	m.SourceDatabaseConnection = s.Model.SourceDatabaseConnection
	m.TargetDatabaseConnection = s.Model.TargetDatabaseConnection
	m.Description = s.Model.Description
	m.DisplayName = s.Model.DisplayName
	m.CreationType = s.Model.CreationType
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.DatabaseCombination = s.Model.DatabaseCombination

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createassessmentdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DatabaseCombination {
	case "ORACLE":
		mm := CreateOracleAssessmentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MYSQL":
		mm := CreateMySqlAssessmentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateAssessmentDetails: %s.", m.DatabaseCombination)
		return *m, nil
	}
}

// GetDescription returns Description
func (m createassessmentdetails) GetDescription() *string {
	return m.Description
}

// GetDisplayName returns DisplayName
func (m createassessmentdetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetCreationType returns CreationType
func (m createassessmentdetails) GetCreationType() CreationTypeEnum {
	return m.CreationType
}

// GetFreeformTags returns FreeformTags
func (m createassessmentdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m createassessmentdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetCompartmentId returns CompartmentId
func (m createassessmentdetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetNetworkSpeedMegabitPerSecond returns NetworkSpeedMegabitPerSecond
func (m createassessmentdetails) GetNetworkSpeedMegabitPerSecond() NetworkSpeedMegabitPerSecondEnum {
	return m.NetworkSpeedMegabitPerSecond
}

// GetAcceptableDowntime returns AcceptableDowntime
func (m createassessmentdetails) GetAcceptableDowntime() AcceptableDowntimeEnum {
	return m.AcceptableDowntime
}

// GetDatabaseDataSize returns DatabaseDataSize
func (m createassessmentdetails) GetDatabaseDataSize() DatabaseDataSizeEnum {
	return m.DatabaseDataSize
}

// GetDdlExpectation returns DdlExpectation
func (m createassessmentdetails) GetDdlExpectation() DdlExpectationEnum {
	return m.DdlExpectation
}

// GetSourceDatabaseConnection returns SourceDatabaseConnection
func (m createassessmentdetails) GetSourceDatabaseConnection() *SourceAssessmentConnection {
	return m.SourceDatabaseConnection
}

// GetTargetDatabaseConnection returns TargetDatabaseConnection
func (m createassessmentdetails) GetTargetDatabaseConnection() *TargetAssessmentConnection {
	return m.TargetDatabaseConnection
}

func (m createassessmentdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createassessmentdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNetworkSpeedMegabitPerSecondEnum(string(m.NetworkSpeedMegabitPerSecond)); !ok && m.NetworkSpeedMegabitPerSecond != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NetworkSpeedMegabitPerSecond: %s. Supported values are: %s.", m.NetworkSpeedMegabitPerSecond, strings.Join(GetNetworkSpeedMegabitPerSecondEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAcceptableDowntimeEnum(string(m.AcceptableDowntime)); !ok && m.AcceptableDowntime != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AcceptableDowntime: %s. Supported values are: %s.", m.AcceptableDowntime, strings.Join(GetAcceptableDowntimeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseDataSizeEnum(string(m.DatabaseDataSize)); !ok && m.DatabaseDataSize != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseDataSize: %s. Supported values are: %s.", m.DatabaseDataSize, strings.Join(GetDatabaseDataSizeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDdlExpectationEnum(string(m.DdlExpectation)); !ok && m.DdlExpectation != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DdlExpectation: %s. Supported values are: %s.", m.DdlExpectation, strings.Join(GetDdlExpectationEnumStringValues(), ",")))
	}

	if _, ok := GetMappingCreationTypeEnum(string(m.CreationType)); !ok && m.CreationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CreationType: %s. Supported values are: %s.", m.CreationType, strings.Join(GetCreationTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
