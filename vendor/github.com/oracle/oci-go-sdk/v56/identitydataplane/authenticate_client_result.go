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

// AuthenticateClientResult The representation of AuthenticateClientResult
type AuthenticateClientResult struct {

	// The original caller's resolved principal object if the authentication succeeds, null otherwise.
	Principal *Principal `mandatory:"false" json:"principal"`

	// If the authentication fails for the original caller (not failing authentication of the calling service, in which case we return 401), we return a 200, but with null principal and an error message
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`
}

func (m AuthenticateClientResult) String() string {
	return common.PointerString(m)
}
