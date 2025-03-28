// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Data Plane API
//
// APIs for managing identity data plane services. For example, use this API to create a scoped-access security token. To manage identity domains (for example, creating or deleting an identity domain) or to manage resources (for example, users and groups) within the default identity domain, see IAM API (https://docs.oracle.com/iaas/api/#/en/identity/).
//

package identitydataplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ThickAuthorizationResponse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
