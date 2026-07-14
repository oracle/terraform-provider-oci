// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Infrastructure Cloud@Customer Service API
//
// API for Database Infrastructure Cloud@Customer Service. Use this API to manage Database Infrastructure VM clusters, Application VMs, and related resources.
//

package datacc

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Infrastructure Details about Database Infrastructure resource.
type Infrastructure struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Infrastructure.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the Database Infrastructure. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Database Infrastructure System Model specification. The system model determines the model of the Database Infrastructure hardware to be used.
	SystemModel SystemModelEnumEnum `mandatory:"true" json:"systemModel"`

	// The shape of the Database Infrastructure. The shape determines the amount of CPU, storage, and memory resources allocated to the instance.
	Shape ShapeEnumEnum `mandatory:"true" json:"shape"`

	// The IP address for the first control plane server.
	CloudControlPlaneServer1 *string `mandatory:"true" json:"cloudControlPlaneServer1"`

	// The IP address for the second control plane server.
	CloudControlPlaneServer2 *string `mandatory:"true" json:"cloudControlPlaneServer2"`

	// The netmask for the control plane network.
	Netmask *string `mandatory:"true" json:"netmask"`

	// The gateway for the control plane network.
	Gateway *string `mandatory:"true" json:"gateway"`

	// The CIDR block for the system network. The system network is a private network in Database Infrastructure and is not connected to your corporate network. The system network is used for storage (ASM) traffic, high-performance interconnect traffic and administration of infrastructure components.
	AdminNetworkcidr *string `mandatory:"true" json:"adminNetworkcidr"`

	// The list of DNS server IP addresses. Maximum of 3 allowed.
	DnsServers []string `mandatory:"true" json:"dnsServers"`

	// The list of NTP server IP addresses. Maximum of 3 allowed.
	NtpServers []string `mandatory:"true" json:"ntpServers"`

	// Database Infrastructure description.
	Description *string `mandatory:"false" json:"description"`

	// The version of the system software on the Database Infrastructure.
	Version *string `mandatory:"false" json:"version"`

	// Percentage of disk space assigned for DATA disk group.
	DataDiskPercentage *int `mandatory:"false" json:"dataDiskPercentage"`

	// Percentage of disk space assigned for RECO disk group.
	RecoDiskPercentage *int `mandatory:"false" json:"recoDiskPercentage"`

	// The list of contacts for the Database Infrastructure.
	Contacts []InfrastructureContact `mandatory:"false" json:"contacts"`

	// The network bonding mode for Client networks for the Database Infrastructure.
	ClientNetworkBondingMode NetworkBondingModeEnum `mandatory:"false" json:"clientNetworkBondingMode,omitempty"`

	// The network bonding mode for Backup networks for the Database Infrastructure.
	BackupNetworkBondingMode NetworkBondingModeEnum `mandatory:"false" json:"backupNetworkBondingMode,omitempty"`

	// The network bonding mode for CPS networks for the Database Infrastructure.
	CpsNetworkBondingMode NetworkBondingModeEnum `mandatory:"false" json:"cpsNetworkBondingMode,omitempty"`

	// The network bonding interface for client network for the Database Infrastructure.
	ClientNetworkBondingInterface NetworkBondingInterfaceEnumEnum `mandatory:"false" json:"clientNetworkBondingInterface,omitempty"`

	// The network bonding interface for backup network for the Database Infrastructure.
	BackupNetworkBondingInterface NetworkBondingInterfaceEnumEnum `mandatory:"false" json:"backupNetworkBondingInterface,omitempty"`

	// The network bonding interface for CPS network for the Database Infrastructure.
	CpsNetworkBondingInterface NetworkBondingInterfaceEnumEnum `mandatory:"false" json:"cpsNetworkBondingInterface,omitempty"`

	// The amount of storage (in GB) in the DATA disk group that is reserved for creating local storage for VM Clusters and application VMs.
	AcfsFileSystemStorageInGbs *float64 `mandatory:"false" json:"acfsFileSystemStorageInGbs"`

	// The amount of storage (in GB) in the DATA disk group that is currently utilized for creating local storage for VM Clusters and application VMs.
	// This attribute is deprecated and will be removed in a subsequent release. Please read from systemStorageCapacity instead.
	AcfsFileSystemUsedStorageInGbs *float64 `mandatory:"false" json:"acfsFileSystemUsedStorageInGbs"`

	// The corporate network proxy for access to the control plane network. Oracle recommends using an HTTPS proxy when possible
	// for enhanced security.
	CorporateProxy *string `mandatory:"false" json:"corporateProxy"`

	// The CPS network VLAN ID.
	VlanId *string `mandatory:"false" json:"vlanId"`

	// Simple key-value pair that is applied without any predefined name, type or scope. This tag option exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{ "orcl-cloud": { "free-tier-retained": "true" } }`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	MaintenanceWindow *MaintenanceWindow `mandatory:"false" json:"maintenanceWindow"`

	// The time that the Database Infrastructure cluster was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time that the Database Infrastructure was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The time, in RFC3339 format, when the lifecycle state was last updated.
	TimeLastStateUpdated *common.SDKTime `mandatory:"false" json:"timeLastStateUpdated"`

	// The time, in RFC3339 format, when the Database Infrastructure was activated.
	TimeActivated *common.SDKTime `mandatory:"false" json:"timeActivated"`

	// The time, in RFC3339 format, when the Database Infrastructure network was validated.
	TimeValidated *common.SDKTime `mandatory:"false" json:"timeValidated"`

	// The current state of the Database Infrastructure.
	LifecycleState InfrastructureLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Lifecycle state details of the Database Infrastructure.
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	// The serial number for the Database Infrastructure.
	RackSerialNumber *string `mandatory:"false" json:"rackSerialNumber"`

	// A list of Database Infrastructure nodes.
	Servers []InfrastructureServer `mandatory:"false" json:"servers"`

	ComputeCapacity *ComputeCapacityDetails `mandatory:"false" json:"computeCapacity"`

	// Capacity details of the Storage disk group.
	// This attribute is deprecated and will be removed in a subsequent release. Please
	// use systemStorageCapacity instead.
	StorageCapacity []StorageCapacityDetails `mandatory:"false" json:"storageCapacity"`

	SystemStorageCapacity *SystemStorageCapacityDetails `mandatory:"false" json:"systemStorageCapacity"`

	// The network adapter, transceiver and cable configuration for the client and backup networks.
	NetworkAdapterConfiguration *string `mandatory:"false" json:"networkAdapterConfiguration"`

	// The unique identifier for the subscription plan number.
	SubscriptionPlanNumber *string `mandatory:"false" json:"subscriptionPlanNumber"`

	// SSD configuration requested for the infrastructure.
	SsdConfigurationRequested ShapeEnumEnum `mandatory:"false" json:"ssdConfigurationRequested,omitempty"`
}

