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

// EntityStatus The representation of EntityStatus
type EntityStatus struct {

	// The entity status.
	Status *string `mandatory:"true" json:"status"`

	// A bit mask showing the reason why the entity is inactive:
	// - bit 0: ACTIVE
	// - bit 1: SUSPENDED
	// - bit 2: DISABLED
	// - bit 3: BLOCKED
	// - bit 4: LOCKED
	InactiveBitMask *int64 `mandatory:"true" json:"inactiveBitMask"`
}

func (m EntityStatus) String() string {
	return common.PointerString(m)
}
