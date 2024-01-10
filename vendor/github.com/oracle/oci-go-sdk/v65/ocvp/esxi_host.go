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

// EsxiHost An ESXi host is a node in a Cluster. At a minimum, each Cluster has 3 ESXi hosts
// that are used to implement a functioning VMware environment.
// In terms of implementation, an ESXi host is a Compute instance that
// is configured with the chosen bundle of VMware software.
// Notice that an `EsxiHost` object has its own OCID (`id`), and a separate
// attribute for the OCID of the Compute instance (`computeInstanceId`).
type EsxiHost struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the ESXi host.
	Id *string `mandatory:"true" json:"id"`

	// A descriptive name for the ESXi host. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the SDDC that the
	// ESXi host belongs to.
	SddcId *string `mandatory:"true" json:"sddcId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Cluster that the
	// ESXi host belongs to.
	ClusterId *string `mandatory:"true" json:"clusterId"`

	// The billing option currently used by the ESXi host.
	// ListSupportedCommitments.
	CurrentCommitment CommitmentEnum `mandatory:"true" json:"currentCommitment"`

	// The billing option to switch to after the current billing cycle ends.
	// If `nextCommitment` is null or empty, `currentCommitment` continues to the next billing cycle.
	// ListSupportedCommitments.
	NextCommitment CommitmentEnum `mandatory:"true" json:"nextCommitment"`

	// Current billing cycle end date. If the value in `currentCommitment` and `nextCommitment` are different, the value specified in `nextCommitment`
	// becomes the new `currentCommitment` when the `contractEndDate` is reached.
	// Example: `2016-08-25T21:10:29.600Z`
	BillingContractEndDate *common.SDKTime `mandatory:"true" json:"billingContractEndDate"`

	// The version of VMware software that Oracle Cloud VMware Solution installed on the ESXi hosts.
	VmwareSoftwareVersion *string `mandatory:"true" json:"vmwareSoftwareVersion"`

	// The availability domain of the ESXi host.
	ComputeAvailabilityDomain *string `mandatory:"true" json:"computeAvailabilityDomain"`

	// The compute shape name of the ESXi host.
	// ListSupportedHostShapes.
	HostShapeName *string `mandatory:"true" json:"hostShapeName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that
	// contains the Cluster.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// In terms of implementation, an ESXi host is a Compute instance that
	// is configured with the chosen bundle of VMware software. The `computeInstanceId`
	// is the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of that Compute instance.
	ComputeInstanceId *string `mandatory:"false" json:"computeInstanceId"`

	// The date and time the ESXi host was created, in the format defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the ESXi host was updated, in the format defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the ESXi host.
	LifecycleState LifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the deleted ESXi Host with LeftOver billing cycle.
	BillingDonorHostId *string `mandatory:"false" json:"billingDonorHostId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the active ESXi Host to swap billing with current host.
	SwapBillingHostId *string `mandatory:"false" json:"swapBillingHostId"`

	// Indicates whether this host is in the progress of billing continuation.
	IsBillingContinuationInProgress *bool `mandatory:"false" json:"isBillingContinuationInProgress"`

	// Indicates whether this host is in the progress of swapping billing.
	IsBillingSwappingInProgress *bool `mandatory:"false" json:"isBillingSwappingInProgress"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the ESXi host that failed.
	FailedEsxiHostId *string `mandatory:"false" json:"failedEsxiHostId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the ESXi host that
	// is created to replace the failed host.
	ReplacementEsxiHostId *string `mandatory:"false" json:"replacementEsxiHostId"`

	// The date and time when the new esxi host should start billing cycle.
	// RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2021-07-25T21:10:29.600Z`
	GracePeriodEndDate *common.SDKTime `mandatory:"false" json:"gracePeriodEndDate"`

	// The version of ESXi software that Oracle Cloud VMware Solution installed on the ESXi hosts.
	EsxiSoftwareVersion *string `mandatory:"false" json:"esxiSoftwareVersion"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the ESXi host that
	// will be upgraded.
	NonUpgradedEsxiHostId *string `mandatory:"false" json:"nonUpgradedEsxiHostId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the ESXi host that
	// is newly created to upgrade the original host.
	UpgradedReplacementEsxiHostId *string `mandatory:"false" json:"upgradedReplacementEsxiHostId"`

	// The OCPU count of the ESXi host.
	HostOcpuCount *float32 `mandatory:"false" json:"hostOcpuCount"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Capacity Reservation.
	CapacityReservationId *string `mandatory:"false" json:"capacityReservationId"`
}

func (m EsxiHost) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EsxiHost) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCommitmentEnum(string(m.CurrentCommitment)); !ok && m.CurrentCommitment != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CurrentCommitment: %s. Supported values are: %s.", m.CurrentCommitment, strings.Join(GetCommitmentEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCommitmentEnum(string(m.NextCommitment)); !ok && m.NextCommitment != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NextCommitment: %s. Supported values are: %s.", m.NextCommitment, strings.Join(GetCommitmentEnumStringValues(), ",")))
	}

	if _, ok := GetMappingLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStatesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
