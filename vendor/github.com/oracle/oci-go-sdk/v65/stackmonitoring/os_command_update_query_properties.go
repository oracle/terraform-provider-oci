// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OsCommandUpdateQueryProperties Query Properties applicable to OS_COMMAND type of collection method
type OsCommandUpdateQueryProperties struct {

	// OS command to execute without arguments
	Command *string `mandatory:"false" json:"command"`

	// Character used to delimit multiple metric values in single line of output
	Delimiter *string `mandatory:"false" json:"delimiter"`

	ScriptDetails *ScriptFileDetails `mandatory:"false" json:"scriptDetails"`

	// Arguments required by either command or script
	Arguments *string `mandatory:"false" json:"arguments"`

	// String prefix used to identify metric output of the OS Command
	StartsWith *string `mandatory:"false" json:"startsWith"`
}

func (m OsCommandUpdateQueryProperties) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OsCommandUpdateQueryProperties) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OsCommandUpdateQueryProperties) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOsCommandUpdateQueryProperties OsCommandUpdateQueryProperties
	s := struct {
		DiscriminatorParam string `json:"collectionMethod"`
		MarshalTypeOsCommandUpdateQueryProperties
	}{
		"OS_COMMAND",
		(MarshalTypeOsCommandUpdateQueryProperties)(m),
	}

	return json.Marshal(&s)
}
