// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// Use the Management Agent API to manage your infrastructure's management agents, including their plugins and install keys.
// For more information, see Management Agent (https://docs.cloud.oracle.com/iaas/management-agents/index.html).
//

package managementagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

	// If set to true, the install key has no expiration date or usage limit. Defaults to false
	IsUnlimited *bool `mandatory:"false" json:"isUnlimited"`
}

func (m CreateManagementAgentInstallKeyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateManagementAgentInstallKeyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
