// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cluster Placement Groups API
//
// API for managing cluster placement groups.
//

package clusterplacementgroups

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// WorkRequestResource The resource created or operated on by a work request.
type WorkRequestResource struct {

	// The resource type that the work request affects.
	EntityType *string `mandatory:"true" json:"entityType"`

	// The way in which the resource is affected by the work tracked in the work request.
	// A resource being created, updated, or deleted remains in the IN_PROGRESS state until
	// work is complete for that resource, at which point the state transitions to CREATED, UPDATED,
	// or DELETED, respectively.
	ActionType ActionTypeEnum `mandatory:"true" json:"actionType"`

	// The identifier of the resource that the work request affects.
	Identifier *string `mandatory:"true" json:"identifier"`

	// The URI path to which a client can send a GET to access the resource metadata.
	EntityUri *string `mandatory:"false" json:"entityUri"`

	// Additional information that helps to explain the resource.
	Metadata map[string]string `mandatory:"false" json:"metadata"`
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
