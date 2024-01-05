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

// ExadataInfrastructure ExadataInfrastructure
type ExadataInfrastructure struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current lifecycle state of the Exadata infrastructure.
	LifecycleState ExadataInfrastructureLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The user-friendly name for the Exadata Cloud@Customer infrastructure. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The shape of the Exadata infrastructure. The shape determines the amount of CPU, storage, and memory resources allocated to the instance.
	Shape *string `mandatory:"true" json:"shape"`

	// The time zone of the Exadata infrastructure. For details, see Exadata Infrastructure Time Zones (https://docs.cloud.oracle.com/Content/Database/References/timezones.htm).
	TimeZone *string `mandatory:"false" json:"timeZone"`

	// The number of enabled CPU cores.
	CpusEnabled *int `mandatory:"false" json:"cpusEnabled"`

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

	// The serial number for the Exadata infrastructure.
	RackSerialNumber *string `mandatory:"false" json:"rackSerialNumber"`

	// The number of Exadata storage servers for the Exadata infrastructure.
	StorageCount *int `mandatory:"false" json:"storageCount"`

	// The requested number of additional storage servers for the Exadata infrastructure.
	AdditionalStorageCount *int `mandatory:"false" json:"additionalStorageCount"`

	// The requested number of additional storage servers activated for the Exadata infrastructure.
	ActivatedStorageCount *int `mandatory:"false" json:"activatedStorageCount"`

	// The number of compute servers for the Exadata infrastructure.
	ComputeCount *int `mandatory:"false" json:"computeCount"`

	// Indicates if deployment is Multi-Rack or not.
	IsMultiRackDeployment *bool `mandatory:"false" json:"isMultiRackDeployment"`

	// The base64 encoded Multi-Rack configuration json file.
	MultiRackConfigurationFile []byte `mandatory:"false" json:"multiRackConfigurationFile"`

	// The requested number of additional compute servers for the Exadata infrastructure.
	AdditionalComputeCount *int `mandatory:"false" json:"additionalComputeCount"`

	// Oracle Exadata System Model specification. The system model determines the amount of compute or storage
	// server resources available for use. For more information, please see System and Shape Configuration Options
	//  (https://docs.oracle.com/en/engineered-systems/exadata-cloud-at-customer/ecccm/ecc-system-config-options.html#GUID-9E090174-5C57-4EB1-9243-B470F9F10D6B)
	AdditionalComputeSystemModel ExadataInfrastructureAdditionalComputeSystemModelEnum `mandatory:"false" json:"additionalComputeSystemModel,omitempty"`

	// The IP address for the first control plane server.
	CloudControlPlaneServer1 *string `mandatory:"false" json:"cloudControlPlaneServer1"`

	// The IP address for the second control plane server.
	CloudControlPlaneServer2 *string `mandatory:"false" json:"cloudControlPlaneServer2"`

	// The netmask for the control plane network.
	Netmask *string `mandatory:"false" json:"netmask"`

	// The gateway for the control plane network.
	Gateway *string `mandatory:"false" json:"gateway"`

	// The CIDR block for the Exadata administration network.
	AdminNetworkCIDR *string `mandatory:"false" json:"adminNetworkCIDR"`

	// The CIDR block for the Exadata InfiniBand interconnect.
	InfiniBandNetworkCIDR *string `mandatory:"false" json:"infiniBandNetworkCIDR"`

	// The corporate network proxy for access to the control plane network.
	CorporateProxy *string `mandatory:"false" json:"corporateProxy"`

	// The list of DNS server IP addresses. Maximum of 3 allowed.
	DnsServer []string `mandatory:"false" json:"dnsServer"`

	// The list of NTP server IP addresses. Maximum of 3 allowed.
	NtpServer []string `mandatory:"false" json:"ntpServer"`

	// The date and time the Exadata infrastructure was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The CSI Number of the Exadata infrastructure.
	CsiNumber *string `mandatory:"false" json:"csiNumber"`

	// The list of contacts for the Exadata infrastructure.
	Contacts []ExadataInfrastructureContact `mandatory:"false" json:"contacts"`

	// A field to capture ‘Maintenance SLO Status’ for the Exadata infrastructure with values ‘OK’, ‘DEGRADED’. Default is ‘OK’ when the infrastructure is provisioned.
	MaintenanceSLOStatus ExadataInfrastructureMaintenanceSLOStatusEnum `mandatory:"false" json:"maintenanceSLOStatus,omitempty"`

	MaintenanceWindow *MaintenanceWindow `mandatory:"false" json:"maintenanceWindow"`

	// The software version of the storage servers (cells) in the Exadata infrastructure.
	StorageServerVersion *string `mandatory:"false" json:"storageServerVersion"`

	// The software version of the database servers (dom0) in the Exadata infrastructure.
	DbServerVersion *string `mandatory:"false" json:"dbServerVersion"`

	// The monthly software version of the database servers (dom0) in the Exadata infrastructure.
	MonthlyDbServerVersion *string `mandatory:"false" json:"monthlyDbServerVersion"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the last maintenance run.
	LastMaintenanceRunId *string `mandatory:"false" json:"lastMaintenanceRunId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the next maintenance run.
	NextMaintenanceRunId *string `mandatory:"false" json:"nextMaintenanceRunId"`

	// Indicates whether cps offline diagnostic report is enabled for this Exadata infrastructure. This will allow a customer to quickly check status themselves and fix problems on their end, saving time and frustration
	// for both Oracle and the customer when they find the CPS in a disconnected state.You can enable offline diagnostic report during Exadata infrastructure provisioning. You can also disable or enable it at any time
	// using the UpdateExadatainfrastructure API.
	IsCpsOfflineReportEnabled *bool `mandatory:"false" json:"isCpsOfflineReportEnabled"`

	NetworkBondingModeDetails *NetworkBondingModeDetails `mandatory:"false" json:"networkBondingModeDetails"`

	// The name of the availability domain that the Exadata infrastructure is located in.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ExadataInfrastructure) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExadataInfrastructure) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExadataInfrastructureLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExadataInfrastructureLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingExadataInfrastructureAdditionalComputeSystemModelEnum(string(m.AdditionalComputeSystemModel)); !ok && m.AdditionalComputeSystemModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AdditionalComputeSystemModel: %s. Supported values are: %s.", m.AdditionalComputeSystemModel, strings.Join(GetExadataInfrastructureAdditionalComputeSystemModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExadataInfrastructureMaintenanceSLOStatusEnum(string(m.MaintenanceSLOStatus)); !ok && m.MaintenanceSLOStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MaintenanceSLOStatus: %s. Supported values are: %s.", m.MaintenanceSLOStatus, strings.Join(GetExadataInfrastructureMaintenanceSLOStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExadataInfrastructureLifecycleStateEnum Enum with underlying type: string
