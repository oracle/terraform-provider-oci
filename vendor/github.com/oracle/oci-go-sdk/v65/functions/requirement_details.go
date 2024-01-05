// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Functions Service API
//
// API for the Functions service.
//

package functions

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RequirementDetails Minimum memory required by this PBF. The user should use memory greater than or equal to this value
// while configuring the Function.
type RequirementDetails struct {

	// Minimum memory required by this PBF. The user should use memory greater than or equal to
	// this value while configuring the Function.
	MinMemoryRequiredInMBs *int64 `mandatory:"true" json:"minMemoryRequiredInMBs"`

	// List of policies required for this PBF execution.
	Policies []PolicyDetails `mandatory:"false" json:"policies"`
}

func (m RequirementDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RequirementDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
