// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateExternalBackupJobDetails The representation of CreateExternalBackupJobDetails
type CreateExternalBackupJobDetails struct {

	// The targeted availability domain for the backup.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment where this backup should be created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name for the backup. This name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// A valid Oracle Database version.
	DbVersion *string `mandatory:"true" json:"dbVersion"`

	// The name of the database from which the backup is being taken.
	DbName *string `mandatory:"true" json:"dbName"`

	// The `DBID` of the Oracle Database being backed up.
	ExternalDatabaseIdentifier *int64 `mandatory:"true" json:"externalDatabaseIdentifier"`

	// The character set for the database.
	CharacterSet *string `mandatory:"true" json:"characterSet"`

	// The national character set for the database.
	NcharacterSet *string `mandatory:"true" json:"ncharacterSet"`

	// The mode (single instance or RAC) of the database being backed up.
	DatabaseMode CreateExternalBackupJobDetailsDatabaseModeEnum `mandatory:"true" json:"databaseMode"`

	// The Oracle Database edition to use for creating a database from this standalone backup.
	// Note that 2-node RAC DB systems require Enterprise Edition - Extreme Performance.
	DatabaseEdition CreateExternalBackupJobDetailsDatabaseEditionEnum `mandatory:"true" json:"databaseEdition"`

	// The `DB_UNIQUE_NAME` of the Oracle Database being backed up.
	DbUniqueName *string `mandatory:"false" json:"dbUniqueName"`

	// The pluggable database name.
	PdbName *string `mandatory:"false" json:"pdbName"`
}

