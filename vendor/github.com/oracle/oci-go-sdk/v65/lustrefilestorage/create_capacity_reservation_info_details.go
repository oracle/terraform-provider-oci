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

// CreateCapacityReservationInfoDetails The data required for creating a Capacity Reservation Info.
type CreateCapacityReservationInfoDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the capacity reservation.
	CapacityReservationId *string `mandatory:"true" json:"capacityReservationId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the LFS service CPG.
	LfsCpgId *string `mandatory:"false" json:"lfsCpgId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the customer CPG.
	CustomerCpgId *string `mandatory:"false" json:"customerCpgId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the customer tenancy.
	CustomerTenancyId *string `mandatory:"false" json:"customerTenancyId"`

	// Provisional cell capacity available for creating new filesystems on the cell. Measured in GB.
	AvailableBlockCapacityInGBs *int64 `mandatory:"false" json:"availableBlockCapacityInGBs"`

	DesiredComputeCount *DesiredComputeCount `mandatory:"false" json:"desiredComputeCount"`

	// If set to true, update capacity requests would not be sent.
	IsUpdateRequestPaused *bool `mandatory:"false" json:"isUpdateRequestPaused"`

	// A list of CPG OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) with block capacity
	// A maximum of 10 is allowed.
	BlockCpgIds []string `mandatory:"false" json:"blockCpgIds"`
}

func (m CreateCapacityReservationInfoDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateCapacityReservationInfoDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
