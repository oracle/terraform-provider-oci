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

// ExadataInfrastructureSummary Details of the Exadata Cloud@Customer infrastructure. Applies to Exadata Cloud@Customer instances only.
// See CloudExadataInfrastructureSummary for details of the cloud Exadata infrastructure resource used by Exadata Cloud Service instances.
type ExadataInfrastructureSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current lifecycle state of the Exadata infrastructure.
	LifecycleState ExadataInfrastructureSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

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
	AdditionalComputeSystemModel ExadataInfrastructureSummaryAdditionalComputeSystemModelEnum `mandatory:"false" json:"additionalComputeSystemModel,omitempty"`

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
	MaintenanceSLOStatus ExadataInfrastructureSummaryMaintenanceSLOStatusEnum `mandatory:"false" json:"maintenanceSLOStatus,omitempty"`

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

func (m ExadataInfrastructureSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExadataInfrastructureSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExadataInfrastructureSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExadataInfrastructureSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingExadataInfrastructureSummaryAdditionalComputeSystemModelEnum(string(m.AdditionalComputeSystemModel)); !ok && m.AdditionalComputeSystemModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AdditionalComputeSystemModel: %s. Supported values are: %s.", m.AdditionalComputeSystemModel, strings.Join(GetExadataInfrastructureSummaryAdditionalComputeSystemModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExadataInfrastructureSummaryMaintenanceSLOStatusEnum(string(m.MaintenanceSLOStatus)); !ok && m.MaintenanceSLOStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MaintenanceSLOStatus: %s. Supported values are: %s.", m.MaintenanceSLOStatus, strings.Join(GetExadataInfrastructureSummaryMaintenanceSLOStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExadataInfrastructureSummaryLifecycleStateEnum Enum with underlying type: string
type ExadataInfrastructureSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for ExadataInfrastructureSummaryLifecycleStateEnum
const (
	ExadataInfrastructureSummaryLifecycleStateCreating               ExadataInfrastructureSummaryLifecycleStateEnum = "CREATING"
	ExadataInfrastructureSummaryLifecycleStateRequiresActivation     ExadataInfrastructureSummaryLifecycleStateEnum = "REQUIRES_ACTIVATION"
	ExadataInfrastructureSummaryLifecycleStateActivating             ExadataInfrastructureSummaryLifecycleStateEnum = "ACTIVATING"
	ExadataInfrastructureSummaryLifecycleStateActive                 ExadataInfrastructureSummaryLifecycleStateEnum = "ACTIVE"
	ExadataInfrastructureSummaryLifecycleStateActivationFailed       ExadataInfrastructureSummaryLifecycleStateEnum = "ACTIVATION_FAILED"
	ExadataInfrastructureSummaryLifecycleStateFailed                 ExadataInfrastructureSummaryLifecycleStateEnum = "FAILED"
	ExadataInfrastructureSummaryLifecycleStateUpdating               ExadataInfrastructureSummaryLifecycleStateEnum = "UPDATING"
	ExadataInfrastructureSummaryLifecycleStateDeleting               ExadataInfrastructureSummaryLifecycleStateEnum = "DELETING"
	ExadataInfrastructureSummaryLifecycleStateDeleted                ExadataInfrastructureSummaryLifecycleStateEnum = "DELETED"
	ExadataInfrastructureSummaryLifecycleStateDisconnected           ExadataInfrastructureSummaryLifecycleStateEnum = "DISCONNECTED"
	ExadataInfrastructureSummaryLifecycleStateMaintenanceInProgress  ExadataInfrastructureSummaryLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
	ExadataInfrastructureSummaryLifecycleStateWaitingForConnectivity ExadataInfrastructureSummaryLifecycleStateEnum = "WAITING_FOR_CONNECTIVITY"
)

var mappingExadataInfrastructureSummaryLifecycleStateEnum = map[string]ExadataInfrastructureSummaryLifecycleStateEnum{
	"CREATING":                 ExadataInfrastructureSummaryLifecycleStateCreating,
	"REQUIRES_ACTIVATION":      ExadataInfrastructureSummaryLifecycleStateRequiresActivation,
	"ACTIVATING":               ExadataInfrastructureSummaryLifecycleStateActivating,
	"ACTIVE":                   ExadataInfrastructureSummaryLifecycleStateActive,
	"ACTIVATION_FAILED":        ExadataInfrastructureSummaryLifecycleStateActivationFailed,
	"FAILED":                   ExadataInfrastructureSummaryLifecycleStateFailed,
	"UPDATING":                 ExadataInfrastructureSummaryLifecycleStateUpdating,
	"DELETING":                 ExadataInfrastructureSummaryLifecycleStateDeleting,
	"DELETED":                  ExadataInfrastructureSummaryLifecycleStateDeleted,
	"DISCONNECTED":             ExadataInfrastructureSummaryLifecycleStateDisconnected,
	"MAINTENANCE_IN_PROGRESS":  ExadataInfrastructureSummaryLifecycleStateMaintenanceInProgress,
	"WAITING_FOR_CONNECTIVITY": ExadataInfrastructureSummaryLifecycleStateWaitingForConnectivity,
}

var mappingExadataInfrastructureSummaryLifecycleStateEnumLowerCase = map[string]ExadataInfrastructureSummaryLifecycleStateEnum{
	"creating":                 ExadataInfrastructureSummaryLifecycleStateCreating,
	"requires_activation":      ExadataInfrastructureSummaryLifecycleStateRequiresActivation,
	"activating":               ExadataInfrastructureSummaryLifecycleStateActivating,
	"active":                   ExadataInfrastructureSummaryLifecycleStateActive,
	"activation_failed":        ExadataInfrastructureSummaryLifecycleStateActivationFailed,
	"failed":                   ExadataInfrastructureSummaryLifecycleStateFailed,
	"updating":                 ExadataInfrastructureSummaryLifecycleStateUpdating,
	"deleting":                 ExadataInfrastructureSummaryLifecycleStateDeleting,
	"deleted":                  ExadataInfrastructureSummaryLifecycleStateDeleted,
	"disconnected":             ExadataInfrastructureSummaryLifecycleStateDisconnected,
	"maintenance_in_progress":  ExadataInfrastructureSummaryLifecycleStateMaintenanceInProgress,
	"waiting_for_connectivity": ExadataInfrastructureSummaryLifecycleStateWaitingForConnectivity,
}

// GetExadataInfrastructureSummaryLifecycleStateEnumValues Enumerates the set of values for ExadataInfrastructureSummaryLifecycleStateEnum
func GetExadataInfrastructureSummaryLifecycleStateEnumValues() []ExadataInfrastructureSummaryLifecycleStateEnum {
	values := make([]ExadataInfrastructureSummaryLifecycleStateEnum, 0)
	for _, v := range mappingExadataInfrastructureSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExadataInfrastructureSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for ExadataInfrastructureSummaryLifecycleStateEnum
func GetExadataInfrastructureSummaryLifecycleStateEnumStringValues() []string {
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

// GetMappingExadataInfrastructureSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadataInfrastructureSummaryLifecycleStateEnum(val string) (ExadataInfrastructureSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingExadataInfrastructureSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExadataInfrastructureSummaryAdditionalComputeSystemModelEnum Enum with underlying type: string
type ExadataInfrastructureSummaryAdditionalComputeSystemModelEnum string

// Set of constants representing the allowable values for ExadataInfrastructureSummaryAdditionalComputeSystemModelEnum
const (
	ExadataInfrastructureSummaryAdditionalComputeSystemModelX7   ExadataInfrastructureSummaryAdditionalComputeSystemModelEnum = "X7"
	ExadataInfrastructureSummaryAdditionalComputeSystemModelX8   ExadataInfrastructureSummaryAdditionalComputeSystemModelEnum = "X8"
	ExadataInfrastructureSummaryAdditionalComputeSystemModelX8m  ExadataInfrastructureSummaryAdditionalComputeSystemModelEnum = "X8M"
	ExadataInfrastructureSummaryAdditionalComputeSystemModelX9m  ExadataInfrastructureSummaryAdditionalComputeSystemModelEnum = "X9M"
	ExadataInfrastructureSummaryAdditionalComputeSystemModelX10m ExadataInfrastructureSummaryAdditionalComputeSystemModelEnum = "X10M"
)

var mappingExadataInfrastructureSummaryAdditionalComputeSystemModelEnum = map[string]ExadataInfrastructureSummaryAdditionalComputeSystemModelEnum{
	"X7":   ExadataInfrastructureSummaryAdditionalComputeSystemModelX7,
	"X8":   ExadataInfrastructureSummaryAdditionalComputeSystemModelX8,
	"X8M":  ExadataInfrastructureSummaryAdditionalComputeSystemModelX8m,
	"X9M":  ExadataInfrastructureSummaryAdditionalComputeSystemModelX9m,
	"X10M": ExadataInfrastructureSummaryAdditionalComputeSystemModelX10m,
}

var mappingExadataInfrastructureSummaryAdditionalComputeSystemModelEnumLowerCase = map[string]ExadataInfrastructureSummaryAdditionalComputeSystemModelEnum{
	"x7":   ExadataInfrastructureSummaryAdditionalComputeSystemModelX7,
	"x8":   ExadataInfrastructureSummaryAdditionalComputeSystemModelX8,
	"x8m":  ExadataInfrastructureSummaryAdditionalComputeSystemModelX8m,
	"x9m":  ExadataInfrastructureSummaryAdditionalComputeSystemModelX9m,
	"x10m": ExadataInfrastructureSummaryAdditionalComputeSystemModelX10m,
}

// GetExadataInfrastructureSummaryAdditionalComputeSystemModelEnumValues Enumerates the set of values for ExadataInfrastructureSummaryAdditionalComputeSystemModelEnum
func GetExadataInfrastructureSummaryAdditionalComputeSystemModelEnumValues() []ExadataInfrastructureSummaryAdditionalComputeSystemModelEnum {
	values := make([]ExadataInfrastructureSummaryAdditionalComputeSystemModelEnum, 0)
	for _, v := range mappingExadataInfrastructureSummaryAdditionalComputeSystemModelEnum {
		values = append(values, v)
	}
	return values
}

// GetExadataInfrastructureSummaryAdditionalComputeSystemModelEnumStringValues Enumerates the set of values in String for ExadataInfrastructureSummaryAdditionalComputeSystemModelEnum
func GetExadataInfrastructureSummaryAdditionalComputeSystemModelEnumStringValues() []string {
	return []string{
		"X7",
		"X8",
		"X8M",
		"X9M",
		"X10M",
	}
}

// GetMappingExadataInfrastructureSummaryAdditionalComputeSystemModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadataInfrastructureSummaryAdditionalComputeSystemModelEnum(val string) (ExadataInfrastructureSummaryAdditionalComputeSystemModelEnum, bool) {
	enum, ok := mappingExadataInfrastructureSummaryAdditionalComputeSystemModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExadataInfrastructureSummaryMaintenanceSLOStatusEnum Enum with underlying type: string
type ExadataInfrastructureSummaryMaintenanceSLOStatusEnum string

// Set of constants representing the allowable values for ExadataInfrastructureSummaryMaintenanceSLOStatusEnum
const (
	ExadataInfrastructureSummaryMaintenanceSLOStatusOk       ExadataInfrastructureSummaryMaintenanceSLOStatusEnum = "OK"
	ExadataInfrastructureSummaryMaintenanceSLOStatusDegraded ExadataInfrastructureSummaryMaintenanceSLOStatusEnum = "DEGRADED"
)

var mappingExadataInfrastructureSummaryMaintenanceSLOStatusEnum = map[string]ExadataInfrastructureSummaryMaintenanceSLOStatusEnum{
	"OK":       ExadataInfrastructureSummaryMaintenanceSLOStatusOk,
	"DEGRADED": ExadataInfrastructureSummaryMaintenanceSLOStatusDegraded,
}

var mappingExadataInfrastructureSummaryMaintenanceSLOStatusEnumLowerCase = map[string]ExadataInfrastructureSummaryMaintenanceSLOStatusEnum{
	"ok":       ExadataInfrastructureSummaryMaintenanceSLOStatusOk,
	"degraded": ExadataInfrastructureSummaryMaintenanceSLOStatusDegraded,
}

// GetExadataInfrastructureSummaryMaintenanceSLOStatusEnumValues Enumerates the set of values for ExadataInfrastructureSummaryMaintenanceSLOStatusEnum
func GetExadataInfrastructureSummaryMaintenanceSLOStatusEnumValues() []ExadataInfrastructureSummaryMaintenanceSLOStatusEnum {
	values := make([]ExadataInfrastructureSummaryMaintenanceSLOStatusEnum, 0)
	for _, v := range mappingExadataInfrastructureSummaryMaintenanceSLOStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetExadataInfrastructureSummaryMaintenanceSLOStatusEnumStringValues Enumerates the set of values in String for ExadataInfrastructureSummaryMaintenanceSLOStatusEnum
func GetExadataInfrastructureSummaryMaintenanceSLOStatusEnumStringValues() []string {
	return []string{
		"OK",
		"DEGRADED",
	}
}

// GetMappingExadataInfrastructureSummaryMaintenanceSLOStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadataInfrastructureSummaryMaintenanceSLOStatusEnum(val string) (ExadataInfrastructureSummaryMaintenanceSLOStatusEnum, bool) {
	enum, ok := mappingExadataInfrastructureSummaryMaintenanceSLOStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
