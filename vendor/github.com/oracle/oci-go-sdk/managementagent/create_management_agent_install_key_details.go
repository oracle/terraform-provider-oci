// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// API for Management Agent Cloud Service
//

package managementagent

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateManagementAgentInstallKeyDetails The information about new Management Agent install Key.
type CreateManagementAgentInstallKeyDetails struct {

	// Management Agent install Key Name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Total number of install for this keys
	AllowedKeyInstallCount *int `mandatory:"false" json:"allowedKeyInstallCount"`

	// date after which key would expire after creation
	TimeExpires *common.SDKTime `mandatory:"false" json:"timeExpires"`
}

func (m CreateManagementAgentInstallKeyDetails) String() string {
	return common.PointerString(m)
}
