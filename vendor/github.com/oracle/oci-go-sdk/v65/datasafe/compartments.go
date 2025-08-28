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

// Compartments Containing the OCID of the compartment and a boolean value to indicates compartmentIdInSubtree.
type Compartments struct {

	// The OCID of the compartment for including target databases to the target database group. All target databases in the compartment will be members of the target database group.
	Id *string `mandatory:"true" json:"id"`

	// This indicates whether the target databases of sub-compartments should also be included in the target database group. By default, this parameter is set to false.
	IsIncludeSubtree *bool `mandatory:"false" json:"isIncludeSubtree"`
}

func (m Compartments) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Compartments) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
