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

// ThickAuthorizationResponse The representation of ThickAuthorizationResponse
type ThickAuthorizationResponse struct {

	// The policy string related to the request
	Policy *string `mandatory:"true" json:"policy"`

	// The duration of how long this policy should be cached. Note that the type is of type java.time.Duration, not
	// string.
	PolicyCacheDuration *string `mandatory:"true" json:"policyCacheDuration"`

	// The policy string related to the request.
	Groups []string `mandatory:"true" json:"groups"`

	// The duration of how long the user's group membership should be cached. Note that the type is of type
	// java.time.Duration, not string.
	GroupMembershipCacheDuration *string `mandatory:"true" json:"groupMembershipCacheDuration"`

	// If set to true, the SDK should clear the caches.
	FlushAllCaches *bool `mandatory:"false" json:"flushAllCaches"`
}

func (m ThickAuthorizationResponse) String() string {
	return common.PointerString(m)
}
