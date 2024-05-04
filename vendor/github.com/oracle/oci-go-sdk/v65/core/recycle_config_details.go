// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RecycleConfigDetails The preferred recycle behavior for Bare Metal hosts associated with the reservation config.
type RecycleConfigDetails struct {

	// Preferred recycle level for hosts associated with the reservation config.
	// * `SKIP_RECYCLE` - Skips host wipe.
	// * `FULL_RECYCLE` - Does not skip host wipe. This is the default behavior.
	RecycleLevel RecycleConfigDetailsRecycleLevelEnum `mandatory:"true" json:"recycleLevel"`
}

func (m RecycleConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RecycleConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRecycleConfigDetailsRecycleLevelEnum(string(m.RecycleLevel)); !ok && m.RecycleLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RecycleLevel: %s. Supported values are: %s.", m.RecycleLevel, strings.Join(GetRecycleConfigDetailsRecycleLevelEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RecycleConfigDetailsRecycleLevelEnum Enum with underlying type: string
type RecycleConfigDetailsRecycleLevelEnum string

// Set of constants representing the allowable values for RecycleConfigDetailsRecycleLevelEnum
const (
	RecycleConfigDetailsRecycleLevelSkipRecycle RecycleConfigDetailsRecycleLevelEnum = "SKIP_RECYCLE"
	RecycleConfigDetailsRecycleLevelFullRecycle RecycleConfigDetailsRecycleLevelEnum = "FULL_RECYCLE"
)

var mappingRecycleConfigDetailsRecycleLevelEnum = map[string]RecycleConfigDetailsRecycleLevelEnum{
	"SKIP_RECYCLE": RecycleConfigDetailsRecycleLevelSkipRecycle,
	"FULL_RECYCLE": RecycleConfigDetailsRecycleLevelFullRecycle,
}

var mappingRecycleConfigDetailsRecycleLevelEnumLowerCase = map[string]RecycleConfigDetailsRecycleLevelEnum{
	"skip_recycle": RecycleConfigDetailsRecycleLevelSkipRecycle,
	"full_recycle": RecycleConfigDetailsRecycleLevelFullRecycle,
}

// GetRecycleConfigDetailsRecycleLevelEnumValues Enumerates the set of values for RecycleConfigDetailsRecycleLevelEnum
func GetRecycleConfigDetailsRecycleLevelEnumValues() []RecycleConfigDetailsRecycleLevelEnum {
	values := make([]RecycleConfigDetailsRecycleLevelEnum, 0)
	for _, v := range mappingRecycleConfigDetailsRecycleLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetRecycleConfigDetailsRecycleLevelEnumStringValues Enumerates the set of values in String for RecycleConfigDetailsRecycleLevelEnum
func GetRecycleConfigDetailsRecycleLevelEnumStringValues() []string {
	return []string{
		"SKIP_RECYCLE",
		"FULL_RECYCLE",
	}
}

// GetMappingRecycleConfigDetailsRecycleLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRecycleConfigDetailsRecycleLevelEnum(val string) (RecycleConfigDetailsRecycleLevelEnum, bool) {
	enum, ok := mappingRecycleConfigDetailsRecycleLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
