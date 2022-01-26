// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// LaunchEligibility Tenant eligibility and other information for launching a PIC image
type LaunchEligibility struct {

	// PIC Image ID
	ImageId *string `mandatory:"true" json:"imageId"`

	// Is the tenant permitted to launch the PIC image
	IsLaunchAllowed *bool `mandatory:"true" json:"isLaunchAllowed"`

	// related meters for the PIC image
	Meters *string `mandatory:"false" json:"meters"`
}

func (m LaunchEligibility) String() string {
	return common.PointerString(m)
}
