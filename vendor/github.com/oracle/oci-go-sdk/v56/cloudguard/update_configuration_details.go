// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// UpdateConfigurationDetails Update cloud guard configuration details for a tenancy.
type UpdateConfigurationDetails struct {

	// The reporting region value
	ReportingRegion *string `mandatory:"true" json:"reportingRegion"`

	// Status of Cloud Guard Tenant
	Status CloudGuardStatusEnum `mandatory:"true" json:"status"`

	// Identifies if Oracle managed resources will be created by customers.
	// If no value is specified false is the default.
	SelfManageResources *bool `mandatory:"false" json:"selfManageResources"`
}

func (m UpdateConfigurationDetails) String() string {
	return common.PointerString(m)
}
