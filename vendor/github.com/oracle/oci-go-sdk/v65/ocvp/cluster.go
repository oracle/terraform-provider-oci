// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// Cluster An Oracle Cloud VMware Solution (https://docs.cloud.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm) Cluster contains the resources required for a
// functional VMware environment. Instances in a Cluster
// (see EsxiHost) run in a virtual cloud network (VCN)
// and are preconfigured with VMware and storage. Use the vCenter utility to manage
// and deploy VMware virtual machines (VMs) in the Cluster.
// The Cluster uses a single management subnet for provisioning the Cluster. It also uses a
// set of VLANs for various components of the VMware environment (vSphere, vMotion,
// vSAN, and so on). See the Core Services API for information about VCN subnets and VLANs.
type Cluster struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Cluster.
	Id *string `mandatory:"true" json:"id"`

	// The availability domain the ESXi hosts are running in. For Multi-AD Cluster, it is `multi-AD`.
	// Example: `Uocm:PHX-AD-1`, `multi-AD`
	ComputeAvailabilityDomain *string `mandatory:"true" json:"computeAvailabilityDomain"`

	// A descriptive name for the Cluster. It must be unique, start with a letter, and contain only letters, digits,
	// whitespaces, dashes and underscores.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// In general, this is a specific version of bundled VMware software supported by
	// Oracle Cloud VMware Solution (see
	// ListSupportedVmwareSoftwareVersions).
	// This attribute is not guaranteed to reflect the version of
	// software currently installed on the ESXi hosts in the Cluster. The purpose
	// of this attribute is to show the version of software that the Oracle
	// Cloud VMware Solution will install on any new ESXi hosts that you *add to this
	// Cluster in the future* with CreateEsxiHost.
	// Therefore, if you upgrade the existing ESXi hosts in the Cluster to use a newer
	// version of bundled VMware software supported by the Oracle Cloud VMware Solution, you
	// should use UpdateCluster to update the Cluster's
	// `vmwareSoftwareVersion` with that new version.
	VmwareSoftwareVersion *string `mandatory:"true" json:"vmwareSoftwareVersion"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that
	// contains the Cluster.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the SDDC that the
	// Cluster belongs to.
	SddcId *string `mandatory:"true" json:"sddcId"`

	// The number of ESXi hosts in the Cluster.
	EsxiHostsCount *int `mandatory:"true" json:"esxiHostsCount"`

	NetworkConfiguration *NetworkConfiguration `mandatory:"true" json:"networkConfiguration"`

	// The date and time the Cluster was created, in the format defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The initial compute shape of the Cluster's ESXi hosts.
	// ListSupportedHostShapes.
	InitialHostShapeName *string `mandatory:"true" json:"initialHostShapeName"`

	// vSphere Cluster types.
	VsphereType VsphereTypesEnum `mandatory:"true" json:"vsphereType"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// A prefix used in the name of each ESXi host and Compute instance in the Cluster.
	// If this isn't set, the Cluster's `displayName` is used as the prefix.
	// For example, if the value is `MyCluster`, the ESXi hosts are named `MyCluster-1`,
	// `MyCluster-2`, and so on.
	InstanceDisplayNamePrefix *string `mandatory:"false" json:"instanceDisplayNamePrefix"`

	// In general, this is a specific version of bundled ESXi software supported by
	// Oracle Cloud VMware Solution (see
	// ListSupportedVmwareSoftwareVersions).
	// This attribute is not guaranteed to reflect the version of
	// software currently installed on the ESXi hosts in the SDDC. The purpose
	// of this attribute is to show the version of software that the Oracle
	// Cloud VMware Solution will install on any new ESXi hosts that you *add to this
	// SDDC in the future* with CreateEsxiHost
	// unless a different version is configured on the ESXi host level.
	// Therefore, if you upgrade the existing ESXi hosts in the Cluster to use a newer
	// version of bundled ESXi software supported by the Oracle Cloud VMware Solution, you
	// should use UpdateCluster to update the Cluster's
	// `esxiSoftwareVersion` with that new version.
	EsxiSoftwareVersion *string `mandatory:"false" json:"esxiSoftwareVersion"`

	// The billing option selected during Cluster creation.
	// ListSupportedCommitments.
	InitialCommitment CommitmentEnum `mandatory:"false" json:"initialCommitment,omitempty"`

	// The CIDR block for the IP addresses that VMware VMs in the SDDC use to run application
	// workloads.
	WorkloadNetworkCidr *string `mandatory:"false" json:"workloadNetworkCidr"`

	// The date and time the Cluster was updated, in the format defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the Cluster.
	LifecycleState LifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The vSphere licenses to use when upgrading the Cluster.
	UpgradeLicenses []VsphereLicense `mandatory:"false" json:"upgradeLicenses"`

	// The links to binary objects needed to upgrade vSphere.
	VsphereUpgradeObjects []VsphereUpgradeObject `mandatory:"false" json:"vsphereUpgradeObjects"`

	// The initial OCPU count of the Cluster's ESXi hosts.
	InitialHostOcpuCount *float32 `mandatory:"false" json:"initialHostOcpuCount"`

	// Indicates whether shielded instance is enabled at the Cluster level.
	IsShieldedInstanceEnabled *bool `mandatory:"false" json:"isShieldedInstanceEnabled"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Capacity Reservation.
	CapacityReservationId *string `mandatory:"false" json:"capacityReservationId"`

	// Datastores used for the Cluster.
	Datastores []DatastoreDetails `mandatory:"false" json:"datastores"`
}

func (m Cluster) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Cluster) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVsphereTypesEnum(string(m.VsphereType)); !ok && m.VsphereType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VsphereType: %s. Supported values are: %s.", m.VsphereType, strings.Join(GetVsphereTypesEnumStringValues(), ",")))
	}

	if _, ok := GetMappingCommitmentEnum(string(m.InitialCommitment)); !ok && m.InitialCommitment != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InitialCommitment: %s. Supported values are: %s.", m.InitialCommitment, strings.Join(GetCommitmentEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStatesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
