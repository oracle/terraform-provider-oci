// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Scheduler API
//
// Use the Resource scheduler API to manage schedules, to perform actions on a collection of resources.
//

package resourcescheduler

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// WorkRequestResource This is a resource created or operated on by a work request.
type WorkRequestResource struct {

	// This is the resource type that the work request affects.
	EntityType *string `mandatory:"true" json:"entityType"`

	// The way in which this resource is affected by the operation tracked in the work request.
	// A resource being created, updated, or deleted remains in the IN_PROGRESS state until
	// work is complete for that resource, at which point it transitions to CREATED, UPDATED,
	// or DELETED, respectively.
	ActionType ActionTypeEnum `mandatory:"true" json:"actionType"`

	// This is an OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) or other unique identifier for the resource.
	Identifier *string `mandatory:"true" json:"identifier"`

	// This is the URI path that you can use for a GET request to access the resource metadata.
	EntityUri *string `mandatory:"false" json:"entityUri"`
}

func (m WorkRequestResource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WorkRequestResource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingActionTypeEnum(string(m.ActionType)); !ok && m.ActionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionType: %s. Supported values are: %s.", m.ActionType, strings.Join(GetActionTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
