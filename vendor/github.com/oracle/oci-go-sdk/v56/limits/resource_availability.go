// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Limits APIs
//
// APIs that interact with the resource limits of a specific resource type.
//

package limits

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ResourceAvailability The availability of a given resource limit, based on the usage, tenant service limits, and quotas set for the tenancy.
// Note: We cannot guarantee this data for all the limits. In such cases, these fields will be empty.
type ResourceAvailability struct {

	// The current usage in the given compartment. To support resources with fractional counts,
	// the field rounds up to the nearest integer.
	Used *int64 `mandatory:"false" json:"used"`

	// The count of available resources. To support resources with fractional counts,
	// the field rounds down to the nearest integer.
	Available *int64 `mandatory:"false" json:"available"`

	// The current most accurate usage in the given compartment.
	FractionalUsage *float32 `mandatory:"false" json:"fractionalUsage"`

	// The most accurate count of available resources.
	FractionalAvailability *float32 `mandatory:"false" json:"fractionalAvailability"`

	// The effective quota value for the given compartment. This field is only present if there is a
	// current quota policy affecting the current resource in the target region or availability domain.
	EffectiveQuotaValue *float32 `mandatory:"false" json:"effectiveQuotaValue"`
}

func (m ResourceAvailability) String() string {
	return common.PointerString(m)
}
