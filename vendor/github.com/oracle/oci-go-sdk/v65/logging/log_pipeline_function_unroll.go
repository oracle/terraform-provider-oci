// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// LogPipelineFunctionUnroll Describes unroll incoming log event.
// It unrolls (explodes) an array of objects within incoming event into individual events, while inheriting the top-level fields from the original event.
type LogPipelineFunctionUnroll struct {

	// Field of log event where we expect array to be exploded. It should be top level array, not subarray. If you want to unroll subarra, then use twoUnroll funcitons.
	SourceField *string `mandatory:"true" json:"sourceField"`

	// Field at unrolled event where data from array will be placed. Rest fields will be inherited from input event.
	DestinationField *string `mandatory:"true" json:"destinationField"`
}

func (m LogPipelineFunctionUnroll) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogPipelineFunctionUnroll) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m LogPipelineFunctionUnroll) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeLogPipelineFunctionUnroll LogPipelineFunctionUnroll
	s := struct {
		DiscriminatorParam string `json:"functionType"`
		MarshalTypeLogPipelineFunctionUnroll
	}{
		"UNROLL",
		(MarshalTypeLogPipelineFunctionUnroll)(m),
	}

	return json.Marshal(&s)
}
