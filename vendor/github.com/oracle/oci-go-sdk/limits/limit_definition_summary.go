// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service limits APIs
//
// APIs that interact with the resource limits of a specific resource type
//

package limits

import (
	"github.com/oracle/oci-go-sdk/common"
)

// LimitDefinitionSummary The metadata specific to a resource limit definition.
type LimitDefinitionSummary struct {

	// The resource limit name. To be used for writing policies (in case of quotas) or other programmatic calls.
	Name *string `mandatory:"false" json:"name"`

	// The service name of the limit.
	ServiceName *string `mandatory:"false" json:"serviceName"`

	// The limit description.
	Description *string `mandatory:"false" json:"description"`

	// Reflects the scope of the resource limit: which can be Global (across all regions), regional or ad specific.
	ScopeType LimitDefinitionSummaryScopeTypeEnum `mandatory:"false" json:"scopeType,omitempty"`

	// If true, quota policies can be created on top of this resource limit.
	AreQuotasSupported *bool `mandatory:"false" json:"areQuotasSupported"`

	// Reflects if the GetResourceAvailability API is supported for this limit or not.
	// If not, the API will return an empty JSON response.
	IsResourceAvailabilitySupported *bool `mandatory:"false" json:"isResourceAvailabilitySupported"`
}

func (m LimitDefinitionSummary) String() string {
	return common.PointerString(m)
}

// LimitDefinitionSummaryScopeTypeEnum Enum with underlying type: string
type LimitDefinitionSummaryScopeTypeEnum string

// Set of constants representing the allowable values for LimitDefinitionSummaryScopeTypeEnum
const (
	LimitDefinitionSummaryScopeTypeGlobal LimitDefinitionSummaryScopeTypeEnum = "GLOBAL"
	LimitDefinitionSummaryScopeTypeRegion LimitDefinitionSummaryScopeTypeEnum = "REGION"
	LimitDefinitionSummaryScopeTypeAd     LimitDefinitionSummaryScopeTypeEnum = "AD"
)

var mappingLimitDefinitionSummaryScopeType = map[string]LimitDefinitionSummaryScopeTypeEnum{
	"GLOBAL": LimitDefinitionSummaryScopeTypeGlobal,
	"REGION": LimitDefinitionSummaryScopeTypeRegion,
	"AD":     LimitDefinitionSummaryScopeTypeAd,
}

// GetLimitDefinitionSummaryScopeTypeEnumValues Enumerates the set of values for LimitDefinitionSummaryScopeTypeEnum
func GetLimitDefinitionSummaryScopeTypeEnumValues() []LimitDefinitionSummaryScopeTypeEnum {
	values := make([]LimitDefinitionSummaryScopeTypeEnum, 0)
	for _, v := range mappingLimitDefinitionSummaryScopeType {
		values = append(values, v)
	}
	return values
}
