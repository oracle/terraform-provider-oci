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

// Sddc A software-defined data center (SDDC) contains the resources required for a
// functional VMware environment. Instances in an SDDC
// (see EsxiHost) run in a virtual cloud network (VCN)
// and are preconfigured with VMware and storage. Use the vCenter utility to manage
// and deploy VMware virtual machines (VMs) in the SDDC.
// The SDDC uses a single management subnet for provisioning the SDDC. It also uses a
// set of VLANs for various components of the VMware environment (vSphere, vMotion,
// vSAN, and so on). See the Core Services API for information about VCN subnets and VLANs.
type Sddc struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the SDDC.
	Id *string `mandatory:"true" json:"id"`

	// The availability domain the ESXi hosts are running in.
	// Example: `Uocm:PHX-AD-1`
	ComputeAvailabilityDomain *string `mandatory:"true" json:"computeAvailabilityDomain"`

	// A descriptive name for the SDDC. It must be unique, start with a letter, and contain only letters, digits,
	// whitespaces, dashes and underscores.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// In general, this is a specific version of bundled VMware software supported by
	// Oracle Cloud VMware Solution (see
	// ListSupportedVmwareSoftwareVersions).
	// This attribute is not guaranteed to reflect the version of
	// software currently installed on the ESXi hosts in the SDDC. The purpose
	// of this attribute is to show the version of software that the Oracle
	// Cloud VMware Solution will install on any new ESXi hosts that you *add to this
	// SDDC in the future* with CreateEsxiHost.
	// Therefore, if you upgrade the existing ESXi hosts in the SDDC to use a newer
	// version of bundled VMware software supported by the Oracle Cloud VMware Solution, you
	// should use UpdateSddc to update the SDDC's
	// `vmwareSoftwareVersion` with that new version.
	VmwareSoftwareVersion *string `mandatory:"true" json:"vmwareSoftwareVersion"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that
	// contains the SDDC.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The number of ESXi hosts in the SDDC.
	EsxiHostsCount *int `mandatory:"true" json:"esxiHostsCount"`

	// FQDN for vCenter
	// Example: `vcenter-my-sddc.sddc.us-phoenix-1.oraclecloud.com`
	VcenterFqdn *string `mandatory:"true" json:"vcenterFqdn"`

	// FQDN for NSX Manager
	// Example: `nsx-my-sddc.sddc.us-phoenix-1.oraclecloud.com`
	NsxManagerFqdn *string `mandatory:"true" json:"nsxManagerFqdn"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the `PrivateIp` object that is
	// the virtual IP (VIP) for vCenter. For information about `PrivateIp` objects, see the
	// Core Services API.
	VcenterPrivateIpId *string `mandatory:"true" json:"vcenterPrivateIpId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the `PrivateIp` object that is
	// the virtual IP (VIP) for NSX Manager. For information about `PrivateIp` objects, see the
	// Core Services API.
	NsxManagerPrivateIpId *string `mandatory:"true" json:"nsxManagerPrivateIpId"`

	// One or more public SSH keys to be included in the `~/.ssh/authorized_keys` file for
	// the default user on each ESXi host. Use a newline character to separate multiple keys.
	// The SSH keys must be in the format required for the `authorized_keys` file.
	// This attribute is not guaranteed to reflect the public SSH keys
	// currently installed on the ESXi hosts in the SDDC. The purpose
	// of this attribute is to show the public SSH keys that Oracle
	// Cloud VMware Solution will install on any new ESXi hosts that you *add to this
	// SDDC in the future* with CreateEsxiHost.
	// Therefore, if you upgrade the existing ESXi hosts in the SDDC to use different
	// SSH keys, you should use UpdateSddc to update
	// the SDDC's `sshAuthorizedKeys` with the new public keys.
	SshAuthorizedKeys *string `mandatory:"true" json:"sshAuthorizedKeys"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the management subnet used
	// to provision the SDDC.
	ProvisioningSubnetId *string `mandatory:"true" json:"provisioningSubnetId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC
	// for the vSphere component of the VMware environment.
	// This attribute is not guaranteed to reflect the vSphere VLAN
	// currently used by the ESXi hosts in the SDDC. The purpose
	// of this attribute is to show the vSphere VLAN that the Oracle
	// Cloud VMware Solution will use for any new ESXi hosts that you *add to this
	// SDDC in the future* with CreateEsxiHost.
	// Therefore, if you change the existing ESXi hosts in the SDDC to use a different VLAN
	// for the vSphere component of the VMware environment, you
	// should use UpdateSddc to update the SDDC's
	// `vsphereVlanId` with that new VLAN's OCID.
	VsphereVlanId *string `mandatory:"true" json:"vsphereVlanId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC
	// for the vMotion component of the VMware environment.
	// This attribute is not guaranteed to reflect the vMotion VLAN
	// currently used by the ESXi hosts in the SDDC. The purpose
	// of this attribute is to show the vMotion VLAN that the Oracle
	// Cloud VMware Solution will use for any new ESXi hosts that you *add to this
	// SDDC in the future* with CreateEsxiHost.
	// Therefore, if you change the existing ESXi hosts in the SDDC to use a different VLAN
	// for the vMotion component of the VMware environment, you
	// should use UpdateSddc to update the SDDC's
	// `vmotionVlanId` with that new VLAN's OCID.
	VmotionVlanId *string `mandatory:"true" json:"vmotionVlanId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC
	// for the vSAN component of the VMware environment.
	// This attribute is not guaranteed to reflect the vSAN VLAN
	// currently used by the ESXi hosts in the SDDC. The purpose
	// of this attribute is to show the vSAN VLAN that the Oracle
	// Cloud VMware Solution will use for any new ESXi hosts that you *add to this
	// SDDC in the future* with CreateEsxiHost.
	// Therefore, if you change the existing ESXi hosts in the SDDC to use a different VLAN
	// for the vSAN component of the VMware environment, you
	// should use UpdateSddc to update the SDDC's
	// `vsanVlanId` with that new VLAN's OCID.
	VsanVlanId *string `mandatory:"true" json:"vsanVlanId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC
	// for the NSX VTEP component of the VMware environment.
	// This attribute is not guaranteed to reflect the NSX VTEP VLAN
	// currently used by the ESXi hosts in the SDDC. The purpose
	// of this attribute is to show the NSX VTEP VLAN that the Oracle
	// Cloud VMware Solution will use for any new ESXi hosts that you *add to this
	// SDDC in the future* with CreateEsxiHost.
	// Therefore, if you change the existing ESXi hosts in the SDDC to use a different VLAN
	// for the NSX VTEP component of the VMware environment, you
	// should use UpdateSddc to update the SDDC's
	// `nsxVTepVlanId` with that new VLAN's OCID.
	NsxVTepVlanId *string `mandatory:"true" json:"nsxVTepVlanId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC
	// for the NSX Edge VTEP component of the VMware environment.
	// This attribute is not guaranteed to reflect the NSX Edge VTEP VLAN
	// currently used by the ESXi hosts in the SDDC. The purpose
	// of this attribute is to show the NSX Edge VTEP VLAN that the Oracle
	// Cloud VMware Solution will use for any new ESXi hosts that you *add to this
	// SDDC in the future* with CreateEsxiHost.
	// Therefore, if you change the existing ESXi hosts in the SDDC to use a different VLAN
	// for the NSX Edge VTEP component of the VMware environment, you
	// should use UpdateSddc to update the SDDC's
	// `nsxEdgeVTepVlanId` with that new VLAN's OCID.
	NsxEdgeVTepVlanId *string `mandatory:"true" json:"nsxEdgeVTepVlanId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC
	// for the NSX Edge Uplink 1 component of the VMware environment.
	// This attribute is not guaranteed to reflect the NSX Edge Uplink 1 VLAN
	// currently used by the ESXi hosts in the SDDC. The purpose
	// of this attribute is to show the NSX Edge Uplink 1 VLAN that the Oracle
	// Cloud VMware Solution will use for any new ESXi hosts that you *add to this
	// SDDC in the future* with CreateEsxiHost.
	// Therefore, if you change the existing ESXi hosts in the SDDC to use a different VLAN
	// for the NSX Edge Uplink 1 component of the VMware environment, you
	// should use UpdateSddc to update the SDDC's
	// `nsxEdgeUplink1VlanId` with that new VLAN's OCID.
	NsxEdgeUplink1VlanId *string `mandatory:"true" json:"nsxEdgeUplink1VlanId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC
	// for the NSX Edge Uplink 2 component of the VMware environment.
	// This attribute is not guaranteed to reflect the NSX Edge Uplink 2 VLAN
	// currently used by the ESXi hosts in the SDDC. The purpose
	// of this attribute is to show the NSX Edge Uplink 2 VLAN that the Oracle
	// Cloud VMware Solution will use for any new ESXi hosts that you *add to this
	// SDDC in the future* with CreateEsxiHost.
	// Therefore, if you change the existing ESXi hosts in the SDDC to use a different VLAN
	// for the NSX Edge Uplink 2 component of the VMware environment, you
	// should use UpdateSddc to update the SDDC's
	// `nsxEdgeUplink2VlanId` with that new VLAN's OCID.
	NsxEdgeUplink2VlanId *string `mandatory:"true" json:"nsxEdgeUplink2VlanId"`

	// The date and time the SDDC was created, in the format defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// A prefix used in the name of each ESXi host and Compute instance in the SDDC.
	// If this isn't set, the SDDC's `displayName` is used as the prefix.
	// For example, if the value is `MySDDC`, the ESXi hosts are named `MySDDC-1`,
	// `MySDDC-2`, and so on.
	InstanceDisplayNamePrefix *string `mandatory:"false" json:"instanceDisplayNamePrefix"`

	// The SDDC includes an administrator username and initial password for vCenter. Make sure
	// to change this initial vCenter password to a different value.
	VcenterInitialPassword *string `mandatory:"false" json:"vcenterInitialPassword"`

	// The SDDC includes an administrator username and initial password for NSX Manager. Make sure
	// to change this initial NSX Manager password to a different value.
	NsxManagerInitialPassword *string `mandatory:"false" json:"nsxManagerInitialPassword"`

	// The SDDC includes an administrator username and initial password for vCenter. You can
	// change this initial username to a different value in vCenter.
	VcenterUsername *string `mandatory:"false" json:"vcenterUsername"`

	// The SDDC includes an administrator username and initial password for NSX Manager. You
	// can change this initial username to a different value in NSX Manager.
	NsxManagerUsername *string `mandatory:"false" json:"nsxManagerUsername"`

	// The CIDR block for the IP addresses that VMware VMs in the SDDC use to run application
	// workloads.
	WorkloadNetworkCidr *string `mandatory:"false" json:"workloadNetworkCidr"`

	// The VMware NSX overlay workload segment to host your application. Connect to workload
	// portgroup in vCenter to access this overlay segment.
	NsxOverlaySegmentName *string `mandatory:"false" json:"nsxOverlaySegmentName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the `PrivateIp` object that is
	// the virtual IP (VIP) for the NSX Edge Uplink. Use this OCID as the route target for
	// route table rules when setting up connectivity between the SDDC and other networks.
	// For information about `PrivateIp` objects, see the Core Services API.
	NsxEdgeUplinkIpId *string `mandatory:"false" json:"nsxEdgeUplinkIpId"`

	// The date and time the SDDC was updated, in the format defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the SDDC.
	LifecycleState LifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m Sddc) String() string {
	return common.PointerString(m)
}
