// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Placement Describes the placement of an instance.
type Placement struct {

	// The affinity setting for the instance on the Dedicated Host.
	Affinity *string `mandatory:"false" json:"affinity"`

	// The Availability Zone of the instance.
	AvailabilityZone *string `mandatory:"false" json:"availabilityZone"`

	// The name of the placement group the instance is in.
	GroupName *string `mandatory:"false" json:"groupName"`

	// The ID of the Dedicated Host on which the instance resides.
	HostKey *string `mandatory:"false" json:"hostKey"`

	// The ARN of the host resource group in which to launch the instances.
	HostResourceGroupArn *string `mandatory:"false" json:"hostResourceGroupArn"`

	// The number of the partition that the instance is in.
	PartitionNumber *int `mandatory:"false" json:"partitionNumber"`

	// Reserved for future use.
	SpreadDomain *string `mandatory:"false" json:"spreadDomain"`

	// The tenancy of the instance (if the instance is running in a VPC).
	Tenancy *string `mandatory:"false" json:"tenancy"`
}

func (m Placement) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Placement) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
