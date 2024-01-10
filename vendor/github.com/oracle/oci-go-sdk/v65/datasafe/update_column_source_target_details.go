// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateColumnSourceTargetDetails Details of the target database to be associated as the column source with a masking policy.
type UpdateColumnSourceTargetDetails struct {

	// The OCID of the target database to be associated as the column source with the masking policy.
	TargetId *string `mandatory:"true" json:"targetId"`
}

func (m UpdateColumnSourceTargetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateColumnSourceTargetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateColumnSourceTargetDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateColumnSourceTargetDetails UpdateColumnSourceTargetDetails
	s := struct {
		DiscriminatorParam string `json:"columnSource"`
		MarshalTypeUpdateColumnSourceTargetDetails
	}{
		"TARGET",
		(MarshalTypeUpdateColumnSourceTargetDetails)(m),
	}

	return json.Marshal(&s)
}
