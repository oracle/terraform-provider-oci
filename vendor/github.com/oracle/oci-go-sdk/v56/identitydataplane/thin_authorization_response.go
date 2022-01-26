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

// ThinAuthorizationResponse The representation of ThinAuthorizationResponse
type ThinAuthorizationResponse struct {

	// The policy string related to the request.
	AuthorizationRequest *AuthorizationRequest `mandatory:"true" json:"authorizationRequest"`

	// The duration of how long this decision should be cached. Note that the type is of type java.time.Duration, not
	// string.
	DecisionCacheDuration *string `mandatory:"true" json:"decisionCacheDuration"`
}

func (m ThinAuthorizationResponse) String() string {
	return common.PointerString(m)
}
