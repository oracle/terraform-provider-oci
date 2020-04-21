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

// ResourceAvailability The availability of a given resource limit, based on the usage, tenant service limits and quotas set for the tenancy.
// Note: We cannot guarantee this data for all the limits. In those cases, these fields will be empty.
type ResourceAvailability struct {

	// The current usage in the given compartment.
	Used *int64 `mandatory:"false" json:"used"`

	// The count of available resources.
	Available *int64 `mandatory:"false" json:"available"`
}

func (m ResourceAvailability) String() string {
	return common.PointerString(m)
}
