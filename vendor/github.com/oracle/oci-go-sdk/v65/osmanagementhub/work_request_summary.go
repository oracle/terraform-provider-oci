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

// WorkRequestSummary The summary of a work request.
type WorkRequestSummary struct {

	// Type of the work request.
	OperationType WorkRequestOperationTypeEnum `mandatory:"true" json:"operationType"`

	// Status of the work request.
	Status OperationStatusEnum `mandatory:"true" json:"status"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the work request.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the work request.
	// Work requests should be scoped to the same compartment as the resource it affects.
	// If the work request affects multiple resources the different compartments, the services selects the compartment of the primary resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the request was created - as described in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// A short description about the work request.
	Description *string `mandatory:"false" json:"description"`

	// A short display name for the work request.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A progress or error message, if there is any.
	Message *string `mandatory:"false" json:"message"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the parent work request, if there is any.
	ParentId *string `mandatory:"false" json:"parentId"`

	// The list of OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for child work requests.
	ChildrenId []string `mandatory:"false" json:"childrenId"`

	// The percentage complete of the operation tracked by this work request.
	PercentComplete *float32 `mandatory:"false" json:"percentComplete"`

	// The scheduled date and time to retry the work request (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	TimeScheduled *common.SDKTime `mandatory:"false" json:"timeScheduled"`

	// Indicates whether this work request is managed by Autonomous Linux
	IsManagedByAutonomousLinux *bool `mandatory:"false" json:"isManagedByAutonomousLinux"`
}

func (m WorkRequestSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WorkRequestSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingWorkRequestOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetWorkRequestOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOperationStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetOperationStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
