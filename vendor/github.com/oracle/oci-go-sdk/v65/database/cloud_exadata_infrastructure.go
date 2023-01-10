// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// CloudExadataInfrastructure Details of the cloud Exadata infrastructure resource. Applies to Exadata Cloud Service instances only.
type CloudExadataInfrastructure struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the cloud Exadata infrastructure resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current lifecycle state of the cloud Exadata infrastructure resource.
	LifecycleState CloudExadataInfrastructureLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The user-friendly name for the cloud Exadata infrastructure resource. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The model name of the cloud Exadata infrastructure resource.
	Shape *string `mandatory:"true" json:"shape"`

	// The name of the availability domain that the cloud Exadata infrastructure resource is located in.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The number of compute servers for the cloud Exadata infrastructure.
	ComputeCount *int `mandatory:"false" json:"computeCount"`

	// The number of storage servers for the cloud Exadata infrastructure.
	StorageCount *int `mandatory:"false" json:"storageCount"`

	// The total storage allocated to the cloud Exadata infrastructure resource, in gigabytes (GB).
	TotalStorageSizeInGBs *int `mandatory:"false" json:"totalStorageSizeInGBs"`

	// The available storage can be allocated to the cloud Exadata infrastructure resource, in gigabytes (GB).
	AvailableStorageSizeInGBs *int `mandatory:"false" json:"availableStorageSizeInGBs"`

	// The total number of CPU cores allocated.
	CpuCount *int `mandatory:"false" json:"cpuCount"`

	// The total number of CPU cores available.
	MaxCpuCount *int `mandatory:"false" json:"maxCpuCount"`

	// The memory allocated in GBs.
	MemorySizeInGBs *int `mandatory:"false" json:"memorySizeInGBs"`

	// The total memory available in GBs.
	MaxMemoryInGBs *int `mandatory:"false" json:"maxMemoryInGBs"`

	// The local node storage allocated in GBs.
	DbNodeStorageSizeInGBs *int `mandatory:"false" json:"dbNodeStorageSizeInGBs"`

	// The total local node storage available in GBs.
	MaxDbNodeStorageInGBs *int `mandatory:"false" json:"maxDbNodeStorageInGBs"`

	// Size, in terabytes, of the DATA disk group.
	DataStorageSizeInTBs *float64 `mandatory:"false" json:"dataStorageSizeInTBs"`

	// The total available DATA disk group size.
	MaxDataStorageInTBs *float64 `mandatory:"false" json:"maxDataStorageInTBs"`

	// The requested number of additional storage servers for the Exadata infrastructure.
	AdditionalStorageCount *int `mandatory:"false" json:"additionalStorageCount"`

	// The requested number of additional storage servers activated for the Exadata infrastructure.
	ActivatedStorageCount *int `mandatory:"false" json:"activatedStorageCount"`

	// The date and time the cloud Exadata infrastructure resource was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	MaintenanceWindow *MaintenanceWindow `mandatory:"false" json:"maintenanceWindow"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the last maintenance run.
	LastMaintenanceRunId *string `mandatory:"false" json:"lastMaintenanceRunId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the next maintenance run.
	NextMaintenanceRunId *string `mandatory:"false" json:"nextMaintenanceRunId"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The list of customer email addresses that receive information from Oracle about the specified OCI Database service resource.
	// Oracle uses these email addresses to send notifications about planned and unplanned software maintenance updates, information about system hardware, and other information needed by administrators.
	// Up to 10 email addresses can be added to the customer contacts for a cloud Exadata infrastructure instance.
	CustomerContacts []CustomerContact `mandatory:"false" json:"customerContacts"`
}

func (m CloudExadataInfrastructure) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CloudExadataInfrastructure) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCloudExadataInfrastructureLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCloudExadataInfrastructureLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CloudExadataInfrastructureLifecycleStateEnum Enum with underlying type: string
type CloudExadataInfrastructureLifecycleStateEnum string

// Set of constants representing the allowable values for CloudExadataInfrastructureLifecycleStateEnum
const (
	CloudExadataInfrastructureLifecycleStateProvisioning          CloudExadataInfrastructureLifecycleStateEnum = "PROVISIONING"
	CloudExadataInfrastructureLifecycleStateAvailable             CloudExadataInfrastructureLifecycleStateEnum = "AVAILABLE"
	CloudExadataInfrastructureLifecycleStateUpdating              CloudExadataInfrastructureLifecycleStateEnum = "UPDATING"
	CloudExadataInfrastructureLifecycleStateTerminating           CloudExadataInfrastructureLifecycleStateEnum = "TERMINATING"
	CloudExadataInfrastructureLifecycleStateTerminated            CloudExadataInfrastructureLifecycleStateEnum = "TERMINATED"
	CloudExadataInfrastructureLifecycleStateFailed                CloudExadataInfrastructureLifecycleStateEnum = "FAILED"
	CloudExadataInfrastructureLifecycleStateMaintenanceInProgress CloudExadataInfrastructureLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
)

var mappingCloudExadataInfrastructureLifecycleStateEnum = map[string]CloudExadataInfrastructureLifecycleStateEnum{
	"PROVISIONING":            CloudExadataInfrastructureLifecycleStateProvisioning,
	"AVAILABLE":               CloudExadataInfrastructureLifecycleStateAvailable,
	"UPDATING":                CloudExadataInfrastructureLifecycleStateUpdating,
	"TERMINATING":             CloudExadataInfrastructureLifecycleStateTerminating,
	"TERMINATED":              CloudExadataInfrastructureLifecycleStateTerminated,
	"FAILED":                  CloudExadataInfrastructureLifecycleStateFailed,
	"MAINTENANCE_IN_PROGRESS": CloudExadataInfrastructureLifecycleStateMaintenanceInProgress,
}

var mappingCloudExadataInfrastructureLifecycleStateEnumLowerCase = map[string]CloudExadataInfrastructureLifecycleStateEnum{
	"provisioning":            CloudExadataInfrastructureLifecycleStateProvisioning,
	"available":               CloudExadataInfrastructureLifecycleStateAvailable,
	"updating":                CloudExadataInfrastructureLifecycleStateUpdating,
	"terminating":             CloudExadataInfrastructureLifecycleStateTerminating,
	"terminated":              CloudExadataInfrastructureLifecycleStateTerminated,
	"failed":                  CloudExadataInfrastructureLifecycleStateFailed,
	"maintenance_in_progress": CloudExadataInfrastructureLifecycleStateMaintenanceInProgress,
}

// GetCloudExadataInfrastructureLifecycleStateEnumValues Enumerates the set of values for CloudExadataInfrastructureLifecycleStateEnum
func GetCloudExadataInfrastructureLifecycleStateEnumValues() []CloudExadataInfrastructureLifecycleStateEnum {
	values := make([]CloudExadataInfrastructureLifecycleStateEnum, 0)
	for _, v := range mappingCloudExadataInfrastructureLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudExadataInfrastructureLifecycleStateEnumStringValues Enumerates the set of values in String for CloudExadataInfrastructureLifecycleStateEnum
func GetCloudExadataInfrastructureLifecycleStateEnumStringValues() []string {
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

// GetMappingCloudExadataInfrastructureLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudExadataInfrastructureLifecycleStateEnum(val string) (CloudExadataInfrastructureLifecycleStateEnum, bool) {
	enum, ok := mappingCloudExadataInfrastructureLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
