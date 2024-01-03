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

// AutonomousVmClusterResourceUsage Autonomous VM Cluster usage details, including the Autonomous Container Databases usage.
type AutonomousVmClusterResourceUsage struct {

	// The user-friendly name for the Autonomous VM cluster. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Autonomous VM cluster.
	Id *string `mandatory:"false" json:"id"`

	// The data disk group size allocated for Autonomous Databases, in TBs.
	AutonomousDataStorageSizeInTBs *float32 `mandatory:"false" json:"autonomousDataStorageSizeInTBs"`

	// The local node storage allocated in GBs.
	DbNodeStorageSizeInGBs *int `mandatory:"false" json:"dbNodeStorageSizeInGBs"`

	// The memory allocated in GBs.
	MemorySizeInGBs *int `mandatory:"false" json:"memorySizeInGBs"`

	// The total number of Autonomous Container Databases that can be created.
	TotalContainerDatabases *int `mandatory:"false" json:"totalContainerDatabases"`

	// The data disk group size available for Autonomous Databases, in TBs.
	AvailableAutonomousDataStorageSizeInTBs *float32 `mandatory:"false" json:"availableAutonomousDataStorageSizeInTBs"`

	// The data disk group size used for Autonomous Databases, in TBs.
	UsedAutonomousDataStorageSizeInTBs *float32 `mandatory:"false" json:"usedAutonomousDataStorageSizeInTBs"`

	// If true, database backup on local Exadata storage is configured for the Autonomous VM cluster. If false, database backup on local Exadata storage is not available in the Autonomous VM cluster.
	IsLocalBackupEnabled *bool `mandatory:"false" json:"isLocalBackupEnabled"`

	// Total exadata storage allocated for the Autonomous VM Cluster. DATA + RECOVERY + SPARSE + any overhead in TBs.
	ExadataStorageInTBs *float64 `mandatory:"false" json:"exadataStorageInTBs"`

	// The amount of memory (in GBs) to be enabled per each CPU core.
	MemoryPerOracleComputeUnitInGBs *int `mandatory:"false" json:"memoryPerOracleComputeUnitInGBs"`

	// The number of CPU cores enabled on the Autonomous VM cluster.
	TotalCpus *float32 `mandatory:"false" json:"totalCpus"`

	// The number of CPU cores alloted to the Autonomous Container Databases in an Autonomous VM cluster.
	UsedCpus *float32 `mandatory:"false" json:"usedCpus"`

	// The number of CPU cores available.
	AvailableCpus *float32 `mandatory:"false" json:"availableCpus"`

	// CPU cores that continue to be included in the count of OCPUs available to the
	// Autonomous Container Database even after one of its Autonomous Database is terminated or scaled down.
	// You can release them to the available OCPUs at its parent AVMC level by restarting the Autonomous Container Database.
	ReclaimableCpus *float32 `mandatory:"false" json:"reclaimableCpus"`

	// The number of CPUs provisioned in an Autonomous VM Cluster.
	ProvisionedCpus *float32 `mandatory:"false" json:"provisionedCpus"`

	// The number of CPUs reserved in an Autonomous VM Cluster.
	ReservedCpus *float32 `mandatory:"false" json:"reservedCpus"`

	// The number of provisionable Autonomous Container Databases in an Autonomous VM Cluster.
	ProvisionableAutonomousContainerDatabases *int `mandatory:"false" json:"provisionableAutonomousContainerDatabases"`

	// The number of provisioned Autonomous Container Databases in an Autonomous VM Cluster.
	ProvisionedAutonomousContainerDatabases *int `mandatory:"false" json:"provisionedAutonomousContainerDatabases"`

	// The number of non-provisionable Autonomous Container Databases in an Autonomous VM Cluster.
	NonProvisionableAutonomousContainerDatabases *int `mandatory:"false" json:"nonProvisionableAutonomousContainerDatabases"`

	// List of autonomous vm cluster resource usages.
	AutonomousVmResourceUsage []AutonomousVmResourceUsage `mandatory:"false" json:"autonomousVmResourceUsage"`
}

func (m AutonomousVmClusterResourceUsage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousVmClusterResourceUsage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
