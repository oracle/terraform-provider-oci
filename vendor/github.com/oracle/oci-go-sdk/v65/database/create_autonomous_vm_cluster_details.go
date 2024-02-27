// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// CreateAutonomousVmClusterDetails Details for the create Autonomous VM cluster operation.
type CreateAutonomousVmClusterDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the Autonomous VM cluster. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
	ExadataInfrastructureId *string `mandatory:"true" json:"exadataInfrastructureId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VM cluster network.
	VmClusterNetworkId *string `mandatory:"true" json:"vmClusterNetworkId"`

	// The time zone to use for the Autonomous VM cluster. For details, see DB System Time Zones (https://docs.cloud.oracle.com/Content/Database/References/timezones.htm).
	TimeZone *string `mandatory:"false" json:"timeZone"`

	// If true, database backup on local Exadata storage is configured for the Autonomous VM cluster. If false, database backup on local Exadata storage is not available in the Autonomous VM cluster.
	IsLocalBackupEnabled *bool `mandatory:"false" json:"isLocalBackupEnabled"`

	// The Oracle license model that applies to the Autonomous VM cluster. The default is BRING_YOUR_OWN_LICENSE.
	LicenseModel CreateAutonomousVmClusterDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The total number of Autonomous Container Databases that can be created.
	TotalContainerDatabases *int `mandatory:"false" json:"totalContainerDatabases"`

	// The number of CPU cores to enable per VM cluster node.
	CpuCoreCountPerNode *int `mandatory:"false" json:"cpuCoreCountPerNode"`

	// The compute model of the Autonomous VM Cluster. ECPU compute model is the recommended model and OCPU compute model is legacy.
	ComputeModel CreateAutonomousVmClusterDetailsComputeModelEnum `mandatory:"false" json:"computeModel,omitempty"`

	// The amount of memory (in GBs) to be enabled per OCPU or ECPU.
	MemoryPerOracleComputeUnitInGBs *int `mandatory:"false" json:"memoryPerOracleComputeUnitInGBs"`

	// The data disk group size to be allocated for Autonomous Databases, in TBs.
	AutonomousDataStorageSizeInTBs *float64 `mandatory:"false" json:"autonomousDataStorageSizeInTBs"`

	MaintenanceWindowDetails *MaintenanceWindow `mandatory:"false" json:"maintenanceWindowDetails"`

	// The list of OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Db servers.
	DbServers []string `mandatory:"false" json:"dbServers"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The SCAN Listener TLS port number. Default value is 2484.
	ScanListenerPortTls *int `mandatory:"false" json:"scanListenerPortTls"`

	// The SCAN Listener Non TLS port number. Default value is 1521.
	ScanListenerPortNonTls *int `mandatory:"false" json:"scanListenerPortNonTls"`

	// Enable mutual TLS(mTLS) authentication for database while provisioning a VMCluster. Default is TLS.
	IsMtlsEnabled *bool `mandatory:"false" json:"isMtlsEnabled"`
}

