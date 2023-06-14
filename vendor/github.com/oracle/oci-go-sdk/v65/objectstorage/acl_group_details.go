// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// AclGroupDetails The details to create or update an ACL Group.
type AclGroupDetails struct {

	// User-friendly name to describe the ACL Group. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A description of the ACL Group. This can be useful for identifying the purpose of the ACL Group.
	Description *string `mandatory:"false" json:"description"`

	// The compartment containing this ACL Group.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Determines whether this ACL Group can be associated with a bucket, compartment, or tenancy.
	AclGroupType AclGroupDetailsAclGroupTypeEnum `mandatory:"false" json:"aclGroupType,omitempty"`

	// Specifies requests that are not subject to any ACLs in the ACL group.
	AclGroupExceptions []AclGroupDetailsAclGroupExceptionsEnum `mandatory:"false" json:"aclGroupExceptions,omitempty"`

	// An ordered list of ACL IDs. The ACLs are evaluated first to last.
	AclIds []string `mandatory:"false" json:"aclIds"`
}

func (m AclGroupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AclGroupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAclGroupDetailsAclGroupTypeEnum(string(m.AclGroupType)); !ok && m.AclGroupType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AclGroupType: %s. Supported values are: %s.", m.AclGroupType, strings.Join(GetAclGroupDetailsAclGroupTypeEnumStringValues(), ",")))
	}
	for _, val := range m.AclGroupExceptions {
		if _, ok := GetMappingAclGroupDetailsAclGroupExceptionsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AclGroupExceptions: %s. Supported values are: %s.", val, strings.Join(GetAclGroupDetailsAclGroupExceptionsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AclGroupDetailsAclGroupTypeEnum Enum with underlying type: string
type AclGroupDetailsAclGroupTypeEnum string

// Set of constants representing the allowable values for AclGroupDetailsAclGroupTypeEnum
const (
	AclGroupDetailsAclGroupTypeTenancy     AclGroupDetailsAclGroupTypeEnum = "TENANCY"
	AclGroupDetailsAclGroupTypeCompartment AclGroupDetailsAclGroupTypeEnum = "COMPARTMENT"
	AclGroupDetailsAclGroupTypeBucket      AclGroupDetailsAclGroupTypeEnum = "BUCKET"
)

var mappingAclGroupDetailsAclGroupTypeEnum = map[string]AclGroupDetailsAclGroupTypeEnum{
	"TENANCY":     AclGroupDetailsAclGroupTypeTenancy,
	"COMPARTMENT": AclGroupDetailsAclGroupTypeCompartment,
	"BUCKET":      AclGroupDetailsAclGroupTypeBucket,
}

var mappingAclGroupDetailsAclGroupTypeEnumLowerCase = map[string]AclGroupDetailsAclGroupTypeEnum{
	"tenancy":     AclGroupDetailsAclGroupTypeTenancy,
	"compartment": AclGroupDetailsAclGroupTypeCompartment,
	"bucket":      AclGroupDetailsAclGroupTypeBucket,
}

// GetAclGroupDetailsAclGroupTypeEnumValues Enumerates the set of values for AclGroupDetailsAclGroupTypeEnum
func GetAclGroupDetailsAclGroupTypeEnumValues() []AclGroupDetailsAclGroupTypeEnum {
	values := make([]AclGroupDetailsAclGroupTypeEnum, 0)
	for _, v := range mappingAclGroupDetailsAclGroupTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAclGroupDetailsAclGroupTypeEnumStringValues Enumerates the set of values in String for AclGroupDetailsAclGroupTypeEnum
func GetAclGroupDetailsAclGroupTypeEnumStringValues() []string {
	return []string{
		"TENANCY",
		"COMPARTMENT",
		"BUCKET",
	}
}

// GetMappingAclGroupDetailsAclGroupTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAclGroupDetailsAclGroupTypeEnum(val string) (AclGroupDetailsAclGroupTypeEnum, bool) {
	enum, ok := mappingAclGroupDetailsAclGroupTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AclGroupDetailsAclGroupExceptionsEnum Enum with underlying type: string
type AclGroupDetailsAclGroupExceptionsEnum string

// Set of constants representing the allowable values for AclGroupDetailsAclGroupExceptionsEnum
const (
	AclGroupDetailsAclGroupExceptionsCopy        AclGroupDetailsAclGroupExceptionsEnum = "COPY"
	AclGroupDetailsAclGroupExceptionsReplication AclGroupDetailsAclGroupExceptionsEnum = "REPLICATION"
)

var mappingAclGroupDetailsAclGroupExceptionsEnum = map[string]AclGroupDetailsAclGroupExceptionsEnum{
	"COPY":        AclGroupDetailsAclGroupExceptionsCopy,
	"REPLICATION": AclGroupDetailsAclGroupExceptionsReplication,
}

var mappingAclGroupDetailsAclGroupExceptionsEnumLowerCase = map[string]AclGroupDetailsAclGroupExceptionsEnum{
	"copy":        AclGroupDetailsAclGroupExceptionsCopy,
	"replication": AclGroupDetailsAclGroupExceptionsReplication,
}

// GetAclGroupDetailsAclGroupExceptionsEnumValues Enumerates the set of values for AclGroupDetailsAclGroupExceptionsEnum
func GetAclGroupDetailsAclGroupExceptionsEnumValues() []AclGroupDetailsAclGroupExceptionsEnum {
	values := make([]AclGroupDetailsAclGroupExceptionsEnum, 0)
	for _, v := range mappingAclGroupDetailsAclGroupExceptionsEnum {
		values = append(values, v)
	}
	return values
}

// GetAclGroupDetailsAclGroupExceptionsEnumStringValues Enumerates the set of values in String for AclGroupDetailsAclGroupExceptionsEnum
func GetAclGroupDetailsAclGroupExceptionsEnumStringValues() []string {
	return []string{
		"COPY",
		"REPLICATION",
	}
}

// GetMappingAclGroupDetailsAclGroupExceptionsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAclGroupDetailsAclGroupExceptionsEnum(val string) (AclGroupDetailsAclGroupExceptionsEnum, bool) {
	enum, ok := mappingAclGroupDetailsAclGroupExceptionsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
