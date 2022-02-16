// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Content and Experience API
//
// Oracle Content and Experience is a cloud-based content hub to drive omni-channel content management and accelerate experience delivery
//

package oce

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// IdentityStripeDetails Details of the identity stripe used for OceInstance
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IdentityStripeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
