// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateAutonomousVmClusterDetails Details for updating the Autonomous VM cluster.
type UpdateAutonomousVmClusterDetails struct {
	MaintenanceWindowDetails *MaintenanceWindow `mandatory:"false" json:"maintenanceWindowDetails"`

	// The Oracle license model that applies to the Autonomous VM cluster. The default is BRING_YOUR_OWN_LICENSE.
	LicenseModel UpdateAutonomousVmClusterDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The new value of autonomous data storage (in TBs) for the Autonomous VM cluster.
	AutonomousDataStorageSizeInTBs *float64 `mandatory:"false" json:"autonomousDataStorageSizeInTBs"`

	// The new value of cpus per Autonomous VM cluster per node for the Autonomous VM cluster.
	CpuCoreCountPerNode *int `mandatory:"false" json:"cpuCoreCountPerNode"`

	// The new value of maximum number of ACDs for the Autonomous VM cluster.
	TotalContainerDatabases *int `mandatory:"false" json:"totalContainerDatabases"`

	// The time zone to use for the Autonomous VM cluster. For details, see DB System Time Zones (https://docs.oracle.com/iaas/Content/Database/References/timezones.htm).
	TimeZone *string `mandatory:"false" json:"timeZone"`

	// The SCAN Listener TLS port number. Default value is 2484.
	ScanListenerPortTls *int `mandatory:"false" json:"scanListenerPortTls"`

	// The SCAN Listener Non TLS port number. Default value is 1521.
	ScanListenerPortNonTls *int `mandatory:"false" json:"scanListenerPortNonTls"`

	// Enable mutual TLS(mTLS) authentication for database. Default is TLS.
	IsMtlsEnabled *bool `mandatory:"false" json:"isMtlsEnabled"`

	// The distribution algorithm used for the Autonomous VM cluster.
	DistributionAlgorithm UpdateAutonomousVmClusterDetailsDistributionAlgorithmEnum `mandatory:"false" json:"distributionAlgorithm,omitempty"`
}

