// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Service
//
// API for the Identity Dataplane
//

package identitydataplane

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// FilterGroupMembershipResult The representation of FilterGroupMembershipResult
type FilterGroupMembershipResult struct {

	// Return passed-in resolved principal object
	Principal *Principal `mandatory:"true" json:"principal"`

	// An array of group or dynamic group Ids which present the intersection between the passed-in group/dynamic-group and the actual group/dynamic-group the resovled principal belongs to.
	GroupIds []string `mandatory:"true" json:"groupIds"`
}

func (m FilterGroupMembershipResult) String() string {
	return common.PointerString(m)
}
