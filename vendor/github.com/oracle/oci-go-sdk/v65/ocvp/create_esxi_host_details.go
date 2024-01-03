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

// CreateEsxiHostDetails Details of the ESXi host to add to the Cluster.
type CreateEsxiHostDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Cluster to add the ESXi host to.
	ClusterId *string `mandatory:"true" json:"clusterId"`

	// A descriptive name for the ESXi host. It's changeable.
	// Esxi Host name requirements are 1-16 character length limit, Must start with a letter,
	// Must be English letters, numbers, - only, No repeating hyphens, Must be unique within the Cluster.
	// If this attribute is not specified, the Cluster's `instanceDisplayNamePrefix` attribute is used
	// to name and incrementally number the ESXi host. For example, if you're creating the fourth
	// ESXi host in the Cluster, and `instanceDisplayNamePrefix` is `MyCluster`, the host's display
	// name is `MyCluster-4`.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the deleted ESXi Host with LeftOver billing cycle.
	BillingDonorHostId *string `mandatory:"false" json:"billingDonorHostId"`

	// The billing option currently used by the ESXi host.
	// ListSupportedCommitments.
	CurrentCommitment CommitmentEnum `mandatory:"false" json:"currentCommitment,omitempty"`

	// The billing option to switch to after the existing billing cycle ends.
	// If `nextCommitment` is null or empty, `currentCommitment` continues to the next billing cycle.
	// ListSupportedCommitments.
	NextCommitment CommitmentEnum `mandatory:"false" json:"nextCommitment,omitempty"`

	// The availability domain to create the ESXi host in.
	// If keep empty, for AD-specific Cluster, new ESXi host will be created in the same availability domain;
	// for multi-AD Cluster, new ESXi host will be auto assigned to the next availability domain following evenly distribution strategy.
	ComputeAvailabilityDomain *string `mandatory:"false" json:"computeAvailabilityDomain"`

	// The compute shape name of the ESXi host.
	// ListSupportedHostShapes.
	HostShapeName *string `mandatory:"false" json:"hostShapeName"`

	// The OCPU count of the ESXi host.
	HostOcpuCount *float32 `mandatory:"false" json:"hostOcpuCount"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Capacity Reservation.
	CapacityReservationId *string `mandatory:"false" json:"capacityReservationId"`

	// The ESXi software bundle to install on the ESXi host.
	// Only versions under the same vmwareSoftwareVersion and have been validate by Oracle Cloud VMware Solution will be accepted.
	// To get a list of the available versions, use
	// ListSupportedVmwareSoftwareVersions.
	EsxiSoftwareVersion *string `mandatory:"false" json:"esxiSoftwareVersion"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateEsxiHostDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateEsxiHostDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCommitmentEnum(string(m.CurrentCommitment)); !ok && m.CurrentCommitment != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CurrentCommitment: %s. Supported values are: %s.", m.CurrentCommitment, strings.Join(GetCommitmentEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCommitmentEnum(string(m.NextCommitment)); !ok && m.NextCommitment != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NextCommitment: %s. Supported values are: %s.", m.NextCommitment, strings.Join(GetCommitmentEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
