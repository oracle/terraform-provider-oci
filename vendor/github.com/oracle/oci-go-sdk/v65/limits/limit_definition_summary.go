// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Limits APIs
//
// APIs that interact with the resource limits of a specific resource type.
//

package limits

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LimitDefinitionSummary The metadata specific to a resource limit definition.
type LimitDefinitionSummary struct {

	// The resource limit name. To be used for writing policies (in case of quotas) or other programmatic calls.
	Name *string `mandatory:"false" json:"name"`

	// The service name of the limit.
	ServiceName *string `mandatory:"false" json:"serviceName"`

	// The limit description.
	Description *string `mandatory:"false" json:"description"`

	// Reflects the scope of the resource limit, whether Global (across all regions), regional, or availability domain-specific.
	ScopeType LimitDefinitionSummaryScopeTypeEnum `mandatory:"false" json:"scopeType,omitempty"`

	// If true, quota policies can be created on top of this resource limit.
	AreQuotasSupported *bool `mandatory:"false" json:"areQuotasSupported"`

	// Reflects whether or not the GetResourceAvailability API is supported for this limit.
	// If not, the API returns an empty JSON response.
	IsResourceAvailabilitySupported *bool `mandatory:"false" json:"isResourceAvailabilitySupported"`

	// Indicates if the limit has been deprecated.
	IsDeprecated *bool `mandatory:"false" json:"isDeprecated"`

	// Indicates if the customer can request a limit increase for this resource.
	IsEligibleForLimitIncrease *bool `mandatory:"false" json:"isEligibleForLimitIncrease"`

	// The limit for this resource has a dynamic value that is based on consumption across all OCI services.
	IsDynamic *bool `mandatory:"false" json:"isDynamic"`

	// An array of subscription types supported by the limit. e,g The type of subscription, such as 'SAAS', 'ERP', 'CRM'.
	SupportedSubscriptions []string `mandatory:"false" json:"supportedSubscriptions"`

	// Supported quota family names for creation of quota policy.
	SupportedQuotaFamilies []string `mandatory:"false" json:"supportedQuotaFamilies"`
}

func (m LimitDefinitionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LimitDefinitionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLimitDefinitionSummaryScopeTypeEnum(string(m.ScopeType)); !ok && m.ScopeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ScopeType: %s. Supported values are: %s.", m.ScopeType, strings.Join(GetLimitDefinitionSummaryScopeTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LimitDefinitionSummaryScopeTypeEnum Enum with underlying type: string
type LimitDefinitionSummaryScopeTypeEnum string

// Set of constants representing the allowable values for LimitDefinitionSummaryScopeTypeEnum
const (
	LimitDefinitionSummaryScopeTypeGlobal LimitDefinitionSummaryScopeTypeEnum = "GLOBAL"
	LimitDefinitionSummaryScopeTypeRegion LimitDefinitionSummaryScopeTypeEnum = "REGION"
	LimitDefinitionSummaryScopeTypeAd     LimitDefinitionSummaryScopeTypeEnum = "AD"
)

var mappingLimitDefinitionSummaryScopeTypeEnum = map[string]LimitDefinitionSummaryScopeTypeEnum{
	"GLOBAL": LimitDefinitionSummaryScopeTypeGlobal,
	"REGION": LimitDefinitionSummaryScopeTypeRegion,
	"AD":     LimitDefinitionSummaryScopeTypeAd,
}

var mappingLimitDefinitionSummaryScopeTypeEnumLowerCase = map[string]LimitDefinitionSummaryScopeTypeEnum{
	"global": LimitDefinitionSummaryScopeTypeGlobal,
	"region": LimitDefinitionSummaryScopeTypeRegion,
	"ad":     LimitDefinitionSummaryScopeTypeAd,
}

// GetLimitDefinitionSummaryScopeTypeEnumValues Enumerates the set of values for LimitDefinitionSummaryScopeTypeEnum
func GetLimitDefinitionSummaryScopeTypeEnumValues() []LimitDefinitionSummaryScopeTypeEnum {
	values := make([]LimitDefinitionSummaryScopeTypeEnum, 0)
	for _, v := range mappingLimitDefinitionSummaryScopeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetLimitDefinitionSummaryScopeTypeEnumStringValues Enumerates the set of values in String for LimitDefinitionSummaryScopeTypeEnum
func GetLimitDefinitionSummaryScopeTypeEnumStringValues() []string {
	return []string{
		"GLOBAL",
		"REGION",
		"AD",
	}
}

// GetMappingLimitDefinitionSummaryScopeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLimitDefinitionSummaryScopeTypeEnum(val string) (LimitDefinitionSummaryScopeTypeEnum, bool) {
	enum, ok := mappingLimitDefinitionSummaryScopeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
