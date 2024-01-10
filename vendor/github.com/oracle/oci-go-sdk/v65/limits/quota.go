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

// Quota Quotas are applied on top of the service limits and inherited through the nested compartment hierarchy.
// Quotas allow compartment admins to limit resource consumption and set boundaries around acceptable resource use.
// The term "quota" can be interpreted as the following:
//   - An individual statement written in the declarative language.
//   - A collection of statements in a single, named "quota" object (which has an Oracle Cloud ID (OCID) assigned to it).
//   - The overall body of quotas your organization uses to control access to resources.
type Quota struct {

	// The OCID of the quota.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the resource this quota applies to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name you assign to the quota during creation. The name must be unique across all quotas
	// in the tenancy and cannot be changed.
	Name *string `mandatory:"true" json:"name"`

	// An array of one or more quota statements written in the declarative quota statement language.
	Statements []string `mandatory:"true" json:"statements"`

	// The description you assign to the quota.
	Description *string `mandatory:"true" json:"description"`

	// Date and time the quota was created, in the format defined by RFC 3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`

	// The quota's current state. After creating a quota, make sure its `lifecycleState` is set to
	// ACTIVE before using it.
	LifecycleState QuotaLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Quota) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Quota) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingQuotaLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetQuotaLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// QuotaLifecycleStateEnum Enum with underlying type: string
type QuotaLifecycleStateEnum string

// Set of constants representing the allowable values for QuotaLifecycleStateEnum
const (
	QuotaLifecycleStateActive QuotaLifecycleStateEnum = "ACTIVE"
)

var mappingQuotaLifecycleStateEnum = map[string]QuotaLifecycleStateEnum{
	"ACTIVE": QuotaLifecycleStateActive,
}

var mappingQuotaLifecycleStateEnumLowerCase = map[string]QuotaLifecycleStateEnum{
	"active": QuotaLifecycleStateActive,
}

// GetQuotaLifecycleStateEnumValues Enumerates the set of values for QuotaLifecycleStateEnum
func GetQuotaLifecycleStateEnumValues() []QuotaLifecycleStateEnum {
	values := make([]QuotaLifecycleStateEnum, 0)
	for _, v := range mappingQuotaLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetQuotaLifecycleStateEnumStringValues Enumerates the set of values in String for QuotaLifecycleStateEnum
func GetQuotaLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
	}
}

// GetMappingQuotaLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingQuotaLifecycleStateEnum(val string) (QuotaLifecycleStateEnum, bool) {
	enum, ok := mappingQuotaLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
