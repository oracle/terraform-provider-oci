// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Control Center Capacity Management API
//
// OCI Control Center (OCC) Capacity Management enables you to manage capacity requests in realms where OCI Control Center Capacity Management is available. For more information, see OCI Control Center (https://docs.cloud.oracle.com/iaas/Content/control-center/home.htm).
//

package capacitymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PlacementDetails Details like building, room and block where the resource was placed after provisioning in the datacenter.
type PlacementDetails struct {

	// The name of the region for which the resources were provisioned.
	Region *string `mandatory:"true" json:"region"`

	// The availability domain (AD) for which the resources were provisioned.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The type of workload to which these resources were provisioned.
	WorkloadType *string `mandatory:"true" json:"workloadType"`

	// The datacenter building where the resource was placed.
	Building *string `mandatory:"true" json:"building"`

	// The name of the room in the dataacenter building where the resource was placed.
	Room *string `mandatory:"true" json:"room"`

	// The block in the datacenter room where the resource was placed.
	Block *string `mandatory:"true" json:"block"`
}

func (m PlacementDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PlacementDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
