// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for the operating system environments in your private data centers through a single management console. For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ScheduledJobOperation Defines an operation in a scheduled job.
type ScheduledJobOperation struct {

	// The type of operation this scheduled job performs.
	OperationType OperationTypesEnum `mandatory:"true" json:"operationType"`

	// The names of the target packages (only if operation type is INSTALL_PACKAGES/UPDATE_PACKAGES/REMOVE_PACKAGES).
	PackageNames []string `mandatory:"false" json:"packageNames"`

	ManageModuleStreamsDetails *ManageModuleStreamsInScheduledJobDetails `mandatory:"false" json:"manageModuleStreamsDetails"`

	SwitchModuleStreamsDetails *ModuleStreamDetails `mandatory:"false" json:"switchModuleStreamsDetails"`

	// The OCIDs for the software sources (only if operation type is ATTACH_SOFTWARE_SOURCES/DETACH_SOFTWARE_SOURCES).
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
