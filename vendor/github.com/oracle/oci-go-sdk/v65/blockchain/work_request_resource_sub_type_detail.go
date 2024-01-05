// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Blockchain Platform Control Plane API
//
// Blockchain Platform Control Plane API
//

package blockchain

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// WorkRequestResourceSubTypeDetail SubType information for a work request resource.
type WorkRequestResourceSubTypeDetail struct {

	// Subtype of the work request resource like osn or peer.
	SubType *string `mandatory:"true" json:"subType"`

	// The identifier of the resource subType.
	SubTypeKey *string `mandatory:"true" json:"subTypeKey"`

	// Status of the resource subType, as a result of the work tracked in this work request.
	// A resource subType would be CREATED, UPDATED or DELETED, after the work request is completed.
	SubTypeStatus WorkRequestResourceSubTypeDetailSubTypeStatusEnum `mandatory:"true" json:"subTypeStatus"`
}

func (m WorkRequestResourceSubTypeDetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WorkRequestResourceSubTypeDetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingWorkRequestResourceSubTypeDetailSubTypeStatusEnum(string(m.SubTypeStatus)); !ok && m.SubTypeStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SubTypeStatus: %s. Supported values are: %s.", m.SubTypeStatus, strings.Join(GetWorkRequestResourceSubTypeDetailSubTypeStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// WorkRequestResourceSubTypeDetailSubTypeStatusEnum Enum with underlying type: string
type WorkRequestResourceSubTypeDetailSubTypeStatusEnum string

// Set of constants representing the allowable values for WorkRequestResourceSubTypeDetailSubTypeStatusEnum
const (
	WorkRequestResourceSubTypeDetailSubTypeStatusCreated WorkRequestResourceSubTypeDetailSubTypeStatusEnum = "CREATED"
	WorkRequestResourceSubTypeDetailSubTypeStatusUpdated WorkRequestResourceSubTypeDetailSubTypeStatusEnum = "UPDATED"
	WorkRequestResourceSubTypeDetailSubTypeStatusDeleted WorkRequestResourceSubTypeDetailSubTypeStatusEnum = "DELETED"
)

var mappingWorkRequestResourceSubTypeDetailSubTypeStatusEnum = map[string]WorkRequestResourceSubTypeDetailSubTypeStatusEnum{
	"CREATED": WorkRequestResourceSubTypeDetailSubTypeStatusCreated,
	"UPDATED": WorkRequestResourceSubTypeDetailSubTypeStatusUpdated,
	"DELETED": WorkRequestResourceSubTypeDetailSubTypeStatusDeleted,
}

var mappingWorkRequestResourceSubTypeDetailSubTypeStatusEnumLowerCase = map[string]WorkRequestResourceSubTypeDetailSubTypeStatusEnum{
	"created": WorkRequestResourceSubTypeDetailSubTypeStatusCreated,
	"updated": WorkRequestResourceSubTypeDetailSubTypeStatusUpdated,
	"deleted": WorkRequestResourceSubTypeDetailSubTypeStatusDeleted,
}

// GetWorkRequestResourceSubTypeDetailSubTypeStatusEnumValues Enumerates the set of values for WorkRequestResourceSubTypeDetailSubTypeStatusEnum
func GetWorkRequestResourceSubTypeDetailSubTypeStatusEnumValues() []WorkRequestResourceSubTypeDetailSubTypeStatusEnum {
	values := make([]WorkRequestResourceSubTypeDetailSubTypeStatusEnum, 0)
	for _, v := range mappingWorkRequestResourceSubTypeDetailSubTypeStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestResourceSubTypeDetailSubTypeStatusEnumStringValues Enumerates the set of values in String for WorkRequestResourceSubTypeDetailSubTypeStatusEnum
func GetWorkRequestResourceSubTypeDetailSubTypeStatusEnumStringValues() []string {
	return []string{
		"CREATED",
		"UPDATED",
		"DELETED",
	}
}

// GetMappingWorkRequestResourceSubTypeDetailSubTypeStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestResourceSubTypeDetailSubTypeStatusEnum(val string) (WorkRequestResourceSubTypeDetailSubTypeStatusEnum, bool) {
	enum, ok := mappingWorkRequestResourceSubTypeDetailSubTypeStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