func (m Infrastructure) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Infrastructure) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSystemModelEnumEnum(string(m.SystemModel)); !ok && m.SystemModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SystemModel: %s. Supported values are: %s.", m.SystemModel, strings.Join(GetSystemModelEnumEnumStringValues(), ",")))
	}
	if _, ok := GetMappingShapeEnumEnum(string(m.Shape)); !ok && m.Shape != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Shape: %s. Supported values are: %s.", m.Shape, strings.Join(GetShapeEnumEnumStringValues(), ",")))
	}

	if _, ok := GetMappingNetworkBondingModeEnum(string(m.ClientNetworkBondingMode)); !ok && m.ClientNetworkBondingMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ClientNetworkBondingMode: %s. Supported values are: %s.", m.ClientNetworkBondingMode, strings.Join(GetNetworkBondingModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingNetworkBondingModeEnum(string(m.BackupNetworkBondingMode)); !ok && m.BackupNetworkBondingMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackupNetworkBondingMode: %s. Supported values are: %s.", m.BackupNetworkBondingMode, strings.Join(GetNetworkBondingModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingNetworkBondingModeEnum(string(m.CpsNetworkBondingMode)); !ok && m.CpsNetworkBondingMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CpsNetworkBondingMode: %s. Supported values are: %s.", m.CpsNetworkBondingMode, strings.Join(GetNetworkBondingModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingNetworkBondingInterfaceEnumEnum(string(m.ClientNetworkBondingInterface)); !ok && m.ClientNetworkBondingInterface != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ClientNetworkBondingInterface: %s. Supported values are: %s.", m.ClientNetworkBondingInterface, strings.Join(GetNetworkBondingInterfaceEnumEnumStringValues(), ",")))
	}
	if _, ok := GetMappingNetworkBondingInterfaceEnumEnum(string(m.BackupNetworkBondingInterface)); !ok && m.BackupNetworkBondingInterface != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackupNetworkBondingInterface: %s. Supported values are: %s.", m.BackupNetworkBondingInterface, strings.Join(GetNetworkBondingInterfaceEnumEnumStringValues(), ",")))
	}
	if _, ok := GetMappingNetworkBondingInterfaceEnumEnum(string(m.CpsNetworkBondingInterface)); !ok && m.CpsNetworkBondingInterface != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CpsNetworkBondingInterface: %s. Supported values are: %s.", m.CpsNetworkBondingInterface, strings.Join(GetNetworkBondingInterfaceEnumEnumStringValues(), ",")))
	}
	if _, ok := GetMappingInfrastructureLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetInfrastructureLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingShapeEnumEnum(string(m.SsdConfigurationRequested)); !ok && m.SsdConfigurationRequested != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SsdConfigurationRequested: %s. Supported values are: %s.", m.SsdConfigurationRequested, strings.Join(GetShapeEnumEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
