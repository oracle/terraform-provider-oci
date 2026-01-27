// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DataIntelligences Control Plane API
//
// Use the DataIntelligences Control Plane API to manage dataIntelligences.
//

package dif

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ComponentValueOverride Component overrides for stack‑specific parameters applied during artifact template rendering.
type ComponentValueOverride struct {

	// Logical name of the grouping independently deployable kubernetes resource artifacts for the current deployment.
	ComponentName *string `mandatory:"true" json:"componentName"`

	// Free-form value overrides for the component. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// Used for overriding the values in value.yaml artifact of the component.
	// Example: `{"WORKER_THREADS": "8"}`
	ValueOverrides map[string]string `mandatory:"true" json:"valueOverrides"`
}

func (m ComponentValueOverride) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComponentValueOverride) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
