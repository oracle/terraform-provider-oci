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

// ManageModuleStreamsOnManagedInstanceDetails The set of changes to make to the state of the modules, streams, and profiles on a managed instance
type ManageModuleStreamsOnManagedInstanceDetails struct {

	// Indicates if this operation is a dry run or if the operation
	// should be commited.  If set to true, the result of the operation
	// will be evaluated but not committed.  If set to false, the
	// operation is committed to the managed instance.  The default is
	// false.
	IsDryRun *bool `mandatory:"false" json:"isDryRun"`

	// The set of module streams to enable.
	Enable []ModuleStreamDetails `mandatory:"false" json:"enable"`

	// The set of module streams to disable.
	Disable []ModuleStreamDetails `mandatory:"false" json:"disable"`

	// The set of module stream profiles to install.
	Install []ModuleStreamProfileDetails `mandatory:"false" json:"install"`

	// The set of module stream profiles to remove.
	Remove []ModuleStreamProfileDetails `mandatory:"false" json:"remove"`
}

func (m ManageModuleStreamsOnManagedInstanceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManageModuleStreamsOnManagedInstanceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
