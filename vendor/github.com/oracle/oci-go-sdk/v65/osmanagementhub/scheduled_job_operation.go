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

// ScheduledJobOperation Defines an operation that is performed by a scheduled job.
type ScheduledJobOperation struct {

	// The type of operation this scheduled job performs.
	OperationType OperationTypesEnum `mandatory:"true" json:"operationType"`

	// The names of the target packages. This parameter only applies when the scheduled job is for installing, updating, or removing packages.
	PackageNames []string `mandatory:"false" json:"packageNames"`

	// Unique identifier for the Windows update. This parameter only applies if the scheduled job is for installing Windows updates.
	// Note that this is not an OCID, but is a unique identifier assigned by Microsoft.
	// For example: '6981d463-cd91-4a26-b7c4-ea4ded9183ed'.
	WindowsUpdateNames []string `mandatory:"false" json:"windowsUpdateNames"`

	ManageModuleStreamsDetails *ManageModuleStreamsInScheduledJobDetails `mandatory:"false" json:"manageModuleStreamsDetails"`

	SwitchModuleStreamsDetails *ModuleStreamDetails `mandatory:"false" json:"switchModuleStreamsDetails"`

	// The software source OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	// This parameter only applies when the scheduled job is for attaching or detaching software sources.
	SoftwareSourceIds []string `mandatory:"false" json:"softwareSourceIds"`
}

func (m ScheduledJobOperation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScheduledJobOperation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOperationTypesEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetOperationTypesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
