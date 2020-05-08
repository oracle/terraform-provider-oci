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

// UpdateSddcDetails The SDDC information to be updated.
// **Important:** Only the `displayName`, `freeFormTags`, and `definedTags` attributes
// affect the existing SDDC. Changing the other attributes affects the `Sddc` object, but not
// the VMware environment currently running on that SDDC. Those other attributes are used
// by the Oracle Cloud VMware Solution *only* for new ESXi hosts that you add to this
// SDDC in the future with CreateEsxiHost.
type UpdateSddcDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the SDDC.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The version of bundled VMware software that the Oracle Cloud VMware Solution will
	// install on any new ESXi hosts that you add to this SDDC in the future.
	// For the list of versions supported by the Oracle Cloud VMware Solution, see
	// ListSupportedVmwareSoftwareVersions).
	VmwareSoftwareVersion *string `mandatory:"false" json:"vmwareSoftwareVersion"`

	// One or more public SSH keys to be included in the `~/.ssh/authorized_keys` file for
	// the default user on each ESXi host, only when adding new ESXi hosts to this SDDC.
	// Use a newline character to separate multiple keys.
	// The SSH keys must be in the format required for the `authorized_keys` file.
	SshAuthorizedKeys *string `mandatory:"false" json:"sshAuthorizedKeys"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN to use for
	// the vSphere component of the VMware environment when adding new ESXi hosts to the SDDC.
	VsphereVlanId *string `mandatory:"false" json:"vsphereVlanId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN to use for
	// the vMotion component of the VMware environment when adding new ESXi hosts to the SDDC.
	VmotionVlanId *string `mandatory:"false" json:"vmotionVlanId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN to use for
	// the vSAN component of the VMware environment when adding new ESXi hosts to the SDDC.
	VsanVlanId *string `mandatory:"false" json:"vsanVlanId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN to use for
	// the NSX VTEP component of the VMware environment when adding new ESXi hosts to the SDDC.
	NsxVTepVlanId *string `mandatory:"false" json:"nsxVTepVlanId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN to use for
	// the NSX Edge VTEP component of the VMware environment when adding new ESXi hosts to the SDDC.
	NsxEdgeVTepVlanId *string `mandatory:"false" json:"nsxEdgeVTepVlanId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN to use for
	// the NSX Edge Uplink 1 component of the VMware environment when adding new ESXi hosts to the SDDC.
	NsxEdgeUplink1VlanId *string `mandatory:"false" json:"nsxEdgeUplink1VlanId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN to use for
	// the NSX Edge Uplink 2 component of the VMware environment when adding new ESXi hosts to the SDDC.
	NsxEdgeUplink2VlanId *string `mandatory:"false" json:"nsxEdgeUplink2VlanId"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateSddcDetails) String() string {
	return common.PointerString(m)
}
