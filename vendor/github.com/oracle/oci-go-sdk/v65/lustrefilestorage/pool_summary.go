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

// PoolSummary A list of pools.
type PoolSummary struct {

	// The id of the pool
	PoolId *string `mandatory:"false" json:"poolId"`

	// The type of pool
	PoolType PoolSummaryPoolTypeEnum `mandatory:"false" json:"poolType,omitempty"`

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

	// List of customer tenancies it is dedicated for
	Accounting *interface{} `mandatory:"false" json:"accounting"`

	// The pools that have affinity with this pool.
	PoolAffinities *interface{} `mandatory:"false" json:"poolAffinities"`
}

func (m PoolSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PoolSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPoolSummaryPoolTypeEnum(string(m.PoolType)); !ok && m.PoolType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PoolType: %s. Supported values are: %s.", m.PoolType, strings.Join(GetPoolSummaryPoolTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PoolSummaryPoolTypeEnum Enum with underlying type: string
type PoolSummaryPoolTypeEnum string

// Set of constants representing the allowable values for PoolSummaryPoolTypeEnum
const (
	PoolSummaryPoolTypeCompute PoolSummaryPoolTypeEnum = "COMPUTE"
	PoolSummaryPoolTypeBlock   PoolSummaryPoolTypeEnum = "BLOCK"
)

var mappingPoolSummaryPoolTypeEnum = map[string]PoolSummaryPoolTypeEnum{
	"COMPUTE": PoolSummaryPoolTypeCompute,
	"BLOCK":   PoolSummaryPoolTypeBlock,
}

var mappingPoolSummaryPoolTypeEnumLowerCase = map[string]PoolSummaryPoolTypeEnum{
	"compute": PoolSummaryPoolTypeCompute,
	"block":   PoolSummaryPoolTypeBlock,
}

// GetPoolSummaryPoolTypeEnumValues Enumerates the set of values for PoolSummaryPoolTypeEnum
func GetPoolSummaryPoolTypeEnumValues() []PoolSummaryPoolTypeEnum {
	values := make([]PoolSummaryPoolTypeEnum, 0)
	for _, v := range mappingPoolSummaryPoolTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPoolSummaryPoolTypeEnumStringValues Enumerates the set of values in String for PoolSummaryPoolTypeEnum
func GetPoolSummaryPoolTypeEnumStringValues() []string {
	return []string{
		"COMPUTE",
		"BLOCK",
	}
}

// GetMappingPoolSummaryPoolTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPoolSummaryPoolTypeEnum(val string) (PoolSummaryPoolTypeEnum, bool) {
	enum, ok := mappingPoolSummaryPoolTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
