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

// CreateCloudAutonomousVmClusterDetails Details for the create cloud Autonomous VM cluster operation.
type CreateCloudAutonomousVmClusterDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet the cloud Autonomous VM Cluster is associated with.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The user-friendly name for the cloud Autonomous VM cluster. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the cloud Exadata infrastructure.
	CloudExadataInfrastructureId *string `mandatory:"true" json:"cloudExadataInfrastructureId"`

	// User defined description of the cloud Autonomous VM cluster.
	Description *string `mandatory:"false" json:"description"`

	// The total number of Autonomous Container Databases that can be created.
	TotalContainerDatabases *int `mandatory:"false" json:"totalContainerDatabases"`

	// The number of CPU cores to be enabled per VM cluster node.
	CpuCoreCountPerNode *int `mandatory:"false" json:"cpuCoreCountPerNode"`

	// The amount of memory (in GBs) to be enabled per OCPU or ECPU.
	MemoryPerOracleComputeUnitInGBs *int `mandatory:"false" json:"memoryPerOracleComputeUnitInGBs"`

	// The data disk group size to be allocated for Autonomous Databases, in TBs.
	AutonomousDataStorageSizeInTBs *float64 `mandatory:"false" json:"autonomousDataStorageSizeInTBs"`

	// The time zone to use for the Cloud Autonomous VM cluster. For details, see DB System Time Zones (https://docs.cloud.oracle.com/Content/Database/References/timezones.htm).
	ClusterTimeZone *string `mandatory:"false" json:"clusterTimeZone"`

	// The compute model of the Cloud Autonomous VM Cluster. ECPU compute model is the recommended model and OCPU compute model is legacy.
	ComputeModel CreateCloudAutonomousVmClusterDetailsComputeModelEnum `mandatory:"false" json:"computeModel,omitempty"`

	// Enable mutual TLS(mTLS) authentication for database at time of provisioning a VMCluster. This is applicable to database TLS Certificates only. Default is TLS
	IsMtlsEnabledVmCluster *bool `mandatory:"false" json:"isMtlsEnabledVmCluster"`

	// The list of database servers.
	DbServers []string `mandatory:"false" json:"dbServers"`

	MaintenanceWindowDetails *MaintenanceWindow `mandatory:"false" json:"maintenanceWindowDetails"`

	// The SCAN Listener TLS port. Default is 2484.
	ScanListenerPortTls *int `mandatory:"false" json:"scanListenerPortTls"`

	// The SCAN Listener Non TLS port. Default is 1521.
	ScanListenerPortNonTls *int `mandatory:"false" json:"scanListenerPortNonTls"`

	// The Oracle license model that applies to the Oracle Autonomous Database. Bring your own license (BYOL) allows you to apply your current on-premises Oracle software licenses to equivalent, highly automated Oracle services in the cloud.
	// License Included allows you to subscribe to new Oracle Database software licenses and the Oracle Database service.
	// Note that when provisioning an Autonomous Database on dedicated Exadata infrastructure (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html), this attribute must be null. It is already set at the
	// Autonomous Exadata Infrastructure level. When provisioning an Autonomous Database Serverless  (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html) database, if a value is not specified, the system defaults the value to `BRING_YOUR_OWN_LICENSE`. Bring your own license (BYOL) also allows you to select the DB edition using the optional parameter.
	// This cannot be updated in parallel with any of the following: cpuCoreCount, computeCount, dataStorageSizeInTBs, adminPassword, isMTLSConnectionRequired, dbWorkload, privateEndpointLabel, nsgIds, dbVersion, dbName, scheduledOperations, dbToolsDetails, or isFreeTier.
	LicenseModel CreateCloudAutonomousVmClusterDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The list of OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs. Setting this to an empty list removes all resources from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm).
	// **NsgIds restrictions:**
	// - A network security group (NSG) is optional for Autonomous Databases with private access. The nsgIds list can be empty.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateCloudAutonomousVmClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateCloudAutonomousVmClusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateCloudAutonomousVmClusterDetailsComputeModelEnum(string(m.ComputeModel)); !ok && m.ComputeModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ComputeModel: %s. Supported values are: %s.", m.ComputeModel, strings.Join(GetCreateCloudAutonomousVmClusterDetailsComputeModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateCloudAutonomousVmClusterDetailsLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetCreateCloudAutonomousVmClusterDetailsLicenseModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateCloudAutonomousVmClusterDetailsComputeModelEnum Enum with underlying type: string
type CreateCloudAutonomousVmClusterDetailsComputeModelEnum string

