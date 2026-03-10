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

// UpdateByolDetails The details to update BYOL.
type UpdateByolDetails struct {

	// A descriptive name for the BYOL.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Total quantity of licensed units for the specified `softwareType`:
	// - VCF, VDEFEND: number of OCPUs
	// - VSAN: storage capacity in TiB (tebibytes)
	// - AVI_LOAD_BALANCER: number of instances
	TotalUnits *int `mandatory:"false" json:"totalUnits"`

	// A description of the BYOL.
	Description *string `mandatory:"false" json:"description"`

	// The type of VMware software the BYOL applies to.
	// Supported values:
	// - VCF (VMware Cloud Foundation)
	// - VSAN (VMware vSAN)
	// - VDEFEND (VMware vDefend Firewall)
	// - AVI_LOAD_BALANCER (VMware Avi Load Balancer)
	SoftwareType ByolSoftwareTypeEnum `mandatory:"false" json:"softwareType,omitempty"`

	// The date and time when the BYOL becomes active. VMware software functionality cannot begin before this time.
	// In the format defined byRFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeTermStart *common.SDKTime `mandatory:"false" json:"timeTermStart"`

	// The date and time when the BYOL expires and becomes inactive.
	// In the format defined byRFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeTermEnd *common.SDKTime `mandatory:"false" json:"timeTermEnd"`

	// The Broadcom-supplied identifier of a BYOL license.
	EntitlementKey *string `mandatory:"false" json:"entitlementKey"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateByolDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateByolDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingByolSoftwareTypeEnum(string(m.SoftwareType)); !ok && m.SoftwareType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SoftwareType: %s. Supported values are: %s.", m.SoftwareType, strings.Join(GetByolSoftwareTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
