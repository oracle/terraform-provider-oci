// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// OceInstance API
//
// Oracle Content and Experience is a cloud-based content hub to drive omni-channel content management and accelerate experience delivery
//

package oce

import (
	"github.com/oracle/oci-go-sdk/common"
)

// IdentityStripeDetails Identity Stripe
type IdentityStripeDetails struct {

	// Name of the Identity Cloud Service instance in My Services to be used.
	// Example: `secondstripe`
	ServiceName *string `mandatory:"true" json:"serviceName"`

	// Value of the Identity Cloud Service tenancy.
	// Example: `idcs-8416ebdd0d674f84803f4193cce026e9`
	Tenancy *string `mandatory:"true" json:"tenancy"`
}

func (m IdentityStripeDetails) String() string {
	return common.PointerString(m)
}
