// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, policies, and identity domains.
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RegionSubscription An object that represents your tenancy's access to a particular region (i.e., a subscription), the status of that
// access, and whether that region is the home region. For more information, see Managing Regions (https://docs.cloud.oracle.com/Content/Identity/regions/managingregions.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access,
// see Get Started with Policies (https://docs.cloud.oracle.com/Content/Identity/policiesgs/get-started-with-policies.htm).
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RegionSubscription) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRegionSubscriptionStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetRegionSubscriptionStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RegionSubscriptionStatusEnum Enum with underlying type: string
type RegionSubscriptionStatusEnum string

// Set of constants representing the allowable values for RegionSubscriptionStatusEnum
const (
	RegionSubscriptionStatusReady      RegionSubscriptionStatusEnum = "READY"
	RegionSubscriptionStatusInProgress RegionSubscriptionStatusEnum = "IN_PROGRESS"
)

var mappingRegionSubscriptionStatusEnum = map[string]RegionSubscriptionStatusEnum{
	"READY":       RegionSubscriptionStatusReady,
	"IN_PROGRESS": RegionSubscriptionStatusInProgress,
}

var mappingRegionSubscriptionStatusEnumLowerCase = map[string]RegionSubscriptionStatusEnum{
	"ready":       RegionSubscriptionStatusReady,
	"in_progress": RegionSubscriptionStatusInProgress,
}

// GetRegionSubscriptionStatusEnumValues Enumerates the set of values for RegionSubscriptionStatusEnum
func GetRegionSubscriptionStatusEnumValues() []RegionSubscriptionStatusEnum {
	values := make([]RegionSubscriptionStatusEnum, 0)
	for _, v := range mappingRegionSubscriptionStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetRegionSubscriptionStatusEnumStringValues Enumerates the set of values in String for RegionSubscriptionStatusEnum
func GetRegionSubscriptionStatusEnumStringValues() []string {
	return []string{
		"READY",
		"IN_PROGRESS",
	}
}

// GetMappingRegionSubscriptionStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRegionSubscriptionStatusEnum(val string) (RegionSubscriptionStatusEnum, bool) {
	enum, ok := mappingRegionSubscriptionStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
