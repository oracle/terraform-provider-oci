// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Functions Service API
//
// API for the Functions service.
//

package functions

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateFunctionUpdateRuntimeConfig FunctionsRuntime configuration for creating a function with 'FUNCTION_UPDATE' Runtime Update strategy.
type CreateFunctionUpdateRuntimeConfig struct {

	// The name of the FunctionsRuntime this function is to be associated with.
	FunctionsRuntimeName *string `mandatory:"true" json:"functionsRuntimeName"`
}

func (m CreateFunctionUpdateRuntimeConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateFunctionUpdateRuntimeConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateFunctionUpdateRuntimeConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateFunctionUpdateRuntimeConfig CreateFunctionUpdateRuntimeConfig
	s := struct {
		DiscriminatorParam string `json:"runtimeConfigType"`
		MarshalTypeCreateFunctionUpdateRuntimeConfig
	}{
		"FUNCTION_UPDATE",
		(MarshalTypeCreateFunctionUpdateRuntimeConfig)(m),
	}

	return json.Marshal(&s)
}
