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

// AutonomousVmCluster Details of the Autonomous VM cluster.
type AutonomousVmCluster struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Autonomous VM cluster.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the Autonomous VM cluster. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the Autonomous VM cluster.
	LifecycleState AutonomousVmClusterLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
	ExadataInfrastructureId *string `mandatory:"true" json:"exadataInfrastructureId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VM cluster network.
	VmClusterNetworkId *string `mandatory:"true" json:"vmClusterNetworkId"`

	// The date and time that the Autonomous VM cluster was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The time zone to use for the Autonomous VM cluster. For details, see DB System Time Zones (https://docs.cloud.oracle.com/Content/Database/References/timezones.htm).
	TimeZone *string `mandatory:"false" json:"timeZone"`

	// If true, database backup on local Exadata storage is configured for the Autonomous VM cluster. If false, database backup on local Exadata storage is not available in the Autonomous VM cluster.
	IsLocalBackupEnabled *bool `mandatory:"false" json:"isLocalBackupEnabled"`

	// The number of enabled CPU cores.
	CpusEnabled *int `mandatory:"false" json:"cpusEnabled"`

	// The number of enabled OCPU cores.
	OcpusEnabled *float32 `mandatory:"false" json:"ocpusEnabled"`

	// The numnber of CPU cores available.
	AvailableCpus *int `mandatory:"false" json:"availableCpus"`

	// The total number of Autonomous Container Databases that can be created.
	TotalContainerDatabases *int `mandatory:"false" json:"totalContainerDatabases"`

	// The amount of memory (in GBs) enabled per each OCPU core.
	MemoryPerOracleComputeUnitInGBs *int `mandatory:"false" json:"memoryPerOracleComputeUnitInGBs"`

	// The number of CPU cores enabled per VM cluster node.
	CpuCoreCountPerNode *int `mandatory:"false" json:"cpuCoreCountPerNode"`

	// The data disk group size allocated for Autonomous Databases, in TBs.
	AutonomousDataStorageSizeInTBs *float64 `mandatory:"false" json:"autonomousDataStorageSizeInTBs"`

	MaintenanceWindow *MaintenanceWindow `mandatory:"false" json:"maintenanceWindow"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the last maintenance run.
	LastMaintenanceRunId *string `mandatory:"false" json:"lastMaintenanceRunId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the next maintenance run.
	NextMaintenanceRunId *string `mandatory:"false" json:"nextMaintenanceRunId"`

	// The memory allocated in GBs.
	MemorySizeInGBs *int `mandatory:"false" json:"memorySizeInGBs"`

	// The local node storage allocated in GBs.
	DbNodeStorageSizeInGBs *int `mandatory:"false" json:"dbNodeStorageSizeInGBs"`

	// The total data storage allocated in TBs
	DataStorageSizeInTBs *float64 `mandatory:"false" json:"dataStorageSizeInTBs"`

	// The total data storage allocated in GBs.
	DataStorageSizeInGBs *float64 `mandatory:"false" json:"dataStorageSizeInGBs"`

	// **Deprecated.** Use `availableAutonomousDataStorageSizeInTBs` for Autonomous Databases' data storage availability in TBs.
	AvailableDataStorageSizeInTBs *float64 `mandatory:"false" json:"availableDataStorageSizeInTBs"`

	// The Oracle license model that applies to the Autonomous VM cluster. The default is LICENSE_INCLUDED.
	LicenseModel AutonomousVmClusterLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// CPU cores that continue to be included in the count of OCPUs available to the Autonomous Container Database even after one of its Autonomous Database is terminated or scaled down. You can release them to the available OCPUs at its parent AVMC level by restarting the Autonomous Container Database.
	ReclaimableCpus *int `mandatory:"false" json:"reclaimableCpus"`

	// The number of Autonomous Container Databases that can be created with the currently available local storage.
	AvailableContainerDatabases *int `mandatory:"false" json:"availableContainerDatabases"`

	// The data disk group size available for Autonomous Databases, in TBs.
	AvailableAutonomousDataStorageSizeInTBs *float64 `mandatory:"false" json:"availableAutonomousDataStorageSizeInTBs"`

	// The SCAN Listener TLS port number. Default value is 2484.
	ScanListenerPortTls *int `mandatory:"false" json:"scanListenerPortTls"`

	// The SCAN Listener Non TLS port number. Default value is 1521.
	ScanListenerPortNonTls *int `mandatory:"false" json:"scanListenerPortNonTls"`

	// Enable mutual TLS(mTLS) authentication for database while provisioning a VMCluster. Default is TLS.
	IsMtlsEnabled *bool `mandatory:"false" json:"isMtlsEnabled"`
}

func (m AutonomousVmCluster) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousVmCluster) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutonomousVmClusterLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAutonomousVmClusterLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingAutonomousVmClusterLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetAutonomousVmClusterLicenseModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutonomousVmClusterLifecycleStateEnum Enum with underlying type: string
type AutonomousVmClusterLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousVmClusterLifecycleStateEnum
const (
	AutonomousVmClusterLifecycleStateProvisioning          AutonomousVmClusterLifecycleStateEnum = "PROVISIONING"
	AutonomousVmClusterLifecycleStateAvailable             AutonomousVmClusterLifecycleStateEnum = "AVAILABLE"
	AutonomousVmClusterLifecycleStateUpdating              AutonomousVmClusterLifecycleStateEnum = "UPDATING"
	AutonomousVmClusterLifecycleStateTerminating           AutonomousVmClusterLifecycleStateEnum = "TERMINATING"
	AutonomousVmClusterLifecycleStateTerminated            AutonomousVmClusterLifecycleStateEnum = "TERMINATED"
	AutonomousVmClusterLifecycleStateFailed                AutonomousVmClusterLifecycleStateEnum = "FAILED"
	AutonomousVmClusterLifecycleStateMaintenanceInProgress AutonomousVmClusterLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
)

var mappingAutonomousVmClusterLifecycleStateEnum = map[string]AutonomousVmClusterLifecycleStateEnum{
	"PROVISIONING":            AutonomousVmClusterLifecycleStateProvisioning,
	"AVAILABLE":               AutonomousVmClusterLifecycleStateAvailable,
	"UPDATING":                AutonomousVmClusterLifecycleStateUpdating,
	"TERMINATING":             AutonomousVmClusterLifecycleStateTerminating,
	"TERMINATED":              AutonomousVmClusterLifecycleStateTerminated,
	"FAILED":                  AutonomousVmClusterLifecycleStateFailed,
	"MAINTENANCE_IN_PROGRESS": AutonomousVmClusterLifecycleStateMaintenanceInProgress,
}

var mappingAutonomousVmClusterLifecycleStateEnumLowerCase = map[string]AutonomousVmClusterLifecycleStateEnum{
	"provisioning":            AutonomousVmClusterLifecycleStateProvisioning,
	"available":               AutonomousVmClusterLifecycleStateAvailable,
	"updating":                AutonomousVmClusterLifecycleStateUpdating,
	"terminating":             AutonomousVmClusterLifecycleStateTerminating,
	"terminated":              AutonomousVmClusterLifecycleStateTerminated,
	"failed":                  AutonomousVmClusterLifecycleStateFailed,
	"maintenance_in_progress": AutonomousVmClusterLifecycleStateMaintenanceInProgress,
}

// GetAutonomousVmClusterLifecycleStateEnumValues Enumerates the set of values for AutonomousVmClusterLifecycleStateEnum
func GetAutonomousVmClusterLifecycleStateEnumValues() []AutonomousVmClusterLifecycleStateEnum {
	values := make([]AutonomousVmClusterLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousVmClusterLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousVmClusterLifecycleStateEnumStringValues Enumerates the set of values in String for AutonomousVmClusterLifecycleStateEnum
func GetAutonomousVmClusterLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"UPDATING",
		"TERMINATING",
		"TERMINATED",
		"FAILED",
		"MAINTENANCE_IN_PROGRESS",
	}
}

// GetMappingAutonomousVmClusterLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousVmClusterLifecycleStateEnum(val string) (AutonomousVmClusterLifecycleStateEnum, bool) {
	enum, ok := mappingAutonomousVmClusterLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousVmClusterLicenseModelEnum Enum with underlying type: string
type AutonomousVmClusterLicenseModelEnum string

// Set of constants representing the allowable values for AutonomousVmClusterLicenseModelEnum
const (
	AutonomousVmClusterLicenseModelLicenseIncluded     AutonomousVmClusterLicenseModelEnum = "LICENSE_INCLUDED"
	AutonomousVmClusterLicenseModelBringYourOwnLicense AutonomousVmClusterLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingAutonomousVmClusterLicenseModelEnum = map[string]AutonomousVmClusterLicenseModelEnum{
	"LICENSE_INCLUDED":       AutonomousVmClusterLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": AutonomousVmClusterLicenseModelBringYourOwnLicense,
}

var mappingAutonomousVmClusterLicenseModelEnumLowerCase = map[string]AutonomousVmClusterLicenseModelEnum{
	"license_included":       AutonomousVmClusterLicenseModelLicenseIncluded,
	"bring_your_own_license": AutonomousVmClusterLicenseModelBringYourOwnLicense,
}

// GetAutonomousVmClusterLicenseModelEnumValues Enumerates the set of values for AutonomousVmClusterLicenseModelEnum
func GetAutonomousVmClusterLicenseModelEnumValues() []AutonomousVmClusterLicenseModelEnum {
	values := make([]AutonomousVmClusterLicenseModelEnum, 0)
	for _, v := range mappingAutonomousVmClusterLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousVmClusterLicenseModelEnumStringValues Enumerates the set of values in String for AutonomousVmClusterLicenseModelEnum
func GetAutonomousVmClusterLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingAutonomousVmClusterLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousVmClusterLicenseModelEnum(val string) (AutonomousVmClusterLicenseModelEnum, bool) {
	enum, ok := mappingAutonomousVmClusterLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
