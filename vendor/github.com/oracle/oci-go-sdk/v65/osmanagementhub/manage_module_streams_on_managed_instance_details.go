// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ManageModuleStreamsOnManagedInstanceDetails The set of changes to make to the state of the modules, streams, and profiles on a managed instance
type ManageModuleStreamsOnManagedInstanceDetails struct {

	// Indicates if this operation is a dry run or if the operation
	// should be committed.  If set to true, the result of the operation
	// will be evaluated but not committed.  If set to false, the
	// operation is committed to the managed instance.  The default is
	// false.
	IsDryRun *bool `mandatory:"false" json:"isDryRun"`

	// The set of module streams to enable. If any streams of a module are already enabled, the service switches from the current stream to the new stream.
	// Once complete, the streams will be in 'ENABLED' status.
	Enable []ModuleStreamDetails `mandatory:"false" json:"enable"`

	// The set of module streams to disable. Any profiles that are installed for the module stream will be removed as part of the operation.
	// Once complete, the streams will be in 'DISABLED' status.
	Disable []ModuleStreamDetails `mandatory:"false" json:"disable"`

	// The set of module stream profiles to install. Any packages that are part of the profile are installed on the managed instance.
	// Once complete, the profile will be in 'INSTALLED' status. The operation will return an error if you attempt to install a profile from a disabled stream, unless enabling the new module stream is included in this operation.
	Install []ModuleStreamProfileDetails `mandatory:"false" json:"install"`

	// The set of module stream profiles to remove. Once complete, the profile will be in 'AVAILABLE' status.
	// The status of packages within the profile after the operation is complete is defined by the package manager on the managed instance group.
	Remove []ModuleStreamProfileDetails `mandatory:"false" json:"remove"`

	WorkRequestDetails *WorkRequestDetails `mandatory:"false" json:"workRequestDetails"`
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
