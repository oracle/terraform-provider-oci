// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// NoneProvisionedConcurrencyConfig Configuration specifying no provisioned concurrency
type NoneProvisionedConcurrencyConfig struct {
}

func (m NoneProvisionedConcurrencyConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NoneProvisionedConcurrencyConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m NoneProvisionedConcurrencyConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeNoneProvisionedConcurrencyConfig NoneProvisionedConcurrencyConfig
	s := struct {
		DiscriminatorParam string `json:"strategy"`
		MarshalTypeNoneProvisionedConcurrencyConfig
	}{
		"NONE",
		(MarshalTypeNoneProvisionedConcurrencyConfig)(m),
	}

	return json.Marshal(&s)
}