func (m UpdateAutonomousVmClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateAutonomousVmClusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateAutonomousVmClusterDetailsLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetUpdateAutonomousVmClusterDetailsLicenseModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateAutonomousVmClusterDetailsDistributionAlgorithmEnum(string(m.DistributionAlgorithm)); !ok && m.DistributionAlgorithm != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DistributionAlgorithm: %s. Supported values are: %s.", m.DistributionAlgorithm, strings.Join(GetUpdateAutonomousVmClusterDetailsDistributionAlgorithmEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateAutonomousVmClusterDetailsLicenseModelEnum Enum with underlying type: string
type UpdateAutonomousVmClusterDetailsLicenseModelEnum string

// Set of constants representing the allowable values for UpdateAutonomousVmClusterDetailsLicenseModelEnum
const (
	UpdateAutonomousVmClusterDetailsLicenseModelLicenseIncluded     UpdateAutonomousVmClusterDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	UpdateAutonomousVmClusterDetailsLicenseModelBringYourOwnLicense UpdateAutonomousVmClusterDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingUpdateAutonomousVmClusterDetailsLicenseModelEnum = map[string]UpdateAutonomousVmClusterDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       UpdateAutonomousVmClusterDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": UpdateAutonomousVmClusterDetailsLicenseModelBringYourOwnLicense,
}

var mappingUpdateAutonomousVmClusterDetailsLicenseModelEnumLowerCase = map[string]UpdateAutonomousVmClusterDetailsLicenseModelEnum{
	"license_included":       UpdateAutonomousVmClusterDetailsLicenseModelLicenseIncluded,
	"bring_your_own_license": UpdateAutonomousVmClusterDetailsLicenseModelBringYourOwnLicense,
}

// GetUpdateAutonomousVmClusterDetailsLicenseModelEnumValues Enumerates the set of values for UpdateAutonomousVmClusterDetailsLicenseModelEnum
func GetUpdateAutonomousVmClusterDetailsLicenseModelEnumValues() []UpdateAutonomousVmClusterDetailsLicenseModelEnum {
	values := make([]UpdateAutonomousVmClusterDetailsLicenseModelEnum, 0)
	for _, v := range mappingUpdateAutonomousVmClusterDetailsLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateAutonomousVmClusterDetailsLicenseModelEnumStringValues Enumerates the set of values in String for UpdateAutonomousVmClusterDetailsLicenseModelEnum
func GetUpdateAutonomousVmClusterDetailsLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingUpdateAutonomousVmClusterDetailsLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateAutonomousVmClusterDetailsLicenseModelEnum(val string) (UpdateAutonomousVmClusterDetailsLicenseModelEnum, bool) {
	enum, ok := mappingUpdateAutonomousVmClusterDetailsLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateAutonomousVmClusterDetailsDistributionAlgorithmEnum Enum with underlying type: string
type UpdateAutonomousVmClusterDetailsDistributionAlgorithmEnum string

// Set of constants representing the allowable values for UpdateAutonomousVmClusterDetailsDistributionAlgorithmEnum
const (
	UpdateAutonomousVmClusterDetailsDistributionAlgorithmResourceOptimized     UpdateAutonomousVmClusterDetailsDistributionAlgorithmEnum = "RESOURCE_OPTIMIZED"
	UpdateAutonomousVmClusterDetailsDistributionAlgorithmDistributionOptimized UpdateAutonomousVmClusterDetailsDistributionAlgorithmEnum = "DISTRIBUTION_OPTIMIZED"
)

var mappingUpdateAutonomousVmClusterDetailsDistributionAlgorithmEnum = map[string]UpdateAutonomousVmClusterDetailsDistributionAlgorithmEnum{
	"RESOURCE_OPTIMIZED":     UpdateAutonomousVmClusterDetailsDistributionAlgorithmResourceOptimized,
	"DISTRIBUTION_OPTIMIZED": UpdateAutonomousVmClusterDetailsDistributionAlgorithmDistributionOptimized,
}

var mappingUpdateAutonomousVmClusterDetailsDistributionAlgorithmEnumLowerCase = map[string]UpdateAutonomousVmClusterDetailsDistributionAlgorithmEnum{
	"resource_optimized":     UpdateAutonomousVmClusterDetailsDistributionAlgorithmResourceOptimized,
	"distribution_optimized": UpdateAutonomousVmClusterDetailsDistributionAlgorithmDistributionOptimized,
}

// GetUpdateAutonomousVmClusterDetailsDistributionAlgorithmEnumValues Enumerates the set of values for UpdateAutonomousVmClusterDetailsDistributionAlgorithmEnum
func GetUpdateAutonomousVmClusterDetailsDistributionAlgorithmEnumValues() []UpdateAutonomousVmClusterDetailsDistributionAlgorithmEnum {
	values := make([]UpdateAutonomousVmClusterDetailsDistributionAlgorithmEnum, 0)
	for _, v := range mappingUpdateAutonomousVmClusterDetailsDistributionAlgorithmEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateAutonomousVmClusterDetailsDistributionAlgorithmEnumStringValues Enumerates the set of values in String for UpdateAutonomousVmClusterDetailsDistributionAlgorithmEnum
func GetUpdateAutonomousVmClusterDetailsDistributionAlgorithmEnumStringValues() []string {
	return []string{
		"RESOURCE_OPTIMIZED",
		"DISTRIBUTION_OPTIMIZED",
	}
}

// GetMappingUpdateAutonomousVmClusterDetailsDistributionAlgorithmEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateAutonomousVmClusterDetailsDistributionAlgorithmEnum(val string) (UpdateAutonomousVmClusterDetailsDistributionAlgorithmEnum, bool) {
	enum, ok := mappingUpdateAutonomousVmClusterDetailsDistributionAlgorithmEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
