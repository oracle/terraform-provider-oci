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

// Claim The representation of Claim
type Claim struct {

	// The key of the claim.
	Key *string `mandatory:"true" json:"key"`

	// The value of the claim.
	Value *string `mandatory:"true" json:"value"`

	// The issuer of the claim.
	Issuer *string `mandatory:"false" json:"issuer"`
}

func (m Claim) String() string {
	return common.PointerString(m)
}