type ExadataInfrastructureLifecycleStateEnum string

// Set of constants representing the allowable values for ExadataInfrastructureLifecycleStateEnum
const (
	ExadataInfrastructureLifecycleStateCreating               ExadataInfrastructureLifecycleStateEnum = "CREATING"
	ExadataInfrastructureLifecycleStateRequiresActivation     ExadataInfrastructureLifecycleStateEnum = "REQUIRES_ACTIVATION"
	ExadataInfrastructureLifecycleStateActivating             ExadataInfrastructureLifecycleStateEnum = "ACTIVATING"
	ExadataInfrastructureLifecycleStateActive                 ExadataInfrastructureLifecycleStateEnum = "ACTIVE"
	ExadataInfrastructureLifecycleStateActivationFailed       ExadataInfrastructureLifecycleStateEnum = "ACTIVATION_FAILED"
	ExadataInfrastructureLifecycleStateFailed                 ExadataInfrastructureLifecycleStateEnum = "FAILED"
	ExadataInfrastructureLifecycleStateUpdating               ExadataInfrastructureLifecycleStateEnum = "UPDATING"
	ExadataInfrastructureLifecycleStateDeleting               ExadataInfrastructureLifecycleStateEnum = "DELETING"
	ExadataInfrastructureLifecycleStateDeleted                ExadataInfrastructureLifecycleStateEnum = "DELETED"
	ExadataInfrastructureLifecycleStateDisconnected           ExadataInfrastructureLifecycleStateEnum = "DISCONNECTED"
	ExadataInfrastructureLifecycleStateMaintenanceInProgress  ExadataInfrastructureLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
	ExadataInfrastructureLifecycleStateWaitingForConnectivity ExadataInfrastructureLifecycleStateEnum = "WAITING_FOR_CONNECTIVITY"
)

