// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// MatchingCriteria Criteria to either include or exclude target databases from the target database group. These criteria can be based on compartments or tags or a list of target databases. See examples below for more details.
// Include: Target databases will be added to the target database group if they match at least one of the include criteria.
// Exclude: Target databases that will be excluded from the target database group (even if they match any of the include criteria).
type MatchingCriteria struct {
	Include *Include `mandatory:"true" json:"include"`

	Exclude *Exclude `mandatory:"false" json:"exclude"`
}

func (m MatchingCriteria) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MatchingCriteria) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
