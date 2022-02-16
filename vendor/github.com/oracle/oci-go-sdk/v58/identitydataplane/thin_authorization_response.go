// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Service
//
// API for the Identity Dataplane
//

package identitydataplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ThinAuthorizationResponse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
