// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Use Object Storage and Archive Storage APIs to manage buckets, objects, and related resources.
// For more information, see Overview of Object Storage (https://docs.cloud.oracle.com/Content/Object/Concepts/objectstorageoverview.htm) and
// Overview of Archive Storage (https://docs.cloud.oracle.com/Content/Archive/Concepts/archivestorageoverview.htm).
//

package objectstorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateAclGroupDetails The details to update an ACL Group.
type UpdateAclGroupDetails struct {

	// User-friendly name to describe the ACL Group. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A description of the ACL Group. This can be useful for identifying the purpose of the ACL Group.
	Description *string `mandatory:"false" json:"description"`

	// The compartment containing this ACL Group.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Determines whether this ACL Group can be associated with a bucket, compartment, or tenancy.
	AclGroupType UpdateAclGroupDetailsAclGroupTypeEnum `mandatory:"false" json:"aclGroupType,omitempty"`

	// Specifies requests that are not subject to any ACLs in the ACL group.
	AclGroupExceptions []UpdateAclGroupDetailsAclGroupExceptionsEnum `mandatory:"false" json:"aclGroupExceptions,omitempty"`

	// An ordered list of ACL IDs. The ACLs are evaluated first to last.
	AclIds []string `mandatory:"false" json:"aclIds"`
}

func (m UpdateAclGroupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateAclGroupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateAclGroupDetailsAclGroupTypeEnum(string(m.AclGroupType)); !ok && m.AclGroupType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AclGroupType: %s. Supported values are: %s.", m.AclGroupType, strings.Join(GetUpdateAclGroupDetailsAclGroupTypeEnumStringValues(), ",")))
	}
	for _, val := range m.AclGroupExceptions {
		if _, ok := GetMappingUpdateAclGroupDetailsAclGroupExceptionsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AclGroupExceptions: %s. Supported values are: %s.", val, strings.Join(GetUpdateAclGroupDetailsAclGroupExceptionsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateAclGroupDetailsAclGroupTypeEnum Enum with underlying type: string
type UpdateAclGroupDetailsAclGroupTypeEnum string

// Set of constants representing the allowable values for UpdateAclGroupDetailsAclGroupTypeEnum
const (
	UpdateAclGroupDetailsAclGroupTypeTenancy     UpdateAclGroupDetailsAclGroupTypeEnum = "TENANCY"
	UpdateAclGroupDetailsAclGroupTypeCompartment UpdateAclGroupDetailsAclGroupTypeEnum = "COMPARTMENT"
	UpdateAclGroupDetailsAclGroupTypeBucket      UpdateAclGroupDetailsAclGroupTypeEnum = "BUCKET"
)

var mappingUpdateAclGroupDetailsAclGroupTypeEnum = map[string]UpdateAclGroupDetailsAclGroupTypeEnum{
	"TENANCY":     UpdateAclGroupDetailsAclGroupTypeTenancy,
	"COMPARTMENT": UpdateAclGroupDetailsAclGroupTypeCompartment,
	"BUCKET":      UpdateAclGroupDetailsAclGroupTypeBucket,
}

var mappingUpdateAclGroupDetailsAclGroupTypeEnumLowerCase = map[string]UpdateAclGroupDetailsAclGroupTypeEnum{
	"tenancy":     UpdateAclGroupDetailsAclGroupTypeTenancy,
	"compartment": UpdateAclGroupDetailsAclGroupTypeCompartment,
	"bucket":      UpdateAclGroupDetailsAclGroupTypeBucket,
}

// GetUpdateAclGroupDetailsAclGroupTypeEnumValues Enumerates the set of values for UpdateAclGroupDetailsAclGroupTypeEnum
func GetUpdateAclGroupDetailsAclGroupTypeEnumValues() []UpdateAclGroupDetailsAclGroupTypeEnum {
	values := make([]UpdateAclGroupDetailsAclGroupTypeEnum, 0)
	for _, v := range mappingUpdateAclGroupDetailsAclGroupTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateAclGroupDetailsAclGroupTypeEnumStringValues Enumerates the set of values in String for UpdateAclGroupDetailsAclGroupTypeEnum
func GetUpdateAclGroupDetailsAclGroupTypeEnumStringValues() []string {
	return []string{
		"TENANCY",
		"COMPARTMENT",
		"BUCKET",
	}
}

// GetMappingUpdateAclGroupDetailsAclGroupTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateAclGroupDetailsAclGroupTypeEnum(val string) (UpdateAclGroupDetailsAclGroupTypeEnum, bool) {
	enum, ok := mappingUpdateAclGroupDetailsAclGroupTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateAclGroupDetailsAclGroupExceptionsEnum Enum with underlying type: string
type UpdateAclGroupDetailsAclGroupExceptionsEnum string

// Set of constants representing the allowable values for UpdateAclGroupDetailsAclGroupExceptionsEnum
const (
	UpdateAclGroupDetailsAclGroupExceptionsCopy        UpdateAclGroupDetailsAclGroupExceptionsEnum = "COPY"
	UpdateAclGroupDetailsAclGroupExceptionsReplication UpdateAclGroupDetailsAclGroupExceptionsEnum = "REPLICATION"
)

var mappingUpdateAclGroupDetailsAclGroupExceptionsEnum = map[string]UpdateAclGroupDetailsAclGroupExceptionsEnum{
	"COPY":        UpdateAclGroupDetailsAclGroupExceptionsCopy,
	"REPLICATION": UpdateAclGroupDetailsAclGroupExceptionsReplication,
}

var mappingUpdateAclGroupDetailsAclGroupExceptionsEnumLowerCase = map[string]UpdateAclGroupDetailsAclGroupExceptionsEnum{
	"copy":        UpdateAclGroupDetailsAclGroupExceptionsCopy,
	"replication": UpdateAclGroupDetailsAclGroupExceptionsReplication,
}

// GetUpdateAclGroupDetailsAclGroupExceptionsEnumValues Enumerates the set of values for UpdateAclGroupDetailsAclGroupExceptionsEnum
func GetUpdateAclGroupDetailsAclGroupExceptionsEnumValues() []UpdateAclGroupDetailsAclGroupExceptionsEnum {
	values := make([]UpdateAclGroupDetailsAclGroupExceptionsEnum, 0)
	for _, v := range mappingUpdateAclGroupDetailsAclGroupExceptionsEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateAclGroupDetailsAclGroupExceptionsEnumStringValues Enumerates the set of values in String for UpdateAclGroupDetailsAclGroupExceptionsEnum
func GetUpdateAclGroupDetailsAclGroupExceptionsEnumStringValues() []string {
	return []string{
		"COPY",
		"REPLICATION",
	}
}

// GetMappingUpdateAclGroupDetailsAclGroupExceptionsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateAclGroupDetailsAclGroupExceptionsEnum(val string) (UpdateAclGroupDetailsAclGroupExceptionsEnum, bool) {
	enum, ok := mappingUpdateAclGroupDetailsAclGroupExceptionsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
