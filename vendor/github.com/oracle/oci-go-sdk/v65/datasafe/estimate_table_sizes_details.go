// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EstimateTableSizesDetails Details required to estimate table sizes for a target database using a subsetting policy.
type EstimateTableSizesDetails struct {
	TargetCredentials *Credentials `mandatory:"true" json:"targetCredentials"`

	// The OCID of the target database to use for estimating table sizes. If it's not provided, the value of the
	// targetId attribute in the SubsettingPolicy resource is used.
	TargetId *string `mandatory:"false" json:"targetId"`
}

func (m EstimateTableSizesDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EstimateTableSizesDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
