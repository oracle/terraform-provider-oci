// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// FleetCollection Results of a Fleet search. Contains FleetSummary items.
type FleetCollection struct {

	// A list of Fleets.
	Items []FleetSummary `mandatory:"true" json:"items"`
}

func (m FleetCollection) String() string {
	return common.PointerString(m)
}