var mappingExadataInfrastructureLifecycleStateEnum = map[string]ExadataInfrastructureLifecycleStateEnum{
	"CREATING":                 ExadataInfrastructureLifecycleStateCreating,
	"REQUIRES_ACTIVATION":      ExadataInfrastructureLifecycleStateRequiresActivation,
	"ACTIVATING":               ExadataInfrastructureLifecycleStateActivating,
	"ACTIVE":                   ExadataInfrastructureLifecycleStateActive,
	"ACTIVATION_FAILED":        ExadataInfrastructureLifecycleStateActivationFailed,
	"FAILED":                   ExadataInfrastructureLifecycleStateFailed,
	"UPDATING":                 ExadataInfrastructureLifecycleStateUpdating,
	"DELETING":                 ExadataInfrastructureLifecycleStateDeleting,
	"DELETED":                  ExadataInfrastructureLifecycleStateDeleted,
	"DISCONNECTED":             ExadataInfrastructureLifecycleStateDisconnected,
	"MAINTENANCE_IN_PROGRESS":  ExadataInfrastructureLifecycleStateMaintenanceInProgress,
	"WAITING_FOR_CONNECTIVITY": ExadataInfrastructureLifecycleStateWaitingForConnectivity,
}

var mappingExadataInfrastructureLifecycleStateEnumLowerCase = map[string]ExadataInfrastructureLifecycleStateEnum{
	"creating":                 ExadataInfrastructureLifecycleStateCreating,
	"requires_activation":      ExadataInfrastructureLifecycleStateRequiresActivation,
	"activating":               ExadataInfrastructureLifecycleStateActivating,
	"active":                   ExadataInfrastructureLifecycleStateActive,
	"activation_failed":        ExadataInfrastructureLifecycleStateActivationFailed,
	"failed":                   ExadataInfrastructureLifecycleStateFailed,
	"updating":                 ExadataInfrastructureLifecycleStateUpdating,
	"deleting":                 ExadataInfrastructureLifecycleStateDeleting,
	"deleted":                  ExadataInfrastructureLifecycleStateDeleted,
	"disconnected":             ExadataInfrastructureLifecycleStateDisconnected,
	"maintenance_in_progress":  ExadataInfrastructureLifecycleStateMaintenanceInProgress,
	"waiting_for_connectivity": ExadataInfrastructureLifecycleStateWaitingForConnectivity,
}

// GetExadataInfrastructureLifecycleStateEnumValues Enumerates the set of values for ExadataInfrastructureLifecycleStateEnum
func GetExadataInfrastructureLifecycleStateEnumValues() []ExadataInfrastructureLifecycleStateEnum {
	values := make([]ExadataInfrastructureLifecycleStateEnum, 0)
	for _, v := range mappingExadataInfrastructureLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExadataInfrastructureLifecycleStateEnumStringValues Enumerates the set of values in String for ExadataInfrastructureLifecycleStateEnum
func GetExadataInfrastructureLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"REQUIRES_ACTIVATION",
		"ACTIVATING",
		"ACTIVE",
		"ACTIVATION_FAILED",
		"FAILED",
		"UPDATING",
		"DELETING",
		"DELETED",
		"DISCONNECTED",
		"MAINTENANCE_IN_PROGRESS",
		"WAITING_FOR_CONNECTIVITY",
	}
}

// GetMappingExadataInfrastructureLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadataInfrastructureLifecycleStateEnum(val string) (ExadataInfrastructureLifecycleStateEnum, bool) {
	enum, ok := mappingExadataInfrastructureLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExadataInfrastructureAdditionalComputeSystemModelEnum Enum with underlying type: string
type ExadataInfrastructureAdditionalComputeSystemModelEnum string

// Set of constants representing the allowable values for ExadataInfrastructureAdditionalComputeSystemModelEnum
const (
	ExadataInfrastructureAdditionalComputeSystemModelX7   ExadataInfrastructureAdditionalComputeSystemModelEnum = "X7"
	ExadataInfrastructureAdditionalComputeSystemModelX8   ExadataInfrastructureAdditionalComputeSystemModelEnum = "X8"
	ExadataInfrastructureAdditionalComputeSystemModelX8m  ExadataInfrastructureAdditionalComputeSystemModelEnum = "X8M"
	ExadataInfrastructureAdditionalComputeSystemModelX9m  ExadataInfrastructureAdditionalComputeSystemModelEnum = "X9M"
	ExadataInfrastructureAdditionalComputeSystemModelX10m ExadataInfrastructureAdditionalComputeSystemModelEnum = "X10M"
)

