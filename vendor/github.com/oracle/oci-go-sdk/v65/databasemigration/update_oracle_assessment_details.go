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

// UpdateOracleAssessmentDetails Update Oracle Assessment resource parameters.
type UpdateOracleAssessmentDetails struct {

	// A user-friendly description. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	SourceDatabaseConnection *SourceAssessmentConnection `mandatory:"false" json:"sourceDatabaseConnection"`

	TargetDatabaseConnection *TargetAssessmentConnection `mandatory:"false" json:"targetDatabaseConnection"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags. Example: {"Department": "Finance"}
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A network speed in Megabits per second.
	NetworkSpeedMegabitPerSecond NetworkSpeedMegabitPerSecondEnum `mandatory:"false" json:"networkSpeedMegabitPerSecond,omitempty"`

	// Time allowed for the application downtime.
	AcceptableDowntime AcceptableDowntimeEnum `mandatory:"false" json:"acceptableDowntime,omitempty"`

	// The size of a source database.
	DatabaseDataSize DatabaseDataSizeEnum `mandatory:"false" json:"databaseDataSize,omitempty"`

	// DDL expectation values.
	DdlExpectation DdlExpectationEnum `mandatory:"false" json:"ddlExpectation,omitempty"`

	// The type of assessment creation.
	CreationType CreationTypeEnum `mandatory:"false" json:"creationType,omitempty"`
}

// GetDescription returns Description
func (m UpdateOracleAssessmentDetails) GetDescription() *string {
	return m.Description
}

// GetDisplayName returns DisplayName
func (m UpdateOracleAssessmentDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetNetworkSpeedMegabitPerSecond returns NetworkSpeedMegabitPerSecond
func (m UpdateOracleAssessmentDetails) GetNetworkSpeedMegabitPerSecond() NetworkSpeedMegabitPerSecondEnum {
	return m.NetworkSpeedMegabitPerSecond
}

// GetAcceptableDowntime returns AcceptableDowntime
func (m UpdateOracleAssessmentDetails) GetAcceptableDowntime() AcceptableDowntimeEnum {
	return m.AcceptableDowntime
}

// GetDatabaseDataSize returns DatabaseDataSize
func (m UpdateOracleAssessmentDetails) GetDatabaseDataSize() DatabaseDataSizeEnum {
	return m.DatabaseDataSize
}

// GetDdlExpectation returns DdlExpectation
func (m UpdateOracleAssessmentDetails) GetDdlExpectation() DdlExpectationEnum {
	return m.DdlExpectation
}

// GetCreationType returns CreationType
func (m UpdateOracleAssessmentDetails) GetCreationType() CreationTypeEnum {
	return m.CreationType
}

// GetSourceDatabaseConnection returns SourceDatabaseConnection
func (m UpdateOracleAssessmentDetails) GetSourceDatabaseConnection() *SourceAssessmentConnection {
	return m.SourceDatabaseConnection
}

// GetTargetDatabaseConnection returns TargetDatabaseConnection
func (m UpdateOracleAssessmentDetails) GetTargetDatabaseConnection() *TargetAssessmentConnection {
	return m.TargetDatabaseConnection
}

// GetFreeformTags returns FreeformTags
func (m UpdateOracleAssessmentDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m UpdateOracleAssessmentDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateOracleAssessmentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateOracleAssessmentDetails) ValidateEnumValue() (bool, error) {
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

// MarshalJSON marshals to json representation
func (m UpdateOracleAssessmentDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateOracleAssessmentDetails UpdateOracleAssessmentDetails
	s := struct {
		DiscriminatorParam string `json:"databaseCombination"`
		MarshalTypeUpdateOracleAssessmentDetails
	}{
		"ORACLE",
		(MarshalTypeUpdateOracleAssessmentDetails)(m),
	}

	return json.Marshal(&s)
}
