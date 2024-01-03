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

// ColumnSourceFromTargetDetails Details of the target database that's used as the source of masking columns.
type ColumnSourceFromTargetDetails struct {

	// The OCID of the target database that's used as the source of masking columns.
	TargetId *string `mandatory:"true" json:"targetId"`
}

func (m ColumnSourceFromTargetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ColumnSourceFromTargetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ColumnSourceFromTargetDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeColumnSourceFromTargetDetails ColumnSourceFromTargetDetails
	s := struct {
		DiscriminatorParam string `json:"columnSource"`
		MarshalTypeColumnSourceFromTargetDetails
	}{
		"TARGET",
		(MarshalTypeColumnSourceFromTargetDetails)(m),
	}

	return json.Marshal(&s)
}
