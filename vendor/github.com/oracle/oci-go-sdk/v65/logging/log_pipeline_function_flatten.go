// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, move and delete
// log groups, log objects, log saved searches, and agent configurations.
// For more information, see Logging Overview (https://docs.oracle.com/iaas/Content/Logging/Concepts/loggingoverview.htm).
//

package logging

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LogPipelineFunctionFlatten Describes flatten incoming log event.
// The Flatten Function pulls fields out of a nested structure, moving nested key-value pairs (fields) up to a highest level in the log entry data field.
type LogPipelineFunctionFlatten struct {

	// Processor will flatten all fields which matches source_field patterns.
	SourceField *string `mandatory:"true" json:"sourceField"`

	// Delimiter String which processor will use flatten between field names. Default is "_"
	Delimiter *string `mandatory:"false" json:"delimiter"`
}

func (m LogPipelineFunctionFlatten) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogPipelineFunctionFlatten) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m LogPipelineFunctionFlatten) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeLogPipelineFunctionFlatten LogPipelineFunctionFlatten
	s := struct {
		DiscriminatorParam string `json:"functionType"`
		MarshalTypeLogPipelineFunctionFlatten
	}{
		"FLATTEN",
		(MarshalTypeLogPipelineFunctionFlatten)(m),
	}

	return json.Marshal(&s)
}
