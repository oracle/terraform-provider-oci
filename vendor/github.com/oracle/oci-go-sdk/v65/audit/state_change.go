// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Audit API
//
// API for the Audit Service. Use this API for compliance monitoring in your tenancy.
// For more information, see Overview of Audit (https://docs.oracle.com/iaas/Content/Audit/Concepts/auditoverview.htm).
// **Tip**: This API is good for queries, but not bulk-export operations.
//

package audit

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// StateChange A container object for state change attributes.
type StateChange struct {

	// Provides the previous state of fields that may have changed during an operation. To determine
	// how the current operation changed a resource, compare the information in this attribute to
	// `current`.
	Previous map[string]interface{} `mandatory:"false" json:"previous"`

	// Provides the current state of fields that may have changed during an operation. To determine
	// how the current operation changed a resource, compare the information in this attribute to
	// `previous`.
	Current map[string]interface{} `mandatory:"false" json:"current"`
}

func (m StateChange) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StateChange) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
