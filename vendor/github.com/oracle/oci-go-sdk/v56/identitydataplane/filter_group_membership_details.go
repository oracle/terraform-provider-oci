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

// FilterGroupMembershipDetails The representation of FilterGroupMembershipDetails
type FilterGroupMembershipDetails struct {

	// A resolved principal object
	Principal *Principal `mandatory:"true" json:"principal"`

	// An array of group or dynamic group Ids the resolved principal potentially belongs to.
	GroupIds []string `mandatory:"true" json:"groupIds"`
}

func (m FilterGroupMembershipDetails) String() string {
	return common.PointerString(m)
}
