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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// Sddc An Oracle Cloud VMware Solution (https://docs.cloud.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm) software-defined data center (SDDC) contains the resources required for a
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

	// The availability domain the ESXi hosts are running in. For Multi-AD SDDC, it is `multi-AD`.
	// Example: `Uocm:PHX-AD-1`, `multi-AD`
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

	// The FQDN for vCenter.
	// Example: `vcenter-my-sddc.sddc.us-phoenix-1.oraclecloud.com`
	VcenterFqdn *string `mandatory:"true" json:"vcenterFqdn"`

	// The FQDN for NSX Manager.
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

	// The billing option selected during SDDC creation.
	// ListSupportedSkus.
	InitialSku SkuEnum `mandatory:"false" json:"initialSku,omitempty"`

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

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC
	// for the vSphere Replication component of the VMware environment.
	ReplicationVlanId *string `mandatory:"false" json:"replicationVlanId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC
	// for the Provisioning component of the VMware environment.
	ProvisioningVlanId *string `mandatory:"false" json:"provisioningVlanId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the `PrivateIp` object that is
	// the virtual IP (VIP) for HCX Manager. For information about `PrivateIp` objects, see the
	// Core Services API.
	HcxPrivateIpId *string `mandatory:"false" json:"hcxPrivateIpId"`

	// The FQDN for HCX Manager.
	// Example: `hcx-my-sddc.sddc.us-phoenix-1.oraclecloud.com`
	HcxFqdn *string `mandatory:"false" json:"hcxFqdn"`

	// The SDDC includes an administrator username and initial password for HCX Manager. Make sure
	// to change this initial HCX Manager password to a different value.
	HcxInitialPassword *string `mandatory:"false" json:"hcxInitialPassword"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC
	// for the HCX component of the VMware environment.
	// This attribute is not guaranteed to reflect the HCX VLAN
	// currently used by the ESXi hosts in the SDDC. The purpose
	// of this attribute is to show the HCX VLAN that the Oracle
	// Cloud VMware Solution will use for any new ESXi hosts that you *add to this
	// SDDC in the future* with CreateEsxiHost.
	// Therefore, if you change the existing ESXi hosts in the SDDC to use a different VLAN
	// for the HCX component of the VMware environment, you
	// should use UpdateSddc to update the SDDC's
	// `hcxVlanId` with that new VLAN's OCID.
	HcxVlanId *string `mandatory:"false" json:"hcxVlanId"`

	// Indicates whether HCX is enabled for this SDDC.
	IsHcxEnabled *bool `mandatory:"false" json:"isHcxEnabled"`

	// The activation keys to use on the on-premises HCX Enterprise appliances you site pair with HCX Manager in your VMware Solution.
	// The number of keys provided depends on the HCX license type. HCX Advanced provides 3 activation keys.
	// HCX Enterprise provides 10 activation keys.
	HcxOnPremKey *string `mandatory:"false" json:"hcxOnPremKey"`

	// Indicates whether HCX Enterprise is enabled for this SDDC.
	IsHcxEnterpriseEnabled *bool `mandatory:"false" json:"isHcxEnterpriseEnabled"`

	// Indicates whether SDDC is pending downgrade from HCX Enterprise to HCX Advanced.
	IsHcxPendingDowngrade *bool `mandatory:"false" json:"isHcxPendingDowngrade"`

	// The activation licenses to use on the on-premises HCX Enterprise appliance you site pair with HCX Manager in your VMware Solution.
	HcxOnPremLicenses []HcxLicenseSummary `mandatory:"false" json:"hcxOnPremLicenses"`

	// The date and time current HCX Enterprise billing cycle ends, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeHcxBillingCycleEnd *common.SDKTime `mandatory:"false" json:"timeHcxBillingCycleEnd"`

	// The date and time the SDDC's HCX on-premise license status was updated, in the format defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeHcxLicenseStatusUpdated *common.SDKTime `mandatory:"false" json:"timeHcxLicenseStatusUpdated"`

	// The date and time the SDDC was updated, in the format defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the SDDC.
	LifecycleState LifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m Sddc) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Sddc) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSkuEnum(string(m.InitialSku)); !ok && m.InitialSku != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InitialSku: %s. Supported values are: %s.", m.InitialSku, strings.Join(GetSkuEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStatesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
