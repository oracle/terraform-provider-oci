// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage with Lustre API
//
// Use the File Storage with Lustre API to manage Lustre file systems and related resources. For more information, see File Storage with Lustre (https://docs.oracle.com/iaas/Content/lustre/home.htm).
//

package lustrefilestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CapacityReservationInfoSummary An object that gives more details about a capacity reservation.
type CapacityReservationInfoSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the capacity reservation.
	CapacityReservationId *string `mandatory:"true" json:"capacityReservationId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the LFS service CPG.
	LfsCpgId *string `mandatory:"true" json:"lfsCpgId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the customer CPG.
	CustomerCpgId *string `mandatory:"true" json:"customerCpgId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the customer tenancy.
	CustomerTenancyId *string `mandatory:"true" json:"customerTenancyId"`

	// Provisional cell capacity available for creating new filesystems on the cell. Measured in GB.
	AvailableBlockCapacityInGBs *int64 `mandatory:"true" json:"availableBlockCapacityInGBs"`

	DesiredComputeCount *DesiredComputeCount `mandatory:"true" json:"desiredComputeCount"`

	AvailableComputeCapacity *AvailableComputeCapacity `mandatory:"true" json:"availableComputeCapacity"`

	CurrentComputeCapacity *CurrentComputeCapacity `mandatory:"true" json:"currentComputeCapacity"`

	// If set to true, update capacity requests would not be sent.
	IsUpdateRequestPaused *bool `mandatory:"true" json:"isUpdateRequestPaused"`

	// A list of CPG OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) with block capacity
	// A maximum of 10 is allowed.
	BlockCpgIds []string `mandatory:"true" json:"blockCpgIds"`

	// The date and time the Capacity Reservation Info was created, expressed
	// in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2024-04-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the Capacity Reservation Info was updated, in the format defined
	// by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2024-04-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`
}

func (m CapacityReservationInfoSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CapacityReservationInfoSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
