// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseFeatureConfiguration The details of an external database feature configuration.
type DatabaseFeatureConfiguration interface {

	// The list of statuses for Database Management features.
	GetFeatureStatus() DatabaseFeatureConfigurationFeatureStatusEnum

	GetConnectorDetails() ConnectorDetails

	GetDatabaseConnectionDetails() *DatabaseConnectionDetails
}

type databasefeatureconfiguration struct {
	JsonData                  []byte
	ConnectorDetails          connectordetails                              `mandatory:"false" json:"connectorDetails"`
	DatabaseConnectionDetails *DatabaseConnectionDetails                    `mandatory:"false" json:"databaseConnectionDetails"`
	FeatureStatus             DatabaseFeatureConfigurationFeatureStatusEnum `mandatory:"true" json:"featureStatus"`
	Feature                   string                                        `json:"feature"`
}

// UnmarshalJSON unmarshals json
func (m *databasefeatureconfiguration) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasefeatureconfiguration databasefeatureconfiguration
	s := struct {
		Model Unmarshalerdatabasefeatureconfiguration
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.FeatureStatus = s.Model.FeatureStatus
	m.ConnectorDetails = s.Model.ConnectorDetails
	m.DatabaseConnectionDetails = s.Model.DatabaseConnectionDetails
	m.Feature = s.Model.Feature

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasefeatureconfiguration) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Feature {
	case "DIAGNOSTICS_AND_MANAGEMENT":
		mm := DatabaseDiagnosticsAndManagementFeatureConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DB_LIFECYCLE_MANAGEMENT":
		mm := DatabaseLifecycleFeatureConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SQLWATCH":
		mm := DatabaseSqlWatchFeatureConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DatabaseFeatureConfiguration: %s.", m.Feature)
		return *m, nil
	}
}

// GetConnectorDetails returns ConnectorDetails
func (m databasefeatureconfiguration) GetConnectorDetails() connectordetails {
	return m.ConnectorDetails
}

// GetDatabaseConnectionDetails returns DatabaseConnectionDetails
func (m databasefeatureconfiguration) GetDatabaseConnectionDetails() *DatabaseConnectionDetails {
	return m.DatabaseConnectionDetails
}

// GetFeatureStatus returns FeatureStatus
func (m databasefeatureconfiguration) GetFeatureStatus() DatabaseFeatureConfigurationFeatureStatusEnum {
	return m.FeatureStatus
}

