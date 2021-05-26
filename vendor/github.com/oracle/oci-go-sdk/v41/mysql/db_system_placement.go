// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"github.com/oracle/oci-go-sdk/v41/common"
)

// DbSystemPlacement The availability domain and fault domain a DB System is placed in.
type DbSystemPlacement struct {

	// The availability domain in which the DB System is placed.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The fault domain in which the DB System is placed.
	FaultDomain *string `mandatory:"false" json:"faultDomain"`
}

func (m DbSystemPlacement) String() string {
	return common.PointerString(m)
}
