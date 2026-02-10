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

// CreateOracleAssessmentDetails Create Oracle Assessment resource parameters.
type CreateOracleAssessmentDetails struct {

	// The OCID of the resource being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	SourceDatabaseConnection *SourceAssessmentConnection `mandatory:"true" json:"sourceDatabaseConnection"`

	TargetDatabaseConnection *TargetAssessmentConnection `mandatory:"true" json:"targetDatabaseConnection"`

	// A user-friendly description. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags. Example: {"Department": "Finance"}
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Database objects to exclude from migration, cannot be specified alongside 'includeObjects'
	ExcludeObjects []OracleDatabaseObject `mandatory:"false" json:"excludeObjects"`

	// Database objects to include from migration, cannot be specified alongside 'excludeObjects'
	IncludeObjects []OracleDatabaseObject `mandatory:"false" json:"includeObjects"`

	// Specifies the database objects to be excluded from the migration in bulk.
	// The definition accepts input in a CSV format, newline separated for each entry.
	// More details can be found in the documentation.
	BulkIncludeExcludeData *string `mandatory:"false" json:"bulkIncludeExcludeData"`

	// A network speed in Megabits per second.
	NetworkSpeedMegabitPerSecond NetworkSpeedMegabitPerSecondEnum `mandatory:"true" json:"networkSpeedMegabitPerSecond"`

	// Time allowed for the application downtime.
	AcceptableDowntime AcceptableDowntimeEnum `mandatory:"true" json:"acceptableDowntime"`

	// The size of a source database.
	DatabaseDataSize DatabaseDataSizeEnum `mandatory:"true" json:"databaseDataSize"`

	// DDL expectation values.
	DdlExpectation DdlExpectationEnum `mandatory:"true" json:"ddlExpectation"`

	// The type of assessment creation.
	CreationType CreationTypeEnum `mandatory:"false" json:"creationType,omitempty"`
}

// GetDescription returns Description
func (m CreateOracleAssessmentDetails) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m CreateOracleAssessmentDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m CreateOracleAssessmentDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetNetworkSpeedMegabitPerSecond returns NetworkSpeedMegabitPerSecond
func (m CreateOracleAssessmentDetails) GetNetworkSpeedMegabitPerSecond() NetworkSpeedMegabitPerSecondEnum {
	return m.NetworkSpeedMegabitPerSecond
}

// GetAcceptableDowntime returns AcceptableDowntime
func (m CreateOracleAssessmentDetails) GetAcceptableDowntime() AcceptableDowntimeEnum {
	return m.AcceptableDowntime
}

// GetDatabaseDataSize returns DatabaseDataSize
func (m CreateOracleAssessmentDetails) GetDatabaseDataSize() DatabaseDataSizeEnum {
	return m.DatabaseDataSize
}

// GetDdlExpectation returns DdlExpectation
func (m CreateOracleAssessmentDetails) GetDdlExpectation() DdlExpectationEnum {
	return m.DdlExpectation
}

// GetCreationType returns CreationType
func (m CreateOracleAssessmentDetails) GetCreationType() CreationTypeEnum {
	return m.CreationType
}

// GetSourceDatabaseConnection returns SourceDatabaseConnection
func (m CreateOracleAssessmentDetails) GetSourceDatabaseConnection() *SourceAssessmentConnection {
	return m.SourceDatabaseConnection
}

// GetTargetDatabaseConnection returns TargetDatabaseConnection
func (m CreateOracleAssessmentDetails) GetTargetDatabaseConnection() *TargetAssessmentConnection {
	return m.TargetDatabaseConnection
}

// GetFreeformTags returns FreeformTags
func (m CreateOracleAssessmentDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateOracleAssessmentDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateOracleAssessmentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOracleAssessmentDetails) ValidateEnumValue() (bool, error) {
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
func (m CreateOracleAssessmentDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateOracleAssessmentDetails CreateOracleAssessmentDetails
	s := struct {
		DiscriminatorParam string `json:"databaseCombination"`
		MarshalTypeCreateOracleAssessmentDetails
	}{
		"ORACLE",
		(MarshalTypeCreateOracleAssessmentDetails)(m),
	}

	return json.Marshal(&s)
}
