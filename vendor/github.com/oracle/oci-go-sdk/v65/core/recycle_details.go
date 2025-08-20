// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RecycleDetails Shows details about the last recycle performed on this host.
type RecycleDetails struct {

	// Preferred recycle level for hosts associated with the reservation config.
	// * `SKIP_RECYCLE` - Skips host wipe.
	// * `FULL_RECYCLE` - Does not skip host wipe. This is the default behavior.
	RecycleLevel RecycleDetailsRecycleLevelEnum `mandatory:"false" json:"recycleLevel,omitempty"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute host group this host was attached to at the time of recycle.
	ComputeHostGroupId *string `mandatory:"false" json:"computeHostGroupId"`
}

func (m RecycleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RecycleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRecycleDetailsRecycleLevelEnum(string(m.RecycleLevel)); !ok && m.RecycleLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RecycleLevel: %s. Supported values are: %s.", m.RecycleLevel, strings.Join(GetRecycleDetailsRecycleLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RecycleDetailsRecycleLevelEnum Enum with underlying type: string
type RecycleDetailsRecycleLevelEnum string

// Set of constants representing the allowable values for RecycleDetailsRecycleLevelEnum
const (
	RecycleDetailsRecycleLevelSkipRecycle RecycleDetailsRecycleLevelEnum = "SKIP_RECYCLE"
	RecycleDetailsRecycleLevelFullRecycle RecycleDetailsRecycleLevelEnum = "FULL_RECYCLE"
)

var mappingRecycleDetailsRecycleLevelEnum = map[string]RecycleDetailsRecycleLevelEnum{
	"SKIP_RECYCLE": RecycleDetailsRecycleLevelSkipRecycle,
	"FULL_RECYCLE": RecycleDetailsRecycleLevelFullRecycle,
}

var mappingRecycleDetailsRecycleLevelEnumLowerCase = map[string]RecycleDetailsRecycleLevelEnum{
	"skip_recycle": RecycleDetailsRecycleLevelSkipRecycle,
	"full_recycle": RecycleDetailsRecycleLevelFullRecycle,
}

// GetRecycleDetailsRecycleLevelEnumValues Enumerates the set of values for RecycleDetailsRecycleLevelEnum
func GetRecycleDetailsRecycleLevelEnumValues() []RecycleDetailsRecycleLevelEnum {
	values := make([]RecycleDetailsRecycleLevelEnum, 0)
	for _, v := range mappingRecycleDetailsRecycleLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetRecycleDetailsRecycleLevelEnumStringValues Enumerates the set of values in String for RecycleDetailsRecycleLevelEnum
func GetRecycleDetailsRecycleLevelEnumStringValues() []string {
	return []string{
		"SKIP_RECYCLE",
		"FULL_RECYCLE",
	}
}

// GetMappingRecycleDetailsRecycleLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRecycleDetailsRecycleLevelEnum(val string) (RecycleDetailsRecycleLevelEnum, bool) {
	enum, ok := mappingRecycleDetailsRecycleLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
