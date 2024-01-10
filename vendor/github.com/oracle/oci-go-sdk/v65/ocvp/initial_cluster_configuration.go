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

// InitialClusterConfiguration Details of the initial Cluster of SDDC.
type InitialClusterConfiguration struct {

	// vSphere Cluster types.
	VsphereType VsphereTypesEnum `mandatory:"true" json:"vsphereType"`

	// The availability domain to create the Cluster's ESXi hosts in. For multi-AD Cluster deployment, set to `multi-AD`.
	ComputeAvailabilityDomain *string `mandatory:"true" json:"computeAvailabilityDomain"`

	// The number of ESXi hosts to create in the Cluster. You can add more hosts later
	// (see CreateEsxiHost). Creating
	// a Cluster with a ESXi host count of 1 will be considered a single ESXi host Cluster.
	// **Note:** If you later delete EXSi hosts from a production Cluster to total less
	// than 3, you are still billed for the 3 minimum recommended ESXi hosts. Also,
	// you cannot add more VMware workloads to the Cluster until it again has at least
	// 3 ESXi hosts.
	EsxiHostsCount *int `mandatory:"true" json:"esxiHostsCount"`

	NetworkConfiguration *NetworkConfiguration `mandatory:"true" json:"networkConfiguration"`

	// A descriptive name for the Cluster.
	// Cluster name requirements are 1-16 character length limit, Must start with a letter, Must be English letters, numbers, - only, No repeating hyphens, Must be unique within the region.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A prefix used in the name of each ESXi host and Compute instance in the Cluster.
	// If this isn't set, the Cluster's `displayName` is used as the prefix.
	// For example, if the value is `myCluster`, the ESXi hosts are named `myCluster-1`,
	// `myCluster-2`, and so on.
	InstanceDisplayNamePrefix *string `mandatory:"false" json:"instanceDisplayNamePrefix"`

	// The billing option selected during Cluster creation.
	// ListSupportedCommitments.
	InitialCommitment CommitmentEnum `mandatory:"false" json:"initialCommitment,omitempty"`

	// The CIDR block for the IP addresses that VMware VMs in the Cluster use to run application
	// workloads.
	WorkloadNetworkCidr *string `mandatory:"false" json:"workloadNetworkCidr"`

	// The initial compute shape of the Cluster's ESXi hosts.
	// ListSupportedHostShapes.
	InitialHostShapeName *string `mandatory:"false" json:"initialHostShapeName"`

	// The initial OCPU count of the Cluster's ESXi hosts.
	InitialHostOcpuCount *float32 `mandatory:"false" json:"initialHostOcpuCount"`

	// Indicates whether shielded instance is enabled for this Cluster.
	IsShieldedInstanceEnabled *bool `mandatory:"false" json:"isShieldedInstanceEnabled"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Capacity Reservation.
	CapacityReservationId *string `mandatory:"false" json:"capacityReservationId"`

	// A list of datastore info for the Cluster.
	// This value is required only when `initialHostShapeName` is a standard shape.
	Datastores []DatastoreInfo `mandatory:"false" json:"datastores"`
}

func (m InitialClusterConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InitialClusterConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVsphereTypesEnum(string(m.VsphereType)); !ok && m.VsphereType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VsphereType: %s. Supported values are: %s.", m.VsphereType, strings.Join(GetVsphereTypesEnumStringValues(), ",")))
	}

	if _, ok := GetMappingCommitmentEnum(string(m.InitialCommitment)); !ok && m.InitialCommitment != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InitialCommitment: %s. Supported values are: %s.", m.InitialCommitment, strings.Join(GetCommitmentEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
