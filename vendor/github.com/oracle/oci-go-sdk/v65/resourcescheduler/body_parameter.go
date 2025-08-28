// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Scheduler API
//
// Use the Resource scheduler API to manage schedules, to perform actions on a collection of resources.
//

package resourcescheduler

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BodyParameter This is an input parameter that will be passed as HTTP request body parameter.
type BodyParameter struct {

	// This is the body parameter value.
	Value *interface{} `mandatory:"false" json:"value"`
}

func (m BodyParameter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BodyParameter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m BodyParameter) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBodyParameter BodyParameter
	s := struct {
		DiscriminatorParam string `json:"parameterType"`
		MarshalTypeBodyParameter
	}{
		"BODY",
		(MarshalTypeBodyParameter)(m),
	}

	return json.Marshal(&s)
}