func (m CreateExternalBackupJobDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateExternalBackupJobDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateExternalBackupJobDetailsDatabaseModeEnum(string(m.DatabaseMode)); !ok && m.DatabaseMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseMode: %s. Supported values are: %s.", m.DatabaseMode, strings.Join(GetCreateExternalBackupJobDetailsDatabaseModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateExternalBackupJobDetailsDatabaseEditionEnum(string(m.DatabaseEdition)); !ok && m.DatabaseEdition != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseEdition: %s. Supported values are: %s.", m.DatabaseEdition, strings.Join(GetCreateExternalBackupJobDetailsDatabaseEditionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateExternalBackupJobDetailsDatabaseModeEnum Enum with underlying type: string
type CreateExternalBackupJobDetailsDatabaseModeEnum string

// Set of constants representing the allowable values for CreateExternalBackupJobDetailsDatabaseModeEnum
const (
	CreateExternalBackupJobDetailsDatabaseModeSi  CreateExternalBackupJobDetailsDatabaseModeEnum = "SI"
	CreateExternalBackupJobDetailsDatabaseModeRac CreateExternalBackupJobDetailsDatabaseModeEnum = "RAC"
)

var mappingCreateExternalBackupJobDetailsDatabaseModeEnum = map[string]CreateExternalBackupJobDetailsDatabaseModeEnum{
	"SI":  CreateExternalBackupJobDetailsDatabaseModeSi,
	"RAC": CreateExternalBackupJobDetailsDatabaseModeRac,
}

var mappingCreateExternalBackupJobDetailsDatabaseModeEnumLowerCase = map[string]CreateExternalBackupJobDetailsDatabaseModeEnum{
	"si":  CreateExternalBackupJobDetailsDatabaseModeSi,
	"rac": CreateExternalBackupJobDetailsDatabaseModeRac,
}

// GetCreateExternalBackupJobDetailsDatabaseModeEnumValues Enumerates the set of values for CreateExternalBackupJobDetailsDatabaseModeEnum
func GetCreateExternalBackupJobDetailsDatabaseModeEnumValues() []CreateExternalBackupJobDetailsDatabaseModeEnum {
	values := make([]CreateExternalBackupJobDetailsDatabaseModeEnum, 0)
	for _, v := range mappingCreateExternalBackupJobDetailsDatabaseModeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateExternalBackupJobDetailsDatabaseModeEnumStringValues Enumerates the set of values in String for CreateExternalBackupJobDetailsDatabaseModeEnum
func GetCreateExternalBackupJobDetailsDatabaseModeEnumStringValues() []string {
	return []string{
		"SI",
		"RAC",
	}
}

// GetMappingCreateExternalBackupJobDetailsDatabaseModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateExternalBackupJobDetailsDatabaseModeEnum(val string) (CreateExternalBackupJobDetailsDatabaseModeEnum, bool) {
	enum, ok := mappingCreateExternalBackupJobDetailsDatabaseModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateExternalBackupJobDetailsDatabaseEditionEnum Enum with underlying type: string
type CreateExternalBackupJobDetailsDatabaseEditionEnum string

// Set of constants representing the allowable values for CreateExternalBackupJobDetailsDatabaseEditionEnum
const (
	CreateExternalBackupJobDetailsDatabaseEditionStandardEdition                     CreateExternalBackupJobDetailsDatabaseEditionEnum = "STANDARD_EDITION"
	CreateExternalBackupJobDetailsDatabaseEditionEnterpriseEdition                   CreateExternalBackupJobDetailsDatabaseEditionEnum = "ENTERPRISE_EDITION"
	CreateExternalBackupJobDetailsDatabaseEditionEnterpriseEditionHighPerformance    CreateExternalBackupJobDetailsDatabaseEditionEnum = "ENTERPRISE_EDITION_HIGH_PERFORMANCE"
	CreateExternalBackupJobDetailsDatabaseEditionEnterpriseEditionExtremePerformance CreateExternalBackupJobDetailsDatabaseEditionEnum = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
)

var mappingCreateExternalBackupJobDetailsDatabaseEditionEnum = map[string]CreateExternalBackupJobDetailsDatabaseEditionEnum{
	"STANDARD_EDITION":                       CreateExternalBackupJobDetailsDatabaseEditionStandardEdition,
	"ENTERPRISE_EDITION":                     CreateExternalBackupJobDetailsDatabaseEditionEnterpriseEdition,
	"ENTERPRISE_EDITION_HIGH_PERFORMANCE":    CreateExternalBackupJobDetailsDatabaseEditionEnterpriseEditionHighPerformance,
	"ENTERPRISE_EDITION_EXTREME_PERFORMANCE": CreateExternalBackupJobDetailsDatabaseEditionEnterpriseEditionExtremePerformance,
}

var mappingCreateExternalBackupJobDetailsDatabaseEditionEnumLowerCase = map[string]CreateExternalBackupJobDetailsDatabaseEditionEnum{
	"standard_edition":                       CreateExternalBackupJobDetailsDatabaseEditionStandardEdition,
	"enterprise_edition":                     CreateExternalBackupJobDetailsDatabaseEditionEnterpriseEdition,
	"enterprise_edition_high_performance":    CreateExternalBackupJobDetailsDatabaseEditionEnterpriseEditionHighPerformance,
	"enterprise_edition_extreme_performance": CreateExternalBackupJobDetailsDatabaseEditionEnterpriseEditionExtremePerformance,
}

// GetCreateExternalBackupJobDetailsDatabaseEditionEnumValues Enumerates the set of values for CreateExternalBackupJobDetailsDatabaseEditionEnum
func GetCreateExternalBackupJobDetailsDatabaseEditionEnumValues() []CreateExternalBackupJobDetailsDatabaseEditionEnum {
	values := make([]CreateExternalBackupJobDetailsDatabaseEditionEnum, 0)
	for _, v := range mappingCreateExternalBackupJobDetailsDatabaseEditionEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateExternalBackupJobDetailsDatabaseEditionEnumStringValues Enumerates the set of values in String for CreateExternalBackupJobDetailsDatabaseEditionEnum
func GetCreateExternalBackupJobDetailsDatabaseEditionEnumStringValues() []string {
	return []string{
		"STANDARD_EDITION",
		"ENTERPRISE_EDITION",
		"ENTERPRISE_EDITION_HIGH_PERFORMANCE",
		"ENTERPRISE_EDITION_EXTREME_PERFORMANCE",
	}
}

// GetMappingCreateExternalBackupJobDetailsDatabaseEditionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateExternalBackupJobDetailsDatabaseEditionEnum(val string) (CreateExternalBackupJobDetailsDatabaseEditionEnum, bool) {
	enum, ok := mappingCreateExternalBackupJobDetailsDatabaseEditionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
