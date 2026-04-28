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

// CreateInfrastructureDetails Request to create Database Infrastructure resource.
type CreateInfrastructureDetails struct {

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

	// The list of DNS server IP addresses. Maximum of 3 allowed.
	DnsServers []string `mandatory:"true" json:"dnsServers"`

	// The list of NTP server IP addresses. Maximum of 3 allowed.
	NtpServers []string `mandatory:"true" json:"ntpServers"`

	// Database Infrastructure description.
	Description *string `mandatory:"false" json:"description"`

	// The CIDR block for the system network. The system network is a private network in Database Infrastructure and is not connected to your corporate network. The system network is used for storage (ASM) traffic, high-performance interconnect traffic and administration of infrastructure components.
	AdminNetworkcidr *string `mandatory:"false" json:"adminNetworkcidr"`

	// Percentage of disk space assigned for DATA disk group. Remaining disk space will get assiged to RECO disk group
	DataDiskPercentage *int `mandatory:"false" json:"dataDiskPercentage"`

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

	// The corporate network proxy for access to the control plane network. Oracle recommends using an HTTPS proxy when possible
	// for enhanced security.
	CorporateProxy *string `mandatory:"false" json:"corporateProxy"`

	// The CPS network VLAN ID.
	VlanId *string `mandatory:"false" json:"vlanId"`

	MaintenanceWindow *MaintenanceWindow `mandatory:"false" json:"maintenanceWindow"`

	// Simple key-value pair that is applied without any predefined name, type or scope. This tag option exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateInfrastructureDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateInfrastructureDetails) ValidateEnumValue() (bool, error) {
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
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