func (m CreateAutonomousVmClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateAutonomousVmClusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateAutonomousVmClusterDetailsLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetCreateAutonomousVmClusterDetailsLicenseModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateAutonomousVmClusterDetailsComputeModelEnum(string(m.ComputeModel)); !ok && m.ComputeModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ComputeModel: %s. Supported values are: %s.", m.ComputeModel, strings.Join(GetCreateAutonomousVmClusterDetailsComputeModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateAutonomousVmClusterDetailsLicenseModelEnum Enum with underlying type: string
type CreateAutonomousVmClusterDetailsLicenseModelEnum string

// Set of constants representing the allowable values for CreateAutonomousVmClusterDetailsLicenseModelEnum
const (
	CreateAutonomousVmClusterDetailsLicenseModelLicenseIncluded     CreateAutonomousVmClusterDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	CreateAutonomousVmClusterDetailsLicenseModelBringYourOwnLicense CreateAutonomousVmClusterDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingCreateAutonomousVmClusterDetailsLicenseModelEnum = map[string]CreateAutonomousVmClusterDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       CreateAutonomousVmClusterDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": CreateAutonomousVmClusterDetailsLicenseModelBringYourOwnLicense,
}

var mappingCreateAutonomousVmClusterDetailsLicenseModelEnumLowerCase = map[string]CreateAutonomousVmClusterDetailsLicenseModelEnum{
	"license_included":       CreateAutonomousVmClusterDetailsLicenseModelLicenseIncluded,
	"bring_your_own_license": CreateAutonomousVmClusterDetailsLicenseModelBringYourOwnLicense,
}

// GetCreateAutonomousVmClusterDetailsLicenseModelEnumValues Enumerates the set of values for CreateAutonomousVmClusterDetailsLicenseModelEnum
func GetCreateAutonomousVmClusterDetailsLicenseModelEnumValues() []CreateAutonomousVmClusterDetailsLicenseModelEnum {
	values := make([]CreateAutonomousVmClusterDetailsLicenseModelEnum, 0)
	for _, v := range mappingCreateAutonomousVmClusterDetailsLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateAutonomousVmClusterDetailsLicenseModelEnumStringValues Enumerates the set of values in String for CreateAutonomousVmClusterDetailsLicenseModelEnum
func GetCreateAutonomousVmClusterDetailsLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingCreateAutonomousVmClusterDetailsLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateAutonomousVmClusterDetailsLicenseModelEnum(val string) (CreateAutonomousVmClusterDetailsLicenseModelEnum, bool) {
	enum, ok := mappingCreateAutonomousVmClusterDetailsLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateAutonomousVmClusterDetailsComputeModelEnum Enum with underlying type: string
type CreateAutonomousVmClusterDetailsComputeModelEnum string

// Set of constants representing the allowable values for CreateAutonomousVmClusterDetailsComputeModelEnum
const (
	CreateAutonomousVmClusterDetailsComputeModelEcpu CreateAutonomousVmClusterDetailsComputeModelEnum = "ECPU"
	CreateAutonomousVmClusterDetailsComputeModelOcpu CreateAutonomousVmClusterDetailsComputeModelEnum = "OCPU"
)

var mappingCreateAutonomousVmClusterDetailsComputeModelEnum = map[string]CreateAutonomousVmClusterDetailsComputeModelEnum{
	"ECPU": CreateAutonomousVmClusterDetailsComputeModelEcpu,
	"OCPU": CreateAutonomousVmClusterDetailsComputeModelOcpu,
}

var mappingCreateAutonomousVmClusterDetailsComputeModelEnumLowerCase = map[string]CreateAutonomousVmClusterDetailsComputeModelEnum{
	"ecpu": CreateAutonomousVmClusterDetailsComputeModelEcpu,
	"ocpu": CreateAutonomousVmClusterDetailsComputeModelOcpu,
}

// GetCreateAutonomousVmClusterDetailsComputeModelEnumValues Enumerates the set of values for CreateAutonomousVmClusterDetailsComputeModelEnum
func GetCreateAutonomousVmClusterDetailsComputeModelEnumValues() []CreateAutonomousVmClusterDetailsComputeModelEnum {
	values := make([]CreateAutonomousVmClusterDetailsComputeModelEnum, 0)
	for _, v := range mappingCreateAutonomousVmClusterDetailsComputeModelEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateAutonomousVmClusterDetailsComputeModelEnumStringValues Enumerates the set of values in String for CreateAutonomousVmClusterDetailsComputeModelEnum
func GetCreateAutonomousVmClusterDetailsComputeModelEnumStringValues() []string {
	return []string{
		"ECPU",
		"OCPU",
	}
}

// GetMappingCreateAutonomousVmClusterDetailsComputeModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateAutonomousVmClusterDetailsComputeModelEnum(val string) (CreateAutonomousVmClusterDetailsComputeModelEnum, bool) {
	enum, ok := mappingCreateAutonomousVmClusterDetailsComputeModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
