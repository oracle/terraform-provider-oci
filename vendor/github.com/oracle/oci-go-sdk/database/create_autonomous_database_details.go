// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateAutonomousDatabaseDetails Details to create an Oracle Autonomous Database.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type CreateAutonomousDatabaseDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment of the autonomous database.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The database name. The name must begin with an alphabetic character and can contain a maximum of 14 alphanumeric characters. Special characters are not permitted. The database name must be unique in the tenancy.
	DbName *string `mandatory:"true" json:"dbName"`

	// The number of CPU Cores to be made available to the database.
	CpuCoreCount *int `mandatory:"true" json:"cpuCoreCount"`

	// The size, in terabytes, of the data volume that will be created and attached to the database. This storage can later be scaled up if needed.
	DataStorageSizeInTBs *int `mandatory:"true" json:"dataStorageSizeInTBs"`

	// The password must be between 12 and 30 characters long, and must contain at least 1 uppercase, 1 lowercase, and 1 numeric character. It cannot contain the double quote symbol (") or the username "admin", regardless of casing.
	AdminPassword *string `mandatory:"true" json:"adminPassword"`

	// The autonomous database workload type.
	DbWorkload CreateAutonomousDatabaseDetailsDbWorkloadEnum `mandatory:"false" json:"dbWorkload,omitempty"`

	// The user-friendly name for the Autonomous Database. The name does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The Oracle license model that applies to the Oracle Autonomous Database. The default is BRING_YOUR_OWN_LICENSE.
	LicenseModel CreateAutonomousDatabaseDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateAutonomousDatabaseDetails) String() string {
	return common.PointerString(m)
}

// CreateAutonomousDatabaseDetailsDbWorkloadEnum Enum with underlying type: string
type CreateAutonomousDatabaseDetailsDbWorkloadEnum string

// Set of constants representing the allowable values for CreateAutonomousDatabaseDetailsDbWorkloadEnum
const (
	CreateAutonomousDatabaseDetailsDbWorkloadOltp CreateAutonomousDatabaseDetailsDbWorkloadEnum = "OLTP"
	CreateAutonomousDatabaseDetailsDbWorkloadDw   CreateAutonomousDatabaseDetailsDbWorkloadEnum = "DW"
)

var mappingCreateAutonomousDatabaseDetailsDbWorkload = map[string]CreateAutonomousDatabaseDetailsDbWorkloadEnum{
	"OLTP": CreateAutonomousDatabaseDetailsDbWorkloadOltp,
	"DW":   CreateAutonomousDatabaseDetailsDbWorkloadDw,
}

// GetCreateAutonomousDatabaseDetailsDbWorkloadEnumValues Enumerates the set of values for CreateAutonomousDatabaseDetailsDbWorkloadEnum
func GetCreateAutonomousDatabaseDetailsDbWorkloadEnumValues() []CreateAutonomousDatabaseDetailsDbWorkloadEnum {
	values := make([]CreateAutonomousDatabaseDetailsDbWorkloadEnum, 0)
	for _, v := range mappingCreateAutonomousDatabaseDetailsDbWorkload {
		values = append(values, v)
	}
	return values
}

// CreateAutonomousDatabaseDetailsLicenseModelEnum Enum with underlying type: string
type CreateAutonomousDatabaseDetailsLicenseModelEnum string

// Set of constants representing the allowable values for CreateAutonomousDatabaseDetailsLicenseModelEnum
const (
	CreateAutonomousDatabaseDetailsLicenseModelLicenseIncluded     CreateAutonomousDatabaseDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	CreateAutonomousDatabaseDetailsLicenseModelBringYourOwnLicense CreateAutonomousDatabaseDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingCreateAutonomousDatabaseDetailsLicenseModel = map[string]CreateAutonomousDatabaseDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       CreateAutonomousDatabaseDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": CreateAutonomousDatabaseDetailsLicenseModelBringYourOwnLicense,
}

// GetCreateAutonomousDatabaseDetailsLicenseModelEnumValues Enumerates the set of values for CreateAutonomousDatabaseDetailsLicenseModelEnum
func GetCreateAutonomousDatabaseDetailsLicenseModelEnumValues() []CreateAutonomousDatabaseDetailsLicenseModelEnum {
	values := make([]CreateAutonomousDatabaseDetailsLicenseModelEnum, 0)
	for _, v := range mappingCreateAutonomousDatabaseDetailsLicenseModel {
		values = append(values, v)
	}
	return values
}
