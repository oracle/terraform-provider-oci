// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use this API to manage the Oracle Cloud VMware Solution.
//

package ocvp

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateSddcDetails Details of the SDDC.
type CreateSddcDetails struct {

	// The availability domain to create the SDDC's ESXi hosts in.
	ComputeAvailabilityDomain *string `mandatory:"true" json:"computeAvailabilityDomain"`

	// The VMware software bundle to install on the ESXi hosts in the SDDC. To get a
	// list of the available versions, use
	// ListSupportedVmwareSoftwareVersions.
	VmwareSoftwareVersion *string `mandatory:"true" json:"vmwareSoftwareVersion"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment to contain the SDDC.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The number of ESXi hosts to create in the SDDC. You can add more hosts later
	// (see CreateEsxiHost).
	// **Note:** If you later delete EXSi hosts from the SDDC to total less than 3,
	// you are still billed for the 3 minimum recommended EXSi hosts. Also,
	// you cannot add more VMware workloads to the SDDC until it again has at least
	// 3 ESXi hosts.
	EsxiHostsCount *int `mandatory:"true" json:"esxiHostsCount"`

	// One or more public SSH keys to be included in the `~/.ssh/authorized_keys` file for
	// the default user on each ESXi host. Use a newline character to separate multiple keys.
	// The SSH keys must be in the format required for the `authorized_keys` file
	SshAuthorizedKeys *string `mandatory:"true" json:"sshAuthorizedKeys"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the management subnet to use
	// for provisioning the SDDC.
	ProvisioningSubnetId *string `mandatory:"true" json:"provisioningSubnetId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN to use for the vSphere
	// component of the VMware environment.
	VsphereVlanId *string `mandatory:"true" json:"vsphereVlanId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN to use for the vMotion
	// component of the VMware environment.
	VmotionVlanId *string `mandatory:"true" json:"vmotionVlanId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN to use for the vSAN
	// component of the VMware environment.
	VsanVlanId *string `mandatory:"true" json:"vsanVlanId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN to use for the NSX VTEP
	// component of the VMware environment.
	NsxVTepVlanId *string `mandatory:"true" json:"nsxVTepVlanId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN to use for the NSX Edge VTEP
	// component of the VMware environment.
	NsxEdgeVTepVlanId *string `mandatory:"true" json:"nsxEdgeVTepVlanId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN to use for the NSX Edge
	// Uplink 1 component of the VMware environment.
	NsxEdgeUplink1VlanId *string `mandatory:"true" json:"nsxEdgeUplink1VlanId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN to use for the NSX Edge
	// Uplink 2 component of the VMware environment.
	NsxEdgeUplink2VlanId *string `mandatory:"true" json:"nsxEdgeUplink2VlanId"`

	// A descriptive name for the SDDC. It must be unique, start with a letter, and contain only letters, digits,
	// whitespaces, dashes and underscores.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A prefix used in the name of each ESXi host and Compute instance in the SDDC.
	// If this isn't set, the SDDC's `displayName` is used as the prefix.
	// For example, if the value is `mySDDC`, the ESXi hosts are named `mySDDC-1`,
	// `mySDDC-2`, and so on.
	InstanceDisplayNamePrefix *string `mandatory:"false" json:"instanceDisplayNamePrefix"`

	// The CIDR block for the IP addresses that VMware VMs in the SDDC use to run application
	// workloads.
	WorkloadNetworkCidr *string `mandatory:"false" json:"workloadNetworkCidr"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateSddcDetails) String() string {
	return common.PointerString(m)
}
