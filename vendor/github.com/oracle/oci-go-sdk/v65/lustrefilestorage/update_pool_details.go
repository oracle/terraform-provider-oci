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

// UpdatePoolDetails The data required for updating a Pool.
type UpdatePoolDetails struct {

	// The type of pool
	PoolType UpdatePoolDetailsPoolTypeEnum `mandatory:"false" json:"poolType,omitempty"`

	// Name of the pool
	PoolName *string `mandatory:"false" json:"poolName"`

	// List of customer tenancies it is dedicated for
	DedicatedCustomerTenancies []string `mandatory:"false" json:"dedicatedCustomerTenancies"`

	// The name of the site group this pool is associated with
	SiteGroup *string `mandatory:"false" json:"siteGroup"`

	// List of customer tenancies it is dedicated for
	Tags []string `mandatory:"false" json:"tags"`

	// List of customer tenancies it is dedicated for
	Resources []interface{} `mandatory:"false" json:"resources"`

	// The pools that have affinity with this pool.
	PoolAffinities *interface{} `mandatory:"false" json:"poolAffinities"`
}

func (m UpdatePoolDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdatePoolDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdatePoolDetailsPoolTypeEnum(string(m.PoolType)); !ok && m.PoolType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PoolType: %s. Supported values are: %s.", m.PoolType, strings.Join(GetUpdatePoolDetailsPoolTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdatePoolDetailsPoolTypeEnum Enum with underlying type: string
type UpdatePoolDetailsPoolTypeEnum string

// Set of constants representing the allowable values for UpdatePoolDetailsPoolTypeEnum
const (
	UpdatePoolDetailsPoolTypeCompute UpdatePoolDetailsPoolTypeEnum = "COMPUTE"
	UpdatePoolDetailsPoolTypeBlock   UpdatePoolDetailsPoolTypeEnum = "BLOCK"
)

var mappingUpdatePoolDetailsPoolTypeEnum = map[string]UpdatePoolDetailsPoolTypeEnum{
	"COMPUTE": UpdatePoolDetailsPoolTypeCompute,
	"BLOCK":   UpdatePoolDetailsPoolTypeBlock,
}

var mappingUpdatePoolDetailsPoolTypeEnumLowerCase = map[string]UpdatePoolDetailsPoolTypeEnum{
	"compute": UpdatePoolDetailsPoolTypeCompute,
	"block":   UpdatePoolDetailsPoolTypeBlock,
}

// GetUpdatePoolDetailsPoolTypeEnumValues Enumerates the set of values for UpdatePoolDetailsPoolTypeEnum
func GetUpdatePoolDetailsPoolTypeEnumValues() []UpdatePoolDetailsPoolTypeEnum {
	values := make([]UpdatePoolDetailsPoolTypeEnum, 0)
	for _, v := range mappingUpdatePoolDetailsPoolTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdatePoolDetailsPoolTypeEnumStringValues Enumerates the set of values in String for UpdatePoolDetailsPoolTypeEnum
func GetUpdatePoolDetailsPoolTypeEnumStringValues() []string {
	return []string{
		"COMPUTE",
		"BLOCK",
	}
}

// GetMappingUpdatePoolDetailsPoolTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdatePoolDetailsPoolTypeEnum(val string) (UpdatePoolDetailsPoolTypeEnum, bool) {
	enum, ok := mappingUpdatePoolDetailsPoolTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
