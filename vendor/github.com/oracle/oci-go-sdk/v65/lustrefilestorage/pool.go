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

// Pool A construct that to represent compute capacity reservation(s) for a purpose (can be a default, or dedicated for customers).
type Pool struct {

	// The id of the pool
	Id *string `mandatory:"true" json:"id"`

	// The type of pool
	PoolType PoolPoolTypeEnum `mandatory:"true" json:"poolType"`

	// Name of the pool
	PoolName *string `mandatory:"true" json:"poolName"`

	// List of customer tenancies it is dedicated for
	DedicatedCustomerTenancies []string `mandatory:"false" json:"dedicatedCustomerTenancies"`

	// The name of the site group this pool is associated with
	SiteGroup *string `mandatory:"false" json:"siteGroup"`

	// List of customer tenancies it is dedicated for
	Tags []string `mandatory:"false" json:"tags"`

	// List of customer tenancies it is dedicated for
	Resources []interface{} `mandatory:"false" json:"resources"`

	// List of customer tenancies it is dedicated for
	Accounting *interface{} `mandatory:"false" json:"accounting"`

	// The pools that have affinity with this pool.
	PoolAffinities *interface{} `mandatory:"false" json:"poolAffinities"`
}

func (m Pool) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Pool) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPoolPoolTypeEnum(string(m.PoolType)); !ok && m.PoolType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PoolType: %s. Supported values are: %s.", m.PoolType, strings.Join(GetPoolPoolTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PoolPoolTypeEnum Enum with underlying type: string
type PoolPoolTypeEnum string

// Set of constants representing the allowable values for PoolPoolTypeEnum
const (
	PoolPoolTypeCompute PoolPoolTypeEnum = "COMPUTE"
	PoolPoolTypeBlock   PoolPoolTypeEnum = "BLOCK"
)

var mappingPoolPoolTypeEnum = map[string]PoolPoolTypeEnum{
	"COMPUTE": PoolPoolTypeCompute,
	"BLOCK":   PoolPoolTypeBlock,
}

var mappingPoolPoolTypeEnumLowerCase = map[string]PoolPoolTypeEnum{
	"compute": PoolPoolTypeCompute,
	"block":   PoolPoolTypeBlock,
}

// GetPoolPoolTypeEnumValues Enumerates the set of values for PoolPoolTypeEnum
func GetPoolPoolTypeEnumValues() []PoolPoolTypeEnum {
	values := make([]PoolPoolTypeEnum, 0)
	for _, v := range mappingPoolPoolTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPoolPoolTypeEnumStringValues Enumerates the set of values in String for PoolPoolTypeEnum
func GetPoolPoolTypeEnumStringValues() []string {
	return []string{
		"COMPUTE",
		"BLOCK",
	}
}

// GetMappingPoolPoolTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPoolPoolTypeEnum(val string) (PoolPoolTypeEnum, bool) {
	enum, ok := mappingPoolPoolTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