var mappingExadataInfrastructureAdditionalComputeSystemModelEnum = map[string]ExadataInfrastructureAdditionalComputeSystemModelEnum{
	"X7":   ExadataInfrastructureAdditionalComputeSystemModelX7,
	"X8":   ExadataInfrastructureAdditionalComputeSystemModelX8,
	"X8M":  ExadataInfrastructureAdditionalComputeSystemModelX8m,
	"X9M":  ExadataInfrastructureAdditionalComputeSystemModelX9m,
	"X10M": ExadataInfrastructureAdditionalComputeSystemModelX10m,
}

var mappingExadataInfrastructureAdditionalComputeSystemModelEnumLowerCase = map[string]ExadataInfrastructureAdditionalComputeSystemModelEnum{
	"x7":   ExadataInfrastructureAdditionalComputeSystemModelX7,
	"x8":   ExadataInfrastructureAdditionalComputeSystemModelX8,
	"x8m":  ExadataInfrastructureAdditionalComputeSystemModelX8m,
	"x9m":  ExadataInfrastructureAdditionalComputeSystemModelX9m,
	"x10m": ExadataInfrastructureAdditionalComputeSystemModelX10m,
}

// GetExadataInfrastructureAdditionalComputeSystemModelEnumValues Enumerates the set of values for ExadataInfrastructureAdditionalComputeSystemModelEnum
func GetExadataInfrastructureAdditionalComputeSystemModelEnumValues() []ExadataInfrastructureAdditionalComputeSystemModelEnum {
	values := make([]ExadataInfrastructureAdditionalComputeSystemModelEnum, 0)
	for _, v := range mappingExadataInfrastructureAdditionalComputeSystemModelEnum {
		values = append(values, v)
	}
	return values
}

// GetExadataInfrastructureAdditionalComputeSystemModelEnumStringValues Enumerates the set of values in String for ExadataInfrastructureAdditionalComputeSystemModelEnum
func GetExadataInfrastructureAdditionalComputeSystemModelEnumStringValues() []string {
	return []string{
		"X7",
		"X8",
		"X8M",
		"X9M",
		"X10M",
	}
}

// GetMappingExadataInfrastructureAdditionalComputeSystemModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadataInfrastructureAdditionalComputeSystemModelEnum(val string) (ExadataInfrastructureAdditionalComputeSystemModelEnum, bool) {
	enum, ok := mappingExadataInfrastructureAdditionalComputeSystemModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExadataInfrastructureMaintenanceSLOStatusEnum Enum with underlying type: string
type ExadataInfrastructureMaintenanceSLOStatusEnum string

// Set of constants representing the allowable values for ExadataInfrastructureMaintenanceSLOStatusEnum
const (
	ExadataInfrastructureMaintenanceSLOStatusOk       ExadataInfrastructureMaintenanceSLOStatusEnum = "OK"
	ExadataInfrastructureMaintenanceSLOStatusDegraded ExadataInfrastructureMaintenanceSLOStatusEnum = "DEGRADED"
)

var mappingExadataInfrastructureMaintenanceSLOStatusEnum = map[string]ExadataInfrastructureMaintenanceSLOStatusEnum{
	"OK":       ExadataInfrastructureMaintenanceSLOStatusOk,
	"DEGRADED": ExadataInfrastructureMaintenanceSLOStatusDegraded,
}

var mappingExadataInfrastructureMaintenanceSLOStatusEnumLowerCase = map[string]ExadataInfrastructureMaintenanceSLOStatusEnum{
	"ok":       ExadataInfrastructureMaintenanceSLOStatusOk,
	"degraded": ExadataInfrastructureMaintenanceSLOStatusDegraded,
}

// GetExadataInfrastructureMaintenanceSLOStatusEnumValues Enumerates the set of values for ExadataInfrastructureMaintenanceSLOStatusEnum
func GetExadataInfrastructureMaintenanceSLOStatusEnumValues() []ExadataInfrastructureMaintenanceSLOStatusEnum {
	values := make([]ExadataInfrastructureMaintenanceSLOStatusEnum, 0)
	for _, v := range mappingExadataInfrastructureMaintenanceSLOStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetExadataInfrastructureMaintenanceSLOStatusEnumStringValues Enumerates the set of values in String for ExadataInfrastructureMaintenanceSLOStatusEnum
func GetExadataInfrastructureMaintenanceSLOStatusEnumStringValues() []string {
	return []string{
		"OK",
		"DEGRADED",
	}
}

// GetMappingExadataInfrastructureMaintenanceSLOStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadataInfrastructureMaintenanceSLOStatusEnum(val string) (ExadataInfrastructureMaintenanceSLOStatusEnum, bool) {
	enum, ok := mappingExadataInfrastructureMaintenanceSLOStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
