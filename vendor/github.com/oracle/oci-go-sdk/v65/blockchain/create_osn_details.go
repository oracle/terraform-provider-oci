// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Blockchain Platform Control Plane API
//
// Blockchain Platform Control Plane API
//

package blockchain

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateOsnDetails The Ordering Service Node details to be added
type CreateOsnDetails struct {

	// Availability Domain to place new OSN
	Ad AvailabilityDomainAdsEnum `mandatory:"true" json:"ad"`

	OcpuAllocationParam *OcpuAllocationNumberParam `mandatory:"false" json:"ocpuAllocationParam"`
}

func (m CreateOsnDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOsnDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAvailabilityDomainAdsEnum(string(m.Ad)); !ok && m.Ad != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Ad: %s. Supported values are: %s.", m.Ad, strings.Join(GetAvailabilityDomainAdsEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
