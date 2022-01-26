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

// ResourceInventory Inventory of JMS resources in a compartment during a specified time period.
type ResourceInventory struct {

	// The number of _active_ fleets.
	ActiveFleetCount *int `mandatory:"true" json:"activeFleetCount"`

	// The number of managed instances.
	ManagedInstanceCount *int `mandatory:"true" json:"managedInstanceCount"`

	// The number of Java Runtimes.
	JreCount *int `mandatory:"true" json:"jreCount"`

	// The number of Java installations.
	InstallationCount *int `mandatory:"true" json:"installationCount"`

	// The number of applications.
	ApplicationCount *int `mandatory:"true" json:"applicationCount"`
}

func (m ResourceInventory) String() string {
	return common.PointerString(m)
}
