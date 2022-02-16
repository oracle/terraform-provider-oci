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

// AuthorizationRequest The representation of AuthorizationRequest
type AuthorizationRequest struct {

	// The id of this request. It is a GUID.
	RequestId *string `mandatory:"true" json:"requestId"`

	// The user principal object
	UserPrincipal *Principal `mandatory:"true" json:"userPrincipal"`

	// The service principal object for service to service calls.
	SvcPrincipal *Principal `mandatory:"true" json:"svcPrincipal"`

	// The name of the service that is making this authorization request
	ServiceName *string `mandatory:"true" json:"serviceName"`

	// A set of permission contexts
	Context []PermissionContext `mandatory:"true" json:"context"`

	// The hash of cached policy on the caller service side. If this is different than what Identity has, it will
	// send the most recent policy statements.
	PolicyHash *string `mandatory:"true" json:"policyHash"`
}

func (m AuthorizationRequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AuthorizationRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
