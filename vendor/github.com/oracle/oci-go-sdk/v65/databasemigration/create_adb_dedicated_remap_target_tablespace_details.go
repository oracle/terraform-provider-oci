// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// CreateAdbDedicatedRemapTargetTablespaceDetails Remap tablespace settings that are valid for Dedicated Autonomous Databases.
type CreateAdbDedicatedRemapTargetTablespaceDetails struct {

	// Name of the tablespace on the target database to which the source database tablespace is to be remapped.
	RemapTarget *string `mandatory:"false" json:"remapTarget"`
}

func (m CreateAdbDedicatedRemapTargetTablespaceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateAdbDedicatedRemapTargetTablespaceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateAdbDedicatedRemapTargetTablespaceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateAdbDedicatedRemapTargetTablespaceDetails CreateAdbDedicatedRemapTargetTablespaceDetails
	s := struct {
		DiscriminatorParam string `json:"targetType"`
		MarshalTypeCreateAdbDedicatedRemapTargetTablespaceDetails
	}{
		"ADB_D_REMAP",
		(MarshalTypeCreateAdbDedicatedRemapTargetTablespaceDetails)(m),
	}

	return json.Marshal(&s)
}
