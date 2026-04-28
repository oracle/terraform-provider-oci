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

// VmInstance VM instance details.
type VmInstance struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM instance.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the VM
	// instance.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Infrastructure.
	InfrastructureId *string `mandatory:"true" json:"infrastructureId"`

	// The number of CPU cores enabled for each VM instance.
	CpusEnabled *int `mandatory:"true" json:"cpusEnabled"`

	// List of public key used for SSH access to the VM instance.
	SshPublicKeys []string `mandatory:"true" json:"sshPublicKeys"`

	// The current state of the VM instance.
	LifecycleState VmInstanceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time that the VM instance was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time that the VM instance was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The time zone to use for the VM instance.
	TimeZone *string `mandatory:"true" json:"timeZone"`

	// VM instance display name. This name does not have to be unique, and is changeable.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// VM instance description.
	Description *string `mandatory:"false" json:"description"`

	// The memory to be allocated in GBs.
	MemorySizeInGBs *float64 `mandatory:"false" json:"memorySizeInGBs"`

	// Boot storage memory to be allocated in GBs.
	BootStorageSizeInGBs *float64 `mandatory:"false" json:"bootStorageSizeInGBs"`

	// Data storage to be allocated in GBs.
	DataStorageSizeInGBs *float64 `mandatory:"false" json:"dataStorageSizeInGBs"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. This tag option exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{ "orcl-cloud": { "free-tier-retained": "true" } }`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM custom instance uploaded.
	ImageId *string `mandatory:"false" json:"imageId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute node on which VM instance should be launched.
	ServerId *string `mandatory:"false" json:"serverId"`

	// Lifecycle state details of the VM instance.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The domain name of the VM instance.
	DomainName *string `mandatory:"false" json:"domainName"`

	// The list of DNS server IP addresses. Maximum of 3 allowed.
	DnsServers []string `mandatory:"false" json:"dnsServers"`

	// The list of NTP server addresses. Maximum of 3 allowed.
	NtpServers []string `mandatory:"false" json:"ntpServers"`

	// The host name of the instance.
	Hostname *string `mandatory:"false" json:"hostname"`

	// The IP address of the instance.
	IpAddress *string `mandatory:"false" json:"ipAddress"`

	// The netmask of the VM instance network.
	Netmask *string `mandatory:"false" json:"netmask"`

	// The gateway IP address of the VM instance network.
	Gateway *string `mandatory:"false" json:"gateway"`

	// The network VLAN ID.
	VlanId *string `mandatory:"false" json:"vlanId"`

	// Base64-encoded data to be used by Cloud-Init to run custom scripts or provide custom Cloud-Init configuration.
	// For information about how to take advantage of user data, see the Cloud-Init Documentation (http://cloudinit.readthedocs.org/en/latest/topics/format.html).
	Userdata *string `mandatory:"false" json:"userdata"`

	// Custom metadata key/value pairs which can be used to:
	// - Provide information to Cloud-Init (https://cloudinit.readthedocs.org/en/latest/) to be used for various system initialization tasks.
	// - Provide additional information which is exposed inside the instance context and can be queried or referenced by user-data scripts for dynamic configuration.
	Metadata map[string]string `mandatory:"false" json:"metadata"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM Network.
	VmNetworkId *string `mandatory:"false" json:"vmNetworkId"`
}

func (m VmInstance) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VmInstance) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVmInstanceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetVmInstanceLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
