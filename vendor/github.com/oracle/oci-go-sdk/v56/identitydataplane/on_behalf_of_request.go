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

// OnBehalfOfRequest The representation of OnBehalfOfRequest
type OnBehalfOfRequest struct {

	// The signed headers of the customer call.
	RequestHeaders map[string][]string `mandatory:"true" json:"requestHeaders"`

	// The name of the target service.
	TargetServiceName *string `mandatory:"true" json:"targetServiceName"`

	// If you have an obo token already, exchange that for a new obo token.
	OboToken *string `mandatory:"false" json:"oboToken"`

	// A duration for which the obo token is requested to be valid.
	Expiration *string `mandatory:"false" json:"expiration"`
}

func (m OnBehalfOfRequest) String() string {
	return common.PointerString(m)
}
