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

// CreatePoolDetails The data required for creating a Pool.
type CreatePoolDetails struct {

	// The type of pool
	PoolType CreatePoolDetailsPoolTypeEnum `mandatory:"true" json:"poolType"`

	// Name of the pool
	PoolName *string `mandatory:"true" json:"poolName"`

	// The name of the site group this pool is associated with
	SiteGroup *string `mandatory:"true" json:"siteGroup"`

	// List of customer tenancies it is dedicated for
	Resources []interface{} `mandatory:"true" json:"resources"`

	// List of customer tenancies it is dedicated for
	DedicatedCustomerTenancies []string `mandatory:"false" json:"dedicatedCustomerTenancies"`

	// List of customer tenancies it is dedicated for
	Tags []string `mandatory:"false" json:"tags"`

	// The pools that have affinity with this pool.
	PoolAffinities *interface{} `mandatory:"false" json:"poolAffinities"`
}

func (m CreatePoolDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreatePoolDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreatePoolDetailsPoolTypeEnum(string(m.PoolType)); !ok && m.PoolType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PoolType: %s. Supported values are: %s.", m.PoolType, strings.Join(GetCreatePoolDetailsPoolTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreatePoolDetailsPoolTypeEnum Enum with underlying type: string
type CreatePoolDetailsPoolTypeEnum string

// Set of constants representing the allowable values for CreatePoolDetailsPoolTypeEnum
const (
	CreatePoolDetailsPoolTypeCompute CreatePoolDetailsPoolTypeEnum = "COMPUTE"
	CreatePoolDetailsPoolTypeBlock   CreatePoolDetailsPoolTypeEnum = "BLOCK"
)

var mappingCreatePoolDetailsPoolTypeEnum = map[string]CreatePoolDetailsPoolTypeEnum{
	"COMPUTE": CreatePoolDetailsPoolTypeCompute,
	"BLOCK":   CreatePoolDetailsPoolTypeBlock,
}

var mappingCreatePoolDetailsPoolTypeEnumLowerCase = map[string]CreatePoolDetailsPoolTypeEnum{
	"compute": CreatePoolDetailsPoolTypeCompute,
	"block":   CreatePoolDetailsPoolTypeBlock,
}

// GetCreatePoolDetailsPoolTypeEnumValues Enumerates the set of values for CreatePoolDetailsPoolTypeEnum
func GetCreatePoolDetailsPoolTypeEnumValues() []CreatePoolDetailsPoolTypeEnum {
	values := make([]CreatePoolDetailsPoolTypeEnum, 0)
	for _, v := range mappingCreatePoolDetailsPoolTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreatePoolDetailsPoolTypeEnumStringValues Enumerates the set of values in String for CreatePoolDetailsPoolTypeEnum
func GetCreatePoolDetailsPoolTypeEnumStringValues() []string {
	return []string{
		"COMPUTE",
		"BLOCK",
	}
}

// GetMappingCreatePoolDetailsPoolTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreatePoolDetailsPoolTypeEnum(val string) (CreatePoolDetailsPoolTypeEnum, bool) {
	enum, ok := mappingCreatePoolDetailsPoolTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
