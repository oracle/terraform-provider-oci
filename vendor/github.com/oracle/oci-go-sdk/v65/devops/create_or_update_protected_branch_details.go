// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateOrUpdateProtectedBranchDetails Information to create a protected branch
type CreateOrUpdateProtectedBranchDetails struct {

	// Name of a branch to protect.
	BranchName *string `mandatory:"true" json:"branchName"`

	// Level of protection to add on a branch.
	ProtectionLevels []ProtectionLevelEnum `mandatory:"false" json:"protectionLevels,omitempty"`
}

func (m CreateOrUpdateProtectedBranchDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOrUpdateProtectedBranchDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.ProtectionLevels {
		if _, ok := GetMappingProtectionLevelEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProtectionLevels: %s. Supported values are: %s.", val, strings.Join(GetProtectionLevelEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
