// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use the Oracle Cloud VMware API to create SDDCs and manage ESXi hosts and software.
// For more information, see Oracle Cloud VMware Solution (https://docs.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ByolAllocationSummary An allocation of Oracle Cloud VMware Solution (https://docs.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm)
// Bring-Your-Own-License (BYOL).
type ByolAllocationSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the BYOL Allocation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that
	// contains the BYOL Allocation.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A descriptive name for the BYOL Allocation.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the BYOL Allocation.
	// Accepted values are:
	// - CREATING
	// - ACTIVE
	// - INACTIVE
	// - UPDATING
	// - DELETING
	// - DELETED
	// - FAILED
	LifecycleState ByolAllocationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The type of VMware software the BYOL Allocation applies to.
	// Supported values:
	// - VCF (VMware Cloud Foundation)
	// - VSAN (VMware vSAN)
	// - VDEFEND (VMware vDefend Firewall)
	// - AVI_LOAD_BALANCER (VMware Avi Load Balancer)
	SoftwareType ByolAllocationSoftwareTypeEnum `mandatory:"true" json:"softwareType"`

	// The quantity of licensed units that allocated to this region.
	AllocatedUnits *int `mandatory:"true" json:"allocatedUnits"`

	// The quantity of licensed units that not yet consumed by resources.
	AvailableUnits *int `mandatory:"true" json:"availableUnits"`

	// The date and time when the BYOL Allocation becomes active. VMware software functionality cannot begin before this time.
	// In the format defined byRFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeTermStart *common.SDKTime `mandatory:"true" json:"timeTermStart"`

	// The date and time when the BYOL Allocation expires and becomes inactive.
	// In the format defined byRFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeTermEnd *common.SDKTime `mandatory:"true" json:"timeTermEnd"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the BYOL resource from which this BYOL Allocation is derived.
	ByolId *string `mandatory:"true" json:"byolId"`

	// The date and time the BYOL Allocation was created, in the format defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the BYOL Allocation was updated, in the format defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"true" json:"systemTags"`
}

func (m ByolAllocationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ByolAllocationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingByolAllocationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetByolAllocationLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingByolAllocationSoftwareTypeEnum(string(m.SoftwareType)); !ok && m.SoftwareType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SoftwareType: %s. Supported values are: %s.", m.SoftwareType, strings.Join(GetByolAllocationSoftwareTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
