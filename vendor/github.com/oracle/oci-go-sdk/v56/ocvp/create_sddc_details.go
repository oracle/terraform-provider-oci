// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use the Oracle Cloud VMware API to create SDDCs and manage ESXi hosts and software.
// For more information, see Oracle Cloud VMware Solution (https://docs.cloud.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CreateSddcDetails Details of the SDDC.
type CreateSddcDetails struct {

	// The availability domain to create the SDDC's ESXi hosts in. For multi-AD SDDC deployment, set to `multi-AD`.
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
	// you are still billed for the 3 minimum recommended ESXi hosts. Also,
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
	// **Note:** This VLAN is reserved for future use to deploy public-facing applications on the VMware SDDC.
	NsxEdgeUplink2VlanId *string `mandatory:"true" json:"nsxEdgeUplink2VlanId"`

	// A descriptive name for the SDDC.
	// SDDC name requirements are 1-16 character length limit, Must start with a letter, Must be English letters, numbers, - only, No repeating hyphens, Must be unique within the region.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A prefix used in the name of each ESXi host and Compute instance in the SDDC.
	// If this isn't set, the SDDC's `displayName` is used as the prefix.
	// For example, if the value is `mySDDC`, the ESXi hosts are named `mySDDC-1`,
	// `mySDDC-2`, and so on.
	InstanceDisplayNamePrefix *string `mandatory:"false" json:"instanceDisplayNamePrefix"`

	// The billing option selected during SDDC creation.
	// ListSupportedSkus.
	InitialSku SkuEnum `mandatory:"false" json:"initialSku,omitempty"`

	// Indicates whether to enable HCX for this SDDC.
	IsHcxEnabled *bool `mandatory:"false" json:"isHcxEnabled"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN to use for the HCX
	// component of the VMware environment. This value is required only when `isHcxEnabled` is true.
	HcxVlanId *string `mandatory:"false" json:"hcxVlanId"`

	// Indicates whether to enable HCX Enterprise for this SDDC.
	IsHcxEnterpriseEnabled *bool `mandatory:"false" json:"isHcxEnterpriseEnabled"`

	// The CIDR block for the IP addresses that VMware VMs in the SDDC use to run application
	// workloads.
	WorkloadNetworkCidr *string `mandatory:"false" json:"workloadNetworkCidr"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC
	// for the vSphere Replication component of the VMware environment.
	ReplicationVlanId *string `mandatory:"false" json:"replicationVlanId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC
	// for the Provisioning component of the VMware environment.
	ProvisioningVlanId *string `mandatory:"false" json:"provisioningVlanId"`

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
