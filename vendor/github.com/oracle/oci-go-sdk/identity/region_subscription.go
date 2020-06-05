// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
)

// RegionSubscription An object that represents your tenancy's access to a particular region (i.e., a subscription), the status of that
// access, and whether that region is the home region. For more information, see Managing Regions (https://docs.cloud.oracle.com/Content/Identity/Tasks/managingregions.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access,
// see Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
type RegionSubscription struct {

	// The region's key. See Regions and Availability Domains (https://docs.cloud.oracle.com/Content/General/Concepts/regions.htm)
	// for the full list of supported 3-letter region codes.
	// Example: `PHX`
	RegionKey *string `mandatory:"true" json:"regionKey"`

	// The region's name. See Regions and Availability Domains (https://docs.cloud.oracle.com/Content/General/Concepts/regions.htm)
	// for the full list of supported region names.
	// Example: `us-phoenix-1`
	RegionName *string `mandatory:"true" json:"regionName"`

	// The region subscription status.
	Status RegionSubscriptionStatusEnum `mandatory:"true" json:"status"`

	// Indicates if the region is the home region or not.
	IsHomeRegion *bool `mandatory:"true" json:"isHomeRegion"`
}

func (m RegionSubscription) String() string {
	return common.PointerString(m)
}

// RegionSubscriptionStatusEnum Enum with underlying type: string
type RegionSubscriptionStatusEnum string

// Set of constants representing the allowable values for RegionSubscriptionStatusEnum
const (
	RegionSubscriptionStatusReady      RegionSubscriptionStatusEnum = "READY"
	RegionSubscriptionStatusInProgress RegionSubscriptionStatusEnum = "IN_PROGRESS"
)

var mappingRegionSubscriptionStatus = map[string]RegionSubscriptionStatusEnum{
	"READY":       RegionSubscriptionStatusReady,
	"IN_PROGRESS": RegionSubscriptionStatusInProgress,
}

// GetRegionSubscriptionStatusEnumValues Enumerates the set of values for RegionSubscriptionStatusEnum
func GetRegionSubscriptionStatusEnumValues() []RegionSubscriptionStatusEnum {
	values := make([]RegionSubscriptionStatusEnum, 0)
	for _, v := range mappingRegionSubscriptionStatus {
		values = append(values, v)
	}
	return values
}
