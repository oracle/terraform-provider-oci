// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
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

	// The number of Clusters in the SDDC.
	ClustersCount *int `mandatory:"true" json:"clustersCount"`

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

	// HCX configuration of the SDDC.
	HcxMode HcxModesEnum `mandatory:"true" json:"hcxMode"`

	InitialConfiguration *InitialConfiguration `mandatory:"true" json:"initialConfiguration"`

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

	// In general, this is a specific version of bundled ESXi software supported by
	// Oracle Cloud VMware Solution (see
	// ListSupportedVmwareSoftwareVersions).
	// This attribute is not guaranteed to reflect the version of
	// software currently installed on the ESXi hosts in the SDDC. The purpose
	// of this attribute is to show the version of software that the Oracle
	// Cloud VMware Solution will install on any new ESXi hosts that you *add to this
	// SDDC in the future* with CreateEsxiHost
	// unless a different version is configured on the Cluster or ESXi host level.
	// Therefore, if you upgrade the existing ESXi hosts in the SDDC to use a newer
	// version of bundled ESXi software supported by the Oracle Cloud VMware Solution, you
	// should use UpdateSddc to update the SDDC's
	// `vmwareSoftwareVersion` with that new version.
	EsxiSoftwareVersion *string `mandatory:"false" json:"esxiSoftwareVersion"`

	// The SDDC includes an administrator username and password for vCenter. You can
	// change this initial username to a different value in vCenter.
	VcenterUsername *string `mandatory:"false" json:"vcenterUsername"`

	// The SDDC includes an administrator username and initial password for NSX Manager. You
	// can change this initial username to a different value in NSX Manager.
	NsxManagerUsername *string `mandatory:"false" json:"nsxManagerUsername"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the `PrivateIp` object that is
	// the virtual IP (VIP) for the NSX Edge Uplink. Use this OCID as the route target for
	// route table rules when setting up connectivity between the SDDC and other networks.
	// For information about `PrivateIp` objects, see the Core Services API.
	NsxEdgeUplinkIpId *string `mandatory:"false" json:"nsxEdgeUplinkIpId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the `PrivateIp` object that is
	// the virtual IP (VIP) for HCX Manager. For information about `PrivateIp` objects, see the
	// Core Services API.
	HcxPrivateIpId *string `mandatory:"false" json:"hcxPrivateIpId"`

	// The FQDN for HCX Manager.
	// Example: `hcx-my-sddc.sddc.us-phoenix-1.oraclecloud.com`
	HcxFqdn *string `mandatory:"false" json:"hcxFqdn"`

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

	// Indicates whether this SDDC is designated for only single ESXi host.
	IsSingleHostSddc *bool `mandatory:"false" json:"isSingleHostSddc"`

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
	if _, ok := GetMappingHcxModesEnum(string(m.HcxMode)); !ok && m.HcxMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for HcxMode: %s. Supported values are: %s.", m.HcxMode, strings.Join(GetHcxModesEnumStringValues(), ",")))
	}

	if _, ok := GetMappingLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStatesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
