// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// WorkRequestResource The resource on which the work request is operating.
type WorkRequestResource struct {

	// How the work request affects the resource.
	ActionType WorkRequestResourceActionTypeEnum `mandatory:"false" json:"actionType,omitempty"`

	// The resource type the work request affects.
	EntityType *string `mandatory:"false" json:"entityType"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the resource that the work request affects.
	Identifier *string `mandatory:"false" json:"identifier"`

	// The URI path used while performing a `GET` to access the resource metadata.
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

	if _, ok := GetMappingWorkRequestResourceActionTypeEnum(string(m.ActionType)); !ok && m.ActionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionType: %s. Supported values are: %s.", m.ActionType, strings.Join(GetWorkRequestResourceActionTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// WorkRequestResourceActionTypeEnum Enum with underlying type: string
type WorkRequestResourceActionTypeEnum string

// Set of constants representing the allowable values for WorkRequestResourceActionTypeEnum
const (
	WorkRequestResourceActionTypeInProgress WorkRequestResourceActionTypeEnum = "IN_PROGRESS"
	WorkRequestResourceActionTypeCreated    WorkRequestResourceActionTypeEnum = "CREATED"
	WorkRequestResourceActionTypeUpdated    WorkRequestResourceActionTypeEnum = "UPDATED"
	WorkRequestResourceActionTypeDeleted    WorkRequestResourceActionTypeEnum = "DELETED"
	WorkRequestResourceActionTypeRelated    WorkRequestResourceActionTypeEnum = "RELATED"
	WorkRequestResourceActionTypePurged     WorkRequestResourceActionTypeEnum = "PURGED"
)

var mappingWorkRequestResourceActionTypeEnum = map[string]WorkRequestResourceActionTypeEnum{
	"IN_PROGRESS": WorkRequestResourceActionTypeInProgress,
	"CREATED":     WorkRequestResourceActionTypeCreated,
	"UPDATED":     WorkRequestResourceActionTypeUpdated,
	"DELETED":     WorkRequestResourceActionTypeDeleted,
	"RELATED":     WorkRequestResourceActionTypeRelated,
	"PURGED":      WorkRequestResourceActionTypePurged,
}

var mappingWorkRequestResourceActionTypeEnumLowerCase = map[string]WorkRequestResourceActionTypeEnum{
	"in_progress": WorkRequestResourceActionTypeInProgress,
	"created":     WorkRequestResourceActionTypeCreated,
	"updated":     WorkRequestResourceActionTypeUpdated,
	"deleted":     WorkRequestResourceActionTypeDeleted,
	"related":     WorkRequestResourceActionTypeRelated,
	"purged":      WorkRequestResourceActionTypePurged,
}

// GetWorkRequestResourceActionTypeEnumValues Enumerates the set of values for WorkRequestResourceActionTypeEnum
func GetWorkRequestResourceActionTypeEnumValues() []WorkRequestResourceActionTypeEnum {
	values := make([]WorkRequestResourceActionTypeEnum, 0)
	for _, v := range mappingWorkRequestResourceActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestResourceActionTypeEnumStringValues Enumerates the set of values in String for WorkRequestResourceActionTypeEnum
func GetWorkRequestResourceActionTypeEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"CREATED",
		"UPDATED",
		"DELETED",
		"RELATED",
		"PURGED",
	}
}

// GetMappingWorkRequestResourceActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestResourceActionTypeEnum(val string) (WorkRequestResourceActionTypeEnum, bool) {
	enum, ok := mappingWorkRequestResourceActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
