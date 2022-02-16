// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Limits APIs
//
// APIs that interact with the resource limits of a specific resource type.
//

package limits

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// LimitValueSummary The value of a specific resource limit.
type LimitValueSummary struct {

	// The resource limit name. To be used for writing policies (in case of quotas) or other programmatic calls.
	Name *string `mandatory:"false" json:"name"`

	// The scope type of the limit.
	ScopeType LimitValueSummaryScopeTypeEnum `mandatory:"false" json:"scopeType,omitempty"`

	// If present, the returned value is only specific to this availability domain.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The resource limit value.
	Value *int64 `mandatory:"false" json:"value"`
}

func (m LimitValueSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LimitValueSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLimitValueSummaryScopeTypeEnum(string(m.ScopeType)); !ok && m.ScopeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ScopeType: %s. Supported values are: %s.", m.ScopeType, strings.Join(GetLimitValueSummaryScopeTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LimitValueSummaryScopeTypeEnum Enum with underlying type: string
type LimitValueSummaryScopeTypeEnum string

// Set of constants representing the allowable values for LimitValueSummaryScopeTypeEnum
const (
	LimitValueSummaryScopeTypeGlobal LimitValueSummaryScopeTypeEnum = "GLOBAL"
	LimitValueSummaryScopeTypeRegion LimitValueSummaryScopeTypeEnum = "REGION"
	LimitValueSummaryScopeTypeAd     LimitValueSummaryScopeTypeEnum = "AD"
)

var mappingLimitValueSummaryScopeTypeEnum = map[string]LimitValueSummaryScopeTypeEnum{
	"GLOBAL": LimitValueSummaryScopeTypeGlobal,
	"REGION": LimitValueSummaryScopeTypeRegion,
	"AD":     LimitValueSummaryScopeTypeAd,
}

// GetLimitValueSummaryScopeTypeEnumValues Enumerates the set of values for LimitValueSummaryScopeTypeEnum
func GetLimitValueSummaryScopeTypeEnumValues() []LimitValueSummaryScopeTypeEnum {
	values := make([]LimitValueSummaryScopeTypeEnum, 0)
	for _, v := range mappingLimitValueSummaryScopeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetLimitValueSummaryScopeTypeEnumStringValues Enumerates the set of values in String for LimitValueSummaryScopeTypeEnum
func GetLimitValueSummaryScopeTypeEnumStringValues() []string {
	return []string{
		"GLOBAL",
		"REGION",
		"AD",
	}
}

// GetMappingLimitValueSummaryScopeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLimitValueSummaryScopeTypeEnum(val string) (LimitValueSummaryScopeTypeEnum, bool) {
	mappingLimitValueSummaryScopeTypeEnumIgnoreCase := make(map[string]LimitValueSummaryScopeTypeEnum)
	for k, v := range mappingLimitValueSummaryScopeTypeEnum {
		mappingLimitValueSummaryScopeTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingLimitValueSummaryScopeTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
