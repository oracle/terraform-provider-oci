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

// ByolAllocation An allocation of Oracle Cloud VMware Solution (https://docs.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm)
// Bring-Your-Own-License (BYOL).
type ByolAllocation struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the BYOL Allocation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that
	// contains the BYOL Allocation.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A descriptive name for the BYOL Allocation.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the BYOL Allocation.
	LifecycleState ByolAllocationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The type of VMware software the BYOL applies to.
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

	// The Broadcom-supplied identifier of a BYOL license.
	EntitlementKey *string `mandatory:"true" json:"entitlementKey"`

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

func (m ByolAllocation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ByolAllocation) ValidateEnumValue() (bool, error) {
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

// ByolAllocationLifecycleStateEnum Enum with underlying type: string
type ByolAllocationLifecycleStateEnum string

// Set of constants representing the allowable values for ByolAllocationLifecycleStateEnum
const (
	ByolAllocationLifecycleStateCreating ByolAllocationLifecycleStateEnum = "CREATING"
	ByolAllocationLifecycleStateActive   ByolAllocationLifecycleStateEnum = "ACTIVE"
	ByolAllocationLifecycleStateInactive ByolAllocationLifecycleStateEnum = "INACTIVE"
	ByolAllocationLifecycleStateUpdating ByolAllocationLifecycleStateEnum = "UPDATING"
	ByolAllocationLifecycleStateDeleting ByolAllocationLifecycleStateEnum = "DELETING"
	ByolAllocationLifecycleStateDeleted  ByolAllocationLifecycleStateEnum = "DELETED"
	ByolAllocationLifecycleStateFailed   ByolAllocationLifecycleStateEnum = "FAILED"
)

var mappingByolAllocationLifecycleStateEnum = map[string]ByolAllocationLifecycleStateEnum{
	"CREATING": ByolAllocationLifecycleStateCreating,
	"ACTIVE":   ByolAllocationLifecycleStateActive,
	"INACTIVE": ByolAllocationLifecycleStateInactive,
	"UPDATING": ByolAllocationLifecycleStateUpdating,
	"DELETING": ByolAllocationLifecycleStateDeleting,
	"DELETED":  ByolAllocationLifecycleStateDeleted,
	"FAILED":   ByolAllocationLifecycleStateFailed,
}

var mappingByolAllocationLifecycleStateEnumLowerCase = map[string]ByolAllocationLifecycleStateEnum{
	"creating": ByolAllocationLifecycleStateCreating,
	"active":   ByolAllocationLifecycleStateActive,
	"inactive": ByolAllocationLifecycleStateInactive,
	"updating": ByolAllocationLifecycleStateUpdating,
	"deleting": ByolAllocationLifecycleStateDeleting,
	"deleted":  ByolAllocationLifecycleStateDeleted,
	"failed":   ByolAllocationLifecycleStateFailed,
}

// GetByolAllocationLifecycleStateEnumValues Enumerates the set of values for ByolAllocationLifecycleStateEnum
func GetByolAllocationLifecycleStateEnumValues() []ByolAllocationLifecycleStateEnum {
	values := make([]ByolAllocationLifecycleStateEnum, 0)
	for _, v := range mappingByolAllocationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetByolAllocationLifecycleStateEnumStringValues Enumerates the set of values in String for ByolAllocationLifecycleStateEnum
func GetByolAllocationLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingByolAllocationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingByolAllocationLifecycleStateEnum(val string) (ByolAllocationLifecycleStateEnum, bool) {
	enum, ok := mappingByolAllocationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ByolAllocationSoftwareTypeEnum Enum with underlying type: string
type ByolAllocationSoftwareTypeEnum string

// Set of constants representing the allowable values for ByolAllocationSoftwareTypeEnum
const (
	ByolAllocationSoftwareTypeVcf             ByolAllocationSoftwareTypeEnum = "VCF"
	ByolAllocationSoftwareTypeVsan            ByolAllocationSoftwareTypeEnum = "VSAN"
	ByolAllocationSoftwareTypeVdefend         ByolAllocationSoftwareTypeEnum = "VDEFEND"
	ByolAllocationSoftwareTypeAviLoadBalancer ByolAllocationSoftwareTypeEnum = "AVI_LOAD_BALANCER"
)

var mappingByolAllocationSoftwareTypeEnum = map[string]ByolAllocationSoftwareTypeEnum{
	"VCF":               ByolAllocationSoftwareTypeVcf,
	"VSAN":              ByolAllocationSoftwareTypeVsan,
	"VDEFEND":           ByolAllocationSoftwareTypeVdefend,
	"AVI_LOAD_BALANCER": ByolAllocationSoftwareTypeAviLoadBalancer,
}

var mappingByolAllocationSoftwareTypeEnumLowerCase = map[string]ByolAllocationSoftwareTypeEnum{
	"vcf":               ByolAllocationSoftwareTypeVcf,
	"vsan":              ByolAllocationSoftwareTypeVsan,
	"vdefend":           ByolAllocationSoftwareTypeVdefend,
	"avi_load_balancer": ByolAllocationSoftwareTypeAviLoadBalancer,
}

// GetByolAllocationSoftwareTypeEnumValues Enumerates the set of values for ByolAllocationSoftwareTypeEnum
func GetByolAllocationSoftwareTypeEnumValues() []ByolAllocationSoftwareTypeEnum {
	values := make([]ByolAllocationSoftwareTypeEnum, 0)
	for _, v := range mappingByolAllocationSoftwareTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetByolAllocationSoftwareTypeEnumStringValues Enumerates the set of values in String for ByolAllocationSoftwareTypeEnum
func GetByolAllocationSoftwareTypeEnumStringValues() []string {
	return []string{
		"VCF",
		"VSAN",
		"VDEFEND",
		"AVI_LOAD_BALANCER",
	}
}

// GetMappingByolAllocationSoftwareTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingByolAllocationSoftwareTypeEnum(val string) (ByolAllocationSoftwareTypeEnum, bool) {
	enum, ok := mappingByolAllocationSoftwareTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
