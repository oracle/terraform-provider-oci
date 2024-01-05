// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ModuleStreamProfile A module stream profile provided by a software source
type ModuleStreamProfile struct {

	// The name of the module that contains the stream profile
	ModuleName *string `mandatory:"true" json:"moduleName"`

	// The name of the stream that contains the profile
	StreamName *string `mandatory:"true" json:"streamName"`

	// The name of the profile
	ProfileName *string `mandatory:"true" json:"profileName"`

	// A list of packages that constitute the profile.  Each element
	// in the list is the name of a package.  The name is suitable to
	// use as an argument to other OS Management APIs that interact
	// directly with packages.
	Packages []string `mandatory:"true" json:"packages"`

	// Indicates if this profile is the default for its module stream.
	IsDefault *bool `mandatory:"false" json:"isDefault"`

	// A description of the contents of the module stream profile
	Description *string `mandatory:"false" json:"description"`
}

func (m ModuleStreamProfile) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ModuleStreamProfile) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