// Set of constants representing the allowable values for CreateCloudAutonomousVmClusterDetailsComputeModelEnum
const (
	CreateCloudAutonomousVmClusterDetailsComputeModelEcpu CreateCloudAutonomousVmClusterDetailsComputeModelEnum = "ECPU"
	CreateCloudAutonomousVmClusterDetailsComputeModelOcpu CreateCloudAutonomousVmClusterDetailsComputeModelEnum = "OCPU"
)

var mappingCreateCloudAutonomousVmClusterDetailsComputeModelEnum = map[string]CreateCloudAutonomousVmClusterDetailsComputeModelEnum{
	"ECPU": CreateCloudAutonomousVmClusterDetailsComputeModelEcpu,
	"OCPU": CreateCloudAutonomousVmClusterDetailsComputeModelOcpu,
}

var mappingCreateCloudAutonomousVmClusterDetailsComputeModelEnumLowerCase = map[string]CreateCloudAutonomousVmClusterDetailsComputeModelEnum{
	"ecpu": CreateCloudAutonomousVmClusterDetailsComputeModelEcpu,
	"ocpu": CreateCloudAutonomousVmClusterDetailsComputeModelOcpu,
}

// GetCreateCloudAutonomousVmClusterDetailsComputeModelEnumValues Enumerates the set of values for CreateCloudAutonomousVmClusterDetailsComputeModelEnum
func GetCreateCloudAutonomousVmClusterDetailsComputeModelEnumValues() []CreateCloudAutonomousVmClusterDetailsComputeModelEnum {
	values := make([]CreateCloudAutonomousVmClusterDetailsComputeModelEnum, 0)
	for _, v := range mappingCreateCloudAutonomousVmClusterDetailsComputeModelEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateCloudAutonomousVmClusterDetailsComputeModelEnumStringValues Enumerates the set of values in String for CreateCloudAutonomousVmClusterDetailsComputeModelEnum
func GetCreateCloudAutonomousVmClusterDetailsComputeModelEnumStringValues() []string {
	return []string{
		"ECPU",
		"OCPU",
	}
}

// GetMappingCreateCloudAutonomousVmClusterDetailsComputeModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateCloudAutonomousVmClusterDetailsComputeModelEnum(val string) (CreateCloudAutonomousVmClusterDetailsComputeModelEnum, bool) {
	enum, ok := mappingCreateCloudAutonomousVmClusterDetailsComputeModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateCloudAutonomousVmClusterDetailsLicenseModelEnum Enum with underlying type: string
type CreateCloudAutonomousVmClusterDetailsLicenseModelEnum string

// Set of constants representing the allowable values for CreateCloudAutonomousVmClusterDetailsLicenseModelEnum
const (
	CreateCloudAutonomousVmClusterDetailsLicenseModelLicenseIncluded     CreateCloudAutonomousVmClusterDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	CreateCloudAutonomousVmClusterDetailsLicenseModelBringYourOwnLicense CreateCloudAutonomousVmClusterDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingCreateCloudAutonomousVmClusterDetailsLicenseModelEnum = map[string]CreateCloudAutonomousVmClusterDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       CreateCloudAutonomousVmClusterDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": CreateCloudAutonomousVmClusterDetailsLicenseModelBringYourOwnLicense,
}

var mappingCreateCloudAutonomousVmClusterDetailsLicenseModelEnumLowerCase = map[string]CreateCloudAutonomousVmClusterDetailsLicenseModelEnum{
	"license_included":       CreateCloudAutonomousVmClusterDetailsLicenseModelLicenseIncluded,
	"bring_your_own_license": CreateCloudAutonomousVmClusterDetailsLicenseModelBringYourOwnLicense,
}

// GetCreateCloudAutonomousVmClusterDetailsLicenseModelEnumValues Enumerates the set of values for CreateCloudAutonomousVmClusterDetailsLicenseModelEnum
func GetCreateCloudAutonomousVmClusterDetailsLicenseModelEnumValues() []CreateCloudAutonomousVmClusterDetailsLicenseModelEnum {
	values := make([]CreateCloudAutonomousVmClusterDetailsLicenseModelEnum, 0)
	for _, v := range mappingCreateCloudAutonomousVmClusterDetailsLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateCloudAutonomousVmClusterDetailsLicenseModelEnumStringValues Enumerates the set of values in String for CreateCloudAutonomousVmClusterDetailsLicenseModelEnum
func GetCreateCloudAutonomousVmClusterDetailsLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// GetMappingCreateCloudAutonomousVmClusterDetailsLicenseModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateCloudAutonomousVmClusterDetailsLicenseModelEnum(val string) (CreateCloudAutonomousVmClusterDetailsLicenseModelEnum, bool) {
	enum, ok := mappingCreateCloudAutonomousVmClusterDetailsLicenseModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
