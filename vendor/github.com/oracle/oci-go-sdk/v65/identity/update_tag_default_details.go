// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, policies, and identity domains.
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateTagDefaultDetails The representation of UpdateTagDefaultDetails
type UpdateTagDefaultDetails struct {

	// The default value for the tag definition. This will be applied to all resources created in the Compartment.
	Value *string `mandatory:"true" json:"value"`

	// If you specify that a value is required, a value is set during resource creation (either by
	// the user creating the resource or another tag defualt). If no value is set, resource
	// creation is blocked.
	// * If the `isRequired` flag is set to "true", the value is set during resource creation.
	// * If the `isRequired` flag is set to "false", the value you enter is set during resource creation.
	// Example: `false`
	IsRequired *bool `mandatory:"false" json:"isRequired"`
}

func (m UpdateTagDefaultDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateTagDefaultDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
