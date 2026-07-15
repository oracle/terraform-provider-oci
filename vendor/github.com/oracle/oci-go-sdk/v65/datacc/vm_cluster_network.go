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

// VmClusterNetwork Details of the VM cluster network on Database Infrastructure.
type VmClusterNetwork struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster network.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the VM cluster network.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Infrastructure.
	InfrastructureId *string `mandatory:"true" json:"infrastructureId"`

	// The user-friendly name for the VM cluster network. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The SCAN details.
	Scans []ScanDetails `mandatory:"true" json:"scans"`

	// Details of the client and backup networks.
	VmNetworks []VmNetworkDetails `mandatory:"true" json:"vmNetworks"`

	// Count of virtual machines in this VM cluster.
	NodeCount *int `mandatory:"false" json:"nodeCount"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated VM cluster.
	BaseVmClusterId *string `mandatory:"false" json:"baseVmClusterId"`

	// Indicates whether Single Client Access Name (SCAN) is enabled on the VM cluster.
	IsScanEnabled *bool `mandatory:"false" json:"isScanEnabled"`

	// The listener TCP/IP port.
	ListenerPort *int `mandatory:"false" json:"listenerPort"`

	// The listener TCP/IP SSL port.
	ListenerPortSsl *int `mandatory:"false" json:"listenerPortSsl"`

	// The list of DNS server IP addresses. Maximum of 3 allowed.
	DnsServers []string `mandatory:"false" json:"dnsServers"`

	// The list of NTP server IP addresses. Maximum of 3 allowed.
	NtpServers []string `mandatory:"false" json:"ntpServers"`

	// The current state of the virtual machine cluster network.
	LifecycleState VmClusterNetworkLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The time that the VM cluster network was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time that the VM cluster network was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Lifecycle state details of the VM cluster network.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated resource.
	AssociatedResourceId *string `mandatory:"false" json:"associatedResourceId"`

	// Consumer type for the VM cluster network.
	ConsumerType VmNetworkConsumerTypeEnum `mandatory:"false" json:"consumerType,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. This tag option exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{ "orcl-cloud": { "free-tier-retained": "true" } }`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m VmClusterNetwork) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VmClusterNetwork) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingVmClusterNetworkLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetVmClusterNetworkLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingVmNetworkConsumerTypeEnum(string(m.ConsumerType)); !ok && m.ConsumerType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConsumerType: %s. Supported values are: %s.", m.ConsumerType, strings.Join(GetVmNetworkConsumerTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