func (m databasefeatureconfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasefeatureconfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseFeatureConfigurationFeatureStatusEnum(string(m.FeatureStatus)); !ok && m.FeatureStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FeatureStatus: %s. Supported values are: %s.", m.FeatureStatus, strings.Join(GetDatabaseFeatureConfigurationFeatureStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseFeatureConfigurationFeatureStatusEnum Enum with underlying type: string
type DatabaseFeatureConfigurationFeatureStatusEnum string

// Set of constants representing the allowable values for DatabaseFeatureConfigurationFeatureStatusEnum
const (
	DatabaseFeatureConfigurationFeatureStatusEnabled             DatabaseFeatureConfigurationFeatureStatusEnum = "ENABLED"
	DatabaseFeatureConfigurationFeatureStatusNotEnabled          DatabaseFeatureConfigurationFeatureStatusEnum = "NOT_ENABLED"
	DatabaseFeatureConfigurationFeatureStatusUnsupported         DatabaseFeatureConfigurationFeatureStatusEnum = "UNSUPPORTED"
	DatabaseFeatureConfigurationFeatureStatusFailedEnabling      DatabaseFeatureConfigurationFeatureStatusEnum = "FAILED_ENABLING"
	DatabaseFeatureConfigurationFeatureStatusFailedDisabling     DatabaseFeatureConfigurationFeatureStatusEnum = "FAILED_DISABLING"
	DatabaseFeatureConfigurationFeatureStatusFailed              DatabaseFeatureConfigurationFeatureStatusEnum = "FAILED"
	DatabaseFeatureConfigurationFeatureStatusEnabledWithWarnings DatabaseFeatureConfigurationFeatureStatusEnum = "ENABLED_WITH_WARNINGS"
	DatabaseFeatureConfigurationFeatureStatusPendingDisable      DatabaseFeatureConfigurationFeatureStatusEnum = "PENDING_DISABLE"
	DatabaseFeatureConfigurationFeatureStatusEnabling            DatabaseFeatureConfigurationFeatureStatusEnum = "ENABLING"
	DatabaseFeatureConfigurationFeatureStatusDisabling           DatabaseFeatureConfigurationFeatureStatusEnum = "DISABLING"
)

var mappingDatabaseFeatureConfigurationFeatureStatusEnum = map[string]DatabaseFeatureConfigurationFeatureStatusEnum{
	"ENABLED":               DatabaseFeatureConfigurationFeatureStatusEnabled,
	"NOT_ENABLED":           DatabaseFeatureConfigurationFeatureStatusNotEnabled,
	"UNSUPPORTED":           DatabaseFeatureConfigurationFeatureStatusUnsupported,
	"FAILED_ENABLING":       DatabaseFeatureConfigurationFeatureStatusFailedEnabling,
	"FAILED_DISABLING":      DatabaseFeatureConfigurationFeatureStatusFailedDisabling,
	"FAILED":                DatabaseFeatureConfigurationFeatureStatusFailed,
	"ENABLED_WITH_WARNINGS": DatabaseFeatureConfigurationFeatureStatusEnabledWithWarnings,
	"PENDING_DISABLE":       DatabaseFeatureConfigurationFeatureStatusPendingDisable,
	"ENABLING":              DatabaseFeatureConfigurationFeatureStatusEnabling,
	"DISABLING":             DatabaseFeatureConfigurationFeatureStatusDisabling,
}

var mappingDatabaseFeatureConfigurationFeatureStatusEnumLowerCase = map[string]DatabaseFeatureConfigurationFeatureStatusEnum{
	"enabled":               DatabaseFeatureConfigurationFeatureStatusEnabled,
	"not_enabled":           DatabaseFeatureConfigurationFeatureStatusNotEnabled,
	"unsupported":           DatabaseFeatureConfigurationFeatureStatusUnsupported,
	"failed_enabling":       DatabaseFeatureConfigurationFeatureStatusFailedEnabling,
	"failed_disabling":      DatabaseFeatureConfigurationFeatureStatusFailedDisabling,
	"failed":                DatabaseFeatureConfigurationFeatureStatusFailed,
	"enabled_with_warnings": DatabaseFeatureConfigurationFeatureStatusEnabledWithWarnings,
	"pending_disable":       DatabaseFeatureConfigurationFeatureStatusPendingDisable,
	"enabling":              DatabaseFeatureConfigurationFeatureStatusEnabling,
	"disabling":             DatabaseFeatureConfigurationFeatureStatusDisabling,
}

// GetDatabaseFeatureConfigurationFeatureStatusEnumValues Enumerates the set of values for DatabaseFeatureConfigurationFeatureStatusEnum
func GetDatabaseFeatureConfigurationFeatureStatusEnumValues() []DatabaseFeatureConfigurationFeatureStatusEnum {
	values := make([]DatabaseFeatureConfigurationFeatureStatusEnum, 0)
	for _, v := range mappingDatabaseFeatureConfigurationFeatureStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseFeatureConfigurationFeatureStatusEnumStringValues Enumerates the set of values in String for DatabaseFeatureConfigurationFeatureStatusEnum
func GetDatabaseFeatureConfigurationFeatureStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"NOT_ENABLED",
		"UNSUPPORTED",
		"FAILED_ENABLING",
		"FAILED_DISABLING",
		"FAILED",
		"ENABLED_WITH_WARNINGS",
		"PENDING_DISABLE",
		"ENABLING",
		"DISABLING",
	}
}

// GetMappingDatabaseFeatureConfigurationFeatureStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseFeatureConfigurationFeatureStatusEnum(val string) (DatabaseFeatureConfigurationFeatureStatusEnum, bool) {
	enum, ok := mappingDatabaseFeatureConfigurationFeatureStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
