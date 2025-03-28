// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Compute Cloud@Customer API
//
// Use the Compute Cloud@Customer API to manage Compute Cloud@Customer infrastructures and upgrade schedules.
// For more information see Compute Cloud@Customer documentation (https://docs.oracle.com/iaas/compute-cloud-at-customer/home.htm).
//

package computecloudatcustomer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CccInfrastructureRoutingStaticDetails Static routing information for a rack.
type CccInfrastructureRoutingStaticDetails struct {

	// The virtual local area network (VLAN) identifier used to connect to the uplink
	// (only access mode is supported).
	UplinkVlan *int `mandatory:"false" json:"uplinkVlan"`

	// The uplink Hot Standby Router Protocol (HSRP) group value for the switch in the
	// Compute Cloud@Customer infrastructure.
	UplinkHsrpGroup *int `mandatory:"false" json:"uplinkHsrpGroup"`
}

func (m CccInfrastructureRoutingStaticDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CccInfrastructureRoutingStaticDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
