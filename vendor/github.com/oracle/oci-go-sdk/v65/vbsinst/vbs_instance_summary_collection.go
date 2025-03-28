// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// VbsControlplaneInstance API
//
// A description of the VbsControlplaneInstance API
//

package vbsinst

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VbsInstanceSummaryCollection Wrapped list of Visual Builder Studio service instances
type VbsInstanceSummaryCollection struct {

	// The Visual Builder Studio service instances
	Items []VbsInstanceSummary `mandatory:"true" json:"items"`
}

func (m VbsInstanceSummaryCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VbsInstanceSummaryCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
