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

// UpdateTargetDefaultsRemapTablespaceDetails Migration tablespace settings valid for TARGET_DEFAULTS_REMAP target type. The service will compute the targetType
// that corresponds to the targetDatabaseConnectionId type, and set the corresponding default values. When target type is ADB_S,
// ADB_D or NON_ADB the default will be set to remap feature ADB_S_REMAP, ADB_D_REMAP or NON_ADB_REMAP.
type UpdateTargetDefaultsRemapTablespaceDetails struct {
}

func (m UpdateTargetDefaultsRemapTablespaceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateTargetDefaultsRemapTablespaceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateTargetDefaultsRemapTablespaceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateTargetDefaultsRemapTablespaceDetails UpdateTargetDefaultsRemapTablespaceDetails
	s := struct {
		DiscriminatorParam string `json:"targetType"`
		MarshalTypeUpdateTargetDefaultsRemapTablespaceDetails
	}{
		"TARGET_DEFAULTS_REMAP",
		(MarshalTypeUpdateTargetDefaultsRemapTablespaceDetails)(m),
	}

	return json.Marshal(&s)
}
