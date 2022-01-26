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

// AuthServiceUser The representation of AuthServiceUser
type AuthServiceUser struct {

	// The id of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The id of the tenant.
	TenantId *string `mandatory:"true" json:"tenantId"`

	// The user's Oracle ID (OCID).
	Id *string `mandatory:"true" json:"id"`

	// The name of the user.
	Name *string `mandatory:"true" json:"name"`

	// The display name of the user.
	DisplayName *string `mandatory:"true" json:"displayName"`
}

func (m AuthServiceUser) String() string {
	return common.PointerString(m)
}
