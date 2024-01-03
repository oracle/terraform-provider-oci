// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateNonAdbRemapTablespaceDetails Migration tablespace settings valid for NON-ADB target type using remap feature.
type CreateNonAdbRemapTablespaceDetails struct {

	// Name of tablespace at target to which the source database tablespace need to be remapped.
	RemapTarget *string `mandatory:"false" json:"remapTarget"`
}

func (m CreateNonAdbRemapTablespaceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateNonAdbRemapTablespaceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateNonAdbRemapTablespaceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateNonAdbRemapTablespaceDetails CreateNonAdbRemapTablespaceDetails
	s := struct {
		DiscriminatorParam string `json:"targetType"`
		MarshalTypeCreateNonAdbRemapTablespaceDetails
	}{
		"NON_ADB_REMAP",
		(MarshalTypeCreateNonAdbRemapTablespaceDetails)(m),
	}

	return json.Marshal(&s)
}
