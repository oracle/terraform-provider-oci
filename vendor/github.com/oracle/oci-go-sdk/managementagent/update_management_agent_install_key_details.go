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

// UpdateManagementAgentInstallKeyDetails Details required to change Management Agent install key.
type UpdateManagementAgentInstallKeyDetails struct {

	// if set to true the install key state would be set to Active and if false to Inactive
	IsKeyActive *bool `mandatory:"false" json:"isKeyActive"`

	// New displayName of Agent install key.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m UpdateManagementAgentInstallKeyDetails) String() string {
	return common.PointerString(m)
}
