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

// ByolRealmAllocationSummary Represents a BYOL allocation for a specific region
type ByolRealmAllocationSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the BYOL resource from which this BYOL Allocation is derived.
	ByolId *string `mandatory:"true" json:"byolId"`

	// The region name for this allocation, in the format defined by
	// Region Identifier (https://docs.oracle.com/en-us/iaas/Content/General/Concepts/regions.htm).
	// Example: `ap-sydney-1`
	Region *string `mandatory:"true" json:"region"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that
	// contains the BYOL Allocation.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The quantity of licensed units that allocated to this region.
	AllocatedUnits *int `mandatory:"true" json:"allocatedUnits"`

	// The date and time the BYOL Allocation was created, in the format defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the BYOL Allocation was updated, in the format defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`
}

func (m ByolRealmAllocationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ByolRealmAllocationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
