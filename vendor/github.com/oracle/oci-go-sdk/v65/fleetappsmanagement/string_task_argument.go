// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// StringTaskArgument A string variable that holds a value
type StringTaskArgument struct {

	// Name of the input variable
	Name *string `mandatory:"true" json:"name"`

	// The task input
	Value *string `mandatory:"false" json:"value"`
}

// GetName returns Name
func (m StringTaskArgument) GetName() *string {
	return m.Name
}

func (m StringTaskArgument) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StringTaskArgument) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m StringTaskArgument) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeStringTaskArgument StringTaskArgument
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeStringTaskArgument
	}{
		"STRING",
		(MarshalTypeStringTaskArgument)(m),
	}

	return json.Marshal(&s)
}
