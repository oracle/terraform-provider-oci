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

// InstallationUsageCollection Results of an installation search. Contains InstallationUsage items.
type InstallationUsageCollection struct {

	// A list of installations.
	Items []InstallationUsage `mandatory:"true" json:"items"`
}

func (m InstallationUsageCollection) String() string {
	return common.PointerString(m)
}
