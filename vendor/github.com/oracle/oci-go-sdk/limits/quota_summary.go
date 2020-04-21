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

// QuotaSummary Entails a subset of all the properties of the corresponding Quota and is recommended to be used in cases requiring
// security of quota details and slightly better API performance.
type QuotaSummary struct {

	// The OCID of the quota.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the resource this quota applies to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name you assign to the quota during creation. The name must be unique across all quotas
	// in the tenancy and cannot be changed.
	Name *string `mandatory:"true" json:"name"`

	// The description you assign to the quota.
	Description *string `mandatory:"true" json:"description"`

	// Date and time the quota was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The quota's current state. After creating a quota, make sure its `lifecycleState` is set to
	// ACTIVE before using it.
	LifecycleState QuotaSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m QuotaSummary) String() string {
	return common.PointerString(m)
}

// QuotaSummaryLifecycleStateEnum Enum with underlying type: string
type QuotaSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for QuotaSummaryLifecycleStateEnum
const (
	QuotaSummaryLifecycleStateActive QuotaSummaryLifecycleStateEnum = "ACTIVE"
)

var mappingQuotaSummaryLifecycleState = map[string]QuotaSummaryLifecycleStateEnum{
	"ACTIVE": QuotaSummaryLifecycleStateActive,
}

// GetQuotaSummaryLifecycleStateEnumValues Enumerates the set of values for QuotaSummaryLifecycleStateEnum
func GetQuotaSummaryLifecycleStateEnumValues() []QuotaSummaryLifecycleStateEnum {
	values := make([]QuotaSummaryLifecycleStateEnum, 0)
	for _, v := range mappingQuotaSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
