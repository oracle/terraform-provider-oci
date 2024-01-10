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

// UpdateTargetDefaultsAutoCreateTablespaceDetails Migration tablespace settings valid for TARGET_DEFAULTS_AUTOCREATE target type. The service will compute
// the targetType that corresponds to the targetDatabaseConnectionId type, and set the corresponding default values. When
// target type is ADB_D or NON_ADB the default will be set to auto-create feature ADB_D_AUTOCREATE or NON_ADB_AUTOCREATE.
type UpdateTargetDefaultsAutoCreateTablespaceDetails struct {
}

func (m UpdateTargetDefaultsAutoCreateTablespaceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateTargetDefaultsAutoCreateTablespaceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateTargetDefaultsAutoCreateTablespaceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateTargetDefaultsAutoCreateTablespaceDetails UpdateTargetDefaultsAutoCreateTablespaceDetails
	s := struct {
		DiscriminatorParam string `json:"targetType"`
		MarshalTypeUpdateTargetDefaultsAutoCreateTablespaceDetails
	}{
		"TARGET_DEFAULTS_AUTOCREATE",
		(MarshalTypeUpdateTargetDefaultsAutoCreateTablespaceDetails)(m),
	}

	return json.Marshal(&s)
}
