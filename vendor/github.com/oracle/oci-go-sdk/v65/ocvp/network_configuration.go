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

// NetworkConfiguration The network configurations used by Cluster, including
// OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the management subnet and VLANs.
type NetworkConfiguration struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the management subnet used
	// to provision the Cluster.
	ProvisioningSubnetId *string `mandatory:"true" json:"provisioningSubnetId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN used by the Cluster
	// for the vMotion component of the VMware environment.
	// This attribute is not guaranteed to reflect the vMotion VLAN
	// currently used by the ESXi hosts in the Cluster. The purpose
	// of this attribute is to show the vMotion VLAN that the Oracle
	// Cloud VMware Solution will use for any new ESXi hosts that you *add to this
	// Cluster in the future* with CreateEsxiHost.
	// Therefore, if you change the existing ESXi hosts in the Cluster to use a different VLAN
	// for the vMotion component of the VMware environment, you
	// should use UpdateCluster to update the Cluster's
	// `vmotionVlanId` with that new VLAN's OCID.
	VmotionVlanId *string `mandatory:"true" json:"vmotionVlanId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN used by the Cluster
	// for the vSAN component of the VMware environment.
	// This attribute is not guaranteed to reflect the vSAN VLAN
	// currently used by the ESXi hosts in the Cluster. The purpose
	// of this attribute is to show the vSAN VLAN that the Oracle
	// Cloud VMware Solution will use for any new ESXi hosts that you *add to this
	// Cluster in the future* with CreateEsxiHost.
	// Therefore, if you change the existing ESXi hosts in the Cluster to use a different VLAN
	// for the vSAN component of the VMware environment, you
	// should use UpdateCluster to update the Cluster's
	// `vsanVlanId` with that new VLAN's OCID.
	VsanVlanId *string `mandatory:"true" json:"vsanVlanId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN used by the Cluster
	// for the NSX VTEP component of the VMware environment.
	// This attribute is not guaranteed to reflect the NSX VTEP VLAN
	// currently used by the ESXi hosts in the Cluster. The purpose
	// of this attribute is to show the NSX VTEP VLAN that the Oracle
	// Cloud VMware Solution will use for any new ESXi hosts that you *add to this
	// Cluster in the future* with CreateEsxiHost.
	// Therefore, if you change the existing ESXi hosts in the Cluster to use a different VLAN
	// for the NSX VTEP component of the VMware environment, you
	// should use UpdateCluster to update the Cluster's
	// `nsxVTepVlanId` with that new VLAN's OCID.
	NsxVTepVlanId *string `mandatory:"true" json:"nsxVTepVlanId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN used by the Cluster
	// for the NSX Edge VTEP component of the VMware environment.
	// This attribute is not guaranteed to reflect the NSX Edge VTEP VLAN
	// currently used by the ESXi hosts in the Cluster. The purpose
	// of this attribute is to show the NSX Edge VTEP VLAN that the Oracle
	// Cloud VMware Solution will use for any new ESXi hosts that you *add to this
	// Cluster in the future* with CreateEsxiHost.
	// Therefore, if you change the existing ESXi hosts in the Cluster to use a different VLAN
	// for the NSX Edge VTEP component of the VMware environment, you
	// should use UpdateCluster to update the Cluster's
	// `nsxEdgeVTepVlanId` with that new VLAN's OCID.
	NsxEdgeVTepVlanId *string `mandatory:"true" json:"nsxEdgeVTepVlanId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC
	// for the vSphere component of the VMware environment. This VLAN is a mandatory attribute
	// for Management Cluster.
	// This attribute is not guaranteed to reflect the vSphere VLAN
	// currently used by the ESXi hosts in the Cluster. The purpose
	// of this attribute is to show the vSphere VLAN that the Oracle
	// Cloud VMware Solution will use for any new ESXi hosts that you *add to this
	// Cluster in the future* with CreateEsxiHost.
	// Therefore, if you change the existing ESXi hosts in the Cluster to use a different VLAN
	// for the vSphere component of the VMware environment, you
	// should use UpdateSddc to update the Cluster's
	// `vsphereVlanId` with that new VLAN's OCID.
	VsphereVlanId *string `mandatory:"false" json:"vsphereVlanId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC
	// for the NSX Edge Uplink 1 component of the VMware environment. This VLAN is a mandatory
	// attribute for Management Cluster.
	// This attribute is not guaranteed to reflect the NSX Edge Uplink 1 VLAN
	// currently used by the ESXi hosts in the Cluster. The purpose
	// of this attribute is to show the NSX Edge Uplink 1 VLAN that the Oracle
	// Cloud VMware Solution will use for any new ESXi hosts that you *add to this
	// Cluster in the future* with CreateEsxiHost.
	// Therefore, if you change the existing ESXi hosts in the Cluster to use a different VLAN
	// for the NSX Edge Uplink 1 component of the VMware environment, you
	// should use UpdateCluster to update the Cluster's
	// `nsxEdgeUplink1VlanId` with that new VLAN's OCID.
	NsxEdgeUplink1VlanId *string `mandatory:"false" json:"nsxEdgeUplink1VlanId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC
	// for the NSX Edge Uplink 2 component of the VMware environment. This VLAN is a mandatory
	// attribute for Management Cluster.
	// This attribute is not guaranteed to reflect the NSX Edge Uplink 2 VLAN
	// currently used by the ESXi hosts in the Cluster. The purpose
	// of this attribute is to show the NSX Edge Uplink 2 VLAN that the Oracle
	// Cloud VMware Solution will use for any new ESXi hosts that you *add to this
	// Cluster in the future* with CreateEsxiHost.
	// Therefore, if you change the existing ESXi hosts in the Cluster to use a different VLAN
	// for the NSX Edge Uplink 2 component of the VMware environment, you
	// should use UpdateCluster to update the Cluster's
	// `nsxEdgeUplink2VlanId` with that new VLAN's OCID.
	NsxEdgeUplink2VlanId *string `mandatory:"false" json:"nsxEdgeUplink2VlanId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN used by the Cluster
	// for the vSphere Replication component of the VMware environment.
	ReplicationVlanId *string `mandatory:"false" json:"replicationVlanId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN used by the Cluster
	// for the Provisioning component of the VMware environment.
	ProvisioningVlanId *string `mandatory:"false" json:"provisioningVlanId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC
	// for the HCX component of the VMware environment. This VLAN is a mandatory attribute
	// for Management Cluster when HCX is enabled.
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
}

func (m NetworkConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NetworkConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
